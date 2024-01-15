package main

import (
    "strings"
    "flag"
    "fmt"
    "io/ioutil"

    _ "github.com/lylemi/passerby"
)

func readCmdline() string {
    data, err := ioutil.ReadFile("/proc/self/cmdline")
    if err != nil {
        fmt.Printf("Error reading /proc/self/cmdline: %v\n", err)
        return ""
    }
    return strings.ReplaceAll(string(data), "\x00", " ")
}

func main() {
    fmt.Printf("/proc/self/cmdline: %s\n", readCmdline())

    var name string
    flag.StringVar(&name, "name", "", "Name to be used")
    flag.Parse()
    fmt.Println("Real Name:", name)
}
