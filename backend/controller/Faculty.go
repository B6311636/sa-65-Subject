package controller

import (
	"net/http"

	"github.com/B6311636/sa-G12-Subject/entity"
	"github.com/gin-gonic/gin"
)

// POST /Facultys
func CreateFaculty(c *gin.Context) {
	var faculty entity.Faculty
	if err := c.ShouldBindJSON(&faculty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&faculty).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": faculty})
}

// GET /faculty/:id
func GetFaculty(c *gin.Context) {
	var faculty entity.Faculty
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&faculty); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "faculty not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": faculty})
}

// GET /facultys
func ListFaculties(c *gin.Context) {
	var faculties []entity.Faculty
	if err := entity.DB().Raw("SELECT * FROM faculties").Scan(&faculties).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": faculties})
}

// DELETE /faculties/:id
func DeleteFaculty(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM faculties WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "faculty not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /faculties
func UpdateFaculty(c *gin.Context) {
	var faculty entity.Faculty
	if err := c.ShouldBindJSON(&faculty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", faculty.ID).First(&faculty); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "faculty not found"})
		return
	}

	if err := entity.DB().Save(&faculty).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": faculty})
}
