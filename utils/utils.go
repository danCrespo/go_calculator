package utils

import (
	"fmt"
	"log"
	"slices"
)

func StringToFloat64(s string) float64 {
	var res float64
	if _, err := fmt.Sscanf(s, "%v\n", &res); err != nil {
		log.Printf("error %v\n", err)
	}
	return res
}

func ReplaceSlice(s []string, i, j int, r string) []string {
	s = slices.Replace[[]string, string](s, i, j, r)
	return s
}
