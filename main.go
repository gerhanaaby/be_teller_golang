package main

import (
	"errors"
	"os"
	"runtime"
	"teller/db"
	"teller/inits"
	"teller/services"
	"time"
)



func init() {
	var err error

	if runtime.GOOS == `windows` {
		inits.Cfg.EnvPath = `.env`
		inits.Cfg.LogPerformancePath = `C:/Users/LENOVO/Desktop/log/performance`
		inits.Cfg.LogReportPath = `C:/Users/LENOVO/Desktop/log/report`
	} else if runtime.GOOS == `linux` {
		inits.Cfg.EnvPath = "/home/golang/app/teller/env"
		inits.Cfg.LogPerformancePath = "/home/golang/log/teller/performance"
		inits.Cfg.LogReportPath = "/home/golang/log/teller/report"
	} else {
		panic(errors.New("error, os not widnows or linux"))
	}

	if err := inits.LoadConfig(inits.Cfg.EnvPath); err != nil{
		panic(err)
	}

	var isLoggerReady bool = false
	go func ()  {
		for {
			CurDate := time.Now()
			Hours := CurDate.Hour()
			Minutes := CurDate.Minute()
			services.LogFileName = services.InitLogFileName(CurDate)
			services.LogPerformace, err = os.OpenFile(
				inits.Cfg.LogPerformancePath+services.LogFileName,
				os.O_WRONLY|os.O_APPEND|os.O_CREATE,
				0777)
			if err != nil {
				panic(err)
			}
			services.LogReport, err = os.OpenFile(
				inits.Cfg.LogReportPath+services.LogFileName,
				os.O_WRONLY|os.O_APPEND|os.O_CREATE,
				0777)
			if err != nil {
				panic(err)
			}

			isLoggerReady = true
			time.Sleep(time.Duration(86400-(Hours+Minutes+time.Now().Second())) * time.Second)
		}
	}()
	for {
		if isLoggerReady{
			break
		}
	}

	if err := db.ConnectDB(); err != nil{
		panic(err)
	}	

	
}


func main() {
	// Init()
	// routes.Routes()	
}
