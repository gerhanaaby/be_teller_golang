package services

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"teller/models"
)

/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-02-16
 * @modify date 2023-02-20
 * @desc [Fungsi consume api]
 */
func ConsumeAPIService(name string, model []byte) ([]byte, error) {

	req, err := http.NewRequest(http.MethodPost, models.ApiMap[name].Url, bytes.NewBuffer(model))
	if err != nil {
		return nil, err
	}

	req.Header.Set(`Content-Type`, `application/json`)
	req.Header.Set(models.ApiMap[name].Key, models.ApiMap[name].Value)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, _ := ioutil.ReadAll(res.Body)

	return respBody, nil
}
