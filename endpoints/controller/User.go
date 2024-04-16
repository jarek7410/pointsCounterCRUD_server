package controller

import (
	"errors"
	"fmt"
	"net/http"
	model2 "pointsCounterCRUD/model"
	"pointsCounterCRUD/model/dto"
	"pointsCounterCRUD/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Register user
func Register(context *gin.Context) {
	var input dto.Register

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model2.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		RoleID:   3,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})

}

// User Login
func Login(context *gin.Context) {
	var input dto.Login

	if err := context.ShouldBindJSON(&input); err != nil {
		var errorMessage string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			validationError := validationErrors[0]
			if validationError.Tag() == "required" {
				errorMessage = fmt.Sprintf("%s not provided", validationError.Field())
			}
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	user, err := model2.GetUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidateUserPassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwt, err := util.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": jwt, "username": input.Username, "message": "Successfully logged in"})

}

// get all users
func GetUsers(context *gin.Context) {
	var user []model2.User
	err := model2.GetUsers(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, user)
}

// get user by id
func GetUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var user model2.User
	err := model2.GetUser(&user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, user)
}

// update user
func UpdateUser(c *gin.Context) {
	//var input model.Update
	var User model2.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := model2.GetUser(&User, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = c.BindJSON(&User)
	if err != nil {
		return
	}
	err = model2.UpdateUser(&User)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, User)
}
