package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func evel(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		// 返回值必要2个变量接收 不想要的可以用 _ 代替
		q, _ := div(a, b)
		return q
	default:
		panic("unsupporten operation: " + op)
	}
}

// 13 / 3 = 4 .. 1
// 多个返回值 需要用 , 隔开 也可以起名称 q int, r int
// 2个返回值 一般第二个用于err
func div(a, b int) (int, int) {
	return a / b, a % b
}

// 参数定义 名写前  类型写后
// 函数式编程 函数可作为参数 (opFunc 类型是func(int, int) int)
// 这的opFunc  和 a, b  是一样为参数的. 所以是有类型的
func apply(opFunc func(int, int) int, a, b int) int {
	p := reflect.ValueOf(opFunc).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args " +
		"(%d, %d) \n", opName, a, b)
	return opFunc(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表 求和
func sumArgs(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}

func main()  {
	fmt.Println(evel(2, 4, "/"))
	fmt.Println(div(5, 4))

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(sumArgs(1, 2, 3, 4, 5))
}
