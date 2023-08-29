package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}
	removeIndex := 2
	fmt.Println("PAYLOAD:", sl)
	fmt.Println("REMOVE INDEX:", removeIndex)
	fmt.Println("RESULT:", HardRemove(sl, removeIndex))
}

// HardRemove быстрое удаление, порядок не сохраняется
// Записываем на место удаляемого элемента последний, а последний обрезаем тем самым меняя порядок
func HardRemove(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
