package setting

import (
	"time"
)

type MySection struct {
	MeyKey time.Duration `json:"mykey" ini:"mykey"`
}

//func Test_str2time(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			str2time()
//		})
//	}
//}
