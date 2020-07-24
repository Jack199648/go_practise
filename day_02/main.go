package main

import (
	"fmt"
)

/*
func 函数名(参数变量 参数类型)(返回值){
    函数体
}

*/

func main() {
	// 函数基础
	fmt.Println("hello_day_02")
	sum := intSum(19, 70)
	fmt.Println(sum)
	sum2 := intSum2(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum2)
	a := someFunc("")
	fmt.Println(a)
	// 函数进阶
	// 定义一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。
	type calculation func(int, int) int
	var c calculation
	c = intSum
	d := c(1, 2)
	fmt.Println(d)
}

// 两数之和
func intSum(x int, y int) int {
	return x + y
}

// 函数的参数中如果相邻变量的类型相同，则可以省略类型
// func intSum(x, y int) int {
// 	return x + y
// }

// 可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识
// 本质上，函数的可变参数是通过切片来实现的。
func intSum2(x ...int) int {
	fmt.Println(x) // x 是一个切片
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 多个返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 返回值命名
func calc1(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片
func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	} else {
		sli := []int{1}
		return sli
	}
}

// 变量作用域
// 全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。
// 局部变量分两种:
// 函数内定义的变量无法在该函数外使用,如果局部变量和全局变量重名，优先访问局部变量
// if条件判断、for循环、switch语句中定义的变量,只在语句块中起作用,语句块外无法访问该变量
