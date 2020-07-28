package main

import (
	"fmt"
	"reflect"
)

func reflectPractise() {
	fmt.Println()
	a := 3.14
	reflectType(a)
	b := []int{}
	reflectType(b)
	var c float32 = 3.15
	reflectType(c)

	var a1 float32 = 3.14
	var b1 int64 = 100
	reflectValue(a1) // type is float32, value is 3.140000
	reflectValue(b1) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	c1 := reflect.ValueOf(10)
	fmt.Printf("type c1 :%T\n", c1) // type c :reflect.Value

	var a2 int64 = 100
	// reflectSetValue1(a2) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a2)
	fmt.Println(a2)
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("typeof:%v\n", v)
	fmt.Printf("kind:%v\ntype:%v\n", v.Kind(), v.Name())
	// Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空
	// 在reflect包中定义的Kind类型
	/*
			type Kind uint
		const (
		    Invalid Kind = iota  // 非法类型
		    Bool                 // 布尔型
		    Int                  // 有符号整型
		    Int8                 // 有符号8位整型
		    Int16                // 有符号16位整型
		    Int32                // 有符号32位整型
		    Int64                // 有符号64位整型
		    Uint                 // 无符号整型
		    Uint8                // 无符号8位整型
		    Uint16               // 无符号16位整型
		    Uint32               // 无符号32位整型
		    Uint64               // 无符号64位整型
		    Uintptr              // 指针
		    Float32              // 单精度浮点数
		    Float64              // 双精度浮点数
		    Complex64            // 64位复数类型
		    Complex128           // 128位复数类型
		    Array                // 数组
		    Chan                 // 通道
		    Func                 // 函数
		    Interface            // 接口
		    Map                  // 映射
		    Ptr                  // 指针
		    Slice                // 切片
		    String               // 字符串
		    Struct               // 结构体
		    UnsafePointer        // 底层指针
		)
	*/
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

// IsNil   报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic
// func (v Value) IsNil() bool

// IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic
// func (v Value) IsValid() bool
