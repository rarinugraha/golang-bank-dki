package handlers

import (
	"go-bank-dki/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ListStock(c *gin.Context) {
	var stocks []models.Stock
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&stocks)
	c.JSON(http.StatusOK, stocks)
}

func CreateStock(c *gin.Context) {
	var input models.Stock

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if userIDFloat, ok := userID.(float64); ok {
		input.CreatedBy = uint(userIDFloat)
		input.UpdatedBy = uint(userIDFloat)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, input)
}

func DetailStock(c *gin.Context) {
	var stock models.Stock
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&stock).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	c.JSON(http.StatusOK, stock)
}

func UpdateStock(c *gin.Context) {
	var stock models.Stock
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&stock).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if userIDFloat, ok := userID.(float64); ok {
		stock.UpdatedBy = uint(userIDFloat)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	stock.UpdatedAt = time.Now()

	if err := db.Save(&stock).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stock)
}

func DeleteStock(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var stock models.Stock

	if err := db.Where("id = ?", c.Param("id")).First(&stock).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	if err := db.Delete(&stock).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock deleted"})
}
