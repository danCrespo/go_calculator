package trigonometry

import (
	"fmt"
	"math"
	"strings"

	"github.com/danCrespo/go_calculator/arithmetic"
	"github.com/danCrespo/go_calculator/utils"
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

func (t *trigonometry) resolveTrigonometryFunctions(elements []string) []string {
	if len(elements) == 1 {
		elements = strings.Split(elements[0], "")
	}
	arg := strings.Join(elements, "")

	match, matches := utils.StringContainsAndWhatContains(arg, Sin, Cos, Sinh, Cosh, Tan, Cot, Sec, Cosec, Hypot, RootN, PotencyN)
	if match {
		for matchInd := 0; matchInd < len(matches); matchInd++ {
			var (
				trigonometryFn string
				firstValue     string
				secondValue    string
				value          = arg
			)
			trigonometryFn = matches[matchInd]
			fmt.Println("Matches", matches)

			for charInd := 0; charInd < len(value); charInd++ {
				if value[charInd] == '(' {

					for nextCharInd := charInd + 1; nextCharInd < len(value); nextCharInd++ {
						if value[nextCharInd] == ')' {
							value = value[charInd+1 : nextCharInd]

							if utils.StringContains(value, ", ") {
								fmt.Println("Matches", arg[strings.Index(arg, trigonometryFn):])
								fmt.Sscanf(arg[strings.Index(arg, trigonometryFn):], trigonometryFn+"(%s, %s)", &firstValue, &secondValue)
								// fmt.Sscanf(arg[strings.Index(arg, trigonometryFn):], trigonometryFn+"("+fmt.Sprint(firstValue)+" %s)", )

								firstValue = firstValue[:len(firstValue)-1]
								secondValue = secondValue[:len(secondValue)-1]
								firstValueResult := t.arithmetic.ResolveOperation(utils.PrepareArguments([]string{firstValue}))
								secondValueResult := t.arithmetic.ResolveOperation(utils.PrepareArguments([]string{secondValue}))

								arg = strings.Replace(arg, fmt.Sprintf("%s(%s, %s)", trigonometryFn, firstValue, secondValue), fmt.Sprintf("%g", functions[trigonometryFn](t, firstValueResult, secondValueResult)), 1)
							} else {
								fmt.Sscanf(arg[strings.Index(arg, trigonometryFn):], trigonometryFn+"(%s)", &firstValue)
								firstValue = firstValue[:len(firstValue)-1]
								result := t.arithmetic.ResolveOperation(utils.PrepareArguments([]string{firstValue}))
								arg = strings.Replace(arg, fmt.Sprintf("%s(%s)", trigonometryFn, firstValue), fmt.Sprintf("%g", functions[trigonometryFn](t, result)), 1)
							}
						}
					}
				}
			}
		}
	}

	return strings.Fields(arg)
}

func (t *trigonometry) sinFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	return math.Sin(radians)
}

func (t *trigonometry) sinHFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	return math.Sinh(radians)
}

func (t *trigonometry) cosineFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	return math.Cos(radians)
}

func (t *trigonometry) cosineHFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	return math.Cosh(radians)
}

func (t *trigonometry) tangentFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	return math.Tan(radians)
}

func (t *trigonometry) cotangentFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	return math.Atan(radians)
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
