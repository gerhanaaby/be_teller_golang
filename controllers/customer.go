package controllers

import (
	"fmt"
	"golang/db"
	"golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCustomer(c *gin.Context) {
	request := models.Customer{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := request.Validate(); err != nil {
		errorMessage := fmt.Sprintf("%s gaboleh Null", err)
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": errorMessage,
		})
		return
	}

	err := db.GetDB().Debug().Create(&request).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)

}
