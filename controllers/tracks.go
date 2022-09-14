package controllers

import (
	"file-share/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CreateTrackInput struct {
	Artist string `json:"artist" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type UpdateTrackInput struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

func GetAllTracks(context *gin.Context) {
	var tracks []models.Track
	models.DB.Find(&tracks)

	context.JSON(http.StatusOK, gin.H{"tracks": tracks})
}

func CreateTrack(context *gin.Context) {
	var input CreateTrackInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	track := models.Track{Artist: input.Artist, Title: input.Title}
	models.DB.Create(&track)

	context.JSON(http.StatusOK, gin.H{"tracks": track})
}

func GetTrack(context *gin.Context) {
	var track models.Track
	if err := models.DB.Where("id = ?", context.Param("id")).First(&track).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"track": track})
}

func UpdateTrack(context *gin.Context) {
	var track models.Track
	if err := models.DB.Where("id = ?", context.Param("id")).First(&track).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateTrackInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&track).Update(input)

	context.JSON(http.StatusOK, gin.H{"tracks": track})
}

func DeleteTrack(context *gin.Context) {
	var track models.Track
	if err := models.DB.Where("id = ?", context.Param("id")).First(&track).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	models.DB.Delete(&track)

	context.JSON(http.StatusOK, gin.H{"tracks": true})
}
