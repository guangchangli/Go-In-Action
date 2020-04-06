package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

/**
引用类型 map[KeyType]ValueType
初始值为 nil make(map[KeyType]ValueType, [cap]) 分配内存
判断是否存在 v,ok:=map[key]
delete(map,key)

*/
func init() {
	carMap := make(map[string]string, 5)
	carMap["audi"] = "A7"
	carMap["Panamera"] = "porsche"
	fmt.Printf("%v", carMap)
	fmt.Printf(carMap["audi"])
	s, ok := carMap["audi"]
	if ok {
		fmt.Println(s)
	}

}
func main() {
	iterMap()
	//iterByCondition()
	mapSlice()
	valueSlice()
	MapDemo()
}
func iterMap() {
	m := make(map[int]string)
	m[1] = "panamera"
	m[2] = "audi"
	for k, v := range m {
		fmt.Printf("%v:%v\n", k, v)
	}
	//delete(map,k)
	delete(m, 2)
	for k := range m {
		fmt.Printf("%v\n", k)
	}
}
func iterByCondition() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
func mapSlice() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "aladdin"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "北京"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
func valueSlice() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	fmt.Printf("%v,value:%v\n", sliceMap, value)
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
func MapDemo() {
	s := "how do you do"
	slice := strings.Split(s, " ")
	m := make(map[string]int)
	for _, k := range slice {
		_, y := m[k]
		if y {
			m[k] += 1
		} else {
			m[k] = 1
		}
	}
	fmt.Printf("%v\n", m)
}
