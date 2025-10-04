package utils

import (
	"fmt"
	"strconv"
)

func ParseUintFromString(idString string) (uint, error) {
	id64, err := strconv.ParseUint(idString, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("failed to parse uint from string '%s': %w", idString, err)
	}
	return uint(id64), nil
}

func ParseIntFromString(idString string) (int, error) {
	id64, err := strconv.ParseInt(idString, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("failed to parse uint from string '%s': %w", idString, err)
	}
	return int(id64), nil
}
