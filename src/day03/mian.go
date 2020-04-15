package main

import (
	"fmt"
)

func main() {
	var a1[3]string
	var a2 = [4]int{1, 2, 3, 4}
	var a3[3]int // [0 0 0]
	var a4 = [...]int{1:3, 4:5}  // 指定索引
	a5 := []string{"北京", "上海", "成都", "广州"}  // name[]T这个用法是切片
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5[2])
	for index, v := range a5 {
		fmt.Println(index, v)
	}
	for _, v := range a5 {
		fmt.Println(v)
	}  // range会返回索引和遍历的数组内的值，不想显示索引时可以用_替代

	// func modifyArray(x [3]int) {
	// 	x[0] = 100
	// }
	// func modifyArray2(x [3][2]int) {
	// 	x[2][0] = 100
	// }
	// func main() {
	// 	a := [3]int{10, 20, 30}
	// 	modifyArray(a) //在modify中修改的是a的副本x
	// 	fmt.Println(a) //[10 20 30]
	// 	b := [3][2]int{
	// 		{1, 1},
	// 		{1, 1},
	// 		{1, 1},
	// 	}
	// 	modifyArray2(b) //在modify中修改的是b的副本x
	// 	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
	// } ***数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	var b1 []string  // slice类型(切片)
	var b2 = []int{}
	var b3 = []bool{true, false}
	// var b4 = []bool{true, false}
	var b4 = []string{"Apple", "Microsoft", "Google"} 
	fmt.Println(b1 == nil)  // 关于nil可见 https://www.jianshu.com/p/174aa63b2cc5
	fmt.Println(b2 == nil)
	fmt.Println(b3 == nil)
	// fmt.Println(b3 == b4) 切片是引用类型，不支持直接比较，只能和nil比较
	fmt.Println(b4, len(b4), cap(b4))  // 我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
	b4 = append(b4, "Facebook")
	fmt.Println(b4, len(b4), cap(b4))  // 容量的用处在哪？在与当你用 appen d扩展长度时，如果新的长度小于容量，不会更换底层数组，否则，go 会新申请一个底层数组，拷贝这边的值过去，把原来的数组丢掉。也就是说，容量的用途是：在数据拷贝和内存申请的消耗与内存占用之间提供一个权衡。
	b4 = append(b4, "Huawei", "Bilibili")
	fmt.Println(b4, len(b4), cap(b4))
	b5 := b4[1:2:6] // 左包含，右不包含。[low : high : max] 切片的三个参数的切片截取的意义为 low为截取的起始下标（含）， high为窃取的结束下标（不含high），max为切片保留的原切片的最大下标（不含max）；即新切片从老切片的low下标元素开始，len = high - low, cap = max - low；high 和 max一旦超出在老切片中越界，就会发生runtime err，slice out of range。
	fmt.Println(b5, len(b5), cap(b5))
	// a[2:]  // 等同于 a[2:len(a)]
	// a[:3]  // 等同于 a[0:3]
	// a[:]   // 等同于 a[0:len(a)]
	b6 := make([]int, 2, 4)  // make([]T, size, cap)
	fmt.Println(len(b6), cap(b6))
	// 切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是nil
	var b7 = []int{1, 2, 3, 4}
	var b8 = b7
	fmt.Println(b7,b8)
	b8[0] = 5
	fmt.Println(b7,b8)  // 拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。
	var b9 = []int{1, 2, 3, 4}
	fmt.Println(b9)
	b9 = append(b9[:1], b9[3:]...)  // 删除切片内的元素
	fmt.Println(b9)
}