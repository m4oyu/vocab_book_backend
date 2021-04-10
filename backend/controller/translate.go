package controller

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"my-vocabulary-book/constant"
	"my-vocabulary-book/controller/uuidgen"
	"my-vocabulary-book/model"
	"net/http"
)

type TranslateHandler interface {
	TranslateText() gin.HandlerFunc
}

type translateHandler struct {
	UserBookModel  model.UserBookModel
	TranslateModel model.TranslateModel
	UUID           uuidgen.UUIDGenerator
}

func NewTranslateHandler(ubm model.UserBookModel, tm model.TranslateModel, uuid uuidgen.UUIDGenerator) TranslateHandler {
	return &translateHandler{
		UserBookModel:  ubm,
		TranslateModel: tm,
		UUID:           uuid,
	}
}

func (h *translateHandler) TranslateText() gin.HandlerFunc {
	return func(c *gin.Context) {
		// user get
		userID := jwt.ExtractClaims(c)[constant.IdentityKey].(string)
		//user, ok := c.Get(constant.IdentityKey)
		//if ok != false {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
		//	return
		//}

		// request receive
		var req TranslateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("translateHandler.TranslateText: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
			return
		}

		// throw translate request to gcp api
		response, err := h.TranslateModel.TranslateAPI("ja", req.Text)
		if err != nil {
			log.Printf("translateHandler.TranslateText: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
		fmt.Print(response)

		// generate uuid
		uuid, err := h.UUID.GenNewRandom()
		if err != nil {
			log.Printf("translateHandler.TranslateText: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// create new data
		userBook := model.UserBook{
			ID:       uuid,
			UserID:   userID,
			English:  req.Text,
			Japanese: response,
		}

		// insert into db
		err = h.UserBookModel.InsertUserBook(&userBook)
		if err != nil {
			log.Printf("translateHandler.TranslateText: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		// return response
		c.JSON(http.StatusOK, gin.H{
			"text": response,
		})
		return
	}
}

type TranslateRequest struct {
	Text string `json:"text" binding:"required"`
}
