package services

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"
)

var LogFileName string
var LogPerformace *os.File
var LogReport *os.File

func WriteLog(prefix,info,filepath,logType string) error {

	var fileInfo os.File

	if logType == "report" {
		fileInfo = *LogReport
	} else if logType == "performance" {
		fileInfo = *LogPerformace
	} else {
		return errors.New("err invalid log type")
	}

	fsinfo, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		WriteLogFile(&fileInfo, prefix, info)
		return err
	} else {
		if fsinfo.Size() >= 20000000 {
			//buat baru
			counter, err := strconv.Atoi(filepath[15:])
			if err != nil {
				return err
			}
			filepath = filepath[:15] + strconv.Itoa(counter+1)

			WriteLogFile(&fileInfo, prefix, info)
			return nil

		} else {
			//append
			WriteLogFile(&fileInfo, prefix, info)
			return nil

		}
	}

}

func WriteLogFile(fileInfo *os.File, prefix,info string) {
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