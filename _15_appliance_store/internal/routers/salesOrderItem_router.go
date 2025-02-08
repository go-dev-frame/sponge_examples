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
		salesOrderItemMiddlewares(c)
	})

	allRouteFns = append(allRouteFns,
		func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc) {
			salesOrderItemRouter(r, groupPathMiddlewares, singlePathMiddlewares, handler.NewSalesOrderItemHandler())
		})
}

func salesOrderItemRouter(
	r *gin.Engine,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iService storeV1.SalesOrderItemLogicer) {
	storeV1.RegisterSalesOrderItemRouter(
		r,
		groupPathMiddlewares,
		singlePathMiddlewares,
		iService,
		storeV1.WithSalesOrderItemLogger(logger.Get()),
		storeV1.WithSalesOrderItemHTTPResponse(),
		storeV1.WithSalesOrderItemErrorToHTTPCode(
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
func salesOrderItemMiddlewares(c *middlewareConfig) {
	// set up group route middleware, group path is left prefix rules,
	// if the left prefix is hit, the middleware will take effect, e.g. group route is /api/v1, route /api/v1/salesOrderItem/:id  will take effect
	// c.setGroupPath("/api/v1/salesOrderItem", middleware.Auth())

	// set up single route middleware, just uncomment the code and fill in the middlewares, nothing else needs to be changed
	//c.setSinglePath("POST", "/api/v1/salesOrderItem", middleware.Auth())    // Create salesOrderItem
	//c.setSinglePath("DELETE", "/api/v1/salesOrderItem/:id", middleware.Auth())    // DeleteByID delete salesOrderItem by id
	//c.setSinglePath("PUT", "/api/v1/salesOrderItem/:id", middleware.Auth())    // UpdateByID update salesOrderItem by id
	//c.setSinglePath("GET", "/api/v1/salesOrderItem/:id", middleware.Auth())    // GetByID get salesOrderItem by id
	//c.setSinglePath("POST", "/api/v1/salesOrderItem/list", middleware.Auth())    // List of salesOrderItem by query parameters
}
