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

func PostInquiryTransfer(c *gin.Context) {
	startTime := time.Now()

	if RefID, ReqTime, Response, err := TransactInquiryTransfer(c); err != nil {
		services.WriteLog(
			"[fail][inquiry-transfer]",
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
			"[done][inquiry-transfer]",
			fmt.Sprintf("REF-ID: %s, Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
			RefID,
			time.Since(startTime).Milliseconds()-ReqTime,
			ReqTime,
			time.Since(startTime).Milliseconds()),
		inits.Cfg.LogPerformancePath+services.LogFileName, "performance")
		c.JSON(http.StatusCreated, Response)
	}
}

func TransactInquiryTransfer(c *gin.Context) (RefID string, reqApiTime int64, dataResponse map[string]interface{}, err error) {
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
		return `NOID`, 0, nil, err
	}

	//--> API REQUEST PROCESS
	request := models.InquiryTransfer{}
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

	reqApiTime, dataResponse, err = PostToAPIInquiry(request)
	if err != nil {
		return request.ReferenceId, 0, nil, err
	}

	return request.ReferenceId, reqApiTime, dataResponse, nil

}

func PostToAPIInquiry(dataInquiry models.InquiryTransfer) (int64, map[string]interface{}, error) {
	start := time.Now()
	data := map[string]interface{}{
		"accountNo":   dataInquiry.AccountNo,
		"referenceId": dataInquiry.ReferenceId,
	}

	requestJson, err := json.Marshal(data)
	if err != nil {
		return 0, nil, err
	}

	body, err := services.ConsumeAPIService("inquirytransfer", requestJson)
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
		"[inquirytransfer-report]",
		"["+dataInquiry.ReferenceId+"]"+dst.String(),
		inits.Cfg.LogReportPath+services.LogFileName, "report")
	return time.Since(start).Milliseconds(), dataResponse, nil
}
