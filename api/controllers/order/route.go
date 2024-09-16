package order

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
	router := rg.Group("orders")

	router.POST("/cart", r.middleware.Validate(), r.controller.AddToCart)
	router.GET("", r.middleware.Validate(), r.controller.GetOrders)
}
