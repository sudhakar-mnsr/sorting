package main

import (
    "syscall"
    "fmt"
)

func main() {
    disk := "/dev/sda"
    var fd, numread int
    var err error

    fd, err = syscall.Open(disk, syscall.O_RDONLY, 0777)

    if err != nil {
        fmt.Print(err.Error(), "\n")
        return
    }

    buffer := make([]byte, 10, 100)

    numread, err = syscall.Read(fd, buffer)
