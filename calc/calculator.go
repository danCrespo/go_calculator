package calculator

import (
	"flag"
	"fmt"
	"math"
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
	Precision string
	Result    float64
}

func NewCalculator() CalculatorInstance {
	c := &Calculator{}
	return c
}

func (c *Calculator) Results(result chan float64) {
	for r := range result {
		c.Result = r
	}
	fmt.Println(c)
}

func (c *Calculator) String() string {
	if c.Precision == "" {
		c.Precision = "%g"
	}
	return fmt.Sprintf("\v Operation result: "+c.Precision+"\n\v", c.Result)
}

func Start(EntriesCh chan []string, PartsCh chan [][]string) {
	signsRegex := regexp.MustCompile(`([\*/\+\-%]){1}`)
	inlineArgRegex := regexp.MustCompile(`(?m)([(\d+(\.\d+)*)]+)([\*/\+\-%\(]*)\b\)*`)

	go func() {
		defer close(EntriesCh)

		var fixedargs []string
		args := flag.Args()

		if len(args) == 1 {
			args = strings.Split(flag.Args()[0], " ")
		}

		inp := strings.Join(args, "")

		if inlineArgRegex.MatchString(inp) {
			group := make([]string, 0)
			for i, arg := range args {
				if strings.Contains(arg, "(") {
					arg = strings.SplitAfter(arg, "(")[1]
					arg = strings.SplitN(arg, ")", -2)[0]
					splitA, splitB, _ := strings.Cut(arg, signsRegex.FindString(arg))
					group = append(group, splitA, signsRegex.FindString(arg), splitB)
					result := FirstOperation(group)
					args = slices.Replace[[]string, string](args, i, i+1, fmt.Sprintf("%g", result))
				}
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
