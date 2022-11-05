package main

import (
	"fmt"
)

type Human interface {
	Info()
	Greet()
	BirthDay()
}

type User struct {
	Name string
	Age  int64
}

//UserがTT型を実装しているか確認できる
// Doc: https://go.dev/doc/faq#guarantee_satisfies_interface
//レシーバがポインタの場合
// var _ Human = (*User)(nil)
//ポインタでない場合
var _ Human = User{}

// func (u *User) Greet() string {
// 	return "Hello" + u.Name
// }

// func (u *User) Info() {
// 	fmt.Printf("%s is % d years old\n", u.Name, u.Age)
// }

// func (u *User) BirthDay() {
// 	u.Age += 1
// }

func (u User) Greet() {
	fmt.Println("Hi " + u.Name)
}

func (u User) Info() {
	fmt.Printf("%s is % d years old\n", u.Name, u.Age)
}

//Userの値を変更する必要があればポインタ型でレシーバを指定
//その他の場合は値を返す
func (u User) BirthDay() {
	u.Age += 1
}

func main() {
	u := User{
		Name: "Tanaka",
		Age:  53,
	}

	u.Greet()
	u.Info()
	u.BirthDay()
	u.Info()
}
