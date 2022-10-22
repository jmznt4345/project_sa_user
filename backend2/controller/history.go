package controller

import (
	"net/http"

	"github.com/jmznt4345/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /historys
func CreateHistory(c *gin.Context) {
	var history entity.History
	if err := c.ShouldBindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&history).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": history})
}

// GET /history/:id
func GetHistory(c *gin.Context) {
	var history entity.History
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM historys WHERE id = ?", id).Scan(&history).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": history})
}

// GET /historys
func ListHistorys(c *gin.Context) {
	var historys []entity.History
	if err := entity.DB().Raw("SELECT * FROM historys").Scan(&historys).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": historys})
}

// DELETE /historys/:id
func DeleteHistory(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM historys WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "history not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /historys
func UpdateHistory(c *gin.Context) {
	var history entity.History
	if err := c.ShouldBindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", history.ID).First(&history); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "history not found"})
		return
	}
	if err := entity.DB().Save(&history).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": history})
}
