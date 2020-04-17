package main

import (
	"fmt"
)

func main() {
	abc := vFunc
	su, sb := calc(7, 3)
	fmt.Println(intSum(3, 4))
	fmt.Println(su, sb)
	fmt.Println("start...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end...")  // defer 通常用于处理资源释放问题，比如：资源清理、文件关闭、解锁及记录时间等。
	fmt.Println(someFunc(""))
	fmt.Println(someFunc("a"))  // 填入的string必须是字符串
	fmt.Printf("%T\n", abc)
	fmt.Println(abc())
	a := wheels(100, 200, ob)
	fmt.Println(a)
}

func intSum(a, b int) int {  // 参数类型简写此时a为 int
	s := a + b
	return s
}  // Golang中是没有参数默认值的。小知识：Google的C++规范里面也禁止使用函数默认参数.

func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func someFunc(x string) []int {
	s := []int{1, 2, 3}
	if x == "" {
		return nil  // 当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。
	}
	return s
}

func vFunc() string {
	return "vFunc"
}

func ob(a, b int) int {
	return a + b
}

func wheels(a, b int, op func(x, y int) int) int { // 可以把函数作为参数， 函数也可以作为返回值
	return op(a, b)
}

