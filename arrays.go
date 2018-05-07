package main

import "fmt"

// 数组是值类型	go 只有一种类型就是值类型.!!!
// 类型 []int 是切片
func printArray(arr [5]int)  {
	arr[0] = 100	// 内部改变. 这是拷贝的数组
	// 遍历函数 range
	// _ 可以省略变量
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray2(arr *[5]int)  {
	arr[0] = 100	// 内部改变. 这是拷贝的数组
	// 遍历函数 range
	// _ 可以省略变量
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	// 定长数组
	var  arr1 [5]int
	arr2 := [3]int {1, 3, 5}
	// 系统识别长度数组
	arr3 := [...]int {2, 4, 6, 8, 10}

	var grid  [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	printArray(arr1)
	//printArray(arr2) 长度不符合
	printArray(arr3)
	// 外部无变化
	fmt.Println(arr1, arr3)

	// 当然也可以使用指针方式,那么外部数组将会改变
	printArray2(&arr1)
	printArray2(&arr3)
	fmt.Println(arr1, arr3)
}