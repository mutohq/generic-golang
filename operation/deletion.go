package operation

import (
	"errors"
	"reflect"
)

//DeleteEle to delete the given element with occurence
func DeleteEle(slice interface{}, item interface{}, optional ...int) (interface{}, int, error) {
	sv := reflect.ValueOf(slice)
	// a := slice.([]reflect.TypeOf(item))
	// fmt.Println(a)
	var err error
	Start := 0
	End := sv.Len()
	PostEnd := sv.Len()
	deleteCount := -1

	if len(optional) == 0 {
		Start = 0
		deleteCount = -1
	} else if len(optional) == 1 {
		Start = 0
		deleteCount = optional[0]
		if optional[0] < 0 {
			deleteCount = -1
		}
	} else if len(optional) == 2 {
		Start = optional[1]
		deleteCount = optional[0]
		if optional[0] < 0 {
			deleteCount = -1
		}

	} else if len(optional) == 3 {
		Start = optional[1]
		PostEnd = optional[2]
		if PostEnd > End {
			PostEnd = End
		}
		deleteCount = optional[0]
		if optional[0] < 0 {
			deleteCount = -1
		}
	} else {
		err = errors.New("You can't use more than 5 aurguments in DeleteEle function")
		return slice, End, err
	}

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

	for i := 0; i < Start; i++ {
		// copy all elements before the starting index  without any deletion

	}
	// fmt.Println("item : ", item)
	// fmt.Printf("%T\n", item)
	// j := 0
	temp := End
	End = PostEnd
	if deleteCount == -1 {
		for i := Start; i < End; i++ {
			if sv.Index(i).Interface() == item {
				// slice = append(slice[:i], slice[i+1:]...)
				continue // deleting item
			} else {
				// copy the elements

			}
		}
	} else {
		for i := Start; i < End; i++ {
			if deleteCount == 0 {
				// don't delete elements just copy the elements

			} else {
				if sv.Index(i).Interface() == item {
					deleteCount = deleteCount - 1 // deleting the item
				} else {
					// copy the elements

				}
			}
		}
	}
	End = temp
	//copy the elements after last given index
	for i := PostEnd; i < End; i++ {

	}

	return 1, 2, err
	// return r_list.([]reflect.TypeOf(item)), End, err
}
