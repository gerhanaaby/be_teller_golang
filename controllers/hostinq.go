package controllers

import (
	"net/http"
	"teller/models"

	"github.com/gin-gonic/gin"
)

func HostInquiry(c *gin.Context) {
	var request models.HostInqRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//convert to iso
}

func GET_HI_ST_01_01(request *models.HostInqRequest) {

	request.Branch = "A"
}
