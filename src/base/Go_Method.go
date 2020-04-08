package main

import "fmt"

/**
	方法 作用于特定类型(接受者)变量的函数，属于特定类型的属性
	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    	函数体
	}
	接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，
	而不是self、this之类的命名。
	例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
	当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。
	在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
	什么时候使用指针接收
		需要修改接收者中的值
		接收者是拷贝代价比较大的大对象
		保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
	嵌套结构体内部可能存在相同的字段名。这个时候为了避免歧义需要指定具体的内嵌结构体的字段。
	结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问
*/
//Person 结构体
type Person struct {
	name string
	age  uint8
}

//匿名字段
type car struct {
	string
	int
}

//嵌套结构体
type house struct {
	string
	car
}

func NewPerson(name string, age uint8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}
func (p Person) Dream() string {
	fmt.Printf("%s%d岁的梦想是学好Go语言！\n", p.name, p.age)
	return "dream return"
}

func (p *Person) setAge(age uint8) {
	p.age = age
}
func main() {
	p1 := NewPerson("Aladdin", 23)
	p1.setAge(24)
	dream := p1.Dream()
	fmt.Printf("%s\n", dream)
	h := house{
		string: "一号",
		car: car{
			"panamera",
			100,
		},
	}
	fmt.Printf("%+v\n", h)
	structInherit()

}

func structInherit() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}
