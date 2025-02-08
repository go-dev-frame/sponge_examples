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

type PromotionLogicer interface {
	Create(ctx context.Context, req *CreatePromotionRequest) (*CreatePromotionReply, error)
	DeleteByID(ctx context.Context, req *DeletePromotionByIDRequest) (*DeletePromotionByIDReply, error)
	UpdateByID(ctx context.Context, req *UpdatePromotionByIDRequest) (*UpdatePromotionByIDReply, error)
	GetByID(ctx context.Context, req *GetPromotionByIDRequest) (*GetPromotionByIDReply, error)
	List(ctx context.Context, req *ListPromotionRequest) (*ListPromotionReply, error)
	BindCouponTemplate(ctx context.Context, req *BindCouponTemplateRequest) (*BindCouponTemplateReply, error)
}

type PromotionOption func(*promotionOptions)

type promotionOptions struct {
	isFromRPC  bool
	responser  errcode.Responser
	zapLog     *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *promotionOptions) apply(opts ...PromotionOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithPromotionHTTPResponse() PromotionOption {
	return func(o *promotionOptions) {
		o.isFromRPC = false
	}
}

func WithPromotionRPCResponse() PromotionOption {
	return func(o *promotionOptions) {
		o.isFromRPC = true
	}
}

func WithPromotionResponser(responser errcode.Responser) PromotionOption {
	return func(o *promotionOptions) {
		o.responser = responser
	}
}

func WithPromotionLogger(zapLog *zap.Logger) PromotionOption {
	return func(o *promotionOptions) {
		o.zapLog = zapLog
	}
}

func WithPromotionErrorToHTTPCode(e ...*errcode.Error) PromotionOption {
	return func(o *promotionOptions) {
		o.httpErrors = e
	}
}

func WithPromotionRPCStatusToHTTPCode(s ...*errcode.RPCStatus) PromotionOption {
	return func(o *promotionOptions) {
		o.rpcStatus = s
	}
}

func WithPromotionWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) PromotionOption {
	return func(o *promotionOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterPromotionRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic PromotionLogicer,
	opts ...PromotionOption) {

	o := &promotionOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog, _ = zap.NewProduction()
	}

	r := &promotionRouter{
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

type promotionRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                PromotionLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *promotionRouter) register() {
	r.iRouter.Handle("POST", "/api/v1/promotion", r.withMiddleware("POST", "/api/v1/promotion", r.Create_34)...)
	r.iRouter.Handle("DELETE", "/api/v1/promotion/:id", r.withMiddleware("DELETE", "/api/v1/promotion/:id", r.DeleteByID_32)...)
	r.iRouter.Handle("PUT", "/api/v1/promotion/:id", r.withMiddleware("PUT", "/api/v1/promotion/:id", r.UpdateByID_32)...)
	r.iRouter.Handle("GET", "/api/v1/promotion/:id", r.withMiddleware("GET", "/api/v1/promotion/:id", r.GetByID_32)...)
	r.iRouter.Handle("POST", "/api/v1/promotion/list", r.withMiddleware("POST", "/api/v1/promotion/list", r.List_34)...)
	r.iRouter.Handle("POST", "/api/v1/promotion/:promotionID/coupons", r.withMiddleware("POST", "/api/v1/promotion/:promotionID/coupons", r.BindCouponTemplate_0)...)

}

func (r *promotionRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
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

func (r *promotionRouter) Create_34(c *gin.Context) {
	req := &CreatePromotionRequest{}
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

func (r *promotionRouter) DeleteByID_32(c *gin.Context) {
	req := &DeletePromotionByIDRequest{}
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

func (r *promotionRouter) UpdateByID_32(c *gin.Context) {
	req := &UpdatePromotionByIDRequest{}
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

func (r *promotionRouter) GetByID_32(c *gin.Context) {
	req := &GetPromotionByIDRequest{}
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

func (r *promotionRouter) List_34(c *gin.Context) {
	req := &ListPromotionRequest{}
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

func (r *promotionRouter) BindCouponTemplate_0(c *gin.Context) {
	req := &BindCouponTemplateRequest{}
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

	out, err := r.iLogic.BindCouponTemplate(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}
