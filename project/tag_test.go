package project

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetTag(t *testing.T) {
	obj := withTagStruct{
		Name: "smiecj",
	}
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	for i := 0; i < val.NumField(); i++ {
		strValue := val.Field(i).Interface().(string)
		tag := typ.Field(i).Tag.Get(tagName)
		fmt.Printf("value: %s, tag: %s\n", strValue, tag)
	}
}
