package main

import (
	"io/ioutil"
	"fmt"
)

func grade(score int) string  {
	g := ""
	// switch 无需使用break 也可以不接条件 相当于其他语言的switch(true){}
	switch {
	case score < 0 || score > 100:
		// 注意使用的打印函数
		panic(fmt.Sprintf("wrong score : %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main()  {
	// 这个abc.txt 文件查找的是gopath 下面的src级目录.非当前运行文件目录中的abc
	const filename  = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(82),
		grade(98),
		grade(100),
		//grade(-100),
	)
}