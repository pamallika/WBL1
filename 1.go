package main

import "fmt"

type Human struct {
	Name string
	Age  uint
}

// Action наследует Human и может использовать его методы
type Action struct {
	Human
}

// Конструкторы для структур
func NewHuman(name string, age uint) Human {
	return Human{Name: name, Age: age}
}

func NewAction(h Human) Action {
	return Action{Human: h}
}

func main() {
	human := NewHuman("alex", 23)
	action := NewAction(human)

	IsAdult := action.isAdult()
	humanName := action.getName()

	fmt.Printf("human: %+v\n", human)
	fmt.Printf("action: %+v\n", action)
	fmt.Printf("Совершеннолетний? %v\n", IsAdult)
	fmt.Printf("Имя: %s\n", humanName)

}

// Методы для структуры Human{} и его наследников(т.е. эти методы можем вызывать и у структуры Action{})
func (h Human) isAdult() bool {
	return h.Age >= 18
}

func (h Human) getName() string {
	return h.Name
}
