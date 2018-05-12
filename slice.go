package main

import "fmt"

func updateSlices(s []int)  {
	// 改变这个slice的下标对应的值 (那么在改变array的值时,可以不使用&传递array地址了,使用slice[:]就可以改变)
	s[1] = 100
}

func main()  {
	arr := [...]int {0, 1, 2, 3, 4, 5, 6, 7, 8}
	s := arr[3:7]	// slice 是对 arr 的view (slice本身没有数据. 是对底层array的一个view)
	fmt.Println("arr[3:7] = ", s) // 半开半闭区间. 前包含 后不包含
	fmt.Println("arr[:] = ", s)

	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	fmt.Println("After updateSlices(s1)")
	updateSlices(s1) // 100 代替了 3  改变slice 就改变了 array
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2) // 0 1 2 100 4
	s2 = s2[2:]
	fmt.Println(s2) // 2 100 4

	// 2 100 4 5
	// slice 可以向后扩展 无法向前扩展 !!!
	fmt.Println("-----------")
	s3 := arr[2:6]	// ptr 展示值
	fmt.Println(s3)
	s4 := s3[3:5]	// 需要理解 view 的含义 (slice本身无数据!!!) 底层数组cap(代表ptr从头到结束.所以不超过cap就可以扩展)
	//s5 := s3[4] // 下标无法取到len以外的值. 越界
	fmt.Println(s4)

	fmt.Println("arr = ", arr)
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n",
		s3, len(s3), cap(s3))
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n",
		s4, len(s4), cap(s4))

	s5 := append(s4, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)
	// 往cap里面替换. 超过cap 系统就会分配一个更大的底层数组.所以必要变量来接受返回值
	fmt.Println("s5, s6, s7 =", s5, s6, s7)
	// s6 s7 不再是 arr 的 view了  是更大的底层的view
	fmt.Println("arr = ", arr)
}