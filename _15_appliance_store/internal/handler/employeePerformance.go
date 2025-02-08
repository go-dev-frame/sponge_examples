package handler

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/copier"

	"github.com/go-dev-frame/sponge/pkg/gin/middleware"
	"github.com/go-dev-frame/sponge/pkg/logger"
	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
	"github.com/go-dev-frame/sponge/pkg/utils"

	storeV1 "store/api/store/v1"
	"store/internal/cache"
	"store/internal/dao"
	"store/internal/database"
	"store/internal/ecode"
	"store/internal/model"
)

var _ storeV1.EmployeePerformanceLogicer = (*employeePerformanceHandler)(nil)
var _ time.Time

type employeePerformanceHandler struct {
	employeePerformanceDao dao.EmployeePerformanceDao
}

// NewEmployeePerformanceHandler create a handler
func NewEmployeePerformanceHandler() storeV1.EmployeePerformanceLogicer {
	return &employeePerformanceHandler{
		employeePerformanceDao: dao.NewEmployeePerformanceDao(
			database.GetDB(), // db driver is mysql
			cache.NewEmployeePerformanceCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *employeePerformanceHandler) Create(ctx context.Context, req *storeV1.CreateEmployeePerformanceRequest) (*storeV1.CreateEmployeePerformanceReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	employeePerformance := &model.EmployeePerformance{}
	err = copier.Copy(employeePerformance, req)
	if err != nil {
		return nil, ecode.ErrCreateEmployeePerformance.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.employeePerformanceDao.Create(ctx, employeePerformance)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("employeePerformance", employeePerformance), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateEmployeePerformanceReply{Id: employeePerformance.ID}, nil
}

// DeleteByID delete a record by id
func (h *employeePerformanceHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteEmployeePerformanceByIDRequest) (*storeV1.DeleteEmployeePerformanceByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.employeePerformanceDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteEmployeePerformanceByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *employeePerformanceHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateEmployeePerformanceByIDRequest) (*storeV1.UpdateEmployeePerformanceByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	employeePerformance := &model.EmployeePerformance{}
	err = copier.Copy(employeePerformance, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDEmployeePerformance.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	employeePerformance.ID = req.Id

	err = h.employeePerformanceDao.UpdateByID(ctx, employeePerformance)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("employeePerformance", employeePerformance), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateEmployeePerformanceByIDReply{}, nil
}

// GetByID get a record by id
func (h *employeePerformanceHandler) GetByID(ctx context.Context, req *storeV1.GetEmployeePerformanceByIDRequest) (*storeV1.GetEmployeePerformanceByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.employeePerformanceDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertEmployeePerformance(record)
	if err != nil {
		logger.Warn("convertEmployeePerformance error", logger.Err(err), logger.Any("employeePerformance", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDEmployeePerformance.Err()
	}

	return &storeV1.GetEmployeePerformanceByIDReply{
		EmployeePerformance: data,
	}, nil
}

// List of records by query parameters
func (h *employeePerformanceHandler) List(ctx context.Context, req *storeV1.ListEmployeePerformanceRequest) (*storeV1.ListEmployeePerformanceReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListEmployeePerformance.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.employeePerformanceDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	employeePerformances := []*storeV1.EmployeePerformance{}
	for _, record := range records {
		data, err := convertEmployeePerformance(record)
		if err != nil {
			logger.Warn("convertEmployeePerformance error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		employeePerformances = append(employeePerformances, data)
	}

	return &storeV1.ListEmployeePerformanceReply{
		Total:                total,
		EmployeePerformances: employeePerformances,
	}, nil
}

// CalculateSalesCommission 计算销售提成
// 实现逻辑：
// 1. 获取员工当月的销售记录
// 2. 应用阶梯提成规则计算奖金
// 3. 扣除已退货订单金额
// 4. 生成提成明细记录
func (h *employeePerformanceHandler) CalculateSalesCommission(ctx context.Context, req *storeV1.CalculateSalesCommissionRequest) (*storeV1.CalculateSalesCommissionReply, error) {
	panic("prompt: 计算销售提成  实现逻辑：  1. 获取员工当月的销售记录  2. 应用阶梯提成规则计算奖金  3. 扣除已退货订单金额  4. 生成提成明细记录")

	// fill in the business logic code here
	// example:
	//
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
	//		    return nil, ecode.InvalidParams.Err()
	//	    }
	//
	//	    reply, err := h.employeePerformanceDao.CalculateSalesCommission(ctx, &model.EmployeePerformance{
	//     	EmployeeID: req.EmployeeID,
	//     	Month: req.Month,
	//     })
	//	    if err != nil {
	//			logger.Warn("CalculateSalesCommission error", logger.Err(err), middleware.CtxRequestIDField(ctx))
	//			return nil, ecode.InternalServerError.Err()
	//		}
	//
	//     return &storeV1.CalculateSalesCommissionReply{
	//     	TotalSales: reply.TotalSales,
	//     	CommissionAmount: reply.CommissionAmount,
	//     	QualifiedOrders: reply.QualifiedOrders,
	//     }, nil
}

func convertEmployeePerformance(record *model.EmployeePerformance) (*storeV1.EmployeePerformance, error) {
	value := &storeV1.EmployeePerformance{}
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
