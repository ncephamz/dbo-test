package customer

import (
	"fmt"
	"net/http"
	"strings"

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

func (c *Controller) GetCustomers(ctx *gin.Context) {
	var (
		page      = utils.StringToInt(ctx.DefaultQuery("page", "1"))
		limit     = utils.StringToInt(ctx.DefaultQuery("limit", "10"))
		keyword   = ctx.DefaultQuery("keyword", "")
		offset    = (page - 1) * limit
		customers []models.Customer
		result    []models.ResponseGetAllCustomer
		count     int64
	)

	if keyword != "" {
		c.DB.Where("name ILIKE ?", fmt.Sprintf("%s%s%s", "%", keyword, "%"))
	}

	results := c.DB.Limit(limit).Offset(offset).Order("name").Find(&customers)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	results = c.DB.Model(&models.Customer{}).Count(&count)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	for _, row := range customers {
		result = append(result, row.ToResponse())
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": result, "count": utils.IntToString(uint64(count))})
}

func (c *Controller) CreateCustomers(ctx *gin.Context) {
	var payload *models.RequestCreateCustomer
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	customer := payload.ToModel()
	customerAddress := payload.Address.ToModel(customer.Id)
	tx := c.DB.Begin()

	result := tx.Create(&customer)
	if result.Error != nil {
		tx.Rollback()

		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Customer with that phone number already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	result = tx.Create(&customerAddress)
	if result.Error != nil {
		tx.Rollback()

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusCreated, gin.H{"status": "success"})
}
