package main

import "fmt"

func Intersect(arr1 []int, arr2 []int) []int {
	var result []int
	// Создаём мапу и записываем как ключи - значения из первого массива
	mapa := make(map[int]struct{})
	for _, v := range arr1 {
		mapa[v] = struct{}{}
	}
	//Если ключ повторяется, записываем в пересечение(result)
	for _, v := range arr2 {
		_, exist := mapa[v]
		if exist {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	var set1 = []int{5, 2, 3, 4, 0, 6}
	var set2 = []int{8, 9, 0, 1, 3, 2}
	intersect := Intersect(set1, set2) // пересечение
	fmt.Println(intersect)
}
