package main

import (
	"fmt"
	"log"
	"os/exec"
)

func test() {
	out, err := exec.Command("worker/always/main").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(out))

}
