package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {

  reader := bufio.NewReader(os.Stdin)

  for {
    fmt.Print("-> ")
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
    fmt.Print(text)
  }

}
