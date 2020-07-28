package main

import (
	"errors"
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

	// 高阶函数演示
	ret := calc2(1, 2, add1)
	fmt.Println(ret)

	fmt.Println("defer start")
	deferTest()
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

// 高阶函数

// 函数作为参数

func add1(x, y int) int {
	return x + y
}
func calc2(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

// 匿名函数
// 函数内部不能定义一般的函数,只能使用匿名函数,匿名函数就是没有函数名的函数,匿名函数多用于实现回调函数和闭包。
func niming() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}

// 闭包
// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
// func calc(base int) (func(int) int, func(int) int) {
// 	add := func(i int) int {
// 		base += i
// 		return base
// 	}

// 	sub := func(i int) int {
// 		base -= i
// 		return base
// 	}
// 	return add, sub
// }

// func main() {
// 	f1, f2 := calc(10)
// 	fmt.Println(f1(1), f2(2)) //11 9
// 	fmt.Println(f1(3), f2(4)) //12 8
// 	fmt.Println(f1(5), f2(6)) //13 7
// }

// defer 语句
// Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
// 在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
// 由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。
func test() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

// 依次输出start,end,3,2,1


func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func deferTest() {
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
}
