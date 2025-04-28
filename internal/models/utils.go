package models

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	errorLoadingJSON = errors.New("error loading JSON from disk.")
)

func ConvertInterfaceArrayToStringArray(aInterface []interface{}) []string {
	aString := make([]string, len(aInterface))
	for i, v := range aInterface {
		aString[i] = v.(string)
	}
	return aString
}

func ObjectAssign(target interface{}, object interface{}) {
	// object atributes values in target atributes values
	// using pattern matching (https://golang.org/pkg/reflect/#Value.FieldByName)
	// https://stackoverflow.com/questions/35590190/how-to-use-the-spread-operator-in-golang
	t := reflect.ValueOf(target).Elem()
	o := reflect.ValueOf(object).Elem()
	for i := 0; i < o.NumField(); i++ {
		for j := 0; j < t.NumField(); j++ {
			if t.Field(j) == o.Field(i) {
				fmt.Printf("Field %s is equal to %s\n", t.Field(j), o.Field(i))
				t.Field(j).Set(o.Field(i))
			}
		}
	}
}

func IsStructEmpty(object interface{}) (bool, error) {
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		// and create an empty copy of the struct object to compare against
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, nil
}
