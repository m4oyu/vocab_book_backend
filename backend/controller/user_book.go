package controller

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"my-vocabulary-book/constant"
	"my-vocabulary-book/model"
	"net/http"
)

type UserBookHandler interface {
	FetchUserBooks() gin.HandlerFunc
}

type userBookHandler struct {
	UserBookModel model.UserBookModel
}

func NewUserBookHandler(ubm model.UserBookModel) UserBookHandler {
	return &userBookHandler{
		UserBookModel: ubm,
	}
}

func (h *userBookHandler) FetchUserBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		// user get
		userID := jwt.ExtractClaims(c)[constant.IdentityKey].(string)
		if userID == "" {
			fmt.Println("controller.FetchUserBook: userID is nil")
			c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		}

		// select user book
		userBooks, err := h.UserBookModel.SelectUserBooksByUserID(userID)
		if err != nil {
			log.Printf("controller.FetchUserBook: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "database connection failed"})
		}

		// create response
		var Books []UserBook
		for _, value := range userBooks {
			Book := UserBook{
				ID:       value.ID,
				UserID:   value.UserID,
				English:  value.English,
				Japanese: value.Japanese,
			}
			Books = append(Books, Book)
		}

		// return response
		c.JSON(http.StatusOK, UserBooksResponse{UserBooks: Books})
		return
	}
}

type UserBooksResponse struct {
	UserBooks []UserBook `json:"userBooks"`
}

type UserBook struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	English  string `json:"english"`
	Japanese string `json:"japanese"`
}
