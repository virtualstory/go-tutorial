package main

import "fmt"

var phrase string

func init() {
    phrase = "hello world!"
}

func main() {
    fmt.Println(phrase)
}
