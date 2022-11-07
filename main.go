package main

import "bytes"

func main() {
	//setting.InitSetting()，every minute
	//get all behavior need to be done
	s := make([]schedule, 3)
	//goroutine doing worker
	for _, schedule := range s {
		go func(schedule) {
			worker_path := schedule.Worker
			source_path := schedule.Argument

		}()
	}

}

//func main() {

//}
cmd := exec.Command("tr", "a-z", "A-Z")
cmd.Stdin = strings.NewReader("some input")

var out bytes.Buffer
cmd.Stdout = &out
if err := cmd.Run(); err != nil {
log.Fatal(err)

}
