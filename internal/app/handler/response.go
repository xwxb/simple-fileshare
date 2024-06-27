package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    ErrCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    ErrCodeSuccess,
		Message: errMsgs[ErrCodeSuccess],
		Data:    data,
	})
}

func Fail(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    ErrCodeInternal,
		Message: err.Error(),
	})
}

func FailWithCode(c *gin.Context, code ErrCode) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    code,
		Message: errMsgs[code],
	})
}
