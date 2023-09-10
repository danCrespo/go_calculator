package arithmetic

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/exp/slices"

	"calculator/utils"
)

const (
	Addition       = "+"
	Sustraction    = "-"
	Multiplication = "*"
	Division       = "/"
	Module         = "%"
	Exponent       = "^"
)

type (
	ArithmeticArea interface {
		ResolveOperation(elements []string) float64
	}

	arithmetic struct{}
)

func Arithmetic() ArithmeticArea {
	return &arithmetic{}
}

func (a *arithmetic) ResolveOperation(elements []string) float64 {
	return result(elements)
}

func result(elements []string) float64 {
	var results float64

	if slices.Contains[[]string, string](elements, Module) {
		results = moduleOperation(elements)
	} else {
		results = pemdasOrder(elements)
	}

	return results
}

func pemdasOrder(args []string) float64 {
	P, E, M, D, A, S := "(", "^", "*", "/", "+", "-"

	args = processParenthesisGroups(args, P)
	args = processOperationElements(args, E)
	args = processOperationElements(args, M)
	args = processOperationElements(args, D)
	args = processOperationElements(args, A)
	args = processOperationElements(args, S)

	return utils.StringToFloat64(args[0])
}

func moduleOperation(operationElements []string) float64 {
	var result float64

	for range operationElements {
		if slices.Contains(operationElements, Module) {
			moduleIndex := slices.Index[[]string, string](operationElements, Module)
			result = utils.StringToFloat64(operationElements[moduleIndex-1]) -
				math.Floor(utils.StringToFloat64(operationElements[moduleIndex-1])/utils.StringToFloat64(operationElements[moduleIndex+1]))*
					utils.StringToFloat64(operationElements[moduleIndex+1])

		}
	}
	return result
}

func processParenthesisGroups(operationElements []string, operator string) (args []string) {
	elementsGroup := make([]string, 0)

	for elementIndex, element := range operationElements {

		if utils.StringContains(element, operator) {
			parenthesisStr := strings.Join(operationElements, " ")
			parenthesisCount := strings.Count(parenthesisStr, operator)
			firstParenthesisIndex := strings.Index(parenthesisStr, operator)

			if parenthesisCount > 1 && parenthesisStr[firstParenthesisIndex+1] == '(' {
				lastParenthesisIndex := strings.LastIndex(parenthesisStr, operator)

				for charIndex := lastParenthesisIndex; charIndex < len(parenthesisStr); charIndex++ {
					for nextCharIndex := charIndex + 1; nextCharIndex < len(parenthesisStr); nextCharIndex++ {

						if parenthesisStr[nextCharIndex] == ')' {
							removeParenthesis := strings.ReplaceAll(parenthesisStr[lastParenthesisIndex:nextCharIndex], operator, "")
							removeParenthesis = strings.ReplaceAll(removeParenthesis, ")", "")
							oldStr := "(" + removeParenthesis + ")"
							parenthesisStr = strings.ReplaceAll(parenthesisStr, oldStr, fmt.Sprintf("%g", resolve(strings.Fields(removeParenthesis))))
							operationElements = strings.Fields(parenthesisStr)
							break
						}
					}
				}

			} else {
				for elementInd := 0; elementInd < len(operationElements); elementInd++ {
					if strings.Contains(operationElements[elementInd], ")") {
						removeParenthesis := strings.ReplaceAll(strings.Join(operationElements[elementIndex:elementInd+1], " "), operator, "")
						removeParenthesis = strings.ReplaceAll(removeParenthesis, ")", "")
						operationElements = utils.ReplaceSlice(operationElements, elementIndex, elementInd+1, fmt.Sprintf("%g", resolve(strings.Fields(removeParenthesis))))
						break
					}
				}
			}
		}

		elementsGroup = slices.Delete[[]string](elementsGroup, 0, len(elementsGroup))
	}

	args = operationElements
	return
}

func processOperationElements(operationElements []string, operator string) (args []string) {
	elementsGroup := make([]string, 0)

	for range operationElements {
		if slices.Contains(operationElements, operator) {
			operatorIndex := slices.Index[[]string, string](operationElements, operator)

			if operator == Sustraction && len(operationElements[operatorIndex]) > 1 {
				elementsGroup = append(elementsGroup, operationElements[operatorIndex:operatorIndex+3]...)
				operationElements = utils.ReplaceSlice(operationElements, operatorIndex, operatorIndex+3, fmt.Sprintf("%g", resolve(elementsGroup)))
			} else {
				elementsGroup = append(elementsGroup, operationElements[operatorIndex-1:operatorIndex+2]...)
				operationElements = utils.ReplaceSlice(operationElements, operatorIndex-1, operatorIndex+2, fmt.Sprintf("%g", resolve(elementsGroup)))
			}
		}

		elementsGroup = slices.Delete[[]string](elementsGroup, 0, len(elementsGroup))
	}

	args = operationElements
	return
}

func resolve(operationElements []string) float64 {
	operationElements = mathOperation(operationElements, Exponent)
	operationElements = mathOperation(operationElements, Multiplication)
	operationElements = mathOperation(operationElements, Division)
	operationElements = mathOperation(operationElements, Addition)
	operationElements = mathOperation(operationElements, Sustraction)

	return utils.StringToFloat64(operationElements[0])
}

func mathOperation(elements []string, operand string) (operationElements []string) {

	for range elements {
		if slices.Contains(elements, operand) {
			operandIndex := slices.Index[[]string, string](elements, operand)
			result := operatorsFunctions[operand](elements, operandIndex)
			elements = utils.ReplaceSlice(elements, operandIndex-1, operandIndex+2, fmt.Sprintf("%g", result))
		}
	}
	operationElements = elements
	return
}

var operatorsFunctions = map[string]func([]string, int) float64{
	Exponent: func(elements []string, operandIndex int) float64 {
		return math.Pow(utils.StringToFloat64(elements[operandIndex-1]), utils.StringToFloat64(elements[operandIndex+1]))
	},
	Multiplication: func(elements []string, operandIndex int) float64 {
		return utils.StringToFloat64(elements[operandIndex-1]) * utils.StringToFloat64(elements[operandIndex+1])
	},
	Division: func(elements []string, operandIndex int) float64 {
		return utils.StringToFloat64(elements[operandIndex-1]) / utils.StringToFloat64(elements[operandIndex+1])
	},
	Addition: func(elements []string, operandIndex int) float64 {
		return utils.StringToFloat64(elements[operandIndex-1]) + utils.StringToFloat64(elements[operandIndex+1])
	},
	Sustraction: func(elements []string, operandIndex int) float64 {
		return utils.StringToFloat64(elements[operandIndex-1]) - utils.StringToFloat64(elements[operandIndex+1])
	},
}
