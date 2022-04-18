package utils

import (
	"fmt"
	"testing"
	"time"
)

func Test_diffTime(t *testing.T) {
	startTime := time.Now()
	time.Sleep(2 * time.Second)
	sendTime := DiffNano(startTime)
	sendTimeStr := fmt.Sprintf("%.3f", float64(sendTime)/1e9)
	t.Log(sendTimeStr)
}

func Test_AddTime(t *testing.T) {
	now := LocChange(time.Now(), location)

	fmt.Println("----day-----")
	fmt.Println(AddDay(now, -1).String())
	fmt.Println(AddDay(now, 1).String())

	fmt.Println("----week-----")
	fmt.Println(AddWeek(now, -1).String())
	fmt.Println(AddWeek(now, 1).String())
}
