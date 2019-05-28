package main

import (
    "fmt"
    "github.com/irgb/go-demo-lib/pkg/dump"
)
func main() {
    fmt.Println("This is app1")
    dump.Dump("hello", "world")
}
