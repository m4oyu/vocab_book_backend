package controller

import (
	"github.com/gin-gonic/gin"
	"log"
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
	UUID      uuidgen.UUIDGenerator
}

func NewUserHandler(um model.UserModel, uuid uuidgen.UUIDGenerator) UserHandler {
	return &userHandler{
		UserModel: um,
		UUID:      uuid,
	}
}

func (h *userHandler) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// receive request
		var req SignUpRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("userHandler.SignUp: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
			return
		}

		// generate uuid
		uuid, err := h.UUID.GenNewRandom()
		if err != nil {
			log.Printf("userHandler.SignUp: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// insert user into db
		err = h.UserModel.InsertUser(&model.User{
			UserID:   uuid,
			Mail:     req.Mail,
			Password: req.Password,
		})
		if err != nil {
			log.Printf("userHandler.SignUp: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// return response
		c.JSON(http.StatusOK, gin.H{"message": "Sign up completed"})
		return
	}
}

type SignUpRequest struct {
	Mail     string `form:"mail" json:"mail" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (h *userHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
