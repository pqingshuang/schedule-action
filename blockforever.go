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





func main() {
	//setting.InitSetting()，every minute, to check if there is any job need to do
	c := cron.New()
	c.AddFunc("* 1 * * * *", setting.InitSetting())
	c.Start()


    // do something
	go func() {
		fmt.Println("Hi~!")
	}()

	timeout := time.After(1 * time.Second)
	pollInt := time.Second
	for {
		select {

		case <-timeout:

			time.Sleep(pollInt)
		default:
			fmt.Println("still waiting")
		}

	}

}
}
