package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"teller/db"
	"teller/inits"
	"teller/routes"
	"teller/services"
	"time"
)



func init() {
	var err error

	if runtime.GOOS == `windows` {
		inits.Cfg.EnvPath = `.env`
		inits.Cfg.LogPath = `C:/Users/LENOVO/Desktop/log/`
	} else if runtime.GOOS == `linux` {
		inits.Cfg.LogPath = "/home/golang/app/teller/env"
		inits.Cfg.LogPath = "/home/golang/log/teller"
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
			services.LogWriter, err = os.OpenFile(
				inits.Cfg.LogPath+services.LogFileName,
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

	time1 := time.Now()
	time2 := time.Now()
	fmt.Println(time2.Sub(time1).Milliseconds())


	services.WriteLog("tesst", "test")
	services.WriteLogFile(services.LogWriter, "[Test]", "ini test 2")

	if err := db.ConnectDB(); err != nil{
		panic(err)
	}	

	
}


func main() {
	// Init()
	routes.Routes()	
}
