package calendar

import "strconv"

func stringToInt32(order string) (int32, error) {
	val64, err := strconv.ParseInt(order, 10, 32)
	if err != nil {
		return 0, err
	}

	val32 := int32(val64)
	return val32, nil
}
