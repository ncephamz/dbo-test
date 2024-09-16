package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/ncephamz/dbo-test/api/pkg/middlewares"
)

type RouteController struct {
	controller Controller
	middleware middlewares.Middleware
}

func NewRouteController(
	controller Controller,
	middleware middlewares.Middleware,
) RouteController {
	return RouteController{
		controller: controller,
		middleware: middleware,
	}
}

func (r *RouteController) Route(rg *gin.RouterGroup) {
	router := rg.Group("customers")

	router.GET("", r.middleware.Validate(), r.controller.GetCustomers)
	router.POST("", r.middleware.Validate(), r.controller.CreateCustomers)
	router.PUT("/:id", r.middleware.Validate(), r.controller.UpdateCustomers)
	router.GET("/:id", r.middleware.Validate(), r.controller.GetDetailCustomers)
	router.DELETE("/:id", r.middleware.Validate(), r.controller.DeleteCustomers)
}
