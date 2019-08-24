# hello world in Go

这是你的第一个go程序

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello world!")
}
```

所有程序都属于某个包，可以用`main`作为包的可执行程序。

你必需导入所有使用的包，包括标准库里的包。在这个例子上，标准库`fmt`是用来格式化I/O的。

在GO中，标准库是你的最好的选择，它可扩展，通常你不再需要第三方依赖包，仅需标准库上的特性功能就能满足你使用。
查看所有标准库：[http://golang.org/pkg/](http://golang.org/pkg/)

`main`包需要一个`main()`函数，当程序启动时调用。

如果存在`init`函数，`init()`函数会在`main()`函数之前调用。

下面是一个加强版的hello world程序，定义了一个全局的字符串类型的变量`phrase`。go编译器会自动检测类型。

```go
package main

import "fmt"

var phrase string

func init() {
	phrase = "Hola Mundo!"
}

func main() {
	fmt.Println(phrase)
}
```

把以上程序保存在`hello.go`文件中，然后在终端上执行`go run`启动该程序：

```bash
$ go run hello.go
Hola Mundo!
```

