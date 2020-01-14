package main

import (
    "sync"
    "fmt"
)

func main(){
	result := GetInstance()
	fmt.Println(result)
}
type singleton struct {}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
