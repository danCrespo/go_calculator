package calculator

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"

	"calculator/flags"
	"calculator/geometry"
	"calculator/utils"
)

const (
	addition       = "+"
	sustraction    = "-"
	multiplication = "*"
	division       = "/"
	module         = "%"
	exponent       = "^"
)

type CalculatorInstance interface {
	StartCalculation(result chan float64, args []string)
	PrintResults(result chan float64)
}

type Calculator struct {
	Precision *flags.CalculatorPrecision
	Hierachy  *flags.CalculatorHierarchy
	Area      *flags.CalculatorFigureArea
	Result    float64
}

func NewCalculator(flags flags.CalculatorFlags) CalculatorInstance {
	c := &Calculator{
		Precision: flags.CalculatorPrecision,
		Hierachy:  flags.CalculatorHierarchy,
		Area:      flags.CalculatorFigureArea,
	}
	return c
}

func (c *Calculator) PrintResults(result chan float64) {
	for r := range result {
		c.Result = r
		fmt.Fprintf(os.Stdout, "\v \033[01;05;32mResult: \033[01;05;36m "+string(*c.Precision)+"\033[00m\n\v", c.Result)
	}
}

func (c *Calculator) StartCalculation(resultOperationCh chan float64, args []string) {
	go func() {
		var result float64
		defer close(resultOperationCh)

		if c.Area.String() != " " {
			fmt.Println("arithmeticOperations results", c.Area)
			result = geometry.Area(args, string(*c.Area))
		} else {
			result = arithmeticOperations(args)
		}
		resultOperationCh <- result
	}()
}

