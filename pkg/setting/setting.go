package setting

import (
	"gopkg.in/ini.v1"
	"log"
	"strings"
	"time"
)

var ScheduleChan = make(chan Schedule, 10)

type Schedule struct {
	Interval time.Duration `ini:"Interval"` //executive Interval
	Worker   string        `ini:"worker"`   //path to worker
	NextTime time.Time     `ini:"nextTime"`
	Argument []string      `ini:"argument"` //path to source configure file, one file may contains many sources
}

type section struct {
	*ini.Section
}
type RawSchedule interface {
	mapTo() *Schedule
}

func (a *section) mapTo() *Schedule {
	schedule1 := &Schedule{}
	//fmt.Println(Cfg.Section(s).Key("nextTime").TimeFormat("2006-01-02 15:04"))
	//a, _ := Cfg.Section(s).Key("nextTime").TimeFormat("2006-01-02 15:04")
	//
	//fmt.Println(Cfg.Section(s).MapTo(schedule1), *schedule1, a)
	schedule1.Worker = a.Key("worker").String()
	schedule1.Argument = strings.Split(a.Key("argument").String(), " ")
	//fmt.Println(a.Key("argument").String())
	//fmt.Println(strings.Split(a.Key("argument").String(), " ")[1])
	idx, _ := a.Key("Interval").Int()
	schedule1.SetInterval(idx)
	return schedule1
}
func formatSchedule(schedule1 RawSchedule) *Schedule {

	return schedule1.mapTo()
}

// func newSchedule
func (s *Schedule) GetInterval() time.Duration {
	return s.Interval
}
func (s *Schedule) SetInterval(intervalNumber int) {
	var interval time.Duration
	interval = time.Duration(intervalNumber) * time.Minute
	s.Interval = interval
}

//
//func (s *Schedule) Worker() string {
//	return s.worker
//}
//
//func (s *Schedule) SetWorker(worker string) {
//	s.worker = worker
//}

func LoadSetting(scheduleDir string) {
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
		section := &section{Cfg.Section(s)}

		schedule1 := formatSchedule(section)

		if section.HasKey("nextTime") {
			timeString := section.Key("nextTime").String()
			schedule1.NextTime = str2time(timeString)
		} else {
			schedule1.NextTime = time.Now()
			section.Key("nextTime").SetValue(schedule1.NextTime.Format(time.RFC3339))

		}
		// Only fetch worker whose next run time is before now, and change nextTime to nextTime+Interval
		// and save to file
		if schedule1.NextTime.Before(time.Now()) {
			// send to schedule channl to trigger worker
			ScheduleChan <- *schedule1
			schedule1.NextTime = time.Now().UTC().Add(schedule1.Interval)
			section.Key("nextTime").SetValue(schedule1.NextTime.Format(time.RFC3339))

		}

		//ScheduleChan <- *schedule1
	}
	Cfg.SaveTo(scheduleDir)

	//return Cfg
}

func str2time(dateString string) time.Time {
	// Parse the date string into Go's time object
	// The 1st param specifies the format,
	// 2nd is our date string
	myDate, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		panic(err)
	}
	//fmt.Println("My Date Reformatted:\t", myDate)
	return myDate
}
