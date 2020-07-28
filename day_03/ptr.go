package main

import (
	"fmt"
)

// 总结： 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
/*
对变量进行取地址（&）操作，可以获得这个变量的指针变量。
指针变量的值是指针地址。
对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值
*/
func ptr() {
	a := 10
	b := &a // 取地址
	c := &b // 指针取值
	d := *b
	e := *c
	fmt.Println(a, b, c, d, e)
	fmt.Printf("%T\n%T\n%T\n%T\n%T\n", a, b, c, d, e)

	// 值类型在声明的时候即分配了内存空间,但是引用类型的变量声明后还不具有内存空间,需要使用new函数或者 make 函数来分配内存空间,方可进行赋值等操作
	/*
		二者都是用来做内存分配的。
		make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
		而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
	*/
	var a1 *int
	a1 = new(int) //对 a1 进行内存分配
	*a1 = 100
	fmt.Println(*a1)
	var b1 map[string]int
	b1 = make(map[string]int, 10) // 对 b1 进行内存分配
	b1["沙河娜扎"] = 100
	fmt.Println(b1)
}
