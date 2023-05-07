package controllers

import (
	"gorest/database"
	"gorest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateComments struct {
	Message string `json:"message"`
	PhotoID uint   `json:"photo_id"`
}

// CreateComment godoc
// @Summary Create a new Comment
// @Description Create Comment with the given information
// @Tags Comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'"
// @Param comment body CreateComments true "Comment has been created"
// @Success 201 {object} models.Comment
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /comment/ [post]
func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	photo := models.Photo{}

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = db.WithContext(ctx).First(&photo, comment.PhotoID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Photo not found",
		})
		return
	}

	err = db.WithContext(ctx).Create(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

// GetOneComment godoc
// @Summary Get One Comment
// @Description Get One Comment by ID Comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Param Authorization header string true "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'"
// @Success 200 {object} models.Comment
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /comment/{id} [get]
func GetOneComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	commentID, _ := strconv.Atoi(ctx.Param("commentID"))

	err := db.WithContext(ctx).First(&comment, commentID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Comment not found",
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

type UpdateComments struct {
	Message string `json:"message"`
	PhotoID uint   `json:"photo_id"`
}

// UpdateComment godoc
// @Summary Update Comment
// @Description Update Comment by ID Comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Param Authorization header string true "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'"
// @Param comment body UpdateComments true "Comment has been Updated"
// @Success 200 {object} models.Comment
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /comment/{id} [put]
func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	commentID, _ := strconv.Atoi(ctx.Param("commentID"))

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = db.WithContext(ctx).Model(&comment).Where("id=?", commentID).Updates(models.Comment{Message: comment.Message}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// DeleteComment godoc
// @Summary Delete Comment
// @Description Delete Comment by ID Comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Param Authorization header string true "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'"
// @Success 200 {object} models.Comment
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /comment/{id} [delete]
func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()
	commentID, _ := strconv.Atoi(ctx.Param("commentID"))

	comment := models.Comment{}
	err := db.WithContext(ctx).First(&comment, commentID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Comment not found",
		})
		return
	}

	err = db.Delete(&models.Comment{}, commentID).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Comment has been deleted",
	})
}

// GetAllComment godoc
// @Summary Get All Comment
// @Description Get All Comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'"
// @Success 200 {object} models.Comment
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /comment/ [get]
func GetAllComment(ctx *gin.Context) {
	db := database.GetDB()
	commentList := []models.Comment{}

	err := db.WithContext(ctx).Find(&commentList).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get Comment data",
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, commentList)
}
