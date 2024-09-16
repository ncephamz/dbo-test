package api

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ncephamz/dbo-test/api/config"
	"github.com/ncephamz/dbo-test/api/controllers/admin"
	"github.com/ncephamz/dbo-test/api/controllers/customer"
	"github.com/ncephamz/dbo-test/api/controllers/order"
	"github.com/ncephamz/dbo-test/api/controllers/product"
	"github.com/ncephamz/dbo-test/api/pkg/database"
	"github.com/ncephamz/dbo-test/api/pkg/middlewares"
)

var (
	server               *gin.Engine
	AdminController      admin.Controller
	AdminRouteController admin.RouteController

	CustomerContoller      customer.Controller
	CustomerRouteContoller customer.RouteController

	ProductController      product.Controller
	ProductRouteController product.RouteController

	OrderConstroller     order.Controller
	OrderRouteController order.RouteController
)

func init() {
	config := config.LoadConfig()

	DB, err := database.ConnectDB(config)
	if err != nil {
		log.Fatal(err)
	}

	jwt := middlewares.Jwt{Secret: config.JwtSecret}
	middleware := middlewares.NewMiddleware(jwt, DB)

	AdminController = admin.NewController(DB, jwt)
	AdminRouteController = admin.NewRouteController(AdminController, middleware)

	CustomerContoller = customer.NewController(DB)
	CustomerRouteContoller = customer.NewRouteController(CustomerContoller, middleware)

	ProductController = product.NewController(DB)
	ProductRouteController = product.NewRouteController(ProductController, middleware)

	OrderConstroller = order.NewController(DB)
	OrderRouteController = order.NewRouteController(OrderConstroller, middleware)

	server = gin.Default()
}

func Run() {
	config := config.LoadConfig()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.AllowCors
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	v1 := router.Group("/v1")
	AdminRouteController.AdminRoute(v1)
	CustomerRouteContoller.Route(v1)
	ProductRouteController.Route(v1)
	OrderRouteController.Route(v1)

	log.Fatal(server.Run(":" + config.Port))
}
