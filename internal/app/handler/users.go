package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xwxb/simple-fileshare/internal/app/service"
	"github.com/xwxb/simple-fileshare/pkg/utils"
)

type RegisterUserReq struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

func UserRegister(c *gin.Context) {
	var req RegisterUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Fail(c, err)
		return
	}

	if !utils.IsValidEmail(req.Email) {
		FailWithCode(c, ErrCodeInvalidParam)
		return
	}

	info, err := service.CreateNewUser(req.Username, req.Email, req.Password)
	if err != nil {
		Fail(c, err)
		return
	}
	Success(c, info)
}

type UserLoginCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserLogin(c *gin.Context) {
	var creds UserLoginCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		Fail(c, err)
		return
	}

	token, err := service.LoginUser(creds.Username, creds.Password)
	if err != nil {
		Fail(c, err)
		return
	}

	Success(c, token)
}
