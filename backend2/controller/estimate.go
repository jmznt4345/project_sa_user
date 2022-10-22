package controller

import (
	"net/http"

	"github.com/jmznt4345/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /estimates
func CreateEstimate(c *gin.Context) {
	var estimate entity.Estimate
	if err := c.ShouldBindJSON(&estimate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&estimate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": estimate})
}

// GET /Estimate
func GetEstimate(c *gin.Context) {
	var estimate entity.Estimate
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM estimates WHERE id = ?", id).Scan(&estimate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": estimate})
}

// GET /Estimates
func ListEstimates(c *gin.Context) {
	var estimates []entity.Estimate
	if err := entity.DB().Raw("SELECT * FROM estimates").Scan(&estimates).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": estimates})
}

// DELETE /estimates/:id
func DeleteEstimate(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM estimates WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "estimate not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /estimates
func UpdateEstimate(c *gin.Context) {
	var estimate entity.Estimate
	if err := c.ShouldBindJSON(&estimate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", estimate.ID).First(&estimate); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "estimate not found"})
		return
	}
	if err := entity.DB().Save(&estimate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": estimate})
}
