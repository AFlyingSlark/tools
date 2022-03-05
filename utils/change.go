package utils

import (
	"bytes"
	"encoding/json"
	"errors"
)

// 根据judge的值返回结果,true:ok false:fail
func BoolToResult(judge bool, ok, fail interface{}) interface{} {
	if !judge {
		return fail
	}

	return ok
}

// 将指定map转为具有jsonTag的指针struct
func MapToStructByJson(source map[string]interface{}, target interface{}) (interface{}, error) {
	if len(source) == 0 {
		return nil, errors.New("源map为空")
	}

	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(source); err != nil {
		return nil, err
	}

	if err := json.NewDecoder(buffer).Decode(target); err != nil {
		return nil, err
	}

	return target, nil
}
