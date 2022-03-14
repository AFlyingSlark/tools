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

func Test_xxx(t *testing.T) {
	source := map[string]interface{}{
		"name":    "xx",
		"UserAge": 18,
	}

	result, _ := MapToStructByJson(source, &TempStruct{})

	fmt.Println(result)
}

func Test_SliceRemoveDuplicate(t *testing.T) {
	s := []interface{}{"1", "1", "3", "5", "3"}

	result := SliceRemoveDuplicate(s)

	fmt.Println(result)
}
