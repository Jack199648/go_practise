package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func mapPractise() {
	fmt.Println()
	map1 := make(map[string]int, 8)
	map1["张三"] = 99
	map1["王五"] = 79
	fmt.Println(map1)
	v, ok := map1["张三"]
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
	for key, val := range map1 {
		fmt.Printf("%v:%v\n", key, val)
	}
	// 注意:遍历map时的元素顺序与添加键值对的顺序无关

	// 删除 map 键值对  delete(map,key)
	delete(map1, "张三")
	fmt.Println(map1)

	// 按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	// 元素为map类型的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	//  值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)

	fmt.Println()
}
