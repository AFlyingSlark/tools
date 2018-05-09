package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func convertToBin(n int) string  {
	result := "aa"
	// 不需要括号  初始条件; 结束条件; 递增表达式
	for ; n >0; n /= 2 {
		lsb := n % 2
		// 需要把int 转为 string 才可以相连起来
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		// 报错停止.
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	// 无初始条件 递增表达式 只有递增条件. 那么 ; 也不必要写出
	// 那么这种写法相当于其他语言的while (while语句的一般表达式为：while（表达式）{循环体})
	// go 中 无while循环 使用 for 是一样的效果
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever()  {
	// 无条件也可以运行. 但这是死循环
	// 写法简单. 是因为用的居多.在通讯中
	/*for {
		fmt.Println("abc")
	}*/
}

func main()  {
	fmt.Println(
		convertToBin(5),
		convertToBin(15),
	)

	printFile("abc.txt")

	forever()
}