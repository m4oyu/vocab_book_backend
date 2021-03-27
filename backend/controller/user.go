package controller

import (
	"github.com/gin-gonic/gin"
	"my-vocabulary-book/controller/uuidgen"
	"my-vocabulary-book/model"
	"net/http"
)

type UserHandler interface {
	SignUp() gin.HandlerFunc
	Update() gin.HandlerFunc
}

type userHandler struct {
	UserModel model.UserModel
	UUID uuidgen.UUIDGenerator
}

func NewUserHandler(um model.UserModel, uuid uuidgen.UUIDGenerator) UserHandler {
	return &userHandler{
		UserModel: um,
		UUID: uuid,
	}
}

func (h *userHandler) SignUp() gin.HandlerFunc{
	return func(c *gin.Context) {
		// receive request
		var req SignUpRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// generate uuid
		uuid, err := h.UUID.GenNewRandom()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		// insert user into db
		err = h.UserModel.InsertUser(&model.User{
			UserID: uuid,
			Mail: req.Mail,
			Password: req.Password,
		})

		// return response
		c.JSON(http.StatusOK, gin.H{"status": "Successful sign-up"})
	}
}

type SignUpRequest struct {
	Mail string `form:"mail" json:"mail" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`}

func (h *userHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

