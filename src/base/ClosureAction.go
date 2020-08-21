package main

func main() {
	f := fa(1)
	g := fa(1)
	println(f(1))
	println(f(1))
	println(g(1))
	println(g(1))
	//0xc0000180b0 1
	//2
	//0xc0000180b0 2
	//3
	//0xc0000180b8 1
	//2
	//0xc0000180b8 2
	//3
}
func fa(a int) func(i int) int {
	return func(i int) int {
		println(&a, a)
		a = a + 1
		return a
	}
}
