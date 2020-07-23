package main

import (
	"fmt"
	"math"
)

func startMath() {
	fmt.Println("start math")
	fmt.Println(math.Pi)
	c := math.Pi
	a := 123
	b := 456
	fmt.Println(!(a < b || a < int(c)))
	fmt.Println(1 + 2 - 4/2 + 9%2)
	d := 8 >> 10 //	右移n位就是除以2的n次方。
	f := 8 << 10 //左移n位就是乘以2的n次方。
	fmt.Println(d)
	fmt.Println(f)
}
