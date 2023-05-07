package controllers

import (
	"gorest/database"
	"gorest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreatePhotos struct {
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_Url string `json:"photo_url"`
}

// CreatePhoto godoc
// @Summary Create a new Photo
// @Description Create Photo with the given information
// @Tags Photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'"
// @Param socialmedia body CreatePhotos true "The social media to create"
// @Success 201 {object} models.Photo
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /photo/ [post]
func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}

	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = db.WithContext(ctx).Create(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, photo)
}

// GetOnePhoto godoc
// @Summary Get One Photo
// @Description Get One Photo by ID Photo
// @Tags Photo
// @Accept json
// @Produce json
// @Param id path int true "Photo ID"
// @Param Authorization header string true "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'"
// @Success 200 {object} models.Photo
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /photo/{id} [get]
func GetOnePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}
	photoID, _ := strconv.Atoi(ctx.Param("photoID"))

	err := db.WithContext(ctx).First(&photo, photoID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Photo not found",
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

type UpdatePhotos struct {
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_Url string `json:"photo_url"`
}

// UpdatePhoto godoc
// @Summary Update Photo
// @Description Update Photo with the given information
// @Tags Photo
// @Accept json
// @Produce json
// @Param id path int true "Photo ID"
// @Param Authorization header string true "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'"
// @Param socialmedia body UpdatePhotos true "The social media to create"
// @Success 200 {object} models.Photo
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /photo/{id} [put]
func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}
	photoID, _ := strconv.Atoi(ctx.Param("photoID"))

	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = db.WithContext(ctx).Model(&photo).Where("id=?", photoID).Updates(models.Photo{Title: photo.Title, Caption: photo.Caption, Photo_Url: photo.Photo_Url}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

// DeletePhoto godoc
// @Summary Delete Photo
// @Description Delete Photo by ID Photo
// @Tags Photo
// @Accept json
// @Produce json
// @Param id path int true "Photo ID"
// @Param Authorization header string true "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'"
// @Success 200 {object} models.Photo
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /photo/{id} [delete]
func DeletePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photoID, _ := strconv.Atoi(ctx.Param("photoID"))

	photo := models.Photo{}
	err := db.WithContext(ctx).First(&photo, photoID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Photo not found",
		})
		return
	}

	err = db.Delete(&models.Photo{}, photoID).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Photo has been deleted",
	})
}

// GetAllPhoto godoc
// @Summary Get All Photo
// @Description Get All Photo by ID Photo
// @Tags Photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'"
// @Success 200 {object} models.Photo
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /photo/ [get]
func GetAllPhoto(ctx *gin.Context) {
	db := database.GetDB()
	photoList := []models.Photo{}

	err := db.WithContext(ctx).Find(&photoList).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get Ptoto Data",
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, photoList)
}
