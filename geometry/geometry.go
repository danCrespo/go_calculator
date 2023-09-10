package geometry

import (
	"calculator/arithmetic"
	"calculator/utils"
	"fmt"
	"slices"
	"strings"
)

type (
	GeometryArea interface {
		CalculateArea(measures []string, figure string) (float64, error)
	}
	geometry struct{ arithmetic arithmetic.ArithmeticArea }
)

func Geometry(a arithmetic.ArithmeticArea) GeometryArea {
	g := geometry{a}
	return &g
}

func (g *geometry) CalculateArea(measures []string, figure string) (float64, error) {
	measures = g.prepareElements(measures)
	return area(measures, figure)
}

func (g *geometry) prepareElements(elements []string) []string {
	for _, arg := range elements {
		if utils.SignsRegex.MatchString(arg) || utils.ParenthesisRegex.MatchString(arg) {
			var backupArg string
			var argSplit []string
			argIndex := slices.Index[[]string](elements, arg)

			if strings.Contains(arg, "=") {
				argSplit = strings.Split(arg, "=")
				fmt.Sscanf(argSplit[0], "%s=", &backupArg)
				arguments := utils.PrepareArguments([]string{argSplit[1]})
				arg = fmt.Sprintf("%s=%g", backupArg, g.arithmetic.ResolveOperation(arguments))
			} else {
				arguments := utils.PrepareArguments([]string{arg})
				arg = fmt.Sprintf("%g", g.arithmetic.ResolveOperation(arguments))
			}

			elements = slices.Replace[[]string, string](elements, argIndex, argIndex+1, arg)
		}
	}
	return elements
}
