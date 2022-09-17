package main

import (
	"Assigment2/config"
	"Assigment2/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.StartDB()
	if db == nil {
		fmt.Println("db failed to run!")
	}
	r := gin.Default()
	orderController := controllers.NewOrderController(db)
	userRoute := r.Group("/orders")
	{
		userRoute.GET("/", orderController.GetOrders)
		userRoute.POST("/", orderController.CreateOrder)
		userRoute.PUT("/:orderId", orderController.UpdateOrders)
		userRoute.DELETE("/:orderId", orderController.DeleteOrders)
	}
	r.Run(":9000")
}
