package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"teller/db"
	"teller/inits"
	"teller/models"
	"teller/services"
	"time"

	"github.com/gin-gonic/gin"
)

func PostInternalTransfer(c *gin.Context) {
	startTime := time.Now()
	var reqApiTime int64 = 0
	var isValid bool = false
	var err error

	//--> TOKEN VALIDATION REQUEST
	_, isValid, err = services.CheckToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		services.WriteLog(
			"[internal-transfer-fail]",
			fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
				time.Since(startTime).Milliseconds()-reqApiTime,
				reqApiTime,
				time.Since(startTime).Milliseconds()),
			inits.Cfg.LogPerformancePath+services.LogFileName, "performance")
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due" + err.Error(),
		})
		return
	}

	if !isValid {
		services.WriteLog(
			"[internal-transfer-fail]",
			fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
				time.Since(startTime).Milliseconds()-reqApiTime,
				reqApiTime,
				time.Since(startTime).Milliseconds()),
			inits.Cfg.LogPerformancePath+services.LogFileName, "performance")
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due invalid login or expired",
		})
		return
	}

	//--> API REQUEST PROCESS
	request := models.InternalTransfer{}
	if err = c.ShouldBindJSON(&request); err != nil {
		services.WriteLog(
			"[internal-transfer-fail]",
			fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
				time.Since(startTime).Milliseconds()-reqApiTime,
				reqApiTime,
				time.Since(startTime).Milliseconds()),
			inits.Cfg.LogPerformancePath+services.LogFileName, "performance")
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due" + err.Error(),
		})
		return
	}

	err = db.GetDB().Create(&request).Error
	if err != nil {
		services.WriteLog(
			"[internal-transfer-fail]",
			fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
				time.Since(startTime).Milliseconds()-reqApiTime,
				reqApiTime,
				time.Since(startTime).Milliseconds()),
			inits.Cfg.LogPerformancePath+services.LogFileName, "performance")
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due" + err.Error(),
		})
		return
	}

	reqApiTime, dataResponse, err := PostToAPIInternal(request)
	if err != nil {
		services.WriteLog(
			"[internal-transfer-fail]",
			fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
				time.Since(startTime).Milliseconds()-reqApiTime,
				reqApiTime,
				time.Since(startTime).Milliseconds()),
			inits.Cfg.LogPerformancePath+services.LogFileName, "performance")
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due" + err.Error(),
		})
		return

	}

	// logTransaction := dataResponse["$errorInfo"].(map[string]interface{})["$pipeline"].(map[string]interface{})["logTransaction"]
	// inputMessage := dataResponse["$errorInfo"].(map[string]interface{})["$pipeline"].(map[string]interface{})["inputMessage"]

	services.WriteLog(
		"[internal-transfer-done]",
		fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
			time.Since(startTime).Milliseconds()-reqApiTime,
			reqApiTime,
			time.Since(startTime).Milliseconds()),
		inits.Cfg.LogPerformancePath+services.LogFileName, "performance")

	// c.JSON(http.StatusCreated, gin.H{
	// 	"logTransaction": logTransaction,
	// 	"inputMessage":   inputMessage,
	// })
	c.JSON(http.StatusCreated, dataResponse)
}

func PostToAPIInternal(dataInternal models.InternalTransfer) (int64, map[string]interface{}, error) {
	start := time.Now()
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
		return 0, nil, err
	}

	body, err := services.ConsumeAPIService("internaltransfer", requestJson)
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
		"[internal-transfer-report]",
		dst.String(),
		inits.Cfg.LogReportPath+services.LogFileName, "report")
	return time.Since(start).Milliseconds(), dataResponse, nil
}
