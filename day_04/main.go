package main // 声明包

/*
一个文件夹下面直接包含的文件只能归属一个package，同样一个package的文件不能在多个文件夹下。
包名可以不和文件夹的名字一样，包名不能包含 - 符号。
包名为main的包为应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含main包的源代码则不会得到可执行文件
*/

import (
	m "day_04/mypackage"
	"fmt"
)

// 包变量可见性

var a = 100 // 首字母小写，外部包不可见，只能在当前包内使用

// 首字母大写外部包可见，可在其他包中使用
const Mode = 1

type person struct { // 首字母小写，外部包不可见，只能在当前包内使用
	name string
}

// 首字母大写，外部包可见，可在其他包中使用
func Add(x, y int) int {
	return x + y
}

func age() { // 首字母小写，外部包不可见，只能在当前包内使用
	var Age = 18 // 函数局部变量，外部包不可见，只能在当前函数内使用
	fmt.Println(Age)
}

// 在Go语言程序执行时导入包语句会自动触发包内部init()函数的调用。需要注意的是： init()函数没有参数也没有返回值。
// init()函数在程序运行时自动被调用执行，不能在代码中主动调用它

func init() {
	fmt.Println("init")
}

// 执行顺序:全局声明-->init()函数-->main()函数
func main() {
	fmt.Println("")
	mysum := m.Test(1, 2)
	fmt.Println(mysum)
}
