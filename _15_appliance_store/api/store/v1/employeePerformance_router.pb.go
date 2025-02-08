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

type EmployeePerformanceLogicer interface {
	Create(ctx context.Context, req *CreateEmployeePerformanceRequest) (*CreateEmployeePerformanceReply, error)
	DeleteByID(ctx context.Context, req *DeleteEmployeePerformanceByIDRequest) (*DeleteEmployeePerformanceByIDReply, error)
	UpdateByID(ctx context.Context, req *UpdateEmployeePerformanceByIDRequest) (*UpdateEmployeePerformanceByIDReply, error)
	GetByID(ctx context.Context, req *GetEmployeePerformanceByIDRequest) (*GetEmployeePerformanceByIDReply, error)
	List(ctx context.Context, req *ListEmployeePerformanceRequest) (*ListEmployeePerformanceReply, error)
	CalculateSalesCommission(ctx context.Context, req *CalculateSalesCommissionRequest) (*CalculateSalesCommissionReply, error)
}

type EmployeePerformanceOption func(*employeePerformanceOptions)

type employeePerformanceOptions struct {
	isFromRPC  bool
	responser  errcode.Responser
	zapLog     *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *employeePerformanceOptions) apply(opts ...EmployeePerformanceOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithEmployeePerformanceHTTPResponse() EmployeePerformanceOption {
	return func(o *employeePerformanceOptions) {
		o.isFromRPC = false
	}
}

func WithEmployeePerformanceRPCResponse() EmployeePerformanceOption {
	return func(o *employeePerformanceOptions) {
		o.isFromRPC = true
	}
}

func WithEmployeePerformanceResponser(responser errcode.Responser) EmployeePerformanceOption {
	return func(o *employeePerformanceOptions) {
		o.responser = responser
	}
}

func WithEmployeePerformanceLogger(zapLog *zap.Logger) EmployeePerformanceOption {
	return func(o *employeePerformanceOptions) {
		o.zapLog = zapLog
	}
}

func WithEmployeePerformanceErrorToHTTPCode(e ...*errcode.Error) EmployeePerformanceOption {
	return func(o *employeePerformanceOptions) {
		o.httpErrors = e
	}
}

func WithEmployeePerformanceRPCStatusToHTTPCode(s ...*errcode.RPCStatus) EmployeePerformanceOption {
	return func(o *employeePerformanceOptions) {
		o.rpcStatus = s
	}
}

func WithEmployeePerformanceWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) EmployeePerformanceOption {
	return func(o *employeePerformanceOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterEmployeePerformanceRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic EmployeePerformanceLogicer,
	opts ...EmployeePerformanceOption) {

	o := &employeePerformanceOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog, _ = zap.NewProduction()
	}

	r := &employeePerformanceRouter{
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

type employeePerformanceRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                EmployeePerformanceLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *employeePerformanceRouter) register() {
	r.iRouter.Handle("POST", "/api/v1/employeePerformance", r.withMiddleware("POST", "/api/v1/employeePerformance", r.Create_14)...)
	r.iRouter.Handle("DELETE", "/api/v1/employeePerformance/:id", r.withMiddleware("DELETE", "/api/v1/employeePerformance/:id", r.DeleteByID_12)...)
	r.iRouter.Handle("PUT", "/api/v1/employeePerformance/:id", r.withMiddleware("PUT", "/api/v1/employeePerformance/:id", r.UpdateByID_12)...)
	r.iRouter.Handle("GET", "/api/v1/employeePerformance/:id", r.withMiddleware("GET", "/api/v1/employeePerformance/:id", r.GetByID_12)...)
	r.iRouter.Handle("POST", "/api/v1/employeePerformance/list", r.withMiddleware("POST", "/api/v1/employeePerformance/list", r.List_14)...)
	r.iRouter.Handle("POST", "/api/v1/employeePerformance/calculate-commission", r.withMiddleware("POST", "/api/v1/employeePerformance/calculate-commission", r.CalculateSalesCommission_0)...)

}

func (r *employeePerformanceRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
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

func (r *employeePerformanceRouter) Create_14(c *gin.Context) {
	req := &CreateEmployeePerformanceRequest{}
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

func (r *employeePerformanceRouter) DeleteByID_12(c *gin.Context) {
	req := &DeleteEmployeePerformanceByIDRequest{}
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

func (r *employeePerformanceRouter) UpdateByID_12(c *gin.Context) {
	req := &UpdateEmployeePerformanceByIDRequest{}
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

func (r *employeePerformanceRouter) GetByID_12(c *gin.Context) {
	req := &GetEmployeePerformanceByIDRequest{}
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

func (r *employeePerformanceRouter) List_14(c *gin.Context) {
	req := &ListEmployeePerformanceRequest{}
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

func (r *employeePerformanceRouter) CalculateSalesCommission_0(c *gin.Context) {
	req := &CalculateSalesCommissionRequest{}
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

	out, err := r.iLogic.CalculateSalesCommission(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}
