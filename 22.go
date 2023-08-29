package main

import "fmt"

type BigNums struct {
	a int
	b int
}

func NewBigNums(a, b int) *BigNums {
	return &BigNums{a: a, b: b}
}
func (nums *BigNums) addition() int {
	return nums.a + nums.b
}
func (nums *BigNums) divisionA() int {
	return nums.a / nums.b
}
func (nums *BigNums) multiplication() int {
	return nums.a * nums.b
}
func (nums *BigNums) subtractionA() int {
	return nums.a - nums.b
}
func (nums *BigNums) divisionB() int {
	return nums.b / nums.a
}
func (nums *BigNums) subtractionB() int {
	return nums.subtractionA() * -1
}

func main() {
	a := 10485760
	b := 104857600
	nums := NewBigNums(a, b)
	fmt.Println("a+b =", nums.addition())
	fmt.Println("a/b =", nums.divisionA())
	fmt.Println("a*b =", nums.multiplication())
	fmt.Println("a-b =", nums.subtractionA())
	fmt.Println("b-a =", nums.subtractionB())
	fmt.Println("b/a =", nums.divisionB())
}
