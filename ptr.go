package main

import "fmt"

// go 只有 值传递!  引用传递可以使用指针处理
func ptr() int {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
	fmt.Println(pa)	// 指针地址
	fmt.Println(*pa)	// 指针的值
	return a
}

func swap(a, b *int)  {
	*b, *a = *a, *b
}

func swapTwo(a, b int) (int, int)  {
	return b, a
}

func main() {
	fmt.Println(ptr())
	fmt.Println(ptr)

	a, b := 3, 4
	// 把 3, 4 调换
	swap(&a, &b)
	fmt.Println(a, b)

	// 把 4, 3 调换
	a, b = swapTwo(a, b)
	fmt.Println(a, b)
}