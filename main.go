package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// fmt.Println("hi~")
	a := [3]string{"hoge", "fuga", "hogehoge"}
	_ = a
	// a := [...]string{"hoge","fuga","hogehoge"}
	// a[2] = "fugafuga"
	// b := [...][...]string{{"hoge","hoge","hoge"},{"fuga","fuga","fuga"},{"pui","pui","pui"}}

	rand.Seed(time.Now().UnixNano())
	// count := rand.Intn(20)

	// for i := 0; i < count; i++ {
	// 	age := rand.Intn(30)
	// 	result := message(age)
	// 	fmt.Println(result)
	// }
	// fmt.Println("ループ数：" + strconv.Itoa(count) + "回")

	var u User
	u.age = 20
	u.favorite = "読書"
	u.gender = "男"
	fmt.Println(u)
}

type User struct {
	gender   string
	age      int
	favorite string
}

func message(age int) string {
	str_age := strconv.Itoa(age)
	if age >= 20 {
		return str_age + "歳なんだ、もう大人だね！"
	} else {
		return str_age + "歳なんだ、大人じゃないのだよ。。。"
	}
}
