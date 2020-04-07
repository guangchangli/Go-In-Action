package main

import "fmt"

/**
recover()必须搭配defer使用。
defer一定要在可能引发panic的语句之前定义。
*/
func main() {

	//fmt.Println(f11())
	//fmt.Println(f22())
	//fmt.Println(f33())
	//fmt.Println(f44())
	funcA()
	funcB()
	funcC()
}
func f11() int {
	x := 5
	defer func() {
		x++
		fmt.Printf("f11")
	}()
	return x //5
}

func f22() (x int) {
	defer func() {
		x++
		fmt.Printf("f22")
	}()
	return 5
}

func f33() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f44() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
