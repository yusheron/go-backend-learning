package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	//变量声明
	var a int = 10
	var b = 20
	c := 30

	fmt.Println(a, b, c)

	//变量类型
	var age int = 20
	var score float64 = 99.5
	var passed bool = true
	var name string = "Go"

	fmt.Println(age, score, passed, name)

	//if语句
	if age >= 18 {
		fmt.Println("成年人")
	}

	if score := 85; score >= 60 {
		fmt.Println("及格")
	}

	//for循环
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	//函数调用
	result := add(2, 3)
	fmt.Println(result)
}
