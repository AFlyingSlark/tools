package main

import "fmt"

func printSlice(s []int)  {
	// 每当cap 装不下时.就自身长度 * 2
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func main()  {
	var s []int	// 定义一个s的0值 为 nil 的 slice

	for i := 0; i < 25; i++ {
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)
}