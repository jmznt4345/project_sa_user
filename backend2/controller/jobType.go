package controller

import (
	"net/http"

	"github.com/jmznt4345/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /JobTypes
func CreateJobType(c *gin.Context) {
	var JobType entity.JobType
	if err := c.ShouldBindJSON(&JobType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&JobType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": JobType})
}

// GET /JobType/:id
func GetJobType(c *gin.Context) {
	var JobType entity.JobType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Job_Types WHERE id = ?", id).Scan(&JobType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": JobType})
}

// GET /JobTypes
func ListJobTypes(c *gin.Context) {
	var JobTypes []entity.JobType
	if err := entity.DB().Raw("SELECT * FROM Job_Types").Scan(&JobTypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": JobTypes})
}


// DELETE /JobTypes/:id
func DeleteJobType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM JobTypes WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JobType not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /JobTypes
func UpdateJobType(c *gin.Context) {
	var JobType entity.JobType
	if err := c.ShouldBindJSON(&JobType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", JobType.ID).First(&JobType); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JobType not found"})
		return
	}
	if err := entity.DB().Save(&JobType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": JobType})
}

