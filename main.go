package main

import (
	"api/database"
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDB()

	r := gin.Default()

	r.POST("/items", controllers.CreateItem)
	r.GET("/items/:id", controllers.GetItemByID)
	r.GET("/items", controllers.GetAllItems)
	r.PUT("/items/:id", controllers.UpdateItem)
	r.DELETE("/items/:id", controllers.DeleteItem)

	r.Run(":8080")
}