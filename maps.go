package main

import "fmt"

func main()  {
	m := map[string] string {
		"name" : "ccmouse",
		"course" :"golang",
		"site" :"imooc",
		"quality":"notbad",
	}

	m2 := make(map[string] int) // m2 == empty map
	var m3 map[string] int // m3 = nil	 nil 可以运算.不想其他语言的null

	fmt.Println(m, m2, m3)

	fmt.Println("traversing map")

	// 这个 map 是无序的. 从打印的值 可以看出
	for k, v := range m {
		fmt.Println(k, v)
	}
}
