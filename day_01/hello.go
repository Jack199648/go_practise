package main // 包声明

import "fmt" // 引入包

// 整型和浮点型变量的默认值为0。 字符串变量的默认值为空字符串。 布尔型变量默认为false。 切片、函数、指针变量的默认为nil
// var _ss string = "qqq"
// var _hh bool = true
// var _123 int = 19
// var dd,bb  = "dsadsa",888
// var dss = 19

// var (
// 	a string
// 	b bool
// 	c int
// 	d float32
// )

// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	n1 = 221212
	n2
	n3
	n4
)

// iota是go语言的常量计数器，只能在常量的表达式中使用。
// iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用
const (
	a1 = iota //0
	a2        //1
	a3        //2
	_         //使用_跳过某些值
	a4        //4
	a5 = 100
	a6 = iota //6
)

// 定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。
// 同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

// 多个iota定义在一行
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func main() { // 函数
	// sss()
	// startMath()
	// process()
	// arrayPractise()
	// slicePractise()
	mapPractise()
	// 在函数内部，可以使用更简略的 := 方式声明并初始化变量
	b := "dsadasd"
	fmt.Println("hello,go")
	fmt.Print("hello,go\n")
	fmt.Println(b)
	test()
	// 在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线_表示
	// 函数外的每个语句都必须以关键字开始（var、const、func等）
	// :=不能使用在函数外。
	// _多用于占位，表示忽略值。
	x, _ := myfun()
	_, y := myfun()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
}

func test() {
	const pi = 3.1415926
	fmt.Println(pi)
	fmt.Printf("test\n")
}
func myfun() (int, string) {
	return 10, "dsadsadsa"
}
