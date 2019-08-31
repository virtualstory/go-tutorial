# 文件操作

## 读文件

用`ioutil.ReadFile`读取整个文件。函数返回两个变量，第一个为文件内容，第二个为出现错误时的错误，如果没有错误，则`err`为`nil`。

```go
filename := "./extras/rabbits.txt"

content, err := ioutil.ReadFile(filename)
if err != nil {
  log.Fatalln("Error reading file", filename)
}
```

返回的文件内容不是字符串，而是`[]byte`。如果需要显示字符串，你可以转化它。使用`string()`转化一个`[]byte`成`string`类型。

```go
fmt.Println(string(content))
```

## 文件存在与否

上面的错误可能是文件不存在。使用`os.Stat`来查看文件是否存在，而不需要试着打开文件。

```go
if _, err := os.Stat("junk.foo"); os.IsNotExist(err) {
  fmt.Println(">>>")
  fmt.Println("File: junk.foo does not exist")
}
```

## 写入成一个新文件

使用`ioutil.WriteFile`新建文件。函数有三个参数，文件名，文件内容(`[]byte`类型) 和文件属性。

```go
outfile := "output.txt"
err = ioutil.WriteFile(outfile, content, 0644)
if err != nil {
  fmt.Prinln("Error writing file: ", err)
} else {
  fmt.Println(">>>")
  fmt.Println("Created: ", outfile)
}
```

## 写入一个存在的文件

写入一个存在的文件，例子如下。

```go
af, err := os.OpenFile(outfile, os.O_APPEND|os.O_WRONLY, 0644)
if err != nil {
  fmt.Println("Error appending to file: ", err)
}
defer af.Close()
if _, err = af.WriteString("Appending this text"); err != nil {
  fmt.Println("Error writing to file: ", err)
}
```

## 使用Filepath

使用`filepath`包来处理可能存在的跨平台的路径。
例如，使用`filepath.Join`来创建一个文件路径。

```go
package main

import (
  "fmt",
  "path/filepath"
)

func main() {
  fmt.Println(filepath.Join("a", "b", "file.ext"))
}
```
