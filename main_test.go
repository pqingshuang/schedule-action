package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func Test_test(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd1, _ := exec.Command("pwd").Output()
			cmd, err := exec.Command("python3", "worker/always/config_influxdb/main.py", "--host", "192.168.100.214", "--port", "8086", "--database", "ArupDemo", "--measurement", "OTP_RealTime").Output()
			a := "worker/always/config_influxdb/main.py --host 192.168.100.214 --port 8086 --database ArupDemo --measurement OTP_RealTime"
			func() {
				fmt.Println(string(cmd))
				fmt.Println(err)
				fmt.Printf(string(cmd1))
				fmt.Println(strings.Split(a, " ")[1])
			}()
		})
	}
}
