package utils

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"slices"
	"strings"
)

var (
	SignsRegex           = regexp.MustCompile(`([\*/\+\-%\^]){1}`)
	ParenthesisRegex     = regexp.MustCompile(`([(\)]){1}`)
	DigitRegex           = regexp.MustCompile(`(\b[\*/\+\-%\^\s]*)([\d]+\.?[\d]*)\b`)
	negativeNumberRegex  = regexp.MustCompile(`[\*/\+\-%\^]-[\d]+\.?[\d]*`)
	negativeNumberRegex2 = regexp.MustCompile(`^-[\d]+\.?[\d]*`)
)

func StringToFloat64(s string) float64 {
	var res float64
	if _, err := fmt.Sscanf(s, "%v", &res); err != nil {
		log.Printf("error %v", err)
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

func StringContains(value string, patterns ...string) bool {
	for _, pattern := range patterns {
		for char := 0; char < len(value); char++ {
			if string(value[char]) == pattern {
				return true
			}
		}
	}
	return false
}

func CheckForNegativeNumber(exp string, negativeNumbers *[]string) {
	negativeNums := negativeNumberRegex.FindAllString(exp, -1)
	negativeNums2 := negativeNumberRegex2.FindAllString(exp, -1)
	if len(negativeNums) > 0 || len(negativeNums2) > 0 {
		*negativeNumbers = append(*negativeNumbers, negativeNums...)
		*negativeNumbers = append(*negativeNumbers, negativeNums2...)
	}
}

func CreateSlice(e []string) (elements []string) {
	str := strings.Join(e, "")
	totalDigits := DigitRegex.FindAllString(str, -1)
	match, matches := StringContainsAndWhatContains(str, "^", "*", "/", "+", "-")
	if match && (len(totalDigits) == len(matches)) && len(matches) == 1 {
		elements = e
		return
	}

	finalStr := ""
	for i := 0; i < len(str); i++ {
		if StringContains(string(str[i]), "^", "*", "/", "+", "-") {
			if i+1 < len(str) && SignsRegex.MatchString(string(str[i+1])) {
				finalStr += " " + string(str[i]) + " "
			} else if i+1 < len(str) && i-1 > 0 && (SignsRegex.MatchString(string(str[i-1])) && DigitRegex.MatchString(string(str[i+1])) || string(str[i-1]) == "" && DigitRegex.MatchString(string(str[i+1]))) {
				finalStr += " " + string(str[i])
			} else if i == 0 && str[i] == '-' && DigitRegex.MatchString(string(str[i+1])) {
				finalStr += string(str[i])
			} else {
				finalStr += " " + string(str[i]) + " "
			}
		} else {
			finalStr += string(str[i])
		}
	}
	elements = strings.Fields(finalStr)
	return
}

func StringContainsAndWhatContains(value string, patterns ...string) (bool, []string) {
	result, patternsMatched := false, make([]string, 0)
	for _, pattern := range patterns {
		for char := 0; char < len(value); char++ {
			if string(value[char]) == pattern {
				result = true
				patternsMatched = append(patternsMatched, pattern)
			}
		}
		if len(patternsMatched) == 0 {
			if strings.Contains(value, pattern) {
				result = true
				patternsMatched = append(patternsMatched, pattern)
			}
		}
	}
	slices.Sort[[]string, string](patternsMatched)
	return result, patternsMatched
}

func SlicesContainsAndWhatContains(s []any, patterns ...string) (bool, []string) {
	result, patternsMatched := false, make([]string, 0)
	for _, pattern := range patterns {
		if slices.Contains[[]any, any](s, pattern) {
			result = true
			patternsMatched = append(patternsMatched, pattern)
		}
	}
	return result, patternsMatched
}

func SlicesContains(s []any, patterns ...string) bool {
	for _, pattern := range patterns {
		if slices.Contains[[]any, any](s, pattern) {
			return true
		}
	}
	return false
}

func PrepareArguments(args []string) []string {
	if len(args) == 1 {
		args = strings.Split(args[0], "")
	}

	inp := strings.Join(args, "")
	var input string

	for i := 0; i < len(inp); i++ {
		if inp[i] != ' ' {
			if inp[i] == '-' && (inp[i-1] != ' ' && !ParenthesisRegex.Match([]byte(string(inp[i-1])))) && (inp[i+1] != ' ' && !ParenthesisRegex.Match([]byte(string(inp[i+1])))) {
				input += " " + string(inp[i]) + " "
			} else {

				if SignsRegex.Match([]byte(string(inp[i]))) {
					if inp[i] == '-' && inp[i+1] != ' ' && !ParenthesisRegex.Match([]byte(string(inp[i+1]))) {
						input += " " + string(inp[i])
					} else {

						input += " " + string(inp[i]) + " "
					}
				} else {
					input += string(inp[i])
				}
			}
		}
	}

	fields := strings.Fields(input)
	return fields
}
