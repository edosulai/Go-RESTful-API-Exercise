package controllers

import (
	"gorest/database"
	"gorest/helpers"
	"gorest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      uint8  `json:"age"`
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the given information
// @Tags User
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "The user to create"
// @Success 201 {object} models.User
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 500 {object} helpers.ErrorResponse
// @Router /user/register [post]
func Register(ctx *gin.Context) {
	db := database.GetDB()
	user := models.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Message: "Internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

type SuccessResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login godoc
// @Summary Login a user
// @Description Authenticates a user with their email and password and returns a token.
// @Tags user
// @Accept json
// @Produce json
// @Param user body LoginRequest true "Login for authentication"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} helpers.ErrorResponse
// @Failure 401 {object} helpers.ErrorResponse
// @Router /user/login [post]
func Login(ctx *gin.Context) {
	db := database.GetDB()
	user := models.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			Message: "Bad request",
			Error:   err.Error(),
		})
		return
	}

	password := user.Password

	err = db.Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			Message: "User not found",
			Error:   "The email and password you provided are not associated with an existing account",
		})
		return
	}

	if !helpers.PasswordValid(user.Password, password) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			Message: "Invalid Email or Password",
			Error:   "You must fill in the correct email and password and have registered",
		})
		return
	}

	token, _ := helpers.GenerateToken(user.ID, user.Email)

	ctx.JSON(http.StatusCreated, SuccessResponse{
		Token: token,
	})
}
