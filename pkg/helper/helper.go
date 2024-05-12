package helper

import (
	"strconv"
)

func Str2Bool(data string) (bool, error) {
	boolValue, err := strconv.ParseBool(data)
	if err != nil {
		return false, err
	}
	return boolValue, nil
}
