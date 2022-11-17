package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var ScheduleChan = make(chan Schedule, 10)

type Schedule struct {
	interval time.Duration `json:"interval,omitempty"` //executive interval
	worker   string        `json:"worker,omitempty"`   //path to worker
	nextTime time.Time
	argument string `json:"argument,omitempty"` //path to source configure file, one file may contains many sources
}
type section ini.Section

func newSchedule 
func (s *Schedule) Interval() time.Duration {
	return s.interval
}

func (s *Schedule) SetInterval(intervalNumber int) {
	var interval time.Duration
	interval = time.Duration(intervalNumber) * time.Minute
	s.interval = interval
}

func (s *Schedule) Worker() string {
	return s.worker
}

func (s *Schedule) SetWorker(worker string) {
	s.worker = worker
}

func Setting(scheduleDir string) {
	var err error
	var Cfg *ini.File
	Cfg, err = ini.Load(scheduleDir)
	if err != nil {
		log.Fatal("Fail to Load ‘conf/app.ini’:", err)
	}

	sections := Cfg.SectionStrings()
	//fmt.Println(server)
	for _, s := range sections {
		//TODO skip default section
		//filter current program that need to run
		schedule1 := Schedule{}

		schedule1.worker = Cfg.Section(s).Key("worker").String()
		schedule1.argument = Cfg.Section(s).Key("argument").String()
		idx, _ := Cfg.Section(s).Key("interval").Int()
		schedule1.SetInterval(idx)
		if Cfg.Section(s).HasKey("nextTime") {
			timeString := Cfg.Section(s).Key("nextTime").String()
			schedule1.nextTime = str2time(timeString)
		} else {
			schedule1.nextTime = time.Now()
			Cfg.Section(s).Key("nextTime").SetValue(schedule1.nextTime.Format("2006-01-02 15:04"))
			Cfg.SaveTo(scheduleDir)

		}
		// Only fetch worker whose next run time is before now, and change nextTime to nextTime+interval
		// and save to file
		if schedule1.nextTime.Before(time.Now()) {
			ScheduleChan <- schedule1
			schedule1.nextTime = schedule1.nextTime.Add(schedule1.interval)
			Cfg.Section(s).Key("nextTime").SetValue(schedule1.nextTime.Format("2006-01-02 15:04"))
			Cfg.SaveTo(scheduleDir)

		}

		//ScheduleChan <- schedule1
	}

	//return Cfg
}

func str2time(dateString string) time.Time {
	// Parse the date string into Go's time object
	// The 1st param specifies the format,
	// 2nd is our date string
	myDate, err := time.Parse("2006-01-02 15:04", dateString)
	if err != nil {
		panic(err)
	}
	fmt.Println("My Date Reformatted:\t", myDate)
	return myDate
}
