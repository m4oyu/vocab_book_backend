package main

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"log"
	"my-vocabulary-book/constant"
	"my-vocabulary-book/controller"
	"my-vocabulary-book/controller/uuidgen"
	"my-vocabulary-book/db"
	"my-vocabulary-book/model"
	"os"
	"time"
)

type login struct {
	Mail     string `form:"mail" json:"mail" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	userModel := model.NewUserModel(db.DB)
	userBookModel := model.NewUserBookModel(db.DB)
	translateModel := model.NewTranslateModel()

	uuid := uuidgen.NewUUIDGenerator()

	userHandle := controller.NewUserHandler(userModel, uuid)
	userBookHandler := controller.NewUserBookHandler(userBookModel)
	translateHandler := controller.NewTranslateHandler(userBookModel, translateModel, uuid)

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: constant.IdentityKey,
		// Authenticatorによる認証後、呼び出し
		// Authenticatorの返り値を引数にとりtoken.Claimに組み込む
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				fmt.Println(v.Mail)
				return jwt.MapClaims{
					constant.IdentityKey: v.UserID,
				}
			}
			return jwt.MapClaims{}
		},
		// コンテキストより、ユーザ関連情報を取得
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				UserID: claims[constant.IdentityKey].(string),
			}
		},
		// 認証処理、返り値はPayloadFuncにてtoken.Claimへ組み込み
		Authenticator: func(c *gin.Context) (interface{}, error) {
			// ログインの値受け取り
			var loginValues login
			if err := c.ShouldBind(&loginValues); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			mail := loginValues.Mail
			password := loginValues.Password

			// user情報取得
			user, err := userModel.SelectUserByMail(mail)
			if err != nil {
				log.Println(fmt.Errorf("middleware.Authenticator: %w", err))
			}

			// 承認
			if password == user.Password {
				fmt.Print(user)
				return &model.User{
					UserID:   user.UserID,
					Mail:     user.Mail,
					Password: user.Password,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		// token が有効かどうか確認
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*model.User); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"http://localhost:8080",
			//"http://35.239.221.204:80",
			//"http://35.239.221.204:8080",
			//"http://35.239.221.204",
			"http://www.vocabulary-book.com",
		},
		// preflight requestで許可した後の接続可能時間
		// https://godoc.org/github.com/gin-contrib/cors#Config の中のコメントに詳細あり
		MaxAge: 24 * time.Hour,
	}))

	// 登録, 認証
	r.POST("/signup", userHandle.SignUp())
	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/userBook", userBookHandler.FetchUserBooks())
		auth.POST("/translate", translateHandler.TranslateText())
		auth.DELETE("/userBook", userBookHandler.DeleteUserBook())
	}

	r.Run(":8000")
}
