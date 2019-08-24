package main

import "fmt"

func main() {
    num := 3
    if num == 1 {
        fmt.Println("One")
    } else if num == 2 {
        fmt.Println("Two")
    } else {
        fmt.Println("Many")
    }

    switch num {
        case 1:
            fmt.Println("One")
        case 2:
            fmt.Println("Two")
        default:
            fmt.Println("Many")
    }

    switch {
        case num == 1:
            fmt.Println("One")
        case num == 2:
            fmt.Println("Two")
        case num > 2:
            fmt.Println("Many")
        default:
            fmt.Println("Less than 1")
    }
}

