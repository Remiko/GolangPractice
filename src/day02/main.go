package main

import (
	"fmt"
	"math"
	"strings"
)

var a = "haha"

func main() {
	var c int = 10
	var b complex64 = 1 + 2i  // 复数类型
	var d bool  // 1、布尔类型变量的默认值为false。2、Go 语言中不允许将整型强制转换为布尔型.3、布尔型无法参与数值运算，也无法与其他类型进行转换。
	// e := 'banana
	// 	apple
	// 	orange' 不知道为什么多行字符串会报错
	var f = "I'm a pig or pig"
	var g = "gy"
	h := 'h'  // 单个字符用''
	i1 := "big"
	i2 := "小白菜"
	l := 'a'
	fmt.Println(len(a))
	fmt.Printf("%d\n", c)
	fmt.Printf("%o\n", c)
	fmt.Printf("%x\n", c)
	fmt.Printf("%.2f\n", math.Pi)
	fmt.Printf("%v\n", b)
	fmt.Println(d)
	fmt.Println(strings.Index(f, "pig"))  // strings.Index(s, str string) 从0开始计数，返回第一次出现的位置
	fmt.Println(strings.Join([]string{f, g}, ""))  // I'm a pig or piggy(别管意思，明白用法就行)
	fmt.Println(h)
	byteS1 := []byte(i1)  // 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))
	runeS2 := []rune(i2)
	runeS2[0] = '大'
	fmt.Println(string(runeS2))
	if j := 85; j >= 90 {  // if条件判断还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断
		fmt.Println("优秀")
	} else if j >= 60 || j < 90 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
	for k := 0; k >= 2; k++ {
		fmt.Println(k)
	}  // 1、for循环的初始语句可以被忽略，但是初始语句后的;必须要写.2、for循环的初始语句和结束语句都可以省略,并且两个;都不用写
	// for {
	// 	循环体语句
	// }  for循环可以通过break、goto、return、panic语句强制退出循环。
	switch {
	case l == 'a':
		fmt.Println("a")
		fallthrough  // fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
	case l == 'b':
		fmt.Println("b")
	case l == 'c':
		fmt.Println("c")
	default:
		fmt.Println("...")
	}  // Go语言规定每个switch只能有一个default分支。一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	
	// func gotoDemo2() {
	// 	for i := 0; i < 10; i++ {
	// 		for j := 0; j < 10; j++ {
	// 			if j == 2 {
	// 				// 设置退出标签
	// 				goto breakTag
	// 			}
	// 			fmt.Printf("%v-%v\n", i, j)
	// 		}
	// 	}
	// 	return
	// 	// 标签
	// breakTag:
	// 	fmt.Println("结束for循环")
	// }  goto的用法
}
