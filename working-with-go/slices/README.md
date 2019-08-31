# 切片

Go有数组类型，但是它限定了大小。Go中的切片更可扩展，是动态大小的。切片是常用的类型。

切片的格式为：`[]T`，T是类型，比如，`[]string`表示字符串集合，再如，`[]int`是整型集合。

## 初始化

初始化一个空的切片。

```go
var empty []int
```

初始化带值的切片。

```go
alphas := []string{"abc", "def", "ghi", "jkl"}
```

## 添加

切片不能修改，只能赋值一个新的值。

```go
var nums []int
nums = append(nums, 203)
nums = append(nums, 302)
fmt.Println(nums)
// [203, 302]
```

一次性增加多个值。

```go
alphas := []string{"abc", "def", "ghi", "jkl"}
alphas = append(alphas, "mno", "pqr", "stu")
```

## 普通数组操作

得到切片的长度。

```go
fmt.Println("Length: ", len(alphas))
```

取切片中的单个值。

```go
fmt.Println(alphas[1])
```

取切片的多个值（也组成一个切片）。

```go
alpha2 := alphas[1:3]
fmt.Println(alpha2)
```

## 查看切片中是否存在某元素

Go中没有内建的函数来查看切片中是否存在某元素。下面是一个实现的例子。

```go
package main

import "fmt"

func main() {
  alphas := []string{"abc", "def", "ghi", "jkl"}

  if elemExists("def", alphas) {
    fmt.Println("Exists!")
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
```
