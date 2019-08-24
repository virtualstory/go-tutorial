# 条件控制结构

在Go中使用If, Else, Switch。

## If, Else

Go条件控制不需要圆括号。

```go
if num > 3 {
    fmt.Println("Many")
}
```

`else`必需和结尾处的`if`大括号在同一行。

```go
if num == 1 {
    fmt.Println("One")
} else if num == 2 {
    fmt.Println("Two")
} else {
    fmt.Println("Many")
}
```

## Switch语句

在Go中使用`switch`和`case`条件控制。注意：Go在每个`case`中自动跳出，避免忘了`break`导致的错误。

```go
switch num {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
    default:
        fmt.Println("Many")
}
```

下面是一个可选的`switch`语法实现同样的功能：

```go
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
```
