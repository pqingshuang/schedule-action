package main

import (
	"fmt"
	"testing"
	"time"
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
			main()
			func() { fmt.Println(time.Now()) }()
		})
	}
}
