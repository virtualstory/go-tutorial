package main

import "fmt"

func main() {
    m := make(map[string]string)
    m["c"] = "Cyan"
    m["y"] = "Yellow"
    m["m"] = "Magenta"
    m["k"] = "Black"

    var n = map[string]string{
        "c": "Cyan",
        "y": "Yellow",
        "m": "Magenta",
        "k": "Black",
    }

    for k, v := range m {
        fmt.Printf("Key: %s, Value: %s\n", k, v)
    }

    fmt.Println(n["c"])

    delete(n, "c")

    fmt.Println(n["c"])

    var c, ok = n["c"]

    fmt.Println(c, ok)
}
