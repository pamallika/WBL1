package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}
	removeIndex := 1
	fmt.Println("PAYLOAD:", sl)
	fmt.Println("REMOVE INDEX:", removeIndex)
	fmt.Println("RESULT:", SoftRemove(sl, removeIndex))
}

// SoftRemove Удаление элемента из слайса и сдвиг остальной для сохранения порядка
func SoftRemove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
