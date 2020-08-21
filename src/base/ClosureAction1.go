package main

func faa(base int) (func(int) int, func(int) int) {
	println(&base, base)
	add := func(i int) int {
		base += 1
		println(&base, base)
		return base
	}
	sub := func(i int) int {
		base -= 1
		println(&base, base)
		return base
	}
	return add, sub
}
func main() {
	f, g := faa(0)
	s, k := faa(0)
	//
	println(f(1), g(3))
	println(s(1), k(3))
	//0xc00008e000 0
	//0xc00008e008 0
	//0xc00008e000 1
	//0xc00008e000 0
	//1 0
	//0xc00008e008 1
	//0xc00008e008 0
	//1 0
}
