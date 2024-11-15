package utils

import (
	"fmt"
	"testing"
)

// 定义一个结构体
type Person struct {
	Name   string
	Age    int
	Gender string
}

func Test_structs(t *testing.T) {
	// 初始化结构体
	p := Person{Name: "Alice", Age: 30, Gender: "Female"}

	// 根据字段名获取值
	field := "Name"
	strValue, err := GetStructFieldValueByName[string](p, field)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Field '%s' value: %v\n", field, strValue)
	}

	// 获取另一个字段
	field = "Age"
	intValue, err := GetStructFieldValueByName[int](p, field)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Field '%s' value: %v\n", field, intValue)
	}

	// 获取不存在的字段
	field = "NonExistent"
	notValue, err := GetStructFieldValueByName[string](p, field)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Field '%s' value: %v\n", field, notValue)
	}
}
