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
	result, err := readConfFile("test", "conf")

	t.Log(result)
	require.NoError(t, err, "文件否存在")
}

func Test_LoadConfing(t *testing.T) {
	conf := &struct{}{}

	err := LoadConfing(conf, "", "")

	t.Log(conf)
	require.NoError(t, err, "加载配置文件")
}

func Test_allLoadConfig(t *testing.T) {
	type loadData struct {
		title     string
		dir       string
		fileName  string
		expectErr bool
	}

	arguments := []loadData{
		{
			title:     "指定路径.指定文件",
			dir:       "./test",
			fileName:  "conf",
			expectErr: true,
		},
		{
			title:     "指定路径.默认文件",
			dir:       "./conf",
			fileName:  "",
			expectErr: true,
		},
		{
			title:     "默认路径.指定文件",
			dir:       "",
			fileName:  "custom",
			expectErr: true,
		},
		{
			title:     "默认路径.默认文件",
			dir:       "",
			fileName:  "",
			expectErr: true,
		},
	}

	conf := &struct{}{}

	for _, value := range arguments {

		err := LoadConfing(conf, value.dir, value.fileName)

		require.NoError(t, err, value.title)
	}
}
