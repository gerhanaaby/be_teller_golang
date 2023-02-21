package utils

import "encoding/base64"

/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-02-14
 * @modify date 2023-02-20
 * @desc [Mengubah byte string menjadi base64]
 */
func Encode(s string) string {
	data := base64.StdEncoding.EncodeToString([]byte(s))
	return string(data)
}

/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-02-14
 * @modify date 2023-02-20
 * @desc [Mengubah base64 menjadi string]
 */
func Decode(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
