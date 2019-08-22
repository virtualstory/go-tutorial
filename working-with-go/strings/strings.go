package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "HI, I'M UPPER CASE"
    // strings.ToLower
    lower := strings.ToLower(str)
    fmt.Println(lower)

    // strings.Contains
    if strings.Contains(lower, "case") {
        fmt.Println("Yes, exists!")
    }

    // Strings as Array of Characters
    str2 := "abcdefghijklmnopqrstuvwxyz"
    fmt.Println("Chars 3-10: " + str2[3:10])

    // strings.Split
    sentence := "I'm a sentence made up of words"
    words := strings.Split(sentence, " ")
    fmt.Printf("%v \n", words)

    // strings.Fields
    fields := strings.Fields(sentence)
    fmt.Printf("%v \n", fields)

    // strings.Join
    lines := []string{ "one", "two", "three" }
    str3 := strings.Join(lines, ",")
    fmt.Println(str3)

    // strings.Replace
    str4 := "The blue whale loves blue fish"
    newstr := strings.Replace(str4, "blue", "red", 1)
    fmt.Println(newstr)

    // strings.ReplaceAll
    // some as: strings.ReplaceAll(str, "blue", "red")
    newstr2 := strings.Replace(str, "blue", "red", -1)
    fmt.Println(newstr2)

    // strings.HasPrefix
    path := "/home/mkaz"
    isAbsolute := strings.HasPrefix(path, "/")
    fmt.Println(isAbsolute)

    // strings.HasSuffix
    dir := "/home/mkaz/"
    hasTrailingSlash := strings.HasSuffix(dir, "/")
    fmt.Println(hasTrailingSlash)
}
