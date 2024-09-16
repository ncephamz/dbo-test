package admin

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

func (rc *RouteController) AdminRoute(rg *gin.RouterGroup) {
	router := rg.Group("admins")

	router.POST("/login", rc.controller.Login)
	router.GET("/profile", rc.middleware.Validate(), rc.controller.GetProfile)
	router.POST("/logout", rc.middleware.Validate(), rc.controller.Logout)
}
