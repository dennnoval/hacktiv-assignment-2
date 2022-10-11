package routes

import (
	conf "hacktiv-assignment-2/config"
	g "github.com/gin-gonic/gin"
	c "hacktiv-assignment-2/controller"
)

func init() {
	r := g.Default()
	r.Use(func(ctx *g.Context) {
		ctx.Set("db", conf.DB)
	})
	r.GET("/orders", c.GetAllOrders)
	r.POST("/orders", c.CreateOrder)
	r.PUT("/orders/:orderId", c.UpdateOrder)
	r.DELETE("/orders/:orderId", c.DeleteOrder)
	r.Run()
}
