package main

import (
	"fmt"
	"math"
)

//数据类型
func main() {
	/**
	基本数据类型
		int8 (-128 到 127)、
		int16 (-32768 到 32767)、
		int32(-2147483648 到 2147483647)、
		int64(-9223372036854775808 到 9223372036854775807)
		(无符号)
		uint8 (0 到 255)、
		uint16 (0 到 65535)、
		uint32 (0 到 4294967295)、
		uint64 (0 到 18446744073709551615)
	特殊类型
		uint	32位操作系统上就是uint32，64位操作系统上就是uint64
		int		32位操作系统上就是int32，64位操作系统上就是int64
		uintptr	无符号整型，用于存放一个指针
		获取对象的长度的内建len()函数返回的长度可以根据不同平台的字节长度进行变化。
		实际使用中，切片或 map 的元素数量等都可以用int来表示。
		在涉及到二进制传输、读写文件的结构描述时，
		为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用int和 uint 使用int64
	Number literals syntax
	方便 二进制 十进制 十六进制
		v := 0b00101101， 代表二进制的 101101，相当于十进制的 45。
		v := 0o377，代表八进制的 377，相当于十进制的 255。
		v := 0x1p-2，代表十六进制的 1 除以 2²，也就是 0.25。
		v:=0x1p2 4
		而且还允许用 _ 来分隔数字，比如说：
		v := 123_456 等于 123456。
	浮点型
		float32 精确位数小数点后7位
		float64 精确位数小数点后15位
	复数 complex64 实部 虚步都是32位 complex128 实部 虚步都是64位
	布尔值
		布尔类型变量的默认值为false。
		Go 语言中不允许将整型强制转换为布尔型.
		布尔型无法参与数值运算，也无法与其他类型进行转换。
	字符串
		内部实现是 utf8 struct
		en(str)	求长度+或fmt.Sprintf	拼接字符串
		strings.Split	分割
		strings.contains	判断是否包含
		strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
		strings.Index(),strings.LastIndex()	子串出现的位置
		strings.Join(a[]string, sep string)	join操作
	要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	转义字符
		\r	回车符（返回行首）
		\n	换行符（直接跳到下一行的同列位置）
		\t	制表符
		\'	单引号
		\"	双引号
		\\	反斜杠
	多行字符串
		s:`类似java
	       中的
		`
	byte rune 类型
		uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
		rune类型，代表一个 UTF-8字符。 int32 处理中文 日文
	类型转换（强转）
		T(表达式)
		T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.
	*/

	changeString()
	Sqrt()
}
func numberLiteralsSyntax() {
	v := 0x1p2
	fmt.Println(v)
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	aa := 010 //八进制
	fmt.Printf("%d \n", aa)
	fmt.Printf("%b \n", aa)
	fmt.Printf("%o \n", aa)
	fmt.Printf("%x \n", aa)

	// 八进制  以0开头
	var b int = 0o77
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF
}

func test() {
	println(math.MaxFloat32)
	println(math.Pi)

	c1 := 1 + 1i
	println(c1)
	s := `类似java
	    中的
		`
	fmt.Printf(s)
}
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

func Sqrt() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}