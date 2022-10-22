package controller

import (
	"net/http"

	"github.com/jmznt4345/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /distributors
func CreateDistributor(c *gin.Context) {
	var distributor entity.Distributor
	if err := c.ShouldBindJSON(&distributor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&distributor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": distributor})
}

// GET /Distributor
func GetDistributor(c *gin.Context) {
	var distributor entity.Distributor
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM distributors WHERE id = ?", id).Scan(&distributor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": distributor})
}

// GET /Distributors
func ListDistributors(c *gin.Context) {
	var distributors []entity.Distributor
	if err := entity.DB().Raw("SELECT * FROM distributors").Scan(&distributors).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": distributors})
}

// DELETE /distributors/:id
func DeleteDistributor(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM distributors WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "distributor not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /distributors
func UpdateDistributor(c *gin.Context) {
	var distributor entity.Distributor
	if err := c.ShouldBindJSON(&distributor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", distributor.ID).First(&distributor); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "distributor not found"})
		return
	}
	if err := entity.DB().Save(&distributor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": distributor})
}
