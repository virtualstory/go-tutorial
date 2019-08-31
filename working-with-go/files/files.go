package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
)

func main() {
  filename := "./files.go"

  content, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println("Error reading file: ", err)
  } else {
    fmt.Println(string(content))
  }

  if _, err := os.Stat("junk.foo"); os.IsNotExist(err) {
  	fmt.Println(">>>")
  	fmt.Println("File: junk.foo does not exist")
  }

  outfile := "output.txt"
  err := ioutil.WriteFile(outfile, nil, 0644)
  if err != nil {
  	fmt.Println("Error writing file: ", err)
  } else {
  	fmt.Println(">>>")
  	fmt.Println("Created: ", outfile)
  }


  af, err := os.OpenFile(outfile, os.O_APPEND|os.O_WRONLY, 0644)
  if err != nil {
  	fmt.Println("Error appending to file:", err)
  }
  defer af.Close()
  if _, err = af.WriteString("Appending this text"); err != nil {
  	fmt.Println("Error writing to file:", err)
  }

  fmt.Println(filepath.Join("a", "b", "file.ext"))
}
