package main

import (
	"DataTransformation/pkg/setting"
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"
	"github.com/robfig/cron"

)

type Schedule struct {
	Interval int    //executive interval
	Worker   string //path to worker
	Argument string //path to source configure file, one file may contains many sources
}





func main() {
	//setting.InitSetting()，every minute
	c := cron.New()
	c.AddFunc("1 * * * * *", setting.InitSetting())
	c.Start()

	//check schedule that need to be done based on interval

	//send it to channel, below use slice to test first

	//get all behavior need to be done
	schedules := make([]Schedule, 1)
	schedules[0] = Schedule{0, "worker/always/main", ""}
	var wg sync.WaitGroup
	//goroutine doing worker
	for _, s := range schedules {
		wg.Add(1)
		go func(s Schedule) {
			defer wg.Done()
			workerPath := s.Worker
			//source_path := s.Argument
			out, err := exec.Command(workerPath).CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf(string(out))

		}(s)
	}
	//closer
	wg.Wait()

	for {
		select {

		case <-timeout:
			fmt.Println("There's no more time to this. Exiting!")
			time.Sleep(pollInt)
		default:
			fmt.Println("still waiting")
		}

	}

}
}

//func main() {

//}
