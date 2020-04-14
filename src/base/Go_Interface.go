package main

import "fmt"

/**
	接口是一种类型 定义行为
 	指针变量都可以赋值给相应类型变量

*/
func main() {
	//接口变量
	var s Sayer
	d := dog{}
	d.say()
	s = d
	s.say()
	testIn()

}
func testIn() {
	var m mover
	d := dog{}
	m = d
	m.move()
	m = &d
	m.move()
	var a animal
	dd := &dog{}
	a = dd
	a.eat()
}

type Sayer interface {
	say()
}
type mover interface {
	move()
}
type animal interface {
	eat()
}

type cat struct {
	color string
}
type dog struct {
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}
func (d dog) say() {
	fmt.Println("汪汪汪")
}
func (d dog) move() {
	fmt.Println("小狗疯了一样 停不下来")
}
func (d *dog) eat() {
	fmt.Println("狗吃屎，自己的屎也吃，不服不行")
}

type People interface {
	Speak(string) string
}

type women struct{}
type mann struct{}

func (w *mann) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大傻比"
	} else {
		talk = "您好"
	}
	return
}
func (w *women) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大傻比"
	} else {
		talk = "您好"
	}
	return
}
func testSb() {
	var w People = &women{}
	var w1 People = &mann{}
	think := "sb"
	fmt.Println(w.Speak(think))
	fmt.Println(w1.Speak(think))
}
