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

var _ storeV1.EmployeeLogicer = (*employeeHandler)(nil)
var _ time.Time

type employeeHandler struct {
	employeeDao dao.EmployeeDao
}

// NewEmployeeHandler create a handler
func NewEmployeeHandler() storeV1.EmployeeLogicer {
	return &employeeHandler{
		employeeDao: dao.NewEmployeeDao(
			database.GetDB(), // db driver is mysql
			cache.NewEmployeeCache(database.GetCacheType()),
		),
	}
}

// Create a record
func (h *employeeHandler) Create(ctx context.Context, req *storeV1.CreateEmployeeRequest) (*storeV1.CreateEmployeeReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	employee := &model.Employee{}
	err = copier.Copy(employee, req)
	if err != nil {
		return nil, ecode.ErrCreateEmployee.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = h.employeeDao.Create(ctx, employee)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("employee", employee), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.CreateEmployeeReply{Id: employee.ID}, nil
}

// DeleteByID delete a record by id
func (h *employeeHandler) DeleteByID(ctx context.Context, req *storeV1.DeleteEmployeeByIDRequest) (*storeV1.DeleteEmployeeByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	err = h.employeeDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Warn("DeleteByID error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.DeleteEmployeeByIDReply{}, nil
}

// UpdateByID update a record by id
func (h *employeeHandler) UpdateByID(ctx context.Context, req *storeV1.UpdateEmployeeByIDRequest) (*storeV1.UpdateEmployeeByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	employee := &model.Employee{}
	err = copier.Copy(employee, req)
	if err != nil {
		return nil, ecode.ErrUpdateByIDEmployee.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	employee.ID = req.Id

	err = h.employeeDao.UpdateByID(ctx, employee)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("employee", employee), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &storeV1.UpdateEmployeeByIDReply{}, nil
}

// GetByID get a record by id
func (h *employeeHandler) GetByID(ctx context.Context, req *storeV1.GetEmployeeByIDRequest) (*storeV1.GetEmployeeByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	record, err := h.employeeDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
			return nil, ecode.NotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	data, err := convertEmployee(record)
	if err != nil {
		logger.Warn("convertEmployee error", logger.Err(err), logger.Any("employee", record), middleware.CtxRequestIDField(ctx))
		return nil, ecode.ErrGetByIDEmployee.Err()
	}

	return &storeV1.GetEmployeeByIDReply{
		Employee: data,
	}, nil
}

// List of records by query parameters
func (h *employeeHandler) List(ctx context.Context, req *storeV1.ListEmployeeRequest) (*storeV1.ListEmployeeReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.ErrListEmployee.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := h.employeeDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
			return nil, ecode.InvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	employees := []*storeV1.Employee{}
	for _, record := range records {
		data, err := convertEmployee(record)
		if err != nil {
			logger.Warn("convertEmployee error", logger.Err(err), logger.Any("id", record.ID), middleware.CtxRequestIDField(ctx))
			continue
		}
		employees = append(employees, data)
	}

	return &storeV1.ListEmployeeReply{
		Total:     total,
		Employees: employees,
	}, nil
}

// Login 员工登录，获取访问令牌
func (h *employeeHandler) Login(ctx context.Context, req *storeV1.LoginRequest) (*storeV1.LoginResponse, error) {
	panic("prompt: 员工登录，获取访问令牌")

	// fill in the business logic code here
	// example:
	//
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
	//		    return nil, ecode.InvalidParams.Err()
	//	    }
	//
	//	    reply, err := h.employeeDao.Login(ctx, &model.Employee{
	//     	Phone: req.Phone,
	//     	Password: req.Password,
	//     })
	//	    if err != nil {
	//			logger.Warn("Login error", logger.Err(err), middleware.CtxRequestIDField(ctx))
	//			return nil, ecode.InternalServerError.Err()
	//		}
	//
	//     return &storeV1.LoginResponse{
	//     	Token: reply.Token,
	//     }, nil
}

// Logout 员工退出登录
func (h *employeeHandler) Logout(ctx context.Context, req *storeV1.Empty) (*storeV1.Empty, error) {
	panic("prompt: 员工退出登录")

	// fill in the business logic code here
	// example:
	//
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
	//		    return nil, ecode.InvalidParams.Err()
	//	    }
	//
	//	    reply, err := h.employeeDao.Logout(ctx, &model.Employee{
	//     })
	//	    if err != nil {
	//			logger.Warn("Logout error", logger.Err(err), middleware.CtxRequestIDField(ctx))
	//			return nil, ecode.InternalServerError.Err()
	//		}
	//
	//     return &storeV1.Empty{
	//     }, nil
}

// ChangePassword 修改密码
func (h *employeeHandler) ChangePassword(ctx context.Context, req *storeV1.ChangePasswordRequest) (*storeV1.Empty, error) {
	panic("prompt: 修改密码")

	// fill in the business logic code here
	// example:
	//
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
	//		    return nil, ecode.InvalidParams.Err()
	//	    }
	//
	//	    reply, err := h.employeeDao.ChangePassword(ctx, &model.Employee{
	//     	OldPassword: req.OldPassword,
	//     	NewPassword: req.NewPassword,
	//     })
	//	    if err != nil {
	//			logger.Warn("ChangePassword error", logger.Err(err), middleware.CtxRequestIDField(ctx))
	//			return nil, ecode.InternalServerError.Err()
	//		}
	//
	//     return &storeV1.Empty{
	//     }, nil
}

// ResetPassword 重置密码
func (h *employeeHandler) ResetPassword(ctx context.Context, req *storeV1.ResetPasswordRequest) (*storeV1.Empty, error) {
	panic("prompt: 重置密码")

	// fill in the business logic code here
	// example:
	//
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
	//		    return nil, ecode.InvalidParams.Err()
	//	    }
	//
	//	    reply, err := h.employeeDao.ResetPassword(ctx, &model.Employee{
	//     	Phone: req.Phone,
	//     	Code: req.Code,
	//     	NewPassword: req.NewPassword,
	//     })
	//	    if err != nil {
	//			logger.Warn("ResetPassword error", logger.Err(err), middleware.CtxRequestIDField(ctx))
	//			return nil, ecode.InternalServerError.Err()
	//		}
	//
	//     return &storeV1.Empty{
	//     }, nil
}

// SendVerificationCode 使用阿里云短信服务，向指定手机号发送验证码，验证码存放在redis中，5分钟内有效
func (h *employeeHandler) SendVerificationCode(ctx context.Context, req *storeV1.SendVerificationCodeRequest) (*storeV1.Empty, error) {
	panic("prompt: 使用阿里云短信服务，向指定手机号发送验证码，验证码存放在redis中，5分钟内有效")

	// fill in the business logic code here
	// example:
	//
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
	//		    return nil, ecode.InvalidParams.Err()
	//	    }
	//
	//	    reply, err := h.employeeDao.SendVerificationCode(ctx, &model.Employee{
	//     	Phone: req.Phone,
	//     })
	//	    if err != nil {
	//			logger.Warn("SendVerificationCode error", logger.Err(err), middleware.CtxRequestIDField(ctx))
	//			return nil, ecode.InternalServerError.Err()
	//		}
	//
	//     return &storeV1.Empty{
	//     }, nil
}

func convertEmployee(record *model.Employee) (*storeV1.Employee, error) {
	value := &storeV1.Employee{}
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
