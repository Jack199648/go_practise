package main

import (
	"fmt"
	"math"
	"reflect"
)

func sss() {
	fmt.Println("hello,world")

	// 整型
	// 整型分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64
	// 其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型
	// 注意： 在使用int和 uint类型时，不能假定它是32位或64位的整型，而是考虑int和uint可能在不同平台上的差异
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制
	fmt.Printf("%x \n", a) // a  占位符%x表示十六进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	// 打印浮点数时，可以使用fmt包配合动词%f
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	// 复数
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	// Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值
	/* 布尔类型变量的默认值为false。
	Go 语言中不允许将整型强制转换为布尔型.
	布尔型无法参与数值运算，也无法与其他类型进行转换。 */

	// 字符串
	// Go 语言里的字符串的内部实现使用UTF-8编码。 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符
	// 字符串转义: \r 回车符（返回行首） ; \n 换行符（直接跳到下一行的同列位置）;\t 制表符;\' 单引号;\" 双引号;\\反斜杠
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"") //str := "c:\Code\lesson1\go.exe"

	// Go语言中要定义一个多行字符串时，就必须使用反引号字符：
	s1 := `第一行
第二行
第三行
`
	fmt.Println(s1)
	// strings.Join() //strings包含一系列字符串操作方法

	// 组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，
	//uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
	// rune类型，代表一个 UTF-8字符。rune类型实际是一个int32
	_a := '中'
	_b := 'x'
	fmt.Println()
	fmt.Println(_a)
	fmt.Println(_b)
	fmt.Println()
	// 遍历字符串
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	//104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153() 230(æ) 178(²) 179(³)
	//104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河)
	// 因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。
	// 字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()

	// 修改字符串
	// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	str1 := "big"
	bytestr := []byte(str1) //byte
	bytestr[0] = 'p'
	fmt.Println(string(bytestr))

	str2 := "白萝卜"
	runestr := []rune(str2) //rune
	runestr[0] = '红'
	fmt.Println(string(runestr))

	// 类型转换
	// Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。
	aa, bb := 3, 4
	cc := int(math.Sqrt(float64(aa*aa + bb*bb)))
	fmt.Println(cc)
	fmt.Println(reflect.TypeOf(cc))
	fmt.Println(reflect.TypeOf('中'))
	fmt.Println(reflect.TypeOf('c'))

}
