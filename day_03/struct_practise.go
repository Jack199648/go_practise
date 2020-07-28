package main

import (
	"encoding/json"
	"fmt"
)

func structPractise() {
	fmt.Println("")
	type NewInt int  // 类型定义
	type Myint = int // 类型别名
	var a NewInt
	var b Myint
	fmt.Printf("typeof a:%T\n", a)
	fmt.Printf("typeof b:%T\n", b)
	// 结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型

	// 定义结构体
	type Person struct {
		name string
		city string
		age  int8
	}
	type Person1 struct {
		name, city string
		age        int8
	}

	// 实例化结构体
	var p Person
	p.age = 18
	p.city = "成都"
	p.name = "dsadsa"
	fmt.Printf("p=%v\n", p)
	fmt.Printf("p=%#v\n", p)

	// 匿名结构体
	var c struct {
		Name string
		age  int8
	}
	c.Name = "hello"
	c.age = 12
	fmt.Printf("c=%v\n", c)

	// 指针结构体
	var p2 = new(Person)
	fmt.Printf("p2=%T\n", p2)
	fmt.Printf("p2=%#v\n", p2)
	fmt.Printf("p2=%v\n", p2)

	// 结构体初始化
	// 没有初始化的结构体，其成员变量都是对应其类型的零值。
	// 使用键值对初始化
	p3 := Person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p3=%#v\n", p3) //p3=main.person{name:"小王子", city:"北京", age:18}

	//使用值的列表初始化
	p4 := &Person{
		"沙河娜扎",
		"北京",
		28,
	}
	fmt.Printf("p4=%#v\n", p4) //p4=&main.person{name:"沙河娜扎", city:"北京", age:28}

	// 结构体占用一块连续的内存
	fmt.Printf("name_ptr:%p\ncity_ptr:%p\nage_ptr:%p\n", &p4.name, &p4.city, &p4.age)
	// name_ptr:0xc000098300
	// city_ptr:0xc000098310
	// age_ptr:0xc000098320

	// 匿名结构体(字段名为类型名,隐藏)
	//Address 地址结构体
	type Address struct {
		Province   string
		City       string
		CreateTime string
	}

	//Email 邮箱结构体
	type Email struct {
		Account    string
		CreateTime string
	}

	//User 用户结构体
	type User struct {
		Name   string
		Gender string
		Address
		Email
	}

	var user3 User
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
	user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime

	// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）

	// JSON序列化
	//Student 学生
	type Student struct {
		ID     int
		Gender string
		Name   string
	}

	//Class 班级
	type Class struct {
		Title    string
		Students []*Student
	}
	c1 := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c1.Students = append(c1.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c2 := &Class{}
	err = json.Unmarshal([]byte(str), c2)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c2)
}
