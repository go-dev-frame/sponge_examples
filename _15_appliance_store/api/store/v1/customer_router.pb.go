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

type CustomerLogicer interface {
	Create(ctx context.Context, req *CreateCustomerRequest) (*CreateCustomerReply, error)
	DeleteByID(ctx context.Context, req *DeleteCustomerByIDRequest) (*DeleteCustomerByIDReply, error)
	UpdateByID(ctx context.Context, req *UpdateCustomerByIDRequest) (*UpdateCustomerByIDReply, error)
	GetByID(ctx context.Context, req *GetCustomerByIDRequest) (*GetCustomerByIDReply, error)
	List(ctx context.Context, req *ListCustomerRequest) (*ListCustomerReply, error)
}

type CustomerOption func(*customerOptions)

type customerOptions struct {
	isFromRPC  bool
	responser  errcode.Responser
	zapLog     *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *customerOptions) apply(opts ...CustomerOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithCustomerHTTPResponse() CustomerOption {
	return func(o *customerOptions) {
		o.isFromRPC = false
	}
}

func WithCustomerRPCResponse() CustomerOption {
	return func(o *customerOptions) {
		o.isFromRPC = true
	}
}

func WithCustomerResponser(responser errcode.Responser) CustomerOption {
	return func(o *customerOptions) {
		o.responser = responser
	}
}

func WithCustomerLogger(zapLog *zap.Logger) CustomerOption {
	return func(o *customerOptions) {
		o.zapLog = zapLog
	}
}

func WithCustomerErrorToHTTPCode(e ...*errcode.Error) CustomerOption {
	return func(o *customerOptions) {
		o.httpErrors = e
	}
}

func WithCustomerRPCStatusToHTTPCode(s ...*errcode.RPCStatus) CustomerOption {
	return func(o *customerOptions) {
		o.rpcStatus = s
	}
}

func WithCustomerWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) CustomerOption {
	return func(o *customerOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterCustomerRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic CustomerLogicer,
	opts ...CustomerOption) {

	o := &customerOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog, _ = zap.NewProduction()
	}

	r := &customerRouter{
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

type customerRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                CustomerLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *customerRouter) register() {
	r.iRouter.Handle("POST", "/api/v1/customer", r.withMiddleware("POST", "/api/v1/customer", r.Create_10)...)
	r.iRouter.Handle("DELETE", "/api/v1/customer/:id", r.withMiddleware("DELETE", "/api/v1/customer/:id", r.DeleteByID_8)...)
	r.iRouter.Handle("PUT", "/api/v1/customer/:id", r.withMiddleware("PUT", "/api/v1/customer/:id", r.UpdateByID_8)...)
	r.iRouter.Handle("GET", "/api/v1/customer/:id", r.withMiddleware("GET", "/api/v1/customer/:id", r.GetByID_8)...)
	r.iRouter.Handle("POST", "/api/v1/customer/list", r.withMiddleware("POST", "/api/v1/customer/list", r.List_10)...)

}

func (r *customerRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
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

func (r *customerRouter) Create_10(c *gin.Context) {
	req := &CreateCustomerRequest{}
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

func (r *customerRouter) DeleteByID_8(c *gin.Context) {
	req := &DeleteCustomerByIDRequest{}
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

func (r *customerRouter) UpdateByID_8(c *gin.Context) {
	req := &UpdateCustomerByIDRequest{}
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

func (r *customerRouter) GetByID_8(c *gin.Context) {
	req := &GetCustomerByIDRequest{}
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

func (r *customerRouter) List_10(c *gin.Context) {
	req := &ListCustomerRequest{}
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
