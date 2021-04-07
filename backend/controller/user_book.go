package controller

import "github.com/gin-gonic/gin"

type UserBookHandler interface {
	FetchUserBook() gin.HandlerFunc
}

type userBookHandler struct {
}

func NewUserBookHandler() UserBookHandler {
	return &userBookHandler{}
}

func (h *userBookHandler) FetchUserBook() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
