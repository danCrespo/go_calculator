package utils

import (
	"fmt"
	"log"
	"math"
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

func DegreesToRadians(degrees float64) float64 {
	result := degrees * (math.Pi / 180)
	return result
}
