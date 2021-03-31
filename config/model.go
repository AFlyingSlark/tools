package config

import (
	"encoding/json"

	fuzz "github.com/google/gofuzz"
)

type Base struct {
	HTTP HTTP `json:"http"`
}

type HTTP struct {
	Port      uint `json:"port"`      // 服务端口
	PprofPort uint `json:"pprofPort"` // pprof端口
}

func EchoTmpl(tmpl interface{}) (string, error) {
	f := fuzz.New()

	f.FuzzNoCustom(tmpl)

	result, err := json.Marshal(tmpl)

	if err != nil {
		return "", err
	}

	return string(result), nil
}
