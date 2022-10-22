package controller

import (
	"net/http"

	"github.com/jmznt4345/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /room_has_devices
func CreateRoom_has_Device(c *gin.Context) {
	var room_has_device entity.Room_has_Device
	if err := c.ShouldBindJSON(&room_has_device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&room_has_device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room_has_device})
}

// GET /room_has_device/:id
func GetRoom_has_Device(c *gin.Context) {
	var room_has_device entity.Room_has_Device
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM room_has_devices WHERE id = ?", id).Scan(&room_has_device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room_has_device})
}

// GET /room_has_device/room/:id
func GetRHD_Device(c *gin.Context) {
	var room_has_device []entity.Room_has_Device
	room_id := c.Param("id")
	if err := entity.DB().Preload("Device").Raw("SELECT * FROM room_has_devices WHERE room_id = ?", room_id).Find(&room_has_device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room_has_device})
}

// GET /room_has_devices
func ListRoom_has_Devices(c *gin.Context) {
	var room_has_devices []entity.Room_has_Device
	if err := entity.DB().Raw("SELECT * FROM room_has_devices").Scan(&room_has_devices).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room_has_devices})
}

// DELETE /room_has_devices/:id
func DeleteRoom_has_Device(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM room_has_devices WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_has_device not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /room_has_devices
func UpdateRoom_has_Device(c *gin.Context) {
	var room_has_device entity.Room_has_Device
	if err := c.ShouldBindJSON(&room_has_device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", room_has_device.ID).First(&room_has_device); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_has_device not found"})
		return
	}
	if err := entity.DB().Save(&room_has_device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room_has_device})
}
