package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"pustaka-api/helper"
	"pustaka-api/jwtUse"
	"pustaka-api/users"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	userService users.Service
}

func NewUserHandler(userService users.Service) *userHandler {
	return &userHandler{userService}
}

func (h userHandler) Login(c *gin.Context) {
	session := sessions.Default(c)
	var userRequest users.UserRequestFindOne
	err := c.ShouldBindJSON(&userRequest)

	var jsonErr *json.UnmarshalTypeError
	if errors.As(err, &jsonErr) {
		ermsg := fmt.Sprintf("error Field %s value %s must be %s ", jsonErr.Field, jsonErr.Value, jsonErr.Type)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ermsg,
		})
		return
	}
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	user, err := h.userService.FindOne(userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid Password",
		})
		return
	}

	token := jwtUse.NewJWTService().GenerateToken(user.Username)
	session.Set("token", token)
	session.Save()

	res := helper.BuildResponse(true, "successed Login", helper.EmptyObj{})
	c.JSON(http.StatusOK, res)
}

func (h userHandler) SignUp(c *gin.Context) {
	//GET body
	var userRequest users.UserRequestCreate
	err := c.ShouldBindJSON(&userRequest)

	//Check validation Body
	var jsonErr *json.UnmarshalTypeError
	if errors.As(err, &jsonErr) {
		ermsg := fmt.Sprintf("error Field %s value %s must be %s ", jsonErr.Field, jsonErr.Value, jsonErr.Type)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ermsg,
		})
		return
	}
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	//Hash Password User
	password := []byte(userRequest.Password)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userRequest.Password = string(hash)

	//Insert data User to Service
	user, err := h.userService.CreateUser(userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Set Resp
	var resp = users.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	})

	return
}
