package main

import "fmt"

/**
array
	长度固定 属于类型属性 不同长度是不一样的
	0 len-1
	自行推断长度
	初始化制定索引值
	for index, value := range a {} 遍历
	二维数组只有第一层可以使用 ... 推导长度
	数组是值类型
	[n]*T表示指针数组，*[n]T表示数组指针 。

*/
func main() {
	arrayInit()
	arrayInit1()
	arrayInit2()
	iterArray()
	array2D()
	a := [...][2]int{{1, 1}, {2, 2}}
	array := changeArray(a)
	fmt.Println(array)
	fmt.Println(a)
	fmt.Println(getArraySum([5]int{1, 3, 5, 7, 8}))
	fmt.Println(getArray([5]int{1, 3, 5, 7, 8}, 8))
}
func arrayInit() {
	test := [3]int{}
	fmt.Println(test)
	//var a [3]int
	//var b [4]int
	//a = b //不可以这样做，因为此时a和b是不同的类型
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
}
func arrayInit1() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}
func arrayInit2() {
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}
func iterArray() {
	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}
func array2D() {
	a := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
	arr := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}
func changeArray(arr [2][2]int) [2][2]int {
	arr[1][1] = 200
	fmt.Println(arr)
	return arr
}

//数组[1, 3, 5, 7, 8]所有元素的和
func getArraySum(arr [5]int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
func getArray(arr [5]int, target int) [4 + 3 + 2 + 1][2]int {
	result := [len(arr) * (len(arr) - 1) / 2][2]int{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				result[i][0] = i
				result[i][1] = j
			}
		}
	}
	return result
}
