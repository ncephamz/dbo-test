package admin

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ncephamz/efishery-be-test/api/models"
	utils "github.com/ncephamz/efishery-be-test/api/pkg"
	"github.com/ncephamz/efishery-be-test/api/pkg/middlewares"
	"gorm.io/gorm"
)

type AdminController struct {
	DB  *gorm.DB
	Jwt middlewares.Jwt
}

func NewAdminController(
	DB *gorm.DB,
	jwt middlewares.Jwt,
) AdminController {
	return AdminController{
		DB:  DB,
		Jwt: jwt,
	}
}

func (ac *AdminController) Login(ctx *gin.Context) {
	var payload *models.AdminLogin

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var admin models.Admin
	result := ac.DB.First(&admin, "username = ?", strings.ToLower(payload.Username))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	if err := utils.VerifyPassword(admin.Password, payload.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	duration := time.Now().Add(time.Hour * 24).Unix()

	token, err := ac.Jwt.Signed(admin, duration)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": models.AdminLoginResponse{Token: token, ExpiredAt: time.Duration(duration)}})
}
