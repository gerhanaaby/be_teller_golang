package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthStatus struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func MakeFailMessages(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, AuthStatus{
		Status:  "Fail",
		Message: "error," + err.Error(),
	})
}