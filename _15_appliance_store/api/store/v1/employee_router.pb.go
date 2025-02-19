// Code generated by https://github.com/go-dev-frame/sponge, DO NOT EDIT.

package v1

import (
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/go-dev-frame/sponge/pkg/errcode"
	"github.com/go-dev-frame/sponge/pkg/gin/middleware"
)

type EmployeeLogicer interface {
	Create(ctx context.Context, req *CreateEmployeeRequest) (*CreateEmployeeReply, error)
	DeleteByID(ctx context.Context, req *DeleteEmployeeByIDRequest) (*DeleteEmployeeByIDReply, error)
	UpdateByID(ctx context.Context, req *UpdateEmployeeByIDRequest) (*UpdateEmployeeByIDReply, error)
	GetByID(ctx context.Context, req *GetEmployeeByIDRequest) (*GetEmployeeByIDReply, error)
	List(ctx context.Context, req *ListEmployeeRequest) (*ListEmployeeReply, error)
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	Logout(ctx context.Context, req *Empty) (*Empty, error)
	ChangePassword(ctx context.Context, req *ChangePasswordRequest) (*Empty, error)
	ResetPassword(ctx context.Context, req *ResetPasswordRequest) (*Empty, error)
	SendVerificationCode(ctx context.Context, req *SendVerificationCodeRequest) (*Empty, error)
}

type EmployeeOption func(*employeeOptions)

type employeeOptions struct {
	isFromRPC  bool
	responser  errcode.Responser
	zapLog     *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *employeeOptions) apply(opts ...EmployeeOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithEmployeeHTTPResponse() EmployeeOption {
	return func(o *employeeOptions) {
		o.isFromRPC = false
	}
}

func WithEmployeeRPCResponse() EmployeeOption {
	return func(o *employeeOptions) {
		o.isFromRPC = true
	}
}

func WithEmployeeResponser(responser errcode.Responser) EmployeeOption {
	return func(o *employeeOptions) {
		o.responser = responser
	}
}

func WithEmployeeLogger(zapLog *zap.Logger) EmployeeOption {
	return func(o *employeeOptions) {
		o.zapLog = zapLog
	}
}

func WithEmployeeErrorToHTTPCode(e ...*errcode.Error) EmployeeOption {
	return func(o *employeeOptions) {
		o.httpErrors = e
	}
}

func WithEmployeeRPCStatusToHTTPCode(s ...*errcode.RPCStatus) EmployeeOption {
	return func(o *employeeOptions) {
		o.rpcStatus = s
	}
}

func WithEmployeeWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) EmployeeOption {
	return func(o *employeeOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterEmployeeRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic EmployeeLogicer,
	opts ...EmployeeOption) {

	o := &employeeOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog, _ = zap.NewProduction()
	}

	r := &employeeRouter{
		iRouter:               iRouter,
		groupPathMiddlewares:  groupPathMiddlewares,
		singlePathMiddlewares: singlePathMiddlewares,
		iLogic:                iLogic,
		iResponse:             o.responser,
		zapLog:                o.zapLog,
		wrapCtxFn:             o.wrapCtxFn,
	}
	r.register()
}

type employeeRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                EmployeeLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *employeeRouter) register() {
	r.iRouter.Handle("POST", "/api/v1/employee", r.withMiddleware("POST", "/api/v1/employee", r.Create_12)...)
	r.iRouter.Handle("DELETE", "/api/v1/employee/:id", r.withMiddleware("DELETE", "/api/v1/employee/:id", r.DeleteByID_10)...)
	r.iRouter.Handle("PUT", "/api/v1/employee/:id", r.withMiddleware("PUT", "/api/v1/employee/:id", r.UpdateByID_10)...)
	r.iRouter.Handle("GET", "/api/v1/employee/:id", r.withMiddleware("GET", "/api/v1/employee/:id", r.GetByID_10)...)
	r.iRouter.Handle("POST", "/api/v1/employee/list", r.withMiddleware("POST", "/api/v1/employee/list", r.List_12)...)
	r.iRouter.Handle("POST", "/api/v1/employee/login", r.withMiddleware("POST", "/api/v1/employee/login", r.Login_0)...)
	r.iRouter.Handle("POST", "/api/v1/employee/logout", r.withMiddleware("POST", "/api/v1/employee/logout", r.Logout_0)...)
	r.iRouter.Handle("POST", "/api/v1/employee/change-password", r.withMiddleware("POST", "/api/v1/employee/change-password", r.ChangePassword_0)...)
	r.iRouter.Handle("POST", "/api/v1/employee/reset-password", r.withMiddleware("POST", "/api/v1/employee/reset-password", r.ResetPassword_0)...)
	r.iRouter.Handle("POST", "/api/v1/employee/send-verification-code", r.withMiddleware("POST", "/api/v1/employee/send-verification-code", r.SendVerificationCode_0)...)

}

func (r *employeeRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
	handlerFns := []gin.HandlerFunc{}

	// determine if a route group is hit or miss, left prefix rule
	for groupPath, fns := range r.groupPathMiddlewares {
		if groupPath == "" || groupPath == "/" {
			handlerFns = append(handlerFns, fns...)
			continue
		}
		size := len(groupPath)
		if len(path) < size {
			continue
		}
		if groupPath == path[:size] {
			handlerFns = append(handlerFns, fns...)
		}
	}

	// determine if a single route has been hit
	key := strings.ToUpper(method) + "->" + path
	if fns, ok := r.singlePathMiddlewares[key]; ok {
		handlerFns = append(handlerFns, fns...)
	}

	return append(handlerFns, fn)
}

func (r *employeeRouter) Create_12(c *gin.Context) {
	req := &CreateEmployeeRequest{}
	var err error

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.Create(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *employeeRouter) DeleteByID_10(c *gin.Context) {
	req := &DeleteEmployeeByIDRequest{}
	var err error

	if err = c.ShouldBindUri(req); err != nil {
		r.zapLog.Warn("ShouldBindUri error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	if err = c.ShouldBindQuery(req); err != nil {
		r.zapLog.Warn("ShouldBindQuery error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.DeleteByID(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *employeeRouter) UpdateByID_10(c *gin.Context) {
	req := &UpdateEmployeeByIDRequest{}
	var err error

	if err = c.ShouldBindUri(req); err != nil {
		r.zapLog.Warn("ShouldBindUri error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.UpdateByID(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *employeeRouter) GetByID_10(c *gin.Context) {
	req := &GetEmployeeByIDRequest{}
	var err error

	if err = c.ShouldBindUri(req); err != nil {
		r.zapLog.Warn("ShouldBindUri error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	if err = c.ShouldBindQuery(req); err != nil {
		r.zapLog.Warn("ShouldBindQuery error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.GetByID(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *employeeRouter) List_12(c *gin.Context) {
	req := &ListEmployeeRequest{}
	var err error

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.List(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *employeeRouter) Login_0(c *gin.Context) {
	req := &LoginRequest{}
	var err error

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.Login(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *employeeRouter) Logout_0(c *gin.Context) {
	req := &Empty{}
	var err error

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.Logout(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *employeeRouter) ChangePassword_0(c *gin.Context) {
	req := &ChangePasswordRequest{}
	var err error

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.ChangePassword(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *employeeRouter) ResetPassword_0(c *gin.Context) {
	req := &ResetPasswordRequest{}
	var err error

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.ResetPassword(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *employeeRouter) SendVerificationCode_0(c *gin.Context) {
	req := &SendVerificationCodeRequest{}
	var err error

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.SendVerificationCode(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}
