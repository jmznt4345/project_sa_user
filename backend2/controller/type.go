package controller

import (
	"net/http"

	"github.com/jmznt4345/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /typeXs
func CreateType(c *gin.Context) {
	var typeX entity.Type
	if err := c.ShouldBindJSON(&typeX); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&typeX).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": typeX})
}

// GET /Type
func GetType(c *gin.Context) {
	var typeX entity.Type
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM typeXs WHERE id = ?", id).Scan(&typeX).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": typeX})
}

// GET /Types
func ListTypes(c *gin.Context) {
	var typeXs []entity.Type
	if err := entity.DB().Raw("SELECT * FROM typeXs").Scan(&typeXs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": typeXs})
}

// DELETE /typeXs/:id
func DeleteType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM typeXs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "typeX not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /typeXs
func UpdateType(c *gin.Context) {
	var typeX entity.Type
	if err := c.ShouldBindJSON(&typeX); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", typeX.ID).First(&typeX); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "typeX not found"})
		return
	}
	if err := entity.DB().Save(&typeX).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": typeX})
}
