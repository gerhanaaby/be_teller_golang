/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-02-23
 * @modify date 2023-02-23
 * @desc [description]
 */
package services

import (
	"errors"
	"strconv"
	"time"
)

func GenTransactID(prefix, sufix string) (string, error) {
	if prefix == "" || sufix == "" {
		return "", errors.New("cant generate transact ID")
	}

	// now := time.Now()
	// d := strconv.Itoa(now.Day())
	// m := now.Month()
	// y := strconv.Itoa(now.Year())
	// h := strconv.Itoa(now.Hour())
	// i := strconv.Itoa(now.Minute())
	// s := strconv.Itoa(now.Second())

	return prefix+strconv.Itoa(int(time.Now().UnixMilli()))+sufix, nil
}