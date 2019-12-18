package main

import "fmt"

func swap(px, py *int){
  temp := *px
  *px = *py
  *py = temp
}

func main(){
  a:=1
  b:=2

  fmt.Println("Before swapping: ",a,b)
  swap(&a,&b)
  fmt.Println("After swapping: ",a,b)
}


