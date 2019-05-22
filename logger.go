package logger

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var L *logrus.Logger

func init() {
	L = logrus.New()
	L.SetLevel(logrus.InfoLevel)
	//L.ReportCaller = true
	ConfigLogger()
}

func ConfigLogger() {
	base, _ := os.Getwd()
	logPath := path.Join(base,"logs")
	_ = os.MkdirAll(logPath, os.ModePerm)
	infoPath := path.Join(logPath, "info.log")
	errorPath := path.Join(logPath, "error.log")
	infoW, err := rotatelogs.New(
		infoPath+".%Y%m%d%H%M%S",
		rotatelogs.WithLinkName(infoPath),
		rotatelogs.WithMaxAge(time.Hour*5208),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		L.Errorf("infoLogger error:%v", errors.WithStack(err))
	}
	errorW, err := rotatelogs.New(
		errorPath+".%Y%m%d%H%M%S",
		rotatelogs.WithLinkName(errorPath),
		rotatelogs.WithMaxAge(time.Hour*5208),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		L.Errorf("errorLogger error:%v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  infoW,
		logrus.WarnLevel:  infoW,
		logrus.ErrorLevel: errorW,
	}, &logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05.000"})
	L.AddHook(lfHook)
}

