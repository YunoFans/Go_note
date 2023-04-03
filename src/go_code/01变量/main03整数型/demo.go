package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// var age int = 299098123123123132
	// fmt.Println(age)

	// var i uint8 = 89
	// fmt.Println(i)

	// fmt.Printf("i的类型是 %U", i)

	var n1 int = 189

	fmt.Printf("n1的类型是%T", n1)

	fmt.Printf("n1的占的大小字节是 %d", unsafe.Sizeof(n1))

}
