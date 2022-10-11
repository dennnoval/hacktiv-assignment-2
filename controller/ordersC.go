package controller

import (
	e "hacktiv-assignment-2/entity"
	g "github.com/gin-gonic/gin"
	h "net/http"
	"gorm.io/gorm"
	s "strconv"
)

func CreateOrder(ctx *g.Context) {
	var o e.Orders
	ctx.BindJSON(&o)
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Create(&o).Error
	if err != nil {
		ctx.JSON(400, g.H{"message": "Failed to add new order!"})
	} else {
		ctx.JSON(201, o)
	}
}

func GetAllOrders(ctx *g.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	o := []e.Orders{}
	db.Preload("Items").Find(&o)
	ctx.JSON(200, g.H{"orders": &o})
}

func UpdateOrder(ctx *g.Context) {
	orderId, _ := s.ParseUint(ctx.Param("orderId"), 10, 32)
	db := ctx.MustGet("db").(*gorm.DB)
	o := e.Orders{}
	err := db.First(&o, "order_id=?", &orderId).Error
	if err != nil {
		ctx.JSON(h.StatusNotFound, g.H{"message": "Failed to update, order doesn't exists!"})
		return
	}
	ctx.BindJSON(&o)
	db.Where("order_id=?", orderId).Delete(&o.Items)
	it := []e.Items{}
	for i := 0; i < len(*o.Items); i++ {
		it = append(it, e.Items{ItemCode: (*o.Items)[i].ItemCode, Description: (*o.Items)[i].Description, 
			Quantity: (*o.Items)[i].Quantity})
	}
	err = db.Model(&o).Where("order_id=?", &orderId).Updates(
		&e.Orders{CustomerName: o.CustomerName, Items: &it},
	).Error
	if err != nil {
		ctx.JSON(h.StatusOK, g.H{"message": "Failed to update product!"})
	} else {
		ctx.JSON(h.StatusOK, o)
	}
}

func DeleteOrder(ctx *g.Context) {
	orderId, _ := s.ParseUint(ctx.Param("orderId"), 10, 32)
	db := ctx.MustGet("db").(*gorm.DB)
	o := e.Orders{}
	err := db.First(&o, "order_id=?", &orderId).Error
	if err != nil {
		ctx.JSON(h.StatusNotFound, g.H{"message": "Failed to delete, order doesn't exists!"})
		return
	}
	err = db.Where("order_id=?", orderId).Delete(&o.Items).Error
	err = db.Where("order_id=?", orderId).Delete(&o).Error
	if err != nil {
		ctx.JSON(h.StatusOK, g.H{"message": "Failed to delete order!"})
	} else {
		ctx.JSON(h.StatusOK, g.H{"message": "Order deleted successfully!"})
	}
}
