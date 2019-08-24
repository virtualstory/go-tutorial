package main

import "fmt"

func main() {
    // initialize
    var nums []int
    alphas := []string{"abc", "def", "gih", "jkl"}

    // append
    nums = append(nums, 200)
    nums = append(nums, 300)
    fmt.Println(nums)
    alphas = append(alphas, "mno", "stu")
    fmt.Println(alphas)

    // len
    fmt.Println("length of alphas: ", len(alphas))

    // retrieve
    fmt.Println(alphas[1])
    fmt.Println(alphas[1:3])

    if elemExists("def", alphas) {
        fmt.Println("def Exists!")
    }
}

func elemExists(s string, a []string) bool {
    for _, v := range a {
        if v == s {
            return true
        }
    }
    return false
}
