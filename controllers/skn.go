package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"teller/db"
	"teller/inits"
	"teller/models"
	"teller/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func SKN(c *gin.Context) {
	startTime := time.Now()

	
	if RefID, ReqTime, Response, err := TransactSKN(c); err != nil {
		services.WriteLog(
			"[fail][skn]",
			fmt.Sprintf("REF-ID: %s, Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
				RefID,
				time.Since(startTime).Milliseconds()-ReqTime,
				ReqTime,
				time.Since(startTime).Milliseconds()),
			inits.Cfg.LogPerformancePath+services.LogFileName, "performance")

		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "error, " + err.Error(),
		})

	} else {
		services.WriteLog(
			"[done][skn]",
			fmt.Sprintf("REF-ID: %s, Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
				RefID,
				time.Since(startTime).Milliseconds()-ReqTime,
				ReqTime,
				time.Since(startTime).Milliseconds()),
			inits.Cfg.LogPerformancePath+services.LogFileName, "performance")
		c.JSON(http.StatusCreated, Response)
	}
}

func TransactSKN(c *gin.Context) (refid string, reqApiTime int64, dataResponse map[string]interface{}, err error) {
	var isValid bool = false
	var claims jwt.MapClaims

	//--> TOKEN VALIDATION REQUEST
	claims, isValid, err = services.CheckToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		return `NOID`, 0, nil, err
	}

	if !isValid {
		return `NOID`, 0, nil, errors.New("error, authentication failure")
	}

	// Convert map to json string
	jsonStr, err := json.Marshal(claims["User"])
	if err != nil {
		return `NOID`, 0, nil, err
	}

	user := models.User{}
	// Convert json string to struct
	if err := json.Unmarshal(jsonStr, &user); err != nil {
		return `NOID`,0, nil, err
	}

	//--> API REQUEST PROCESS
	request := models.Skn{}
	if err = c.ShouldBindJSON(&request); err != nil {
		return `NOID`, 0, nil, err
	}

	request.ReferenceId, err = services.GenTransactID("MDLN-", user.Nik)
	if err != nil {
		return `NOID`, 0, nil, err
	}

	err = db.GetDB().Create(&request).Error
	if err != nil {
		return request.ReferenceId, 0, nil, err
	}

	reqApiTime, dataResponse, err = PostToAPIdev(request)
	if err != nil {
		return request.ReferenceId, 0, nil, err
	}

	dataResponse["caseID"] = request.CaseID

	return request.ReferenceId, reqApiTime, dataResponse, nil

}

func PostToAPIdev(dataSKN models.Skn) (int64, map[string]interface{}, error) {
	start := time.Now()
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
		return 0, nil, err
	}

	body, err := services.ConsumeAPIService("skn", requestJson)
	if err != nil {
		return 0, nil, err
	}

	var dataResponse map[string]interface{}
	err = json.Unmarshal(body, &dataResponse)

	if err != nil {
		return 0, nil, err
	}

	dst := &bytes.Buffer{}
	if err := json.Compact(dst, body); err != nil {
		return 0, nil, err
	}
	services.WriteLog(
		"[skn-report]",
		"["+dataSKN.ReferenceId+"]"+dst.String(),
		inits.Cfg.LogReportPath+services.LogFileName, "report")
	return time.Since(start).Milliseconds(), dataResponse, nil
}
