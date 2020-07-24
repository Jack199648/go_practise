package main

import (
	"fmt"
)

/*
Go语言中最常用的流程控制有if和for，而switch和goto主要是为了简化代码、降低重复代码而生的结构，属于扩展类的流程控制。
for循环可以通过break、goto、return、panic语句强制退出循环
*/
func process() {
	fmt.Println()

	fmt.Println("start process")
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
		if true {
			fmt.Println("hello,world")
		}
	}
	if true {
		fmt.Println("hello,world")
	} else {
		fmt.Println("")
	}
	if true {
		fmt.Println("ss")
	} else if 1 < 2 {
		fmt.Println("dd")
	}
	var1 := "哈哈莎婚纱"
	// for range 遍历
	// 数组、切片、字符串返回索引和值。
	// map返回键和值。
	// 通道（channel）只返回通道内的值。
	for _, var1 := range var1 {
		fmt.Println(var1)
	}
	fig := 6
	switch fig {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("无效输入")
	}
	// 一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
	// 分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量
	// fallthrough语法可以执行满足条件的case的下一个case
	age := 5
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
		fallthrough // 输出 "好好学习吧","好好工作吧" 和"好好享受吧"
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
		fallthrough // 输出 "好好工作吧" 和"好好享受吧"
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	// goto跳转到指定标签
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
	fmt.Println()

	// break语句可以结束for、switch和select的代码块。
	// break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和 select的代码块上
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")

	// continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用
	// 在 continue语句后添加标签时，表示开始标签对应的循环
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}

BOOKUP:
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%v*%v=%v\n", i, j, i*j)
			if i == j {
				continue BOOKUP
			}
		}
	}

	fmt.Println()
}
