package controller

import (
	"net/http"

	"github.com/jmznt4345/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /dmg_levels
func CreateDMGLevel(c *gin.Context) {
	var dmg_level entity.DMGLevel
	if err := c.ShouldBindJSON(&dmg_level); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&dmg_level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dmg_level})
}

// GET /DMGLevel
func GetDMGLevel(c *gin.Context) {
	var dmg_level entity.DMGLevel
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM dmg_levels WHERE id = ?", id).Scan(&dmg_level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dmg_level})
}

// GET /DMGLevels
func ListDMGLevels(c *gin.Context) {
	var dmg_levels []entity.DMGLevel
	if err := entity.DB().Raw("SELECT * FROM dmg_levels").Scan(&dmg_levels).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dmg_levels})
}

// DELETE /dmg_levels/:id
func DeleteDMGLevel(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM dmg_levels WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dmg_level not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /dmg_levels
func UpdateDMGLevel(c *gin.Context) {
	var dmg_level entity.DMGLevel
	if err := c.ShouldBindJSON(&dmg_level); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", dmg_level.ID).First(&dmg_level); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dmg_level not found"})
		return
	}
	if err := entity.DB().Save(&dmg_level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dmg_level})
}
