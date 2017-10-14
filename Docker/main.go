package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("hello...")
    fmt.Println("MYVAR:", os.Getenv("MYVAR"))
    fmt.Println("...good bye")
}

