package calculator

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const (
	plus     = "+"
	minus    = "-"
	multiply = "*"
	division = "/"
	module   = "%"
)

type CalculatorInstance interface {
	Results(result chan float64)
}

type Calculator struct {
	Precision *CalculatorPrecision
	Result    float64
}

type CalculatorPrecision string

func NewCalculator(precision *CalculatorPrecision) CalculatorInstance {
	c := &Calculator{Precision: precision}
	return c
}

func (c *Calculator) Results(result chan float64) {
	for r := range result {
		c.Result = r
		fmt.Fprintf(os.Stdout, "\v \033[01;05;32mResult: \033[01;05;36m "+string(*c.Precision)+"\033[00m\n\v", c.Result)
	}
}

func (c CalculatorPrecision) String() string {
	return string(c)
}

func Start(EntriesCh chan []string, PartsCh chan [][]string, args []string) {
	// signsRegex := regexp.MustCompile(`([\*/\+\-%]){1}`)
	inlineArgRegex := regexp.MustCompile(`(?m)([(\d+(\.\d+)*)]+)([\*/\+\-%\(]*)\b\)*`)

	go func() {
		defer close(EntriesCh)

		var fixedargs []string

		if len(args) == 1 {
			args = strings.Split(args[0], " ")
		}

		inp := strings.Join(args, "")
		if inlineArgRegex.MatchString(inp) {
			group := make([]string, 0)
			for i, arg := range args {
				var result float64
				if strings.Contains(arg, "(") {
					firstDigit := strings.SplitAfter(arg, "(")[1]
					secondDigit := strings.SplitN(args[i+2], ")", -1)[0]
					group = append(group, firstDigit, args[i+1], secondDigit)
					result = FirstOperation(group)
					args = slices.Replace[[]string, string](args, i, i+3, fmt.Sprintf("%g", result))
				}
				group = slices.Delete[[]string](group, 0, len(group))
			}
			fixedargs = append(fixedargs, args...)
		}
		EntriesCh <- fixedargs
	}()

	go func() {
		parts := make([][]string, 0)

		defer close(PartsCh)

		for arg := range EntriesCh {
			if len(arg) == 3 {
				parts = append(parts, arg)
			} else {
				for i := 3; i < len(arg); i++ {
					if i == 3 {
						parts = append(parts, arg[:i])
					}
					parts = append(parts, arg[i:])
				}
			}
			PartsCh <- parts
		}
	}()
}

func Operations(PartsCh chan [][]string, result chan float64) {
	var (
		part1 []string
		part2 []string
		res1  float64
	)

	defer close(result)

	for p := range PartsCh {
		part1 = p[0]
		if len(p) > 1 && len(p[1]) > 1 {
			part2 = p[1]
		}
	}

	// fmt.Println("parts1", part1)
	// fmt.Println("parts2", part2)

	res1 = FirstOperation(part1)

	for i, p := range part2 {
		if p == "" {
			break
		}

		switch p {
		case plus:
			res1 += fatoi64(part2[i+1])
		case minus:
			res1 -= fatoi64(part2[i+1])
		case division:
			res1 /= fatoi64(part2[i+1])
		case multiply:
			res1 *= fatoi64(part2[i+1])
		case module:
			res1 = float64(int(res1) % atoi(part2[i+1]))
		}
	}
	result <- res1
}

var fatoi64 = func(s string) float64 {
	var res float64
	if _, err := fmt.Sscanf(s, "%v", &res); err != nil {
		fmt.Printf("error %v\n", err)
	}
	return res
}

var atoi = func(s string) int {
	d, _ := strconv.Atoi(s)
	return d
}

func FirstOperation(elements []string) float64 {
	result := 0.0

	for range elements {
		sign := elements[1]
		switch sign {
		case plus:
			result = fatoi64(elements[0]) + fatoi64(elements[2])
		case minus:
			result = fatoi64(elements[0]) - fatoi64(elements[2])
		case division:
			result = fatoi64(elements[0]) / fatoi64(elements[2])
		case multiply:
			result = fatoi64(elements[0]) * fatoi64(elements[2])
		case module:
			res := atoi(elements[0]) - int(math.Floor(fatoi64(elements[0])/fatoi64(elements[2])))*atoi(elements[2])
			result = float64(res)
		}
	}

	return result
}
