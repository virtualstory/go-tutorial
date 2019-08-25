# 字符串

`strings`标准库中包含一些你能操作字符串的函数。

创建并赋值给一个变量在一行中完成用`:=`操作符。Go将在赋值的时候自动判断类型。

```go
str := "This is an example string"
```

通常，Go没有用标准的面向对象标示去操作值，跟javascript和python不同。相反，在Go中，你通常会将正在处理的变量传递给函数。
下面是从`strings`包中使用`Contains`函数的一个例子：

```go
exists := strings.Contains(str, "example")
```

下面是一个完整的例子，用`ToLower`函数把一个字符串转成小写，用`Contains`去检测一个字符串中是否包含其他字符串。

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // 创建一个字符串变量
    str := "HI, I'M UPPER CASE"

    // 转化成小写
    lower := strings.ToLower(str)

    // 输出查看是否变成小写了
    fmt.Println(lower)

    // 检测字符串中是否包含其他字符串
    if strings.Contains(lower, "case") {
        fmt.Println("Yes, exists!")
    }
}
```

## 字符串当作字符数组

字符串在Go中可以如下面一样操作。

```go
str := "abcdefghijklmnopqrstuvwxyz"
fmt.Println("Char 3-10: " + str[3:10])

// 打印开始的5个字符
fmt.Println("First Five: " + str[:5])

// 打印从13到结尾的字符
fmt.Println("From 13 on: " + str[13:])
```

Go Playground是一个有用的在线环境，可以在浏览器中运行Go代码。 请参阅Go Playground中的上述示例：
[https://play.golang.org/p/DC7R7XKzZ5G](https://play.golang.org/p/DC7R7XKzZ5G)。

## 拆字符串

根据特定字符或单词拆解字符串

```go
sentence := "I'm a sentence made up of words"
words := strings.Split(sentence, " ")
fmt.Printf("%v \n", words)
```

如果你要以空白作为分割，使用`Field`是更好的方式，因为它将以所有空白字符分割字符串，不仅仅是空格。

```go
sentence := "I'm a sentence made up of words"
fields := strings.Fields(sentence)
fmt.Printf("%v \n", fields)
```

## 加入一个字符串数组

在OOP世界中，Join函数将在数组包中定义。 但是，在Go中，它是strings包的一部分，您传入一个字符串数组和要加入它们的分隔符。

```go
lines := []string{ "one", "two", "three" }
str := string.Join(lines, ",")
fmt.Println(str)
```

## 字符替换

字符替换使用`strings.Replace`

```go
str := "The blue whale loves blue fish."
newstr := strings.Replace(str, "blue", "red", 1)
```

`strings.Replace`函数需要传递需要替换多少个，传递`-1`将替换所有匹配上的。

```go
str := "The blue whale loves blue fish."
newstr := strings.Replace(str, "blue", "red", -1)
```

Go1.12版本引入了`strings.ReplaceAll`函数，没有第四个参数-1，正如名字所描述的那样，将替换所有匹配上的。

## 字符串的前缀和后缀

Go函数检测一个字符串开始或结尾是否以某字符匹配。`strings.HasPrefix()`检测前缀，`strings.HasSuffix()`检测后缀。

```go
path := "/home/mkaz"
isAbsolute := strings.HasPrefix(path, "/")

dir := "/home/mkaz/"
hasTrailingSlash := strings.HasSuffix(path, "/")
```

