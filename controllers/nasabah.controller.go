/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-03-01
 * @modify date 2023-03-01
 * @desc [description]
 */
package controllers

import (
	"net/http"
	"teller/db"
	"teller/models"

	"github.com/gin-gonic/gin"
)

func GetNasabahByCIF(c *gin.Context){

	result := models.Nasabah{}

	cif := c.Param("cif")
	if cif == "" {
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status: "Fail", 
			Message: "error, empty CIF",
		})
		return
	}

	err := db.GetDB().Where(`"CIF" = ? `, cif).Find(&result).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status: "Fail", 
			Message: "error, "+err.Error(),
		})
		return
	}

	if result.CIF == `` {
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status: "Fail", 
			Message: "invalid, nasabah not found",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}