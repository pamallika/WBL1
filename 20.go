package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	//Делим строку на слайс слов по разделителю - пробелу
	words := strings.Split(str, " ")
	//Переворачиваем слайс
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	str := "Hello world hello world"
	fmt.Printf("PAYLOAD: %s\nRESULT: %s", str, reverseString(str))
}
