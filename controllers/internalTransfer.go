package controllers

import (
	"encoding/json"
	"net/http"
	"teller/db"
	"teller/models"
	"teller/services"

	"github.com/gin-gonic/gin"
)

func PostInternalTransfer(c *gin.Context) {
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
	request := models.InternalTransfer{}

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

	dataResponse, err := PostToAPIInternal(request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due" + err.Error(),
		})
		return

	}

	logTransaction := dataResponse["$errorInfo"].(map[string]interface{})["$pipeline"].(map[string]interface{})["logTransaction"]
	inputMessage := dataResponse["$errorInfo"].(map[string]interface{})["$pipeline"].(map[string]interface{})["inputMessage"]

	c.JSON(http.StatusCreated, gin.H{
		"logTransaction": logTransaction,
		"inputMessage":   inputMessage,
	})
}

func PostToAPIInternal(dataInternal models.InternalTransfer) (map[string]interface{}, error) {

	data := map[string]interface{}{
		"referenceId":      dataInternal.ReferenceId,
		"debitAccountNo":   dataInternal.DebitAccountNo,
		"creditAccountNo":  dataInternal.CreditAccountNo,
		"creditAmount":     dataInternal.CreditAmount,
		"creditCurrency":   dataInternal.CreditCurrency,
		"transactionDate":  dataInternal.TransactionDate,
		"remark":           dataInternal.Remark,
		"beneficiaryName":  dataInternal.BeneficiaryName,
		"debitAccountName": dataInternal.DebitAccountName,
	}

	requestJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	body, err := services.ConsumeAPIService("internaltransfer", requestJson)
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
