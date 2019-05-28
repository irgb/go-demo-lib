package util

import (
    "go.uber.org/zap"
    "github.com/davecgh/go-spew/spew"
)

var Logger, _ = zap.NewDevelopment()

func Dump(a ...interface {}) {
    Logger.Info("This is info")
    Logger.Warn("This is warn")
    spew.Dump(a...)
}
