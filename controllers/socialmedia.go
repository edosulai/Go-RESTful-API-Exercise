package controllers

import (
	"gorest/database"
	"gorest/helpers"
	"gorest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateSocial struct {
	Name             string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
}

// CreateSocialMedia godoc
// @Summary Create a new Social Media
// @Description CreateSocialMedia a new socialmedia with the given information
// @Tags Social Media
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'"
// @Param socialmedia body CreateSocial true "The social media to create"
// @Success 201 {object} models.SocialMedia
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /socialmedia/ [post]
func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialmedia := models.SocialMedia{}

	err := ctx.ShouldBindJSON(&socialmedia)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
		return
	}

	err = db.WithContext(ctx).Create(&socialmedia).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Message: "Internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, socialmedia)
}

// GetOneSocialMedia godoc
// @Summary Get One Social Media
// @Description Get One Social Media by ID Social Media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param id path int true "Social Media ID"
// @Param Authorization header string true "Bearer {access_token}"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /socialmedia/{id} [get]
func GetOneSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialmedia := models.SocialMedia{}
	socialmediaID, _ := strconv.Atoi(ctx.Param("socialmediaID"))

	err := db.WithContext(ctx).First(&socialmedia, socialmediaID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Social Media not found",
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, socialmedia)
}

// GetAllSocialMedia godoc
// @Summary Get All Social Media
// @Description Get All Social Media Which Has Been Made
// @Tags Social Media
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token}"
// @Success 200 {array} models.SocialMedia
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /socialmedia [get]
func GetAllSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialmediaList := []models.SocialMedia{}

	err := db.WithContext(ctx).Find(&socialmediaList).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get Social Media data",
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, socialmediaList)
}

type UpdateSocial struct {
	Name             string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
}

// UpdateSocialMedia godoc
// @Summary Update Social Media
// @Description Update Social Media by ID Social Media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param id path int true "Social Media ID"
// @Param Authorization header string true "Bearer {access_token}"
// @Param socialmedia body UpdateSocial true "The Social Media Has Been Updated"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /socialmedia/{id} [put]
func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialmedia := models.SocialMedia{}
	socialmediaID, _ := strconv.Atoi(ctx.Param("socialmediaID"))

	err := ctx.ShouldBindJSON(&socialmedia)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = db.WithContext(ctx).Model(&socialmedia).Where("id=?", socialmediaID).Updates(models.SocialMedia{Name: socialmedia.Name, Social_Media_Url: socialmedia.Social_Media_Url}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, socialmedia)
}

// DeleteSocialMedia godoc
// @Summary Delete Social Media
// @Description Delete Social Media by ID Social Media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token}"
// @Param id path int true "Social Media ID"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /socialmedia/{id} [delete]
func DeleteSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialmediaID, _ := strconv.Atoi(ctx.Param("socialmediaID"))

	socialmedia := models.SocialMedia{}
	err := db.WithContext(ctx).First(&socialmedia, socialmediaID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Social media not found",
		})
		return
	}

	err = db.Delete(&models.SocialMedia{}, socialmediaID).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Social media has been deleted",
	})
}
