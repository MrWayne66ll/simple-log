package simple_log

import (
	"testing"
	"time"
)

func TestMain(t *testing.M) {
	t.Run()
}

func TestLog(t *testing.T) {
	if err := InitLogger(SetFileWriter("./log/")); err != nil {
		t.Error(err)
		return
	}
	defer Sync()

	t.Log("start testing logger ......")

	Info("test info")
	Infof("test %v info at %v", "xx", time.Now().Format(time.RFC3339))

	Error("test error")
	Errorf("test %v error at %v", "xxx-err", time.Now().Format(time.RFC3339))

	//Panicf("test %v panic at %v", "xx-panic", time.Now().Format(time.RFC3339))

	Fatalf("test %v fatal at %v", "xx-fatal", time.Now().Format(time.RFC3339))
	makeError()
}

func makeError() {
	Errorf("test internal error %v", time.Now().Format(time.RFC3339))
}
