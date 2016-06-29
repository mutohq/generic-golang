package operation

import (
	"errors"
	"reflect"
)

//PopBack pops the last element from slice/array and it returns new slice with the popped item and error
func PopBack(slice interface{}) (interface{}, interface{}, error) {
	sv := reflect.ValueOf(slice)
	length := sv.Len()
	item := sv.Index(length - 1).Interface()
	// fmt.Printf("%T\n", item)
	// fmt.Println(item)

	var err error
	// Check for first argument that must be slice or array
	if reflect.TypeOf(slice).Kind() != reflect.Array && reflect.TypeOf(slice).Kind() != reflect.Slice {
		err = errors.New("array parameter's type is neither array nor slice\n")
	}

	return 1, item, err
}
