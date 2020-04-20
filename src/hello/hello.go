package main

import (
	"fmt"
)

func main() {
	var pwd string
	for {
		fmt.Println("请输入：")
		fmt.Scanln(&pwd)
		if pwd == "abcdef" {
			fmt.Println("输入串")
		}
	}
}
