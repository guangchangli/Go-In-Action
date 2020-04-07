package main

import (
	"fmt"
	"reflect"
)

//全局变量
const cons = "my const"

var boo = true

const (
	n1 = iota //0
	n2        //1
	_         //2
	n4        //3
)
const n5 = iota //0
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func main() {
	/**
	类型推断
	var variable type = realValue
	匿名变量不占用命名空间
	不占用内存
	_ 占位 忽略值
	iota 计数器 这个比较神奇
	*/
	fmt.Println(reflect.TypeOf(boo))
}
