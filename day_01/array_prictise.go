package main

import (
	"fmt"
	"reflect"
)

/*
数组是同一种数据类型元素的集合。 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变
var a [5]int， 数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。 [5]int和[10]int是不同的类型
*/
func arrayPractise() {
	fmt.Println()
	fmt.Println("hello,world")
	var a [3]int
	var b [4]int
	c := 123
	// a = b //不可以这样做，因为此时a和b是不同的类型
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(a)
	fmt.Println(b)
	arr1 := [3]string{"dds", "dsa", "答"}
	var arr2 [6]string // 初始化数组会用空值来填充
	fmt.Println(len(arr1))
	fmt.Println(arr2)

	// 编译器根据初始值的个数自行推断数组的长度
	arr3 := []int{1, 2, 3, 4, 5}
	fmt.Println(arr3)
	fmt.Println(len(arr3))
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for index, value := range arr3 {
		fmt.Println(index, value)
	}

}
