package controllers

import (
	"net/http"
	"teller/db"
	"teller/models"

	"github.com/gin-gonic/gin"
)

func PostSkn(c *gin.Context) {
	request := models.Skn{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.GetDB().Debug().Create(&request).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Yeay Berhasil!",
	})
}
