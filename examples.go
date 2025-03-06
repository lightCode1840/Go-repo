package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.基础语法结构
	fmt.Println("Hello, Go!")

	// 2.变量声明与类型推导
	var a int = 10
	b := "自动类型推导"
	fmt.Printf("a类型:%T 值:%d, b类型:%T 值:%s\n", a, a, b, b)

	// 3.控制结构
	// if语句
	if num := 9; num < 0 {
		fmt.Println("负数")
	} else if num < 10 {
		fmt.Println("个位数")
	}

	// for循环
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 4.切片和map
	// 切片
	nums := []int{1, 2, 3}
	nums = append(nums, 4)
	fmt.Println("切片:", nums)

	// map
	langs := map[string]string{
		"Go":   "高效",
		"Java": "跨平台",
	}
	fmt.Println("编程语言:", langs)

	// 5.协程和通道
	ch := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "通道消息"
	}()
	msg := <-ch
	fmt.Println("收到:", msg)
}
