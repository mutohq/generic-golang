package operation

import (
	"errors"
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
