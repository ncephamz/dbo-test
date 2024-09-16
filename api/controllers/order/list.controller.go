package order

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ncephamz/dbo-test/api/models"
	utils "github.com/ncephamz/dbo-test/api/pkg"
)

func (c *Controller) GetOrders(ctx *gin.Context) {
	var (
		page     = utils.StringToInt(ctx.DefaultQuery("page", "1"))
		limit    = utils.StringToInt(ctx.DefaultQuery("limit", "10"))
		keyword  = ctx.DefaultQuery("keyword", "")
		offset   = (page - 1) * limit
		orders   []models.OrderAssosiationToCustomer
		response []models.ResponseGetOrders
		count    int64
		db       = c.DB
	)

	if keyword != "" {
		keyword = fmt.Sprintf("%s%s%s", "%", keyword, "%")
		db = db.Where("name ILIKE ?", keyword)
	}

	result := db.Limit(limit).Offset(offset).Order("created_at").Preload("Customer").Find(&orders)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	result = db.Model(&models.OrderAssosiationToCustomer{}).Count(&count)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error})
		return
	}

	for _, row := range orders {
		response = append(response, row.ToResponse())
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": response, "count": utils.IntToString(uint64(count))})
}

func (c *Controller) GetDetailOrder(ctx *gin.Context) {
	var (
		id    = ctx.Param("id")
		order models.OrderDetailAssosiation
	)

	result := c.DB.Preload("Customer").Preload("CustomerAddress").Preload("Details").Find(&order, "id = ?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": order.ToResponse()})
}
