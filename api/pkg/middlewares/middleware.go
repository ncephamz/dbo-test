package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ncephamz/dbo-test/api/models"
	"gorm.io/gorm"
)

type Middleware struct {
	jwt Jwt
	DB  *gorm.DB
}

func NewMiddleware(
	jwt Jwt,
	DB *gorm.DB,
) Middleware {
	return Middleware{
		jwt: jwt,
		DB:  DB,
	}
}

func (m Middleware) Validate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		}

		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "unauthorized"})
			return
		}

		claims, err := m.jwt.Validate(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		var admin models.Admin
		result := m.DB.First(&admin, "id = ? AND token is not null", claims["id"])
		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "unauthorized"})
			return
		}

		ctx.Set("admin", admin)
		ctx.Next()
	}
}
