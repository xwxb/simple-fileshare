package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/xwxb/simple-fileshare/internal/config"
	"github.com/xwxb/simple-fileshare/internal/global"
	"github.com/xwxb/simple-fileshare/internal/model"
	"github.com/xwxb/simple-fileshare/pkg/utils"
	"time"
)

type UserRegInfo struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func CreateNewUser(username, email, passwd string) (user *UserRegInfo, err error) {
	passwdHash, err := utils.HashPassword(passwd)
	if err != nil {
		return
	}

	var uid int64
	uid, err = model.InsertOneUser(&model.User{
		Username:     username,
		Email:        email,
		PasswordHash: passwdHash,
	})
	if err != nil {
		return
	}

	user = &UserRegInfo{
		Id:       uid,
		Username: username,
		Email:    email,
	}
	return
}

type UserAuthClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func LoginUser(username, passwd string) (token string, err error) {
	passwdHash, err := GetUserPasswordHash(username)
	if err != nil {
		return
	}

	if !utils.CheckPasswordHash(passwd, passwdHash) {
		err = errors.New("password error")
		return
	}

	expTime := time.Now().Add(global.DefaultUserLoginExpire)
	claims := &UserAuthClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenObj.SignedString([]byte(config.GlobalConfig.JWTSecretKey))
	if err != nil {
		return
	}
	return
}

func GetUserPasswordHash(username string) (passwdHash string, err error) {
	user, err := model.FindFirstUserByUsername(username)
	if err != nil {
		return
	}
	passwdHash = user.PasswordHash
	return
}
