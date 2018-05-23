package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	s := "Yes我爱我学习!"
	for _, b := range []byte(s) {
		// 英文1字节 中文3字节
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		// 把utf-8 进行了解码. 放入了rune中
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s)  {
		// 这是解码 重新开的内存 来储存 rune 占4个字节
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
