package models

import (
	"errors"
	"fmt"
	"reflect"
)

type Model interface {
	PK() string
	TableName() string
}

func GetModel(model interface{}) (Model, error) {
	if m, ok := model.(Model); ok {
		return m, nil
	}

	mValue := reflect.ValueOf(model)
	fmt.Printf("%#v\n", mValue)
	if mValue.Kind() != reflect.Ptr {
		return nil, errors.New(fmt.Sprintf("model argument must pass a pointer, not a value %#v", model))
	}

	if mValue.IsNil() {
		return nil, errors.New(fmt.Sprintf("%#v cannot be nil pointer passed", model))
	}

	if mValue.Elem().Type().Kind() != reflect.Slice {
		return nil, errors.New(fmt.Sprintf("model argument must slice, but get %#v", model))
	}

	iModel := reflect.New(mValue.Elem().Type().Elem()).Elem().Interface()
	if m, ok := iModel.(Model); ok {
		return m, nil
	}

	return nil, errors.New(fmt.Sprintf("%#v need to implement Model", model))
}
