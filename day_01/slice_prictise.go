package main

import (
	"fmt"
)

/*
切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
*/
func slicePractise() {
	// 声明切片类型
	var arr1 = [4]int{}
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", arr1)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	a1 := []string{"sa", "dds", "saa", "dsdsd"}
	// 基于数组构造切片
	sli := a1[1:3]
	fmt.Println(sli, len(sli), cap(sli))

	// 使用make()函数构造切片
	sli2 := make([]int, 4, 10)
	fmt.Println(sli2)
	fmt.Println(len(sli2))
	fmt.Println(cap(sli2))
	str1 := "ssssss"
	sli3 := str1[:4]
	fmt.Println(sli3)
	fmt.Println((len(sli3) != 0))
	fmt.Println(nil)
	// 切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）
	// 切片不能进行比较
	// var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	// s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	// s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	s2 := make([]int, 5)
	s3 := s2 //将s1直接赋值给s2，s1和s2共用一个底层数组,对一个切片的修改会影响另一个切片的内容
	s2[0] = 100
	fmt.Println(s2)
	fmt.Println(s3)
	var s []int // 声明一个切片
	s = append(s, 1, 2, 3, 4, 5)
	// 每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。
	// “扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。
	fmt.Println(s)
	var sli4 []int
	for i := 0; i < 10; i++ {
		sli4 = append(sli4, i)
		fmt.Printf("sli4:%v,len(sli4):%d,cap(sli4):%d,ptr:%p,type:%T\n", sli4, len(sli4), cap(sli4), sli4, sli4)
	}
	var sli5 []int
	sli6 := sli5
	fmt.Printf("sli5_ptr:%p,sli6_ptr:%p\n", sli5, sli6) //sli5_ptr:0x0,sli6_ptr:0x0
	// 由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化
	// Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中 copy(destSlice, srcSlice []T)
	sli7 := []int{}
	copy(sli7, sli5)
	fmt.Printf("sli7_ptr:%p\n", sli7) // sli7_ptr:0x11ccc50
	var sli8 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sli8 = append(sli8[:2], sli8[3:]...) // 删除索引为 2 的元素
	fmt.Println(sli8)
}
