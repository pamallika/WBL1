package main

import "fmt"

func reverseStr(s string) string {
	//Превращаем строку в слайс рун
	runes := []rune(s)
	//Переворачиваем слайс
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// showSize Показывает разммер каждого символа
func showSize(str string) {
	runes := []rune(str)
	for i := range runes {
		fmt.Println(string(runes[i]), "size: ", len(string(runes[i])))
	}
}

func main() {
	s1 := "约大 ୮Ḩ@ሎỢℾ💓i ,iH"
	showSize(s1)
	fmt.Printf("PAYLOAD: %s\nRESULT: %s\n", s1, reverseStr(s1))
}
