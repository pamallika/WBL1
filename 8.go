package main

import (
	"fmt"
)

// пока решение такое, потому что не до конца вкурил как работают сдвиги(нам надо сделать маску, потом сдвинуть биты,
// отреверсить все значения кроме изменяемого чтобы вернуть то что сдвинулось)
// поэтому вот так

// editByte в тупую превращает число в строку из бит, из строки в руны, руну под нужным индексом меняет на противоположную
// и выводит строку бит(я бы не назвал это решением , но с задачей справляется)
func editByte(a int64, i int) {
	byteStr := fmt.Sprintf("%b", a)
	fmt.Println(byteStr)
	runes := []rune(byteStr)
	if runes[i] == 48 {
		runes[i] = 49
	} else {
		runes[i] = 48
	}
	byteStr = string(runes)
	fmt.Println(byteStr)
}

func main() {
	var a int64
	var i int
	a = 123
	i = 1
	editByte(a, i)
}
