package main

import (
	"fmt"
	"io"
	"os"
)

func init() {
	defer func() {
		println("first")
	}()
	defer func() { println("second") }()
	println("function body")
}
func main() {
	defer tes(1)
	g()
	w, err := CopyFile("/Users/lgc/Desktop/g/g", "/Users/lgc/Desktop/gt")
	copyFile("/Users/lgc/Desktop/g/gt", "/Users/lgc/Desktop/gt")
	if err != nil {
		println(err)
	}
	println("fileCopy %s", w)
}
func g() int {
	a := 0
	defer func(i int) {
		println("defer i=", i)
	}(a)
	a++
	return a
}
func tes(a int) {
	println(a)
}

func copyFile(dstName, srcName string) (w int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	new, err := os.Create(dstName)
	if err != nil {
		src.Close()
		return
	}
	w, err = io.Copy(new, src)
	new.Close()
	src.Close()
	return
}
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)

	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	defer srcFile.Close()

	//打开dstFileName

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)

}
