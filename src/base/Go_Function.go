package main

import (
	"errors"
	"fmt"
	"time"
)

/**
可变参数 是通过切片实现的 要放在后面
局部变量和全局变量重名 优先访问局部变量
函数类型与变量 type typeName func(paramType) returnType
函数作为参数 函数作为返回值
匿名函数 变量保存或者直接调用
闭包 TODO
defer 最先defer 的后执行
	它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前
*/

var sum = 0

type calcu func(int, int) int

func main() {
	fmt.Printf("%T\n",do)
	println(funcDemo(0, 10, 20, 30))
	ret2 := calc(10, 20, ad)
	fmt.Println(ret2) //30
	//函数参数
	println(calc(1, 2, add))
	time.Sleep(1)
	//函数返回值
	fmt.Printf("函数作为返回值\n")
	f2, err := do("+")
	if err != nil {
		println(f2(1, 2))
	}
	//匿名函数
	println("匿名函数\n")
	f := func(x, y int) int {
		return x + y
	}
	println(f(3, 2))
	//直接调用
	println(func(x, y int) int {
		return x + y
	}(3, 2))
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")

}
func testFunc() {
	var c calcu
	c = sub
	println(c(2, 3))
}
func add(x int, y int) int {
	return x + y
}
func sub(x int, y int) int {
	return x - y
}
func funcDemo(x int, y ...int) int {
	fmt.Println(y)
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

func ad(x, y int) int {
	return x + y
}

//函数作为参数
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return ad, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}
