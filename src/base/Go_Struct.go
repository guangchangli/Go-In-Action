package main

import (
	"fmt"
)

/**
类型别名 只存在编译 编译后不存在
自定义类型一直存在
结构体
	实例化后才会分配内存,没有初始化默认类型零值
	new 实例化 返回地址 指针可以通过 。 访问属性
	键值初始化保证顺序 全部初始化
	占用连续内存，空结构体不占用内存空间
结构体很复杂构造返回指针
TODO 内存对齐
*/
type person struct {
	name string
	age  uint8
	tel  string
}

func init() {
	type myInt int64
	type mint = int32 //别名
	var a myInt
	var b mint
	fmt.Printf("%T\n", a) //main.myInt
	fmt.Printf("%T\n", b) //int32

	p := person{
		name: "jasmine",
		age:  20,
		tel:  "155156",
	}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	var p1 person
	p1.age = 16
	p1.name = "jasmine"
	p1.tel = "155156"
	fmt.Printf("%#v\n", p1)
	//匿名结构体
	var user struct {
		name string
		age  uint8
	}
	user.name = "aladdin"
	user.age = 16
	fmt.Printf("%+v\n", user)
	//new 实例化 获得地址
	p2 := new(person)
	p2.age = 18
	p2.name = "jasmine"
	p2.tel = "155"
	fmt.Printf("%v\n", p2) //&{jasmine 18 155}
	fmt.Printf("%v\n", *p2)
	p3 := &person{}
	p3.age = 20
	p3.name = "jasmine"
	p3.tel = "155156"
	fmt.Printf("%v\n", *p3)
	fmt.Printf("%v\n", p3)
}
func main() {
	print()
	fmt.Printf("%+v\n", *newPerson("jasmine", "zz", 16, "155"))
}
func print() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	println("----")
	for _, v := range stus {
		fmt.Printf("%v\n", &v)
	}
	//他妈的 遍历到最后 被修改了
	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

type student struct {
	name string
	age  int
}

//自定义构造函数
func newPerson(name, city string, age uint8, tel string) *person {
	return &person{
		name: name,
		age:  age,
		tel:  tel,
	}
}
