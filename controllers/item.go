package controllers

import (
	"api/database"
	"api/model"
	"github.com/gin-gonic/gin"
)


func CreateItem(c *gin.Context) {
	
	var item model.Item
	if err := c.BindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	database.DB.Create(&item)
	c.JSON(201, item)	
	
}

func GetItemByID(c *gin.Context) {
	
	id := c.Param("id")
	var item model.Item
	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(200, item)
	
}

func GetAllItems(c *gin.Context) {
	var items []model.Item
	database.DB.Find(&items)
	c.JSON(200, items)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")

	var item model.Item

	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not Found"})
		return
	}

	if err := c.BindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	database.DB.Save(&item)
	c.JSON(200, item)
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")

	var item model.Item

	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Item not Found"})
		return
	}
	database.DB.Delete(&item, id)
	
	c.JSON(200, gin.H{"message": "Item deleted successfully"})
}