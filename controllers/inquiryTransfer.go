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

func PostInquiryTransfer(c *gin.Context) {
	startTime := time.Now()
	var reqApiTime int64 = 0
	var isValid bool = false
	var err error

	//--> TOKEN VALIDATION REQUEST
	isValid, err = services.CheckToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		services.WriteLog(
			"[inquiry-transfer-fail]", 
			fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
				time.Since(startTime).Milliseconds()-reqApiTime, 
				reqApiTime,
				time.Since(startTime).Milliseconds()),
			inits.Cfg.LogPerformancePath+services.LogFileName,"performance")
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due" + err.Error(),
		})
		return
	}

	if !isValid {
		services.WriteLog(
			"[inquiry-transfer-fail]", 
			fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
				time.Since(startTime).Milliseconds()-reqApiTime, 
				reqApiTime,
				time.Since(startTime).Milliseconds()),
			inits.Cfg.LogPerformancePath+services.LogFileName,"performance")
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due invalid login or expired",
		})
		return
	}

	//--> API REQUEST PROCESS
	request := models.InquiryTransfer{}
	if err = c.ShouldBindJSON(&request); err != nil {
		services.WriteLog(
			"[inquiry-transfer-fail]", 
			fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
				time.Since(startTime).Milliseconds()-reqApiTime, 
				reqApiTime,
				time.Since(startTime).Milliseconds()),
			inits.Cfg.LogPerformancePath+services.LogFileName,"performance")
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due" + err.Error(),
		})
		return
	}

	err = db.GetDB().Debug().Create(&request).Error
	if err != nil {
		services.WriteLog(
			"[inquiry-transfer-fail]", 
			fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
				time.Since(startTime).Milliseconds()-reqApiTime, 
				reqApiTime,
				time.Since(startTime).Milliseconds()),
			inits.Cfg.LogPerformancePath+services.LogFileName,"performance")
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due" + err.Error(),
		})
		return
	}

	reqApiTime, dataResponse, err := PostToAPIInquiry(request)
	if err != nil {
		services.WriteLog(
			"[inquiry-transfer-fail]", 
			fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
				time.Since(startTime).Milliseconds()-reqApiTime, 
				reqApiTime,
				time.Since(startTime).Milliseconds()),
			inits.Cfg.LogPerformancePath+services.LogFileName,"performance")
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status:  "Fail",
			Message: "unable to precess due" + err.Error(),
		})
		return

	}
	services.WriteLog(
		"[inquiry-transfer-done]", 
		fmt.Sprintf("Go-Time: %dms, Api-Time: %dms, Total-TIme: %dms",
			time.Since(startTime).Milliseconds()-reqApiTime, 
			reqApiTime,
			time.Since(startTime).Milliseconds()),
		inits.Cfg.LogPerformancePath+services.LogFileName,"performance")
	c.JSON(http.StatusCreated, dataResponse)
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
		"[inquiry-transfer-report]", 
		dst.String(),
		inits.Cfg.LogPerformancePath+services.LogFileName,"report")
	return time.Since(start).Milliseconds(),dataResponse, nil
}
