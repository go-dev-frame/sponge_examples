// Code generated by https://github.com/zhufuyi/sponge, DO NOT EDIT.

package v1

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	errcode "github.com/zhufuyi/sponge/pkg/errcode"
	middleware "github.com/zhufuyi/sponge/pkg/gin/middleware"
	zap "go.uber.org/zap"
	strings "strings"
)

// import packages: strings. context. errcode. middleware. zap. gin.

type RelationServiceLogicer interface {
	BatchGetRelation(ctx context.Context, req *BatchGetRelationRequest) (*BatchGetRelationReply, error)
	Follow(ctx context.Context, req *FollowRequest) (*FollowReply, error)
	ListFollower(ctx context.Context, req *ListFollowerRequest) (*ListFollowerReply, error)
	ListFollowing(ctx context.Context, req *ListFollowingRequest) (*ListFollowingReply, error)
	Unfollow(ctx context.Context, req *UnfollowRequest) (*UnfollowReply, error)
}

type RelationServiceOption func(*relationServiceOptions)

type relationServiceOptions struct {
	isFromRPC  bool
	responser  errcode.Responser
	zapLog     *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *relationServiceOptions) apply(opts ...RelationServiceOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithRelationServiceHTTPResponse() RelationServiceOption {
	return func(o *relationServiceOptions) {
		o.isFromRPC = false
	}
}

func WithRelationServiceRPCResponse() RelationServiceOption {
	return func(o *relationServiceOptions) {
		o.isFromRPC = true
	}
}

func WithRelationServiceResponser(responser errcode.Responser) RelationServiceOption {
	return func(o *relationServiceOptions) {
		o.responser = responser
	}
}

func WithRelationServiceLogger(zapLog *zap.Logger) RelationServiceOption {
	return func(o *relationServiceOptions) {
		o.zapLog = zapLog
	}
}

func WithRelationServiceErrorToHTTPCode(e ...*errcode.Error) RelationServiceOption {
	return func(o *relationServiceOptions) {
		o.httpErrors = e
	}
}

func WithRelationServiceRPCStatusToHTTPCode(s ...*errcode.RPCStatus) RelationServiceOption {
	return func(o *relationServiceOptions) {
		o.rpcStatus = s
	}
}

func WithRelationServiceWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) RelationServiceOption {
	return func(o *relationServiceOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterRelationServiceRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic RelationServiceLogicer,
	opts ...RelationServiceOption) {

	o := &relationServiceOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog, _ = zap.NewProduction()
	}

	r := &relationServiceRouter{
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

type relationServiceRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                RelationServiceLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *relationServiceRouter) register() {
	r.iRouter.Handle("POST", "/api/v1/relation/follow", r.withMiddleware("POST", "/api/v1/relation/follow", r.Follow_0)...)
	r.iRouter.Handle("POST", "/api/v1/relation/unfollow", r.withMiddleware("POST", "/api/v1/relation/unfollow", r.Unfollow_0)...)
	r.iRouter.Handle("GET", "/api/v1/relation/following/list", r.withMiddleware("GET", "/api/v1/relation/following/list", r.ListFollowing_0)...)
	r.iRouter.Handle("GET", "/api/v1/relation/follower/list", r.withMiddleware("GET", "/api/v1/relation/follower/list", r.ListFollower_0)...)
	r.iRouter.Handle("POST", "/api/v1/relation/check/list", r.withMiddleware("POST", "/api/v1/relation/check/list", r.BatchGetRelation_0)...)

}

func (r *relationServiceRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
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

func (r *relationServiceRouter) Follow_0(c *gin.Context) {
	req := &FollowRequest{}
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

	out, err := r.iLogic.Follow(ctx, req)
	if err != nil {
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *relationServiceRouter) Unfollow_0(c *gin.Context) {
	req := &UnfollowRequest{}
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

	out, err := r.iLogic.Unfollow(ctx, req)
	if err != nil {
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *relationServiceRouter) ListFollowing_0(c *gin.Context) {
	req := &ListFollowingRequest{}
	var err error

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

	out, err := r.iLogic.ListFollowing(ctx, req)
	if err != nil {
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *relationServiceRouter) ListFollower_0(c *gin.Context) {
	req := &ListFollowerRequest{}
	var err error

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

	out, err := r.iLogic.ListFollower(ctx, req)
	if err != nil {
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *relationServiceRouter) BatchGetRelation_0(c *gin.Context) {
	req := &BatchGetRelationRequest{}
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

	out, err := r.iLogic.BatchGetRelation(ctx, req)
	if err != nil {
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}
