package admin

import (
	"github.com/gin-gonic/gin"
)

type AdminRouteController struct {
	controller AdminController
}

func NewAdminRouteController(controller AdminController) AdminRouteController {
	return AdminRouteController{controller}
}

func (rc *AdminRouteController) AdminRoute(rg *gin.RouterGroup) {
	router := rg.Group("admins")

	router.POST("/login", rc.controller.Login)
}
