package utils

import (
	"io"
	"io/ioutil"
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
