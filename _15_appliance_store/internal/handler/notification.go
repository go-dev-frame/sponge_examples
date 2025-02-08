package handler

import (
	"context"
	"errors"
	"github.com/go-dev-frame/sponge/pkg/utils"
	"strings"
	"time"

	"github.com/jinzhu/copier"

	"github.com/go-dev-frame/sponge/pkg/gin/middleware"
	"github.com/go-dev-frame/sponge/pkg/logger"
	"github.com/go-dev-frame/sponge/pkg/sgorm/query"

	storeV1 "store/api/store/v1"
	"store/internal/cache"
	"store/internal/dao"
	"store/internal/database"
	"store/internal/ecode"
	"store/internal/model"
)

var _ storeV1.NotificationLogicer = (*notificationHandler)(nil)
var _ time.Time

type notificationHandler struct {
	notificationDao dao.NotificationDao
}

// NewNotificationHandler create a handler
func NewNotificationHandler() storeV1.NotificationLogicer {
	return &notificationHandler{
		notificationDao: dao.NewNotificationDao(
			database.GetDB(), // db driver is mysql
			cache.NewNotificationCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *notificationHandler) Create(ctx context.Context, req *storeV1.CreateNotificationRequest) (*storeV1.CreateNotificationReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	notification := &model.Notification{}
	err = copier.Copy(notification, req)
	if err != nil {
		return nil, ecode.ErrCreateNotification.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.notificationDao.Create(ctx, notification)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("notification", notification), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateNotificationReply{Id: notification.ID}, nil
}

// DeleteByID delete a record by id
func (h *notificationHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteNotificationByIDRequest) (*storeV1.DeleteNotificationByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.notificationDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteNotificationByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *notificationHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateNotificationByIDRequest) (*storeV1.UpdateNotificationByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	notification := &model.Notification{}
	err = copier.Copy(notification, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDNotification.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	notification.ID = req.Id

	err = h.notificationDao.UpdateByID(ctx, notification)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("notification", notification), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateNotificationByIDReply{}, nil
}

// GetByID get a record by id
func (h *notificationHandler) GetByID(ctx context.Context, req *storeV1.GetNotificationByIDRequest) (*storeV1.GetNotificationByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.notificationDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertNotification(record)
	if err != nil {
		logger.Warn("convertNotification error", logger.Err(err), logger.Any("notification", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDNotification.Err()
	}

	return &storeV1.GetNotificationByIDReply{
		Notification: data,
	}, nil
}

// List of records by query parameters
func (h *notificationHandler) List(ctx context.Context, req *storeV1.ListNotificationRequest) (*storeV1.ListNotificationReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListNotification.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.notificationDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	notifications := []*storeV1.Notification{}
	for _, record := range records {
		data, err := convertNotification(record)
		if err != nil {
			logger.Warn("convertNotification error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		notifications = append(notifications, data)
	}

	return &storeV1.ListNotificationReply{
		Total:         total,
		Notifications: notifications,
	}, nil
}

// SendRealTimeNotification 发送实时通知
func (h *notificationHandler) SendRealTimeNotification(ctx context.Context, req *storeV1.SendNotificationRequest) (*storeV1.SendNotificationResponse, error) {
	panic("prompt: 发送实时通知")

	// fill in the business logic code here
	// example:
	//
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
	//		    return nil, ecode.InvalidParams.Err()
	//	    }
	//
	//	    reply, err := h.notificationDao.SendRealTimeNotification(ctx, &model.Notification{
	//     	Channel: req.Channel,
	//     	Recipient: req.Recipient,
	//     	TemplateID: req.TemplateID,
	//     	Params: req.Params,
	//     })
	//	    if err != nil {
	//			logger.Warn("SendRealTimeNotification error", logger.Err(err), middleware.CtxRequestIDField(ctx))
	//			return nil, ecode.InternalServerError.Err()
	//		}
	//
	//     return &storeV1.SendNotificationResponse{
	//     	MessageID: reply.MessageID,
	//     	SendTime: reply.SendTime,
	//     }, nil
}

func convertNotification(record *model.Notification) (*storeV1.Notification, error) {
	value := &storeV1.Notification{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.Id = record.ID
	value.CreatedAt = utils.FormatDateTimeRFC3339(*record.CreatedAt)
	value.UpdatedAt = utils.FormatDateTimeRFC3339(*record.UpdatedAt)

	return value, nil
}
