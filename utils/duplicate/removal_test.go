package duplicate

import "testing"

func Test_int64Slice(t *testing.T) {
	req := []int64{1, 2, 3, 4, 5, 6, 7, 6, 4, 3, 2}
	t.Log(RemoveRespSlice(req))
}
