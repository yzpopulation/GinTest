package controllers

import (
	"GinTest/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type CreateProductInput struct {
	Code  string `json:"code" binding:"required"`
	Price uint `json:"price" binding:"required"`
}

type UpdateProductInput struct {
	Code  string `json:"code"`
	Price uint `json:"price"`
}

func FindProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}
func FindProduct(c *gin.Context) {
	var product models.Product
	fmt.Println(c.Params)
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}
func CreateProduct(c *gin.Context) {
	// Validate input
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product := models.Product{Code: input.Code, Price: input.Price}
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}
func UpdateProduct(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": product})
}
func DeleteProduct(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}