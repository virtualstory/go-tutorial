package main

import (
    "fmt"
    "math"
)

func main() {
    Echo("go")
    fmt.Println(Say("go"))
    fmt.Println(Say2("go2"))
    Divide(4.2, 2.1)
    sum := Sum(1, 3, 5, 7)
    fmt.Println(sum)
    nums := []int{1, 3, 5, 7}
    sum2 := Sum(nums...)
    fmt.Println(sum2)
}

// basic function
func Echo(s string) {
    fmt.Println(s)
}

// function with return value
func Say(s string) string {
    phrase := "Hello " + s
    return phrase
}

// named returned value
func Say2(s string) (phrase string) {
    phrase = "Hello " + s
    return
}

// multiple parameters
func Divide(x, y float64) (float64, float64) {
    q := math.Trunc(x / y)
    r := math.Mod(x, y)
    return q, r
}
func Divide2(x, y float64) (q, r float64) {
    q = math.Trunc(x / y)
    r = math.Mod(x, y)
    return
}

// variadic parameters
func Sum(x ...int) int {
    sum := 0
    for _, v := range x {
        sum += v
    }
    return sum
}

