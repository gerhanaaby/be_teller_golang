package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"teller/db"
	"teller/models"

	"github.com/gin-gonic/gin"
)

func PostInquiryTransfer(c *gin.Context) {
	request := models.InquiryTransfer{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.GetDB().Debug().Create(&request).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("===============================================")
	fmt.Println("request inquiry", request)

	fmt.Println("===============================================")
	PostToAPIInquiry(request)

	dataResponse := PostToAPIInquiry(request)

	c.JSON(http.StatusCreated, dataResponse)

}

func PostToAPIInquiry(dataInquiry models.InquiryTransfer) map[string]interface{} {

	data := map[string]interface{}{
		"accountNo":   dataInquiry.AccountNo,
		"referenceId": dataInquiry.ReferenceId,
	}

	requestJson, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}

	request, err := http.NewRequest("POST", "https://apidev.banksinarmas.com/internal/transactions/transfer/v2.0/inquiry", bytes.NewBuffer(requestJson))

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
