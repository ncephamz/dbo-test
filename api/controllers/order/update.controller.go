package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ncephamz/dbo-test/api/models"
	utils "github.com/ncephamz/dbo-test/api/pkg"
)

func (c *Controller) UpdateStatus(ctx *gin.Context) {
	var (
		id     = utils.StringToUint64(ctx.Param("id"))
		status = ctx.Param("status")
	)

	result := c.DB.Updates(&models.Orders{Id: id, Status: status})
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (c *Controller) Delete(ctx *gin.Context) {
	var (
		id = utils.StringToUint64(ctx.Param("id"))
	)

	result := c.DB.Delete(&models.Orders{Id: id})
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
