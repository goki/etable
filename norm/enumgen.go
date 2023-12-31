// Code generated by "goki generate"; DO NOT EDIT.

package norm

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _StdNormsValues = []StdNorms{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

// StdNormsN is the highest valid value
// for type StdNorms, plus one.
const StdNormsN StdNorms = 12

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _StdNormsNoOp() {
	var x [1]struct{}
	_ = x[L1-(0)]
	_ = x[L2-(1)]
	_ = x[SumSquares-(2)]
	_ = x[N-(3)]
	_ = x[Sum-(4)]
	_ = x[Mean-(5)]
	_ = x[Var-(6)]
	_ = x[Std-(7)]
	_ = x[Max-(8)]
	_ = x[MaxAbs-(9)]
	_ = x[Min-(10)]
	_ = x[MinAbs-(11)]
}

var _StdNormsNameToValueMap = map[string]StdNorms{
	`L1`:         0,
	`l1`:         0,
	`L2`:         1,
	`l2`:         1,
	`SumSquares`: 2,
	`sumsquares`: 2,
	`N`:          3,
	`n`:          3,
	`Sum`:        4,
	`sum`:        4,
	`Mean`:       5,
	`mean`:       5,
	`Var`:        6,
	`var`:        6,
	`Std`:        7,
	`std`:        7,
	`Max`:        8,
	`max`:        8,
	`MaxAbs`:     9,
	`maxabs`:     9,
	`Min`:        10,
	`min`:        10,
	`MinAbs`:     11,
	`minabs`:     11,
}

var _StdNormsDescMap = map[StdNorms]string{
	0:  ``,
	1:  ``,
	2:  ``,
	3:  ``,
	4:  ``,
	5:  ``,
	6:  ``,
	7:  ``,
	8:  ``,
	9:  ``,
	10: ``,
	11: ``,
}

var _StdNormsMap = map[StdNorms]string{
	0:  `L1`,
	1:  `L2`,
	2:  `SumSquares`,
	3:  `N`,
	4:  `Sum`,
	5:  `Mean`,
	6:  `Var`,
	7:  `Std`,
	8:  `Max`,
	9:  `MaxAbs`,
	10: `Min`,
	11: `MinAbs`,
}

// String returns the string representation
// of this StdNorms value.
func (i StdNorms) String() string {
	if str, ok := _StdNormsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the StdNorms value from its
// string representation, and returns an
// error if the string is invalid.
func (i *StdNorms) SetString(s string) error {
	if val, ok := _StdNormsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _StdNormsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type StdNorms")
}

// Int64 returns the StdNorms value as an int64.
func (i StdNorms) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the StdNorms value from an int64.
func (i *StdNorms) SetInt64(in int64) {
	*i = StdNorms(in)
}

// Desc returns the description of the StdNorms value.
func (i StdNorms) Desc() string {
	if str, ok := _StdNormsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// StdNormsValues returns all possible values
// for the type StdNorms.
func StdNormsValues() []StdNorms {
	return _StdNormsValues
}

// Values returns all possible values
// for the type StdNorms.
func (i StdNorms) Values() []enums.Enum {
	res := make([]enums.Enum, len(_StdNormsValues))
	for i, d := range _StdNormsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type StdNorms.
func (i StdNorms) IsValid() bool {
	_, ok := _StdNormsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i StdNorms) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *StdNorms) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}
