package main

import (
	"fmt"
)

func printType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("channel int")
	case chan bool:
		fmt.Println("channel bool")
	case chan string:
		fmt.Println("channel string")
	}

}

func main() {
	var v interface{} = "s"
	printType(v)
	v = 1
	printType(v)
	v = true
	printType(v)
	v = make(chan int)
	printType(v)
	v = make(chan bool)
	printType(v)
	v = make(chan string)
	printType(v)
}
