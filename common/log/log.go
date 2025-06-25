package log

import (
	"fmt"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func Info(msg ...any) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	log.WithFields(log.Fields{
		"meta": fmt.Sprintf("%s:%d", srcFile, numLine),
	}).Info(msg...)
}

func Warning(msg ...any) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	log.WithFields(log.Fields{
		"meta": fmt.Sprintf("%s:%d", srcFile, numLine),
	}).Warning(msg...)
}

func Error(err ...any) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	log.WithFields(log.Fields{
		"meta": fmt.Sprintf("%s:%d", srcFile, numLine),
	}).Error(err...)
}

func Debug(value ...any) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	log.WithFields(log.Fields{
		"meta": fmt.Sprintf("%s:%d", srcFile, numLine),
	}).Debug(value...)
}

func Fatal(value ...any) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	log.WithFields(log.Fields{
		"meta": fmt.Sprintf("%s:%d", srcFile, numLine),
	}).Fatal(value...)
}

func Println(value ...any) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	log.WithFields(log.Fields{
		"meta": fmt.Sprintf("%s:%d", srcFile, numLine),
	}).Println(value...)
}

func Infof(format string, msg ...any) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	log.WithFields(log.Fields{
		"meta": fmt.Sprintf("%s:%d", srcFile, numLine),
	}).Infof(format, msg...)
}

func Warningf(format string, msg ...any) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	log.WithFields(log.Fields{
		"meta": fmt.Sprintf("%s:%d", srcFile, numLine),
	}).Warningf(format, msg...)
}

func Errorf(format string, err ...any) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	log.WithFields(log.Fields{
		"meta": fmt.Sprintf("%s:%d", srcFile, numLine),
	}).Errorf(format, err...)
}

func Debugf(format string, value ...any) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	log.WithFields(log.Fields{
		"meta": fmt.Sprintf("%s:%d", srcFile, numLine),
	}).Debugf(format, value...)
}

func Fatalf(format string, value ...any) {
	_, path, numLine, _ := runtime.Caller(1)
	srcFile := filepath.Base(path)
	log.WithFields(log.Fields{
		"meta": fmt.Sprintf("%s:%d", srcFile, numLine),
	}).Fatalf(format, value...)
}
