package main

import "fmt"

func createSet(data []string) map[string]struct{} {
	//Создаём мапу, ключи не могут дублироваться
	set := make(map[string]struct{})
	for i := range data {
		set[data[i]] = struct{}{}
	}
	return set
}

func main() {
	var data = []string{"cat", "cat", "dog", "cat", "tree"}
	res := createSet(data)
	for i, _ := range res {
		fmt.Println(i)
	}
}
