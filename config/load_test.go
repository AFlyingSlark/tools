package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	dir = "../utils"
)

// go test -timeout 300s  github.com/base-tools/config -v -count=1 -run=Test_getDefaultFilePath
func Test_existFilePath(t *testing.T) {
	result, err := existFilePath(dir)

	t.Log(result)
	require.NoError(t, err, "指定路径是否存在")
}

func Test_getDefaultFilePath(t *testing.T) {
	result, err := getDefaultFilePath()

	t.Log(result)
	require.NoError(t, err, "默认路径是否存在")
}

func Test_readConfFile(t *testing.T) {
	result, err := readConfFile("", "")

	t.Log(result)
	require.NoError(t, err, "文件否存在")
}
