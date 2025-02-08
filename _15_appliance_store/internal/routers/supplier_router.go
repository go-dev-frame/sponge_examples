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
		supplierMiddlewares(c)
	})

	allRouteFns = append(allRouteFns,
		func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc) {
			supplierRouter(r, groupPathMiddlewares, singlePathMiddlewares, handler.NewSupplierHandler())
		})
}

func supplierRouter(
	r *gin.Engine,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iService storeV1.SupplierLogicer) {
	storeV1.RegisterSupplierRouter(
		r,
		groupPathMiddlewares,
		singlePathMiddlewares,
		iService,
		storeV1.WithSupplierLogger(logger.Get()),
		storeV1.WithSupplierHTTPResponse(),
		storeV1.WithSupplierErrorToHTTPCode(
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
func supplierMiddlewares(c *middlewareConfig) {
	// set up group route middleware, group path is left prefix rules,
	// if the left prefix is hit, the middleware will take effect, e.g. group route is /api/v1, route /api/v1/supplier/:id  will take effect
	// c.setGroupPath("/api/v1/supplier", middleware.Auth())

	// set up single route middleware, just uncomment the code and fill in the middlewares, nothing else needs to be changed
	//c.setSinglePath("POST", "/api/v1/supplier", middleware.Auth())    // Create supplier
	//c.setSinglePath("DELETE", "/api/v1/supplier/:id", middleware.Auth())    // DeleteByID delete supplier by id
	//c.setSinglePath("PUT", "/api/v1/supplier/:id", middleware.Auth())    // UpdateByID update supplier by id
	//c.setSinglePath("GET", "/api/v1/supplier/:id", middleware.Auth())    // GetByID get supplier by id
	//c.setSinglePath("POST", "/api/v1/supplier/list", middleware.Auth())    // List of supplier by query parameters
}
