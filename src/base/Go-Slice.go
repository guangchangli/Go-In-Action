package main

import (
	"fmt"
	"sort"
)

/**
Slice
	相同类型元素的可变长度的序列,基于数组的封装
	引用类型 内部结构包括 地址 长度 容量 只能和 nil 比较
	define var name []Type
		[>=low:<height] 再切片 high 上限是容量不是长度
		a[2:]  // 等同于 a[2:len(a)]
		a[:3]  // 等同于 a[0:3]
		a[:]   // 等同于 a[0:len(a)]
		a[low : high : max] low 可以省略 0 cap=max-low
	判断是否为空
		len(s)==0
	比较
		可以和 nil 比较 nil 底层没有数组 nil值的切片 len cap 都是 0
	赋值拷贝
		引用类型 值变化
		copy(dest,src) 复制内容到另一个新的slice
	append 添加 并返回切片 二倍扩容 声明后可以添加
	扩容针对不同的数据类型处理不一样
		首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
		否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
		否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，
		即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
		如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
	删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)

*/
func init() {
	//define slice
	var s []string
	s1 := [4]string{"slice define", "123", "456"}
	//[>=low:<height]
	//a[2:]  // 等同于 a[2:len(a)]
	//a[:3]  // 等同于 a[0:3]
	//a[:]   // 等同于 a[0:len(a)]
	s2 := s1[:2]
	fmt.Printf("%s\n", s)
	fmt.Println(s1)
	fmt.Println(len(s), cap(s))
	fmt.Printf("s1:%v ,len(s1): %v,cap(s1): %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v ,len(s2): %v,cap(s2): %v,pointer: %v\n", s2, len(s2), cap(s2), &s2)
	s3 := s2[:3]
	fmt.Printf("s3:%v ,len(s3): %v,cap(s3): %v,pointer: %v\n", s3, len(s3), cap(s3), &s3)
	/**
	s1:[slice define 123 456 ] ,
		len(s1): 4,
		cap(s1): 4
	s2:[slice define 123] ,
		len(s2): 2,
		cap(s2): 4,
		pointer: &[slice define 123]
	再切片 high 上限是容量不是长度
	s3:[slice define 123 456] ,
		len(s3): 3,
		cap(s3): 4,
		pointer: &[slice define 123 456]

	*/
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
	println("init out -----------")
}
func main() {
	makeSlice()
	copySlice()
	deleteElement()
	demo()
	sortSlice()
}
func makeSlice() {
	s := make([]string, 3, 5)
	s = append(s, "append")
	fmt.Printf("s:%v, len(s): %v, cap(s): %v\n", s, len(s), cap(s))
	s1 := s
	fmt.Printf("s1:%v, len(s1): %v, cap(s1): %v\n", s1, len(s1), cap(s1))
	s1[1] = "change"
	fmt.Printf("s:%v, len(s): %v, cap(s): %v\n", s, len(s), cap(s))
	fmt.Printf("s1:%v, len(s1): %v, cap(s1): %v\n", s1, len(s1), cap(s1))
}

func copySlice() {
	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]
}
func deleteElement() {
	s := []int{1, 2, 3, 4, 5}
	s = append(s[:3], s[4:]...)
	fmt.Printf("%v\n", s)
}

func demo() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Printf("a: len(a):%v,cap(a):%v\n", len(a), cap(a))
}

func sortSlice() {
	var a = [...]int{3, 7, 8, 9, 1}
	//s:=make([]int, len(a))
	//for i:=0;i<len(a) ;i++  {
	//	s= append(s, a[i])
	//}
	//sort.Ints(s)
	//s= append(s[len(a):])
	//fmt.Printf("%v",s)
	// fmt.Println(a[:]) //将数组切片
	//sort包内部实现了内部数据类型的排序
	//升序排列
	sort.Ints(a[:])
	fmt.Println(a)
	//降序排列
	sort.Sort(sort.Reverse(sort.IntSlice(a[:])))
	fmt.Println(a)
}
