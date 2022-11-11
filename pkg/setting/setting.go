package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

var Cfg *ini.File

type schedule struct {
	Interval int    //executive interval
	Worker   string //path to worker
	Argument string //path to source configure file, one file may contains many sources
}

// func getAddress([]*ini.Section,int )
// var scheduleChan chan schedule
func InitSetting(scheduleDir string) *ini.File {
	var err error

	Cfg, err = ini.Load(scheduleDir)
	if err != nil {
		log.Fatal("Fail to Load ‘conf/app.ini’:", err)
	}

	server := Cfg.Sections()
	for _, v := range server {
		//
		fmt.Println(v.Name())
		//for b, a := range v.Keys() {
		//	fmt.Println(b, a)
		//}
	}
	//fmt.Println(Map(server,))
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
