package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

func main() {
	a1 := make(map[string]int, 8)  // 8为容量
	a2 := map[int]string{
		1: "Apple",
		2: "Orange",
		3: "straberry",
	}
	a1["小米"] = 18
	fmt.Println(a1, a2)
	a1["小李"] = 38
	a1["李四"] = 25
	fmt.Println(a1, a2)  // cap()不支持map类型作为参数
	_, ok := a1["小天"]
	if ok {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
	for k, v := range a1 {
		fmt.Println(k,v)
	}
	for k := range a1 {
		fmt.Println(k)
	}  //  遍历map时的元素顺序与添加键值对的顺序无关。
	delete(a2, 1)
	fmt.Println(a2)
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreMap = make(map[string]int, 20)
	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(20)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 20)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}  // 按照指定顺序遍历map
	// 元素为map类型的切片
	var mapSlice = make([]map[string]string, 3)  // 只写一个3表示容量和长度都一样
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)  // 要加入map还需在声明
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	// 值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
	testFunc("小李子")
	fmt.Println(intSum(10, 22))
	fmt.Println(intSpliceSum(5, 10, 20, 30))
}

func testFunc(name string) {
	fmt.Println("Hello", name)
}

func intSum(a int, b int) (s int) {
	s = a + b
	return  // 或者写return s 推荐写完整
}

func intSpliceSum(a int, b ...int) int {  // 固定参数和可变参数同时出现时，可变参数一定要放在最后面，并且参数a必须指定，参数b可以不写
	s := 0
	for _, v := range b {
		s = s + v
	}
	return s + a
}