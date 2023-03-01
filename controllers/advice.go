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

func PostAdvice(c *gin.Context) {
	startTime := time.Now()

	if ReqTime, Response, err := TransactAdvice(c); err != nil {
		services.WriteLog(
			"[fail][advice]",
			fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
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
			"[done][advice]",
			fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
				time.Since(startTime).Milliseconds()-ReqTime,
				ReqTime,
				time.Since(startTime).Milliseconds()),
			inits.Cfg.LogPerformancePath+services.LogFileName, "performance")
		c.JSON(http.StatusCreated, Response)
	}
}

func TransactAdvice(c *gin.Context) (reqApiTime int64, dataResponse map[string]interface{}, err error) {
	var isValid bool = false
	var claims jwt.MapClaims

	//--> TOKEN VALIDATION REQUEST
	claims, isValid, err = services.CheckToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		return 0, nil, err
	}

	if !isValid {
		return 0, nil, errors.New("error, authentication failure")
	}

	// Convert map to json string
	jsonStr, err := json.Marshal(claims["User"])
	if err != nil {
		return 0, nil, err
	}

	user := models.User{}
	// Convert json string to struct
	if err := json.Unmarshal(jsonStr, &user); err != nil {
		return 0, nil, err
	}

	//--> API REQUEST PROCESS
	request := models.Advice{}
	if err = c.ShouldBindJSON(&request); err != nil {
		return 0, nil, err
	}

	request.ReferenceId, err = services.GenTransactID("", user.Nik)
	if err != nil {
		return 0, nil, err
	}

	err = db.GetDB().Create(&request).Error
	if err != nil {
		return 0, nil, err
	}

	reqApiTime, dataResponse, err = PostToAPIAdvice(request)
	if err != nil {
		return 0, nil, err
	}

	return reqApiTime, dataResponse, nil

}

func PostToAPIAdvice(dataAdvice models.Advice) (int64, map[string]interface{}, error) {
	start := time.Now()
	data := map[string]interface{}{
		"referenceId":         dataAdvice.ReferenceId,
		"debitAccountNo":      dataAdvice.DebitAccountNo,
		"creditAccountNo":     dataAdvice.CreditAccountNo,
		"transactionDate":     dataAdvice.TransactionDate,
		"originalReferenceId": dataAdvice.OriginalReferenceId,
		"clientId":            dataAdvice.ClientId,
		"transactionTime":     dataAdvice.TransactionDate,
	}

	requestJson, err := json.Marshal(data)
	if err != nil {
		return 0, nil, err
	}

	body, err := services.ConsumeAPIService("advice", requestJson)
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
		"[advice-report]",
		dst.String(),
		inits.Cfg.LogReportPath+services.LogFileName, "report")
	return time.Since(start).Milliseconds(), dataResponse, nil
}
