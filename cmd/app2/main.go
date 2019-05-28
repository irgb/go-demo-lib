package main

import (
    "fmt"
    "github.com/irgb/go-demo-lib/internal/dump"
)
func main() {
    fmt.Println("This is app2")
    dump.Dump("hello", "world")
}
