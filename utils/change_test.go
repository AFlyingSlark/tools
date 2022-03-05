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
