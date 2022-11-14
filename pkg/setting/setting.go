package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

var Cfg *ini.File

type Schedule struct {
	interval int    `json:"interval,omitempty"` //executive interval
	worker   string `json:"worker,omitempty"`   //path to worker
	argument string `json:"argument,omitempty"` //path to source configure file, one file may contains many sources
}

var ScheduleChan chan Schedule

func InitSetting(scheduleDir string) *ini.File {
	var err error

	Cfg, err = ini.Load(scheduleDir)
	if err != nil {
		log.Fatal("Fail to Load ‘conf/app.ini’:", err)
	}

	sections := Cfg.SectionStrings()
	//fmt.Println(server)
	for _, s := range sections {
		//get section name
		//fmt.Println(s)
		schedule1 := Schedule{}

		schedule1.worker = Cfg.Section(s).Key("worker").String()
		schedule1.argument = Cfg.Section(s).Key("argument").String()
		fmt.Println(schedule1, s)
		go func() { ScheduleChan <- schedule1 }()
	}
	//close(ScheduleChan)
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
