# Maps

map是关联数组数据类型。 您可以定义使用键或值的任何数据类型的map。

## 创建Map

Map使用内置的`make`函数初始化，创建一个空的没有键的map。

```go
m := make(map[string]string)
m["c"] = "Cyan"
m["y"] = "Yellow"
m["m"] = "Magenta"
m["k"] = "Black"
```

你也可以在新建的时候初始化它。

```go
var m = map[string]string {
    "c": "Cyan",
    "y": "Yellow",
    "m": "Magenta",
    "k": "Black"
}
```

## Map的迭代

使用`range`来遍历map，它返回`key`, `value`键值对。

```go
for k, v := range m {
    fmt.Printf("Key: %s, Value: %s", k, v)
}
```

## 取值

使用中括号取值。

```go
c = m["c"]
```

## 删除某项

使用`delete`内建函数，传递map和需要删除的key。

```go
delete(m, "c")
```

## 查看某项是否存在

通过检查提取项目时返回的第二个值来测试map是否包含项目。 如果该项目不存在，则为false。

如果该项不存在，则返回的值将为零值，但该项可能在映射中作为零值存在，因此最好检查返回的第二个值。

```go
c, ok = m["P"]
```

