package main

import (
	"fmt"
	"math"
	"strings"
)

var a = "haha"

func main(){
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
}
