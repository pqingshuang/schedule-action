package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"time"
)

var Cfg *ini.File

type schedule struct {
	Interval int    //executive interval
	Worker   string //path to worker
	Argument string //path to source configure file, one file may contains many sources
}

func InitSetting() *ini.File {
	var err error
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	Cfg, err = ini.Load("../../config/schedule/schedule.init")
	if err != nil {
		log.Fatal("Fail to Load ‘conf/app.ini’:", err)
	}
	timeout := time.After(5 * time.Second)
	pollInt := time.Second

	for {
		select {

		case <-timeout:
			fmt.Println("There's no more time to this. Exiting!")
			time.Sleep(pollInt)
		default:
			fmt.Println("still waiting")
		}

	}

	return Cfg
	////直接读取
	//RunMode := Cfg.Section("").Key("RUN_MODE").MustString("debug")
	//
	////读取内部配置
	//server, err := Cfg.GetSection("server")
	//if err != nil {
	//	log.Fatal("Fail to load section 'server': ", err)
	//}
	//HttpPort = server.Key("HTTP_PORT").MustUint(8080)
	//ReadTimeout = time.Duration(server.Key("READ_TIMEOUT").MustUint(60)) * time.Second
	//WriteTimeout = time.Duration(server.Key("WRITE_TIMEOUT").MustUint(60)) * time.Second
}
