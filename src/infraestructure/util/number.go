package util

import "strconv"

func ConvertStringToInt64(number string) (int64, error) {

	integer, err := strconv.Atoi(number)
	if err != nil {
		return 0, err
	}

	return int64(integer), nil
}
