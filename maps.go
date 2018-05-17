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

	fmt.Println("----traversing map----")

	// 这个 map 是无序的. 从打印的值 可以看出
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("----Getting values------")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	couName, ok := m["cou"]
	fmt.Println(couName, ok) // 依然可以打印出来. 不过是 empty

	if couName, ok := m["cou"]; ok {
		fmt.Println(couName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("--delete values------")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"] //  需要重新赋值下.  要不然就是上面的值
	fmt.Println(name, ok)
}
