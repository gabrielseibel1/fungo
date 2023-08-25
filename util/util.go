package util

import (
	"reflect"
	"strconv"
)

type Data struct {
	V string
	S []Data
}

var (
	Data1 = Data{
		V: "1",
		S: nil,
	}
	Data2 = Data{
		V: "2",
		S: []Data{Data1},
	}
)

type Record struct {
	N int
	S []Record
}

var (
	Record1 = Record{
		N: 1,
		S: []Record{{N: 3, S: nil}},
	}
	Record2 = Record{
		N: 2,
		S: []Record{{N: 4, S: nil}},
	}
)

func PassAll[T any](T) bool { return true }

func PassNo[T any](T) bool { return false }

func Is[T any](t T) func(T) bool {
	return func(u T) bool {
		return reflect.DeepEqual(t, u)
	}
}

func DataToRecord(d Data) Record {
	if d.V == Data1.V {
		return Record1
	} else if d.V == Data2.V {
		return Record2
	} else {
		return Record{}
	}
}

func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
