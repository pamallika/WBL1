package main

import "fmt"

// Чтобы не путаться я бы разнёс по пакетам, но одно решение - один файл
// OutputData Требуется реализовать этот интерфейс и вызвать printData не имея доступа к конструктору outsourceUser
// и не переопределяя метод
type OutputData interface {
	toString() string
}

func printData(o OutputData) {
	fmt.Println(o.toString())
}

type OutsourceUser struct {
	name  string
	age   string
	phone string
}

// Представим что не имеем доступа к этому конструктору
func newOutsourceUser(name, age, phone string) *OutsourceUser {
	return &OutsourceUser{name: name, age: age, phone: phone}
}

// Для примера возьмём что у другого сервиса метод который складывает поля структуры пользователя так втупую
// и если мы попробуем передать свою структуру пользователя - вылетит ошибка т.к. мы попытаемся сложить стрингу с интом (age int)
func (u *OutsourceUser) toString() string {
	return u.name + u.age + u.phone
}

func (u *OutsourceUser) printData() string {
	return u.toString()
}

type User struct {
	name  string
	age   int
	phone string
}

func NewUser(name, phone string, age int) *User {
	return &User{name: name, age: age, phone: phone}
}

type UserAdapter struct {
	outsourceUser *OutsourceUser
}

// В адаптере превращаем age int в age string и возвращаем объект UserAdapter в котором OutsourceUser через который можем вызвать нужный метод
func NewUserAdapter(u User) *UserAdapter {
	return &UserAdapter{
		outsourceUser: &OutsourceUser{
			name:  u.name,
			age:   fmt.Sprintf("%v", u.age),
			phone: u.phone,
		},
	}
}

func main() {
	outsourceUser := newOutsourceUser("Alex", "20", "8800553535")
	user := NewUser("Alex", "8800553535", 20)
	userAdapter := NewUserAdapter(*user)
	printData(outsourceUser)
	printData(userAdapter.outsourceUser)
}
