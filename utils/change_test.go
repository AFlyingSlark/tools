package utils

import (
	"fmt"
	"testing"
)

type TempStruct struct {
	Name    string `json:"name"`
	UserID  int64
	UserAge int64
}

func Test_MapToStructByJson(t *testing.T) {
	source := map[string]interface{}{
		"name":    "xx",
		"UserAge": 18,
	}

	result, _ := MapToStructByJson(source, &TempStruct{})

	fmt.Println(result)
}

func Test_SliceIntersection(t *testing.T) {
	first := []string{"1", "2", "3", "4", "5"}
	last := []string{"1", "3", "5", "7"}

	result := SliceIntersection(first, last)

	fmt.Println(result)
}
