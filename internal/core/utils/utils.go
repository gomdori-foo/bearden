package utils

import (
	"reflect"
	"runtime"
	"strings"
)

func IsConstructor(constructor interface{}) bool {
	constructorType := reflect.TypeOf(constructor)
	constructorValue := reflect.ValueOf(constructor)
	constructorName := runtime.FuncForPC(constructorValue.Pointer()).Name()

	if lastDot := strings.LastIndex(constructorName, "."); lastDot >= 0 {
		constructorName = constructorName[lastDot+1:]
	}

	return constructorType.Kind() == reflect.Func && strings.HasPrefix(constructorName, "New")
}