// Code generated by "stringer -type=Type"; DO NOT EDIT.

package etensor

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NULL-0]
	_ = x[BOOl-1]
	_ = x[UINT8-2]
	_ = x[INT8-3]
	_ = x[UINT16-4]
	_ = x[INT16-5]
	_ = x[UINT32-6]
	_ = x[INT32-7]
	_ = x[UINT64-8]
	_ = x[INT64-9]
	_ = x[FLOAT16-10]
	_ = x[FLOAT32-11]
	_ = x[FLOAT64-12]
	_ = x[STRING-13]
	_ = x[INT-14]
}

const _Type_name = "NULLBOOlUINT8INT8UINT16INT16UINT32INT32UINT64INT64FLOAT16FLOAT32FLOAT64STRINGINT"

var _Type_index = [...]uint8{0, 4, 8, 13, 17, 23, 28, 34, 39, 45, 50, 57, 64, 71, 77, 80}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}

func (i *Type) FromString(s string) error {
	for j := 0; j < len(_Type_index)-1; j++ {
		if s == _Type_name[_Type_index[j]:_Type_index[j+1]] {
			*i = Type(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Type")
}
