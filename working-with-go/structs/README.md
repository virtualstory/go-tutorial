# 结构体

结构体在Go中是一种集合数据类型。结构体类似于对象。

## 定义结构体

```go
// 定义Dog结构
type Dog struct {
    Name string
    Color string
}
```

## 创建实例

创建结构体的一个实例

```go
Spot := Dog{Name: "Spot", Color: "brown"}
```

一个可选的方式是：创建一个空的实例，然后通过点符号来获取或者设置值。

```go
var Rover Dog
Rover.Name = "Rover"
Rover.Color = "light tan with dark patches"
```

