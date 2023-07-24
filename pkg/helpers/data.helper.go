package helpers

import "strconv"

func ConvertStringToUInt64(data string) (uint64) {
	number, _ := strconv.ParseUint(data, 10, 64)
	return number
}
