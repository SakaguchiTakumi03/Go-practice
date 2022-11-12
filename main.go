package main

import (
	"fmt"
	"math"
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

	// u.age = 20
	// u.favorite = "読書"
	// u.gender = "男"
	// u := User{"男", 20, "読書"} //上のコメントでもこっちでもどちらでも大丈夫
	// u := User{20, "男", "読書"} structの構造に合わせた順番にしないといけない
	// fmt.Println(u)
	// result := Student{"hogehoge", 55, 60}
	// s := Student{"hogehoge", 60, 60}
	// avg_result(s)
	BMI_result(43, 163)
}

type Student struct {
	name                 string
	math_point, eg_point float64
}

// type BMI_Point struct {
// 	weight, height float64
// }

type User struct {
	gender   string
	age      int
	favorite string
}

func BMI_result(weight float64, height float64) {
	result_height := height / 100.0
	var result = weight / math.Pow(result_height, 2)
	fmt.Println("身長：", height)
	fmt.Println("体重：", weight)
	fmt.Println("BMIの結果は", math.Floor(result*100)/100, "です")
}

func avg_result(s Student) {
	fmt.Println(s.name, "君の点数は", (s.eg_point+s.math_point)/2, "点です")
}

func message(age int) string {
	str_age := strconv.Itoa(age)
	if age >= 20 {
		return str_age + "歳なんだ、もう大人だね！"
	} else {
		return str_age + "歳なんだ、大人じゃないのだよ。。。"
	}
}
