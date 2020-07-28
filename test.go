package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}



func main() {
	var peo People = &Student{} // 保证值接收者和指针接受者都能正常运行
	think := "bitch"
	fmt.Println(peo.Speak(think))
	
}
