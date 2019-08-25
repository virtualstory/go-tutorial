# 循环

Go使用`for`来循环，没有`while`和`do-while`循环。

## For循环

```go
alphas := []string{"abc", "def", "ghi"}

// 标准for循环
for i := 0; i < len(alphas); i++ {
    fmt.Printf("%d: %s \n", i, alphas[i])
}

// 向后循环
for i := len(alphas) - 1; i >= 0; i-- {
    fmt.Printf("%d: %s \n", i, alphas[i])
}
```

## Range

使用`range`函数实现数组遍历：

```go
for i, val := range alphas {
    fmt.Printf("%d: %s \n", i, val)
}
```

仅取索引：

```go
for i := range alphas {
    fmt.Println(i)
}
```

仅取值，不需要索引时，你需要使用`_`字符：

```go
for _, val := range alphas {
    fmt.Println(val)
}
```

## 使用for实现类while循环

```go
x := 0
for x < 10 {
    fmt.Println(x)
    x++
}
```

## 无限循环（死循环）

```go
for {
    fmt.Print(".")
}
```

