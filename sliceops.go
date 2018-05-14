package main

import "fmt"

func printSlice(s []int)  {
	// 每当cap 装不下时.就自身长度 * 2
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main()  {
	fmt.Println("Creating slice")
	var s []int	// 定义一个s的0值 为 nil 的 slice

	for i := 0; i < 25; i++ {
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	s1 := []int {2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 8, 8)
	s3 := make([]int,10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("copying slice")
	copy(s2, s1)
	println(s2)

	fmt.Println("deleteing elements from slice")
	// 可变参数 加上 ...
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}