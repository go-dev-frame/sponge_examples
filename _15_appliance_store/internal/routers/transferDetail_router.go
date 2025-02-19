// Code generated by https://github.com/go-dev-frame/sponge

package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/go-dev-frame/sponge/pkg/logger"
	//"github.com/go-dev-frame/sponge/pkg/middleware"

	storeV1 "store/api/store/v1"
	"store/internal/handler"
)

func init() {
	allMiddlewareFns = append(allMiddlewareFns, func(c *middlewareConfig) {
		transferDetailMiddlewares(c)
	})

	allRouteFns = append(allRouteFns,
		func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc) {
			transferDetailRouter(r, groupPathMiddlewares, singlePathMiddlewares, handler.NewTransferDetailHandler())
		})
}

func transferDetailRouter(
	r *gin.Engine,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iService storeV1.TransferDetailLogicer) {
	storeV1.RegisterTransferDetailRouter(
		r,
		groupPathMiddlewares,
		singlePathMiddlewares,
		iService,
		storeV1.WithTransferDetailLogger(logger.Get()),
		storeV1.WithTransferDetailHTTPResponse(),
		storeV1.WithTransferDetailErrorToHTTPCode(
		// Set some error codes to standard http return codes,
		// by default there is already ecode.InternalServerError and ecode.ServiceUnavailable
		// example:
		// 	ecode.Forbidden, ecode.LimitExceed,
		),
	)
}

// you can set the middleware of a route group, or set the middleware of a single route,
// or you can mix them, pay attention to the duplication of middleware when mixing them,
// it is recommended to set the middleware of a single route in preference
func transferDetailMiddlewares(c *middlewareConfig) {
	// set up group route middleware, group path is left prefix rules,
	// if the left prefix is hit, the middleware will take effect, e.g. group route is /api/v1, route /api/v1/transferDetail/:id  will take effect
	// c.setGroupPath("/api/v1/transferDetail", middleware.Auth())

	// set up single route middleware, just uncomment the code and fill in the middlewares, nothing else needs to be changed
	//c.setSinglePath("POST", "/api/v1/transferDetail", middleware.Auth())    // Create transferDetail
	//c.setSinglePath("DELETE", "/api/v1/transferDetail/:transferID", middleware.Auth())    // DeleteByTransferID delete transferDetail by transferID
	//c.setSinglePath("PUT", "/api/v1/transferDetail/:transferID", middleware.Auth())    // UpdateByTransferID update transferDetail by transferID
	//c.setSinglePath("GET", "/api/v1/transferDetail/:transferID", middleware.Auth())    // GetByTransferID get transferDetail by transferID
	//c.setSinglePath("POST", "/api/v1/transferDetail/list", middleware.Auth())    // List of transferDetail by query parameters
}
