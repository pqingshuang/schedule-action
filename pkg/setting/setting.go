package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var Cfg *ini.File
var ScheduleChan = make(chan Schedule, 10)

type Schedule struct {
	interval int    `json:"interval,omitempty"` //executive interval
	worker   string `json:"worker,omitempty"`   //path to worker
	nextTime time.Time
	argument string `json:"argument,omitempty"` //path to source configure file, one file may contains many sources
}

func (s *Schedule) Worker() string {
	return s.worker
}

func (s *Schedule) SetWorker(worker string) {
	s.worker = worker
}

func InitSetting(scheduleDir string) {

	go func() {
		c := time.Tick(1 * time.Minute)
		for range c {
			setting(scheduleDir)
		}
	}()

	//fmt.Println(Map(server,))d
	time.Sleep(5 * time.Second)

}

func setting(scheduleDir string) {
	var err error

	Cfg, err = ini.Load(scheduleDir)
	if err != nil {
		log.Fatal("Fail to Load ‘conf/app.ini’:", err)
	}

	sections := Cfg.SectionStrings()
	//fmt.Println(server)
	for _, s := range sections {

		//filter current program that need to run
		schedule1 := Schedule{}

		schedule1.worker = Cfg.Section(s).Key("worker").String()
		schedule1.argument = Cfg.Section(s).Key("argument").String()
		fmt.Println(schedule1, s)
		ScheduleChan <- schedule1
		//ScheduleChan <- schedule1
	}

	//return Cfg
}

func str2time(dateString string) {
	// Parse the date string into Go's time object
	// The 1st param specifies the format,
	// 2nd is our date string
	myDate, err := time.Parse("2006-01-02 15:04", dateString)
	if err != nil {
		panic(err)
	}
	fmt.Println("My Date Reformatted:\t", myDate)

}
