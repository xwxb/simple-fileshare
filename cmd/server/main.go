package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xwxb/simple-fileshare/internal/app/handler"
	"github.com/xwxb/simple-fileshare/internal/config"
	"github.com/xwxb/simple-fileshare/internal/model"
)

func init() {
	config.LoadConfig()
	model.InitDB()
}

func main() {
	defer model.DeferDBClose()

	r := gin.Default()
	initRouters(r)
	portStr := fmt.Sprintf(":%d", config.GlobalConfig.ServerPort)
	r.Run(portStr)
}

func initRouters(r *gin.Engine) {
	userRoute := r.Group("/users")
	{
		userRoute.POST("/register", handler.UserRegister)
		userRoute.POST("/login", handler.UserLogin)
		//userRoute.POST("/logout", handler.UserLogout)
	}

	//fileRoute := r.Group("/files")
	//fileRoute.Use(middleware.AuthMiddleware)
	//{
	//	fileRoute.POST("/upload", UploadFile)
	//	fileRoute.GET("/list", ListFiles)
	//}
}
