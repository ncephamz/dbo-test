package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ncephamz/dbo-test/api/models"
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

func (c *Controller) AddToCart(ctx *gin.Context) {
	var payload *models.RequestAddToCart
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var tx = c.DB.Begin()

	var customerAddress models.CustomerAddress
	result := tx.Find(&customerAddress, "id = ?", payload.CustomerAddressId)
	if result.Error != nil {
		tx.Rollback()

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	var storeWarehouseWithCoverage models.StoreWarehouseAssosiationToWareHouseCoverage
	result = tx.Preload("Coverage").Find(&storeWarehouseWithCoverage,
		"product_id = ?",
		payload.ProductId)
	if result.Error != nil {
		tx.Rollback()

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	if storeWarehouseWithCoverage.StoresWarehouses.Qty < payload.Qty {
		tx.Rollback()

		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "stock not avaible"})
		return
	}

	var order models.Orders
	result = tx.Find(&order, "customer_address_id = ? AND customer_id = ? AND status = ?", payload.CustomerAddressId, payload.CustomerId, models.CART)
	if result.Error != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	if order.Id == 0 {
		order = payload.ToModelOrder()
		result = tx.Create(&order)
		if result.Error != nil {
			tx.Rollback()
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
			return
		}
	}

	var orderDetails models.OrdersDetails
	result = tx.Find(&orderDetails, "order_id = ? AND store_warehouse_id = ?", order.Id, storeWarehouseWithCoverage.StoresWarehouses.Id)
	if result.Error != nil {
		tx.Rollback()

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	if orderDetails.Id != 0 {
		orderDetails.Qty += payload.Qty
		result = tx.Updates(&orderDetails)
		if result.Error != nil {
			tx.Rollback()

			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
			return
		}
	} else {
		orderDetails = models.ToModelOrderDetail(order.Id, storeWarehouseWithCoverage.StoresWarehouses.Id, payload.Qty)
		result = tx.Create(&orderDetails)
		if result.Error != nil {
			tx.Rollback()

			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
			return
		}
	}

	tx.Commit()

	ctx.JSON(http.StatusCreated, gin.H{"status": "success"})
}
