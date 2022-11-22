package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
	"testing"
	"time"
)

type MySection struct {
	MeyKey time.Duration `json:"mykey" ini:"mykey"`
}

func TestInitialSetting(t *testing.T) {
	Setting("../../config/schedule/schedule.init")
}
func TestMaptotest(t *testing.T) {
	cfg, _ := ini.Load([]byte(`
[mysection]
mykey=90
`))
	mySec := &MySection{}
	err := cfg.Section("mysection").MapTo(mySec)
	fmt.Println(err, mySec)

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
