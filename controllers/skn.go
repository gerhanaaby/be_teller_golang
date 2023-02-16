package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"teller/db"
	"teller/models"

	"github.com/gin-gonic/gin"
)

func PostSkn(c *gin.Context) {
	var isValid bool
	var err error
	reqToken := c.Request.Header.Get("Authorization")
	if reqToken == ""{
		c.AbortWithError(http.StatusInternalServerError, errors.New("error, empty token"))
		c.JSON(http.StatusBadRequest, AuthStatus{
		Status: "Fail", 
		Message: "error, empty token"})
	}
	
	token := strings.Replace(reqToken, "Bearer ", "", 1)

	if isValid, err = ValidateToken(token); err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status: "Fail", 
			Message: "unable to process due to error processing token, "+err.Error(),
		})
		return 
	}

	if !isValid {
		c.AbortWithError(http.StatusBadRequest, errors.New("inValid/expired User Login"))
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status: "Fail", 
			Message: "unable to precess due to invalid or epired login"+err.Error(),
		})
		return
	}
	//


	request := models.Skn{}

	if err = c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.GetDB().Debug().Create(&request).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("===============================================")
	fmt.Println("request", request)

	fmt.Println("===============================================")
	PostToAPIdev(request)

	dataResponse := PostToAPIdev(request)

	c.JSON(http.StatusCreated, dataResponse)

}

func PostToAPIdev(dataSKN models.Skn) map[string]interface{} {

	data := map[string]interface{}{
		"creditAccountNo":           dataSKN.CreditAccountNo,
		"amount":                    dataSKN.Amount,
		"beneficiaryResidentStatus": dataSKN.BeneficiaryResidentStatus,
		"clearingCode":              dataSKN.ClearingCode,
		"remark":                    dataSKN.Remark,
		"transactionDate":           dataSKN.TransactionDate,
		"transactionTime":           dataSKN.TransactionTime,
		"clearingTransactionCode":   dataSKN.ClearingTransactionCode,
		"referenceId":               dataSKN.ReferenceId,
		"paymentDetails1":           dataSKN.PaymentDetails1,
		"senderName":                dataSKN.SenderName,
		"paymentDetails2":           dataSKN.PaymentDetails2,
		"paymentDetails3":           dataSKN.PaymentDetails3,
		"debitAccountNo":            dataSKN.DebitAccountNo,
		"beneficiaryNationStatus":   dataSKN.BeneficiaryNationStatus,
		"beneficiaryType":           dataSKN.BeneficiaryType,
		"beneficiaryName":           dataSKN.BeneficiaryName,
		"chargeAmount":              dataSKN.ChargeAmount,
		"currency":                  dataSKN.Currency,
	}

	requestJson, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}

	request, err := http.NewRequest("POST", "https://apidev.banksinarmas.com/internal/transactions/transfer/v2.0/skn", bytes.NewBuffer(requestJson))

	request.Header.Set("Content-type", "application/json")
	request.Header.Set("x-gateway-apikey", "97817cac-d589-4d9c-b9bf-a874f0ff943d")

	if err != nil {
		log.Fatalln(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)

	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var dataResponse map[string]interface{}
	err = json.Unmarshal(body, &dataResponse)

	// Check your errors!
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("===============================================")

	fmt.Println(string(body))
	return dataResponse

}
