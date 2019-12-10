package main

import "fmt"

func findDuplicate(a []int) {
  for i:=0; i<len(a); i++ {
    for j:=i+1; j<len(a); j++ {
     if a[i]==a[j] {
       fmt.Println(a[i])
     }
    }
  }
}
func main(){
 i:=[]int{1,2,3,4,4,5}
 findDuplicate(i)
}
