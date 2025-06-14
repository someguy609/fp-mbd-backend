package utils

import "fmt"

func StringToUint(s string) (uint, error) {
	var result uint
	n, err := fmt.Sscanf(s, "%d", &result)
	if n != 1 || err != nil {
		return 0, fmt.Errorf("invalid string to uint conversion: %s", s)
	}
	return result, nil
}
