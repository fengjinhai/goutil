package array

import (
	"reflect"
)

// Distinct returns the unique vals of a slice
//
// [1, 1, 2, 3] >> [1, 2, 3]
func Distinct(arr interface{}) (reflect.Value, bool) {
	// create a slice from our input interface
	slice, ok := takeArg(arr, reflect.Slice)
	if !ok {
		return reflect.Value{}, ok
	}

	// put the values of our slice into a map
	// the key's of the map will be the slice's unique values
	c := slice.Len()
	m := make(map[interface{}]bool)
	for i := 0; i < c; i++ {
		m[slice.Index(i).Interface()] = true
	}
	mapLen := len(m)

	// create the output slice and populate it with the map's keys
	out := reflect.MakeSlice(reflect.TypeOf(arr), mapLen, mapLen)
	i := 0
	for k := range m {
		v := reflect.ValueOf(k)
		o := out.Index(i)
		o.Set(v)
		i++
	}

	return out, ok
}

// Intersect returns a slice of values that are present in all of the input slices
//
// [1, 1, 3, 4, 5, 6] & [2, 3, 6] >> [3, 6]
//
// [1, 1, 3, 4, 5, 6] >> [1, 3, 4, 5, 6]
func Intersect(arrs ...interface{}) (reflect.Value, bool) {
	// create a map to count all the instances of the slice elems
	arrLength := len(arrs)
	var kind reflect.Kind

	tempMap := make(map[interface{}]int)
	for i, arg := range arrs {
		tempArr, ok := Distinct(arg)
		if !ok {
			return reflect.Value{}, ok
		}

		// check to be sure the type hasn't changed
		if i > 0 && tempArr.Index(0).Kind() != kind {
			return reflect.Value{}, false
		}
		kind = tempArr.Index(0).Kind()

		c := tempArr.Len()
		for idx := 0; idx < c; idx++ {
			// how many times have we encountered this elem?
			if _, ok := tempMap[tempArr.Index(idx).Interface()]; ok {
				tempMap[tempArr.Index(idx).Interface()]++
			} else {
				tempMap[tempArr.Index(idx).Interface()] = 1
			}
		}
	}

	// find the keys equal to the length of the input args
	numElems := 0
	for _, v := range tempMap {
		if v == arrLength {
			numElems++
		}
	}
	out := reflect.MakeSlice(reflect.TypeOf(arrs[0]), numElems, numElems)
	i := 0
	for key, val := range tempMap {
		if val == arrLength {
			v := reflect.ValueOf(key)
			o := out.Index(i)
			o.Set(v)
			i++
		}
	}

	return out, true
}

// Union returns a slice that contains the unique values of all the input slices
//
// [1, 2, 2, 4, 6] & [2, 4, 5] >> [1, 2, 4, 5, 6]
//
// [1, 1, 3, 4, 5, 6] >> [1, 3, 4, 5, 6]
func Union(arrs ...interface{}) (reflect.Value, bool) {
	// create a temporary map to hold the contents of the arrays
	tempMap := make(map[interface{}]uint8)
	var kind reflect.Kind

	// write the contents of the arrays as keys to the map. The map values don't matter
	for i, arg := range arrs {
		tempArr, ok := Distinct(arg)
		if !ok {
			return reflect.Value{}, ok
		}

		// check to be sure the type hasn't changed
		if i > 0 && tempArr.Index(0).Kind() != kind {
			return reflect.Value{}, false
		}
		kind = tempArr.Index(0).Kind()

		c := tempArr.Len()
		for idx := 0; idx < c; idx++ {
			tempMap[tempArr.Index(idx).Interface()] = 0
		}
	}

	// the map keys are now unique instances of all of the array contents
	mapLen := len(tempMap)
	out := reflect.MakeSlice(reflect.TypeOf(arrs[0]), mapLen, mapLen)
	i := 0
	for key := range tempMap {
		v := reflect.ValueOf(key)
		o := out.Index(i)
		o.Set(v)
		i++
	}

	return out, true
}

// Difference returns a slice of values that are only present in one of the input slices
//
// [1, 2, 2, 4, 6] & [2, 4, 5] >> [1, 5, 6]
//
// [1, 1, 3, 4, 5, 6] >> [1, 3, 4, 5, 6]
func Difference(arrs ...interface{}) (reflect.Value, bool) {
	// create a temporary map to hold the contents of the arrays
	tempMap := make(map[interface{}]int)
	var kind reflect.Kind

	for i, arg := range arrs {
		tempArr, ok := Distinct(arg)
		if !ok {
			return reflect.Value{}, ok
		}

		// check to be sure the type hasn't changed
		if i > 0 && tempArr.Index(0).Kind() != kind {
			return reflect.Value{}, false
		}
		kind = tempArr.Index(0).Kind()

		c := tempArr.Len()
		for idx := 0; idx < c; idx++ {
			// how many times have we encountered this elem?
			if _, ok := tempMap[tempArr.Index(idx).Interface()]; ok {
				tempMap[tempArr.Index(idx).Interface()]++
			} else {
				tempMap[tempArr.Index(idx).Interface()] = 1
			}
		}
	}

	// write the final val of the diffMap to an array and return
	numElems := 0
	for _, v := range tempMap {
		if v == 1 {
			numElems++
		}
	}
	out := reflect.MakeSlice(reflect.TypeOf(arrs[0]), numElems, numElems)
	i := 0
	for key, val := range tempMap {
		if val == 1 {
			v := reflect.ValueOf(key)
			o := out.Index(i)
			o.Set(v)
			i++
		}
	}

	return out, true
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
