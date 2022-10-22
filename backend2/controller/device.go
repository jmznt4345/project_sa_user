package controller

import (
	"net/http"

	"github.com/jmznt4345/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /devices
func CreateDevice(c *gin.Context) {
	var device entity.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": device})
}

// GET /device/:id
func GetDevice(c *gin.Context) {
	var device entity.Device
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM devices WHERE id = ?", id).Scan(&device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": device})
}

// GET /devices
func ListDevices(c *gin.Context) {
	var devices []entity.Device
	if err := entity.DB().Raw("SELECT * FROM devices").Scan(&devices).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": devices})
}

// DELETE /devices/:id
func DeleteDevice(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM devices WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /devices
func UpdateDevice(c *gin.Context) {
	var device entity.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", device.ID).First(&device); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device not found"})
		return
	}
	if err := entity.DB().Save(&device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": device})
}
