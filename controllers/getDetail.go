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

func PostGetDetail(c *gin.Context) {
	request := models.GetDetail{}

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
	PostToAPIGetDetail(request)

	dataResponse := PostToAPIGetDetail(request)

	c.JSON(http.StatusCreated, dataResponse)

}

func PostToAPIGetDetail(dataGetDetail models.GetDetail) map[string]interface{} {

	data := map[string]interface{}{
		"transactionID": dataGetDetail.TransactionID,
		"accountNumber": dataGetDetail.AccountNumber,
	}

	requestJson, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}

	request, err := http.NewRequest("POST", "https://apidev.banksinarmas.com/internal/accounts/management/v2.0/getDetail", bytes.NewBuffer(requestJson))

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
