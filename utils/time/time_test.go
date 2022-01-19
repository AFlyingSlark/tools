package times

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
