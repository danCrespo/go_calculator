package arithmetic

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/danCrespo/go_calculator/utils"
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
	firstMatch := checkOrder(args, A, S)
	if firstMatch == A {
		args = processOperationElements(args, A)
		args = processOperationElements(args, S)
	} else {
		args = processOperationElements(args, S)
		args = processOperationElements(args, A)
	}

	return utils.StringToFloat64(args[0])
}

func checkOrder(args []string, obj1, obj2 string) string {
	var firstMatch string
	for index := 0; index < len(args); index++ {
		if args[index] == obj1 {
			firstMatch = obj1
			break
		} else if args[index] == obj2 {
			firstMatch = obj2
			break
		}
	}

	return firstMatch
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
	elementStr := strings.Join(operationElements, "")
	openParenthesisIndexes := make([]int, 0)

	for char := 0; char < len(elementStr); char++ {
		if elementStr[char] == '(' {
			openParenthesisIndexes = append(openParenthesisIndexes, char)
		}
	}

	for parenthesisIndex := len(openParenthesisIndexes) - 1; parenthesisIndex >= 0; parenthesisIndex-- {
		openParenthesis := openParenthesisIndexes[parenthesisIndex]
		elementsSlice := elementStr[openParenthesis:]

		for char := 0; char < len(elementsSlice); char++ {
			if elementsSlice[char] == ')' {
				portion := elementsSlice[1:char]
				parenthesisResult := resolve(strings.Fields(portion))
				elementStr = strings.ReplaceAll(elementStr, elementStr[openParenthesis:openParenthesis+char+1], fmt.Sprintf("%g", parenthesisResult))
				break
			}
		}
	}

	args = strings.Fields(elementStr)
	return
}

func processOperationElements(operationElements []string, operator string) (args []string) {
	elementsStr := strings.Join(operationElements, "")

	if len(operationElements) == 1 && utils.SignsRegex.MatchString(elementsStr) {
		operationElements = utils.CreateSlice(operationElements)
	}

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
	elementsStr := strings.Join(elements, "")

	if len(elements) == 1 && utils.SignsRegex.MatchString(elementsStr) {
		elements = utils.CreateSlice(elements)
	}

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
