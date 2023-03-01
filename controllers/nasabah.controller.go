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

	"github.com/gin-gonic/gin"
)

func GetNasabahByCIF(c *gin.Context){


	cif := c.Param("cif")


	if len(cif) < 1 {
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status: "Fail", 
			Message: "error, empty CIF",
		})
		return
	}

	var newResult map[string]interface{}

	err := db.GetDB().Table("nasabahs").Select(
	`"CIF"`,
	`"mnemonics"`,
	`"kseisid"`,	
	`"Nama"`,
	`"birthPlace"`,
	`"TanggalLahir"`,
	`"documentFlag"`,
	`"gender"`,
	`"titleBeforeName"`,
	`"titleAfterName"`,
	`"religion"`,
	`"resident"`,
	`"nationality"`,
	`"Createdby"`,
	`"CreationDate"`,
	`"lastchange"`,
	`"supervisorId"`,
	`"custStatus"`,
	).Where(`"CIF" = ? `, cif).Find(&newResult).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status: "Fail", 
			Message: "error, "+err.Error(),
		})
		return
	}

	if len(newResult) == 0 {
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status: "Fail", 
			Message: "invalid, nasabah not found",
		})
		return
	}

	c.JSON(http.StatusOK, newResult)
	return
}