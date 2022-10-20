package controller

import (
	"net/http"

	"github.com/B6311636/sa-G12-Subject/entity"
	"github.com/gin-gonic/gin"
)

// GET /officers
// List all officers
func ListOfficers(c *gin.Context) {
	var officers []entity.Officer
	if err := entity.DB().Raw("SELECT * FROM officers").Scan(&officers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": officers})
}

// GET /officer/:id
// Get officer by id
func GetOfficer(c *gin.Context) {
	var officer entity.Officer
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&officer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "officer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": officer})
}

// PATCH /officers
func UpdateOfficer(c *gin.Context) {
	var officer entity.Officer
	if err := c.ShouldBindJSON(&officer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", officer.ID).First(&officer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "officer not found"})
		return
	}

	if err := entity.DB().Save(&officer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": officer})
}

// DELETE /officers/:id
func DeleteOfficer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM officers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "officer not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.officer{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}
