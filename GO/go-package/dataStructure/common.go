package datastructure

import "reflect"

func checkDataType(datatype string, value interface{}) bool {
	valType := reflect.TypeOf(value)
	return valType.String() == datatype
}