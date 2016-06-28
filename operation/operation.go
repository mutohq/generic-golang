package operation

import (
	"errors"
	"fmt"
	"reflect"
)

//Presence to check presence
func Presence(slice interface{}, val interface{}, startoptional ...int) (bool, int, error) {
	var (
		present bool
		errs    error
		index   int
	)
	sv := reflect.ValueOf(slice)
	f := 0
	present = false
	index = -1
	End := sv.Len()
	Start := 0
	errs = nil
	if len(startoptional) == 0 {
		Start = 0
	} else if len(startoptional) == 1 {
		// fmt.Println(len(startoptional))
		Start = startoptional[0]
	} else if len(startoptional) == 2 {
		Start = startoptional[0]
		if startoptional[1] <= End {
			End = startoptional[1]
		}
	} else {
		errs = errors.New("You can't use more than 4 aurguments in presence function")
		return false, -1, errs
	}
	// Check for first argument that must be slice or array
	if reflect.TypeOf(slice).Kind() != reflect.Array && reflect.TypeOf(slice).Kind() != reflect.Slice {
		errs = errors.New("array parameter's type is neither array nor slice")
	}
	// check for type of val is same as that of slice element
	if reflect.TypeOf(slice).Elem() != reflect.TypeOf(val) {
		errs = errors.New("mismatch of slice type and item type is found ")
	}
	// code for checking presence of element in given slice
	// time complexity O(N) and space complexity O(1)
	for i := Start; i < End; i++ {
		if sv.Index(i).Interface() == val {
			// if item is found in slice return true
			present = true
			index = i
			f = 1
			break
		}
	}
	//  Item is not found in slice return false
	if f == 0 {
		present = false
		index = -1
	}
	return present, index, errs
}

// DeleteEle to delete the given element with occurence
func DeleteEle(slice interface{}, item interface{}, optional ...int) (interface{}, int, error) {
	sv := reflect.ValueOf(slice)
	var err error
	Start := 0
	End := sv.Len()

	// fmt.Printf("%T\n", slice)
	// var r_list []interface{}
	// r_list := make([]reflect.TypeOf(item),0)
	// xx := reflect.TypeOf(item)
	// fmt.Printf(xx)
	// Check for first argument that must be slice or array
	if reflect.TypeOf(slice).Kind() != reflect.Array && reflect.TypeOf(slice).Kind() != reflect.Slice {
		err = errors.New("array parameter's type is neither array nor slice\n")
	}
	// check for type of val is same as that of slice element
	if reflect.TypeOf(slice).Elem() != reflect.TypeOf(item) {
		err = errors.New("mismatch of slice type and item type is found \n")
	}
	// fmt.Println("item : ", item)
	fmt.Printf("%T\n", item)
	j := 0
	for i := Start; i < End; i++ {
		// fmt.Printf("i : %T \n", sv.Index(i).Interface())
		if sv.Index(i).Interface() == item {
			continue
		} else {
			tmp := sv.Index(i).Interface()
			fmt.Println(tmp)
			// r_list = append(r_list, tmp)
			// sv.Index(j).Interface() = tmp
			j++
		}
	}
	return 1, 2, err
	// return r_list.([]reflect.TypeOf(item)), End, err
}

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
