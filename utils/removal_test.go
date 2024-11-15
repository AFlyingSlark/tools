package utils

import (
	"fmt"
	"testing"
)

func Test_RemoveSlice(t *testing.T) {
	req := []int64{1, 2, 3, 4, 5, 4, 3, 2, 1}
	//req := []string{"1", "2", "3", "4", "5", "4", "3", "2", "1"}

	result := SliceDistinct(req)

	t.Log(result)
}

func Test_SliceRemoveDuplicate(t *testing.T) {
	s := []string{"1", "1", "3", "5", "3"}

	result := SliceRemoveDuplicate(s)

	fmt.Println(result)
}

func Test_RemoveSpecificBySlice(t *testing.T) {
	s := []string{"1", "1", "3", "5", "3"}

	var param string = "1"

	result := SliceRemoveSpecific(s, param)

	fmt.Println(result)
}
