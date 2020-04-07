package main

import "fmt"

/**
指针地址、数值、类型
指针不能进行偏移和运算
& 取地址  * 根据地址取值
值类型（int、float、bool、string、array、struct）都有对应指针类型
对变量进行取地址（&）操作，可以获得这个变量的指针变量。
指针变量的值是指针地址。
对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值
在Go语言中对于引用类型的变量，
	我们在使用的时候不仅要声明它，还要为它分配内存空间，
	否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，
	是因为它们在声明的时候已经默认分配好了内存空间
func new(Type) *Type
	Type表示类型，new函数只接受一个参数，这个参数是一个类型
	*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
new make
	二者都是用来做内存分配的。
	make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/
func init() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)
	c := *b
	fmt.Printf("c:%d,type:%T\n", c, c)
}
func main() {
	b := 10
	modify1(b)
	println(b)
	modify2(&b)
	println(b)

	var a *int
	a = new(int)
	*a = 100
	fmt.Println(*a)

	var c map[string]int
	c = make(map[string]int, 10)
	c["aladdin"] = 100
	fmt.Println(c)
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}
