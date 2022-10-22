package controller

import (
	"net/http"

	"github.com/jmznt4345/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /requests
func CreateRequest(c *gin.Context) {

	var Request entity.Request
	var Room_has_Device entity.Room_has_Device
	var JobType entity.JobType
	var User entity.User

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร watchVideo
	if err := c.ShouldBindJSON(&Request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา Room_has_Device ด้วย id
	if tx := entity.DB().Where("id = ?", Request.Room_has_Device_ID).First(&Room_has_Device); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room_has_Device in room not found"})
		return
	}

	// 10: ค้นหา JobType ด้วย id
	if tx := entity.DB().Where("id = ?", Request.JobTypeID).First(&JobType); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JobType not found"})
		return
	}

	// 11: ค้นหา User ด้วย id
	if tx := entity.DB().Where("id = ?", Request.UserID).First(&User); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	// 12: สร้าง WatchVideo
	rq := entity.Request{
		Explain:    Request.Explain,
		Date_Start: Request.Date_Start,

		User:            User,
		Room_has_Device: Room_has_Device,
		JobType:         JobType,
	}

	// 13: บันทึก
	if err := entity.DB().Create(&rq).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": rq})
}

// GET /request/:id
func GetRequest(c *gin.Context) {
	var request entity.Request
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM requests WHERE id = ?", id).Scan(&request).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": request})
}

// GET /requests
func ListRequests(c *gin.Context) {
	var requests []entity.Request
	if err := entity.DB().Preload("Room_has_Device").Preload("Room_has_Device.Device").Preload("Room_has_Device.Room").Preload("Room_has_Device.Room.Building").Raw("SELECT * FROM requests").Find(&requests).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": requests})
}

// DELETE /requests/:id
func DeleteRequest(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM requests WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /requests
func UpdateRequest(c *gin.Context) {
	var request entity.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", request.ID).First(&request); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request not found"})
		return
	}
	if err := entity.DB().Save(&request).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": request})
}
