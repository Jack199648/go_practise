package main

import "fmt"

// 接口（interface）是一种类型,Go语言提倡面向接口编程
// 接口区别于所有的具体类型，接口是一种抽象的类型。当你看到一个接口类型的值时，你不知道它是什么，唯一知道的是通过它的方法能做什么
/*
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略
*/

// Sayer 接口
type Sayer interface {
	say()
}

// 定义两个结构体
type Dog struct{}
type Cat struct{}

// dog实现 Sayer接口
func (d Dog) say() {
	fmt.Println("汪汪汪")
}

// cat 实现 Sayer 接口
func (c Cat) say() {
	fmt.Println("喵喵喵")
}

// 空接口作为函数参数,空接口类型的变量可以存储任意类型的变量
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func interfacePractise() {
	var x Sayer // 声明一个Sayer类型的变量x
	a := Cat{}  // 实例化一个Cat
	b := Dog{}  // 实例化一个Dog
	x = a       // 可以把Cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把Dog实例直接赋值给x
	x.say()     // 汪汪汪

	// 空接口作为函数参数传递任意类型
	show(1)
	show([]int{1, 2, 3})

	// 空接口作为 map 值,保存任意值的字典
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

	// 一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值
	// 类型断言x.(T)  x 为类型为interface{}的变量,T：表示断言x可能是的类型;该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败
	var x1 interface{}
	x1 = "Hello 沙河"
	v, ok := x1.(int)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
