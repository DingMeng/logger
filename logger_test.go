package logger

import "testing"

func TestLog(t *testing.T){
	L.Info("info")
	L.Infof("info:%s","f")
	L.Warn("warn")
	L.Error("error")
}


