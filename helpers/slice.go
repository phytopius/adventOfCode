package adventOfCode

import (
	"fmt"
	"strconv"
)

func ConvertStringSliceToIntSlice(stringSlice []string) ([]int, error) {
	var intSlice []int
	for _, str := range stringSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to int: %v", str, err)
		}
		intSlice = append(intSlice, num)
	}
	return intSlice, nil
}

func FindIndex[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}