func arithmeticOperations(args []string) float64 {
	var results float64

	signsRegex := regexp.MustCompile(`([\*/\+\-%\^]){1}`)
	parenthesisRegex := regexp.MustCompile(`([(\)]){1}`)

	if len(args) == 1 {
		args = strings.Split(args[0], "")
	}

	inp := strings.Join(args, "")
	var input string

	for i := 0; i < len(inp); i++ {
		if inp[i] != ' ' {
			if inp[i] == '-' && (inp[i-1] != ' ' && !parenthesisRegex.Match([]byte(string(inp[i-1])))) && (inp[i+1] != ' ' && !parenthesisRegex.Match([]byte(string(inp[i+1])))) {
				input += " " + string(inp[i]) + " "
			} else {

				if signsRegex.Match([]byte(string(inp[i]))) {
					if inp[i] == '-' && inp[i+1] != ' ' && !parenthesisRegex.Match([]byte(string(inp[i+1]))) {
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

	if slices.Contains[[]string, string](fields, module) {
		results = moduleOperation(fields)
	} else {
		results = PEMDASOrder(fields)
	}
	fmt.Println("arithmeticOperations results", results)

	return results
}

func PEMDASOrder(args []string) float64 {
	P, E, M, D, A, S := "(", "^", "*", "/", "+", "-"

	args = processParenthesisGroups(args, P)
	args = processOperationElements(args, E)
	args = processOperationElements(args, M)
	args = processOperationElements(args, D)
	args = processOperationElements(args, A)
	args = processOperationElements(args, S)

	return utils.StringToFloat64(args[0])
}

func resolveOperations(operationElements []string) float64 {

	for range operationElements {
		if slices.Contains(operationElements, exponent) {
			exponentIndex := slices.Index[[]string, string](operationElements, exponent)
			result := math.Pow(utils.StringToFloat64(operationElements[exponentIndex-1]), utils.StringToFloat64(operationElements[exponentIndex+1]))
			operationElements = utils.ReplaceSlice(operationElements, exponentIndex-1, exponentIndex+2, fmt.Sprintf("%g", result))
		}
	}

	for range operationElements {
		if slices.Contains[[]string, string](operationElements, multiplication) {
			multiplyIndex := slices.Index[[]string, string](operationElements, multiplication)
			result := utils.StringToFloat64(operationElements[multiplyIndex-1]) * utils.StringToFloat64(operationElements[multiplyIndex+1])
			operationElements = utils.ReplaceSlice(operationElements, multiplyIndex-1, multiplyIndex+2, fmt.Sprintf("%g", result))
		}
	}

	for range operationElements {
		if slices.Contains(operationElements, division) {
			divisionIndex := slices.Index[[]string, string](operationElements, division)
			result := utils.StringToFloat64(operationElements[divisionIndex-1]) / utils.StringToFloat64(operationElements[divisionIndex+1])
			operationElements = utils.ReplaceSlice(operationElements, divisionIndex-1, divisionIndex+2, fmt.Sprintf("%g", result))
		}
	}

	for range operationElements {
		if slices.Contains(operationElements, addition) {
			additionIndex := slices.Index[[]string, string](operationElements, addition)
			result := utils.StringToFloat64(operationElements[additionIndex-1]) + utils.StringToFloat64(operationElements[additionIndex+1])
			operationElements = utils.ReplaceSlice(operationElements, additionIndex-1, additionIndex+2, fmt.Sprintf("%g", result))
		}
	}

	for range operationElements {
		if slices.Contains(operationElements, sustraction) {
			sustractionIndex := slices.Index[[]string, string](operationElements, sustraction)
			result := utils.StringToFloat64(operationElements[sustractionIndex-1]) - utils.StringToFloat64(operationElements[sustractionIndex+1])
			operationElements = utils.ReplaceSlice(operationElements, sustractionIndex-1, sustractionIndex+2, fmt.Sprintf("%g", result))
		}
	}

	return utils.StringToFloat64(operationElements[0])
}

func moduleOperation(operationElements []string) float64 {
	var result float64

	for range operationElements {
		if slices.Contains(operationElements, module) {
			moduleIndex := slices.Index[[]string, string](operationElements, module)
			result = utils.StringToFloat64(operationElements[moduleIndex-1]) -
				math.Floor(utils.StringToFloat64(operationElements[moduleIndex-1])/utils.StringToFloat64(operationElements[moduleIndex+1]))*
					utils.StringToFloat64(operationElements[moduleIndex+1])

		}
	}
	return result
}

func processParenthesisGroups(operationElements []string, operator string) (args []string) {
	groups := make([]string, 0)

	for i, arg := range operationElements {
		var result float64

		if strings.Contains(arg, operator) {
			parenthesisStr := strings.Join(operationElements, " ")
			parenthesisCount := strings.Count(parenthesisStr, operator)
			firstParenthesisIndex := strings.Index(parenthesisStr, operator)

			if parenthesisCount > 1 && parenthesisStr[firstParenthesisIndex+1] == '(' {
				lastParenthesisIndex := strings.LastIndex(parenthesisStr, operator)

				for j := lastParenthesisIndex; j < len(parenthesisStr); j++ {
					for k := j + 1; k < len(parenthesisStr); k++ {

						if parenthesisStr[k] == ')' {
							removeParenthesis := strings.ReplaceAll(parenthesisStr[lastParenthesisIndex:k], operator, "")
							removeParenthesis = strings.ReplaceAll(removeParenthesis, ")", "")
							fields := strings.Fields(removeParenthesis)
							result = resolveOperations(fields)
							oldStr := "(" + removeParenthesis + ")"
							parenthesisStr = strings.ReplaceAll(parenthesisStr, oldStr, fmt.Sprintf("%g", result))
							operationElements = strings.Fields(parenthesisStr)
							break
						}
					}
				}

			} else {
				for j := 0; j < len(operationElements); j++ {
					if strings.Contains(operationElements[j], ")") {
						removeParenthesis := strings.ReplaceAll(strings.Join(operationElements[i:j+1], " "), operator, "")
						removeParenthesis = strings.ReplaceAll(removeParenthesis, ")", "")
						result = resolveOperations(strings.Fields(removeParenthesis))
						operationElements = utils.ReplaceSlice(operationElements, i, j+1, fmt.Sprintf("%g", result))
						break
					}
				}
			}
		}

		groups = slices.Delete[[]string](groups, 0, len(groups))
	}

	args = operationElements
	return
}

func processOperationElements(operationElements []string, operator string) (args []string) {
	groups := make([]string, 0)

	for range operationElements {
		var result float64

		if slices.Contains(operationElements, operator) {
			operatorIndex := slices.Index[[]string, string](operationElements, operator)

			if operator == sustraction && len(operationElements[operatorIndex]) > 1 {
				groups = append(groups, operationElements[operatorIndex:operatorIndex+3]...)
				result = resolveOperations(groups)
				operationElements = utils.ReplaceSlice(operationElements, operatorIndex, operatorIndex+3, fmt.Sprintf("%g", result))
			} else {
				groups = append(groups, operationElements[operatorIndex-1:operatorIndex+2]...)
				result = resolveOperations(groups)
				operationElements = utils.ReplaceSlice(operationElements, operatorIndex-1, operatorIndex+2, fmt.Sprintf("%g", result))
			}
		}

		groups = slices.Delete[[]string](groups, 0, len(groups))
	}

	args = operationElements
	return
}
