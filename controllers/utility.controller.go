/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-02-28
 * @modify date 2023-02-28
 * @desc [description]
 */
package controllers

import (
	"net/http"
	"os"
	"teller/inits"

	"github.com/gin-gonic/gin"
)


type Base64Data struct {
	CIf   string `json:"cif" binding:"required"`
	Base64Data string `json:"base64data"`
}

func GetBase64ByCif(c *gin.Context) {

	// cif2 := c.Request.URL.Query().Get("cif")
	// cif, exist := c.GetQuery("cif")

	cif := c.Param("cif")
	if len(cif) < 1 {
		c.JSON(http.StatusBadRequest, AuthStatus{
			Status: "Fail", 
			Message: "error, empty CIF",
		})
	} else {
		if content, err := os.ReadFile(inits.Cfg.Base64Path + cif); err != nil {
			c.JSON(http.StatusBadRequest, AuthStatus{
				Status: "Fail", 
				Message: "error, "+err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, AuthStatus{
				Status: "Success", 
				Message: string(content),
		})
		}	
	}
}