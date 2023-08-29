package main

import (
	"fmt"
	"reflect"
)

func reflectPrintType(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
}

func main() {
	var v interface{} = "s"
	reflectPrintType(v)
	v = 1
	reflectPrintType(v)
	v = true
	reflectPrintType(v)
	v = make(chan int)
	reflectPrintType(v)
	v = make(chan bool)
	reflectPrintType(v)
	v = make(chan string)
	reflectPrintType(v)
}
