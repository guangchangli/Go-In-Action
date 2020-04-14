package main

import "fmt"

type people interface {
	Speak(string) string
}

type man struct{}

func (stu *man) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = &man{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
