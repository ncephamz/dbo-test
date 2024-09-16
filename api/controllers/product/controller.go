package product

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ncephamz/dbo-test/api/models"
	utils "github.com/ncephamz/dbo-test/api/pkg"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func NewController(
	DB *gorm.DB,
) Controller {
	return Controller{
		DB: DB,
	}
}

func (c *Controller) GetProducts(ctx *gin.Context) {
	var (
		page     = utils.StringToInt(ctx.DefaultQuery("page", "1"))
		limit    = utils.StringToInt(ctx.DefaultQuery("limit", "10"))
		keyword  = ctx.DefaultQuery("keyword", "")
		offset   = (page - 1) * limit
		products []models.ProductsAssosiationsToStoreWarehouse
		response []models.ResponseGetProducts
		count    int64
		db       = c.DB
	)

	if keyword != "" {
		keyword = fmt.Sprintf("%s%s%s", "%", keyword, "%")
		db = db.Where("product_name ILIKE ?", keyword)
	}

	result := db.Limit(limit).Offset(offset).Order("product_name").Preload("StoreWarehouse").Find(&products)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	result = db.Model(&models.ProductsAssosiationsToStoreWarehouse{}).Count(&count)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error})
		return
	}

	for _, row := range products {
		response = append(response, row.ToResponse())
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": response, "count": utils.IntToString(uint64(count))})
}
