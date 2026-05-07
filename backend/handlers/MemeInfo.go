package handlers

import (
	"memeboard/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MemeInfo struct {
	DB *gorm.DB
}

func NewMemeInfo(db *gorm.DB) *MemeInfo {
	return &MemeInfo{DB: db}
}

func (m *MemeInfo) GetMemeInfo(c *gin.Context) {
	var Memes []models.Meme
	m.DB.Find(&Memes)
	c.JSON(http.StatusOK, Memes)
}
func (m *MemeInfo) UploadNewMeme(c *gin.Context) {
	var req models.Meme
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := m.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create meme",
		})
		return
	}
	c.JSON(http.StatusCreated, req)
}
func (m *MemeInfo) UpdateMemeInfo(c *gin.Context) {
	var req models.Meme
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if req.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing meme id",
		})
		return
	}
	var updatingMeme models.Meme
	if err := m.DB.First(&updatingMeme, req.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Meme not found",
		})
		return
	}
	if err := m.DB.Model(&updatingMeme).Update("content", req.Content).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update meme",
		})
		return
	}
	c.JSON(http.StatusOK, updatingMeme)
}
func (m *MemeInfo) DeleteMeme(c *gin.Context) {
	id := c.Param("id")
	var deletingMeme models.Meme
	if err := m.DB.First(&deletingMeme, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Meme not found",
		})
		return
	}
	if err := m.DB.Delete(&deletingMeme).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete meme",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successful delete",
	})
}
