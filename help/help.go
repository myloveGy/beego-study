package help

import (
	"reflect"
	"strconv"
)

func MapArray(arr []interface{}, key, val string) map[string]interface{} {
	m := make(map[string]interface{})
	for i, iL := 0, len(arr); i < iL; i++ {
		tmp := reflect.ValueOf(arr[i])
		id := strconv.FormatInt(tmp.FieldByName(key).Int(), 10)
		m[id] = tmp.FieldByName(val).Interface()
	}

	return m
}
