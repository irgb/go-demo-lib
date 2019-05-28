package dump

import (
    "fmt"
    "go.uber.org/zap"
    "github.com/davecgh/go-spew/spew"
)

var Logger, _ = zap.NewDevelopment()

func Dump(a ...interface {}) {
    fmt.Println("This is PkgDump")
    Logger.Info("This is pkg info")
    spew.Dump(a...)
}
