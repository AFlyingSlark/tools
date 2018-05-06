package main

import (
	"math/cmplx"
	"math"
	"fmt"
)

func euler()  {
	// 保留小数后三位.
	fmt.Printf("%.3f \n",
	cmplx.Pow(math.E, 1i * math.Pi) +1)
}

func triangle()  {
	var a, b int = 3, 4
	var c int
	// 类型是强制转换的.
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

const (
	// 常量  非必须大写. 因为在go中 首字母大写 表示的是 public
	filename = "abc.txt"
	a, b  = 3, 4
)

func  enums()  {
	const (
		// 没有特定的枚举类型. 用const块定义. iota 代表自增值
		cpp = iota // 0
		_			// 1
		python		// 2
		golang		// 3
		javascript	// 4
	)
	fmt.Println(cpp, javascript, python, golang)
}
// 没有char类型.  只有rune类型 而且为32位.

func main()  {
	euler()
	triangle()
	enums()
}