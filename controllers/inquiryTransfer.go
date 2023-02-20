package controllers

import (
	"encoding/json"
	"net/http"
	"teller/db"
	"teller/models"
	"teller/services"

	"github.com/gin-gonic/gin"
)

func PostInquiryTransfer(c *gin.Context) {
	var isValid bool = false
	var err error

	//--> TOKEN VALIDATION REQUEST
	isValid, err = services.CheckToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due" + err.Error(),
		})
		return
	}

	if !isValid {
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due invalid login or expired",
		})
		return
	}

	//--> API REQUEST PROCESS
	request := models.InquiryTransfer{}

	if err = c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due" + err.Error(),
		})
		return
	}

	err = db.GetDB().Debug().Create(&request).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due" + err.Error(),
		})
		return
	}

	dataResponse, err := PostToAPIInquiry(request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due" + err.Error(),
		})
		return

	}
	c.JSON(http.StatusCreated, dataResponse)
}

func PostToAPIInquiry(dataInquiry models.InquiryTransfer) (map[string]interface{}, error) {

	data := map[string]interface{}{
		"accountNo":   dataInquiry.AccountNo,
		"referenceId": dataInquiry.ReferenceId,
	}

	requestJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	body, err := services.ConsumeAPIService("inquirytransfer", requestJson)
	if err != nil {
		return nil, err
	}

	var dataResponse map[string]interface{}
	err = json.Unmarshal(body, &dataResponse)

	if err != nil {
		return nil, err
	}

	return dataResponse, nil
}
