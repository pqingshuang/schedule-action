package main

import (
	"DataTransformation/pkg/setting"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

//type Schedule struct {
//	Interval int    //executive interval
//	Worker   string //path to worker
//	Argument string //path to source configure file, one file may contains many sources
//}

func main() {
	//setting.InitSetting()，every minute
	var wg sync.WaitGroup
	wg.Add(1)

	ticker := time.NewTicker(time.Second)
	quit := make(chan struct{})

	go func() {
		defer func() {
			wg.Done()
			ticker.Stop()
		}()
		for {
			select {
			case <-ticker.C:
				// do stuff
				setting.Setting("config/schedule/schedule.init")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	//goroutine doing worker
	for s := range setting.ScheduleChan {
		wg.Add(1)
		go func(s setting.Schedule) {
			defer wg.Done()
			workerPath := s.Worker
			//source_path := s.Argument
			//fmt.Println('a')
			//worker.A()
			_, err := exec.Command(workerPath).CombinedOutput()

			if err != nil {
				//log.Fatal(err)
				fmt.Println(err)

			}

		}(s)
	}
	//closer

	wg.Wait()
	//close(setting.ScheduleChan)

	return

}

//func main() {

//}
