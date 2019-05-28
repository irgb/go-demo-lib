package dump

import (
    "fmt"
    "go.uber.org/zap"
    "github.com/davecgh/go-spew/spew"
)

var Logger, _ = zap.NewDevelopment()

func Dump(a ...interface {}) {
    fmt.Println("This is InternalDump")
    Logger.Info("This is internal info")
    spew.Dump(a...)
}
