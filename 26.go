package main

import (
	"fmt"
	"strings"
)

func checkUniqueChar(str string) bool {
	charMap := make(map[rune]int)
	//приведем к нижнему регистру
	//Если под ключом(руной) в мапе есть запись значит это повторяющийся символ
	for _, v := range strings.ToLower(str) {
		if charMap[v] != 0 {
			return false
		}
		charMap[v]++
	}
	return true
}

func main() {
	fmt.Println("Введите строку: ")
	var str string
	fmt.Scanf("%s", &str)
	fmt.Println(str, " - ", checkUniqueChar(str))
}
