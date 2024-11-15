package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// 针对http读取body信息
func GetStrByHttpBody(r io.ReadCloser) string {
	buf := make([]byte, 1024)
	n, _ := r.Read(buf)
	return string(buf[0:n])
}

// 针对reader读取数据
func GetStrByReader(r io.Reader) string {
	body, _ := ioutil.ReadAll(r)
	return string(body)
}

// 定义一个泛型函数，将切片合并为字符串
func Join[T any](arr []T, sep string) string {
	var strArr []string
	for _, item := range arr {
		strArr = append(strArr, fmt.Sprintf("%v", item))
	}
	return strings.Join(strArr, sep)
}

/*
不会修改底层数组的容量（即底层数组仍然存在原始的第一个元素，只是不会被访问到）。
如果你想真正释放切片中被移除的元素（尤其在长切片中节省内存），可以将不需要的部分切片置空，例如：
```go
// 获取队列的第一个元素
 num := nums[0]
 // 移除第一个元素
nums = nums[1:]
nums = append([]int(nil), nums...) // 拷贝新切片以去掉原始引用
```
这样可以确保移除的元素不会再被引用，允许 Go 的垃圾回收机制回收这些元素的内存
*/
