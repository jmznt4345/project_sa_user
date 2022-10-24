package controller

import (
	"net/http"

	"github.com/jmznt4345/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /Educational_backgrounds
func CreateEducational_background(c *gin.Context) {
	var Educational_background entity.Educational_background
	if err := c.ShouldBindJSON(&Educational_background); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&Educational_background).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Educational_background})
}

// GET /Educational_background
func GetEducational_background(c *gin.Context) {
	var Educational_background entity.Educational_background
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Educational_backgrounds WHERE id = ?", id).Scan(&Educational_background).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Educational_background})
}

// GET /Educational_backgrounds
func ListEducational_backgrounds(c *gin.Context) {
	var Educational_backgrounds []entity.Educational_background
	if err := entity.DB().Raw("SELECT * FROM Educational_backgrounds").Scan(&Educational_backgrounds).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Educational_backgrounds})
}

// DELETE /Educational_backgrounds/:id
func DeleteEducational_background(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Educational_backgrounds WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Educational_background not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Educational_backgrounds
func UpdateEducational_background(c *gin.Context) {
	var Educational_background entity.Educational_background
	if err := c.ShouldBindJSON(&Educational_background); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", Educational_background.ID).First(&Educational_background); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Educational_background not found"})
		return
	}
	if err := entity.DB().Save(&Educational_background).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Educational_background})
}
