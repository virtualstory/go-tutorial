# 函数

一些展示如何在Go中使用函数的例子。

## 基本函数

一个参数和没有返回值的例子。在函数定义中参数类型必需指定。

```go
func Echo(s string) {
  fmt.Println(s)
}
```

## 函数带返回值

定义一个带返回值的函数，必需指定返回值类型。

```go
func Say(s string) string {
  phrase := "Hello " + s
  return phrase
}
```

## 命名返回值

可以命名一个函数的返回值。使用变量名初始化它。有命名返回值的函数在函数体中不需要再指定返回，它会自动返回当前命名的返回值。

```go
func Say(s string) (phrase string) {
  phrase = "Hello " + s
  return
}
```

## 多个参数

多个参数的多个返回值的函数。

```go
func Divide(x, y float64) (float64, float64) {
  q := math.Trunc(x / y)
  r := math.Mod(x, y)
  return q, r
}
```

函数带多个参数或返回值时。如果类型都相同，可以在末尾一次性指定。

```go
func Divide2(x, y float64) (q, r float64) {
  q = math.Trunc(x / y)
  r = math.Mod(x, y)
  return
}
```

## 可变参数

可变参数函数可以带入不定个数的参数，参数传递进去组成一个切片。下面是一个例子。

```go
func Sum(x ...int) int {
  sum := 0
  for _, v := range x {
    sum += v
  }
  return sum
}
```

函数调用时你可以传递多个参数:

```go
sum := Sum(1,3,5,7)
fmt.Println(sum)
```

你也可以使用展开运算符：

```go
nums := []int{1, 2, 3, 4, 5}
sum := Sum(nums...)
fmt.Println(sum)
```
