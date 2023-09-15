package utils

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"slices"
	"strings"
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
		if strings.Contains(value, pattern) {
			return true
		}
	}
	return false
}

func CreateSlice(e []string) (elements []string) {

	str := strings.Join(e, "")
	negativeNumberRegex := regexp.MustCompile(`[\*/\+\-%\^]-[\d]+\.[\d]*`)
	digitRegex := regexp.MustCompile(`[\d]+\.?[\d]*`)
	totalDigits := digitRegex.FindAllString(str, -1)
	negativeNumber := negativeNumberRegex.FindString(str)
	negativeNumberInd := strings.Index(str, negativeNumber)
	charArray := ""

	if negativeNumber != "" {
		str = negativeNumberRegex.ReplaceAllString(str, "")

	}

	for char := 0; char < len(str); char++ {
		if negativeNumber != "" && char == negativeNumberInd-1 {
			charArray += string(str[char]) + " " + string(negativeNumber[0]) + " " + negativeNumber[1:]
		}
		if StringContains(string(str[char]), "^", "*", "/", "+", "-") && len(totalDigits) > 1 {
			charArray += " " + string(str[char]) + " "
		} else {
			charArray += string(str[char])
		}
	}

	elements = strings.Fields(charArray)
	return
}

func StringContainsAndWhatContains(value string, patterns ...string) (bool, []string) {
	result := false
	patternsMatched := make([]string, 0)
	for i, pattern := range patterns {
		if strings.Contains(value, pattern) {
			result = true
			patternsMatched = append(patternsMatched, pattern)
		}

		if i == len(patterns)-1 {
			break
		}
		continue
	}

	slices.Sort[[]string, string](patternsMatched)
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

var (
	SignsRegex       = regexp.MustCompile(`([\*/\+\-%\^]){1}`)
	ParenthesisRegex = regexp.MustCompile(`([(\)]){1}`)
)

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
