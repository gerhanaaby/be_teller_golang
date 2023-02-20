package services

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var LogFileName string
var LogPerformace *os.File
var LogReport *os.File

func WriteLog(prefix string, info string) error {

	// filename := LogFileName

	fsinfo, err := os.Stat(LogFileName)
	if os.IsNotExist(err) {
		fmt.Println("notexist")
		// WriteLogFile(prefix, info)
		return err
	} else {
		if fsinfo.Size() >= 20000000 {
			//buat baru
			counter, err := strconv.Atoi(LogFileName[15:])
			if err != nil {
				return err
			}
			LogFileName = LogFileName[:15] + strconv.Itoa(counter+1)

			// WriteLogFile(prefix, info)
			return nil

		} else {
			//append
			// WriteLogFile(prefix, info)
			return nil

		}
	}

}

func WriteLogFile(fileInfo *os.File, prefix string, info string) {
	// log.New(fileInfo, prefix, log.LstdFlags|log.Lshortfile|log.Lmicroseconds).Println(info)
	log.New(
		fileInfo, 
		prefix, 
		log.LstdFlags|
		// log.Ldate|
		log.Lmicroseconds).Println(info)

}

func InitLogFileName(curdate time.Time) string {
	return "LOG-" + curdate.Format("02-01-2006") + "-1"
}