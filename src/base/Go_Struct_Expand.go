package main

import (
	"encoding/json"
	"fmt"
)

/**
slice map struct 引用类型 使用需要拷贝 后者变量引用
*/
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

func main() {
	jsonTag()
	jsonTest()
	structTest()
}

func jsonTag() {
	type person struct {
		Name string `json:"name"` //通过指定tag实现json序列化该字段时的key
		age  uint8  //私有不能被json包访问
		Sex  string //json序列化是默认使用字段名作为key
	}
	p := person{
		Name: "jasmine",
		age:  23,
		Sex:  "女",
	}
	marshal, err := json.Marshal(p)
	if err == nil {
		fmt.Printf("json:%s\n", marshal)
	}
	return
}
func jsonTest() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
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
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}
func structTest() {
	p := &Phone{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p.dreams)
	//拷贝一份
	bak := make([]string, len(data))
	copy(bak, data)
	p.SetDreams(bak)
	bak[1] = "学go"
	fmt.Printf("%+v\n", bak)
	fmt.Printf("%+v\n", data)
}
func (p *Phone) SetDreams(dreams []string) {
	p.dreams = dreams
}

type Phone struct {
	name   string
	age    int8
	dreams []string
}
