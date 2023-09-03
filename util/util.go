package util

import (
	"reflect"
	"strconv"
)

// Data is some data to use in tests
type Data struct {
	V string
	S []Data
}

var (
	// Data1 is one example of Data
	Data1 = Data{
		V: "1",
		S: nil,
	}
	// Data2 is another example of Data
	Data2 = Data{
		V: "2",
		S: []Data{Data1},
	}
)

// Record is another type of data to use in tests
type Record struct {
	N int
	S []Record
}

var (
	// Record1 is an example of Record
	Record1 = Record{
		N: 1,
		S: []Record{{N: 3, S: nil}},
	}
	// Record2 is another example of Record
	Record2 = Record{
		N: 2,
		S: []Record{{N: 4, S: nil}},
	}
)

// PassAll is a function that always returns true for any type and value provided, useful for tests
func PassAll[T any](T) bool { return true }

// PassNo is a function that always returns false for any type and value provided, useful for tests
func PassNo[T any](T) bool { return false }

// Is can be used to know whether an element is equal to another by reflect.DeepEqual
func Is[T any](t T) func(T) bool {
	return func(u T) bool {
		return reflect.DeepEqual(t, u)
	}
}

// DataToRecord gets a Record1 if provided Data1, or Record2 if provided Data2
func DataToRecord(d Data) Record {
	if d.V == Data1.V {
		return Record1
	} else if d.V == Data2.V {
		return Record2
	} else {
		return Record{}
	}
}

// IsInt returns whether a string is an integer representation by strconv.Atoi
func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
