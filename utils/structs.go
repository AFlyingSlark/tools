package utils

import (
	"fmt"
	"reflect"
)

func GetStructFieldValueByName[T comparable](s any, fieldName string) (val T, err error) {
	// 获取传入对象的反射值
	value := reflect.ValueOf(s)

	// 检查是否是指针，如果是指针需要获取其指向的元素
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// 检查是否是结构体
	if value.Kind() != reflect.Struct {
		return val, fmt.Errorf("provided value is not a struct")
	}

	// 获取字段值
	fieldValue := value.FieldByName(fieldName)
	if !fieldValue.IsValid() {
		return val, fmt.Errorf("no such field: %s", fieldName)
	}

	// 检查字段是否可以被导出（是否是大写字母开头）
	// 非导出字段错误: panic: reflect.Value.Interface: cannot return value obtained from unexported field or method [recovered]
	if !fieldValue.CanInterface() {
		return val, fmt.Errorf("cannot access unexported field: %s", fieldName)
	}

	return fieldValue.Interface().(T), nil
}
