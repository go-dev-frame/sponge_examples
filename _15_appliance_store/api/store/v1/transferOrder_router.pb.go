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

type TransferOrderLogicer interface {
	Create(ctx context.Context, req *CreateTransferOrderRequest) (*CreateTransferOrderReply, error)
	DeleteByID(ctx context.Context, req *DeleteTransferOrderByIDRequest) (*DeleteTransferOrderByIDReply, error)
	UpdateByID(ctx context.Context, req *UpdateTransferOrderByIDRequest) (*UpdateTransferOrderByIDReply, error)
	GetByID(ctx context.Context, req *GetTransferOrderByIDRequest) (*GetTransferOrderByIDReply, error)
	List(ctx context.Context, req *ListTransferOrderRequest) (*ListTransferOrderReply, error)
	PrecheckTransfer(ctx context.Context, req *PrecheckTransferRequest) (*PrecheckTransferReply, error)
}

type TransferOrderOption func(*transferOrderOptions)

type transferOrderOptions struct {
	isFromRPC  bool
	responser  errcode.Responser
	zapLog     *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *transferOrderOptions) apply(opts ...TransferOrderOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithTransferOrderHTTPResponse() TransferOrderOption {
	return func(o *transferOrderOptions) {
		o.isFromRPC = false
	}
}

func WithTransferOrderRPCResponse() TransferOrderOption {
	return func(o *transferOrderOptions) {
		o.isFromRPC = true
	}
}

func WithTransferOrderResponser(responser errcode.Responser) TransferOrderOption {
	return func(o *transferOrderOptions) {
		o.responser = responser
	}
}

func WithTransferOrderLogger(zapLog *zap.Logger) TransferOrderOption {
	return func(o *transferOrderOptions) {
		o.zapLog = zapLog
	}
}

func WithTransferOrderErrorToHTTPCode(e ...*errcode.Error) TransferOrderOption {
	return func(o *transferOrderOptions) {
		o.httpErrors = e
	}
}

func WithTransferOrderRPCStatusToHTTPCode(s ...*errcode.RPCStatus) TransferOrderOption {
	return func(o *transferOrderOptions) {
		o.rpcStatus = s
	}
}

func WithTransferOrderWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) TransferOrderOption {
	return func(o *transferOrderOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterTransferOrderRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic TransferOrderLogicer,
	opts ...TransferOrderOption) {

	o := &transferOrderOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog, _ = zap.NewProduction()
	}

	r := &transferOrderRouter{
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

type transferOrderRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                TransferOrderLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *transferOrderRouter) register() {
	r.iRouter.Handle("POST", "/api/v1/transferOrder", r.withMiddleware("POST", "/api/v1/transferOrder", r.Create_52)...)
	r.iRouter.Handle("DELETE", "/api/v1/transferOrder/:id", r.withMiddleware("DELETE", "/api/v1/transferOrder/:id", r.DeleteByID_48)...)
	r.iRouter.Handle("PUT", "/api/v1/transferOrder/:id", r.withMiddleware("PUT", "/api/v1/transferOrder/:id", r.UpdateByID_48)...)
	r.iRouter.Handle("GET", "/api/v1/transferOrder/:id", r.withMiddleware("GET", "/api/v1/transferOrder/:id", r.GetByID_48)...)
	r.iRouter.Handle("POST", "/api/v1/transferOrder/list", r.withMiddleware("POST", "/api/v1/transferOrder/list", r.List_52)...)
	r.iRouter.Handle("POST", "/api/v1/transferOrder/precheck", r.withMiddleware("POST", "/api/v1/transferOrder/precheck", r.PrecheckTransfer_0)...)

}

func (r *transferOrderRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
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

func (r *transferOrderRouter) Create_52(c *gin.Context) {
	req := &CreateTransferOrderRequest{}
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

func (r *transferOrderRouter) DeleteByID_48(c *gin.Context) {
	req := &DeleteTransferOrderByIDRequest{}
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

func (r *transferOrderRouter) UpdateByID_48(c *gin.Context) {
	req := &UpdateTransferOrderByIDRequest{}
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

func (r *transferOrderRouter) GetByID_48(c *gin.Context) {
	req := &GetTransferOrderByIDRequest{}
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

func (r *transferOrderRouter) List_52(c *gin.Context) {
	req := &ListTransferOrderRequest{}
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

func (r *transferOrderRouter) PrecheckTransfer_0(c *gin.Context) {
	req := &PrecheckTransferRequest{}
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

	out, err := r.iLogic.PrecheckTransfer(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}
