package trigonometry

import (
	"fmt"
	"math"
	"strings"

	"github.com/danCrespo/go_calculator/arithmetic"
	"github.com/danCrespo/go_calculator/utils"
	"golang.org/x/exp/slices"
)

const (
	Sin      = "sin"
	Sinh     = "sinh"
	Cos      = "cos"
	Cosh     = "cosh"
	Tan      = "tan"
	Cot      = "cot"
	Sec      = "sec"
	Cosec    = "cosec"
	Hypot    = "hypot"
	RootN    = "rootN"
	PotencyN = "^N"
)

type (
	TrigonometryArea interface {
		sinFn(degrees float64) float64
		sinHFn(degrees float64) float64
		cosineFn(degrees float64) float64
		cosineHFn(degrees float64) float64
		tangentFn(degrees float64) float64
		cotangentFn(degrees float64) float64
		secantFn(degrees float64) float64
		cosecantFn(degrees float64) float64
		hypotenuseFn(catheiA, catheiB float64) float64
		rootNFn(value float64, root float64) float64
		potencyNFn(value, exponent float64) float64
		Calculate(args []string) float64
	}

	trigonometry struct{ arithmetic arithmetic.ArithmeticArea }
)

func Trigonometry(a arithmetic.ArithmeticArea) TrigonometryArea {
	t := trigonometry{a}
	return &t
}

func (t *trigonometry) Calculate(elements []string) float64 {
	elements = t.resolveTrigonometryFunctions(elements)
	return t.arithmetic.ResolveOperation(elements)
}

var functionNames = []string{Sin, Cos, Sinh, Cosh, Tan, Cot, Sec, Cosec, Hypot, RootN, PotencyN}

func sustractFunctionsNames(input string) ([]string, bool) {
	match, matches := false, make([]string, 0)
	for _, name := range functionNames {
		if strings.Contains(input, name) {
			match = true
			matches = append(matches, name)
		}
	}
	return matches, match
}

func processSubmatchFunctions(t *trigonometry, input string) string {
	matches, match := sustractFunctionsNames(input)
	var result string

	if match {
		slices.Reverse[[]string](matches)
		for matchIndex := 0; matchIndex < len(matches); matchIndex++ {
			var (
				trigonometryFn string
				firstValue     string
			)
			trigonometryFn = matches[matchIndex]
			enclosedValue := input[strings.Index(input, trigonometryFn):]
			parenthesisOpenIndex, parenthesisCloseIndex := strings.Index(enclosedValue, "("), strings.LastIndex(enclosedValue, ")")

			if enclosedValue[parenthesisCloseIndex-1] == ')' {
				enclosedValue = strings.Replace(enclosedValue, string(enclosedValue[parenthesisCloseIndex]), "", 1)
				parenthesisCloseIndex = strings.LastIndex(enclosedValue, ")")
			}

			firstValue = enclosedValue[parenthesisOpenIndex+1 : parenthesisCloseIndex]
			result := t.arithmetic.ResolveOperation(utils.PrepareArguments([]string{firstValue}))
			input = strings.Replace(input, enclosedValue[:parenthesisCloseIndex+1], fmt.Sprintf("%g", functions[trigonometryFn](t, result)), 1)

		}
	}

	result = input
	return result
}

func (t *trigonometry) resolveTrigonometryFunctions(elements []string) []string {
	elementsStr := strings.Join(elements, "")

	if len(elements) == 1 && utils.SignsRegex.MatchString(elementsStr) {
		elements = utils.CreateSlice(elements)
	}
	arg := strings.Join(elements, "")

	parenthesisOpenIndex, parenthesisCloseIndex := strings.Index(arg, "("), strings.LastIndex(arg, ")")
	principalFunction := arg[:parenthesisOpenIndex]
	enclosedValue := arg[parenthesisOpenIndex+1 : parenthesisCloseIndex]
	_, match := sustractFunctionsNames(enclosedValue)

	if match {
		enclosedValue = processSubmatchFunctions(t, enclosedValue)
	}
	substitutedValue := functions[principalFunction](t, t.arithmetic.ResolveOperation(utils.CreateSlice(strings.Fields(enclosedValue))))
	arg = strings.Replace(arg, arg, fmt.Sprintf("%g", substitutedValue), 1)
	return strings.Fields(arg)
}

func (t *trigonometry) sinFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	return math.Sin(radians)
}

func (t *trigonometry) sinHFn(degrees float64) float64 {
	return math.Sinh(degrees)
}

func (t *trigonometry) cosineFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	return math.Cos(radians)
}

func (t *trigonometry) cosineHFn(degrees float64) float64 {
	return math.Cosh(degrees)
}

func (t *trigonometry) tangentFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	return math.Tan(radians)
}

func (t *trigonometry) cotangentFn(degrees float64) float64 {
	return math.Atan(degrees)
}

func (t *trigonometry) secantFn(degrees float64) float64 {
	return 1 / t.cosineFn(degrees)
}

func (t *trigonometry) cosecantFn(degrees float64) float64 {
	return 1 / t.sinFn(degrees)
}

func (t *trigonometry) hypotenuseFn(catheiA, catheiB float64) float64 {
	return math.Hypot(catheiA, catheiB)
}

func (t *trigonometry) rootNFn(value float64, root float64) float64 {
	var result = value
	counter := root
	for counter != 0 {
		counter--
		result /= root
	}
	return result
}

func (t *trigonometry) potencyNFn(value, exponent float64) float64 {
	return math.Pow(value, exponent)
}

var functions = map[string]func(t *trigonometry, value float64, extraValues ...float64) float64{
	Sin: func(t *trigonometry, value float64, extraValues ...float64) float64 {
		return t.sinFn(value)
	},
	Sinh: func(t *trigonometry, value float64, extraValues ...float64) float64 {
		return t.sinHFn(value)
	},
	Cos: func(t *trigonometry, value float64, extraValues ...float64) float64 {
		return t.cosineFn(value)
	},
	Cosh: func(t *trigonometry, value float64, extraValues ...float64) float64 {
		return t.cosineHFn(value)
	},
	Tan: func(t *trigonometry, value float64, extraValues ...float64) float64 {
		return t.tangentFn(value)
	},
	Cot: func(t *trigonometry, value float64, extraValues ...float64) float64 {
		return t.cotangentFn(value)
	},
	Sec: func(t *trigonometry, value float64, extraValues ...float64) float64 {
		return t.secantFn(value)
	},
	Cosec: func(t *trigonometry, value float64, extraValues ...float64) float64 {
		return t.cosecantFn(value)
	},
	Hypot: func(t *trigonometry, value float64, extraValues ...float64) float64 {
		return t.hypotenuseFn(value, extraValues[0])
	},
	RootN: func(t *trigonometry, value float64, extraValues ...float64) float64 {
		return t.rootNFn(value, extraValues[0])
	},
	PotencyN: func(t *trigonometry, value float64, extraValues ...float64) float64 {
		return t.potencyNFn(value, extraValues[0])
	},
}
