package main

import "fmt"

var (
	name string
	age int
)

const (
	n1 = 100
	n2
	n3
)  // const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。

const (
	z1 = iota  // 0
	z2  // 1
	z3  // 2
	_  // 使用_跳过某些值
	z4  // 4
)  // iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。

func foo() (int, string) {
	return 10, "Q1mi"
}

func main(){
	name = "Andew"
	age = 18
	fmt.Printf("My name is %v, I'm %v\n", name, age)  // %v会自动识别类型
	fmt.Println("Hello",name)  // Printf和Println用法不同
	var a string = "str a"  // should omit type string from declaration of var a; it will be inferred from the right-hand side
	fmt.Println(a)
	var b = 16
	fmt.Println("b is", b)
	var c, d = "attack", 18
	fmt.Println(c, d)
	m := "短变量声明"
	fmt.Println(m)
	x, _ := foo()  //匿名变量,用_表示,匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。 
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println(n1, n2, n3)
	fmt.Println(z1, z2, z3, z4)
}
