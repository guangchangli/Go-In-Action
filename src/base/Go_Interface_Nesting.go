package main

import (
	"fmt"
)

/**
接口嵌套 新的接口
空接口
	空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
	空接口类型的变量可以存储任意类型的变量。
	应用
		作为参数
		map值
*/
func main() {
	var a apple
	m := &me{}
	a = m
	a.call()
	a.listen()
	a.read()
	a.work()
	var null interface{}
	fmt.Println(null)
	inter := testNullInter("null interface apply")
	fmt.Printf("%+v\n", inter)
}
func testNullInter(i interface{}) interface{} {

	switch i.(type) {
	case map[string]string:
		fmt.Printf("%+v\n\n", i)
	case string:
		fmt.Printf("%v\n", i)
	default:
		fmt.Printf("%s\n", i)
	}
	m := make(map[string]string)
	m["m"] = "f"
	return m
}

type me struct {
	name string
}

func (m *me) call() {
	fmt.Println(m.name + "use iphone call")
}
func (m *me) work() {
	fmt.Println(m.name + "use iphone word")
}
func (m *me) read() {
	fmt.Println(m.name + "use iphone read")
}
func (m *me) listen() {
	fmt.Println(m.name + "use iphone listen")
}

type iphone interface {
	call()
}
type ipad interface {
	read()
}
type mac interface {
	work()
}
type airPods interface {
	listen()
}
type apple interface {
	ipad
	iphone
	mac
	airPods
}
