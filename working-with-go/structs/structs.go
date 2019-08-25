package main

import "fmt"

func main() {
    type Dog struct {
        Name string
        Color string
    }

    Spot := Dog{Name: "Spot", Color: "brown"}
    fmt.Println(Spot)

    var Rover Dog
    Rover.Name = "Rover"
    Rover.Color = "light tan with dark patches"
    fmt.Println(Rover)
}
