package controllers

import (
	"golang/db"
	"golang/models"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostOrder(c *gin.Context) {
	request := models.Order{}
	randim := rand.Intn(10000000000)
	request.OrderNo = randim

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.GetDB().Debug().Select("order_no", "date_order", "tracking_no", "total_price", "discount_tag", "customer_id").Where("order_no = ?", request.OrderNo).Create(&request).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// result := models.Order{}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Yeay Berhasil!",
		"order_no": request.OrderNo,
		"refID":    request.ID,
	})
}

func GetOrder(c *gin.Context) {
	result := []models.Orderjadi{}
	orderNoStr := c.Param("orderNo")
	orderNo, err := strconv.Atoi(orderNoStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.GetDB().Debug().Preload("Customers").Preload("Items").Where("order_no= ?", orderNo).Find(&result).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateOrder(c *gin.Context) {
	request := models.Order{}
	randim := rand.Intn(10000000000)
	request.OrderNo = randim

	orderNoStr := c.Param("orderNo")
	orderNo, err := strconv.Atoi(orderNoStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.GetDB().Debug().Model(&models.Order{}).Where("order_no = ?", orderNo).Updates(&request).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order_no": request.OrderNo,
	})

}
