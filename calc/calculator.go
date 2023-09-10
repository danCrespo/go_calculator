package calculator

import (
	"fmt"
	"os"

	"calculator/arithmetic"
	"calculator/flags"
	"calculator/geometry"
	"calculator/trigonometry"
	"calculator/utils"
)

type (
	CalculatorInstance interface {
		StartCalculation(resultOperationCh chan map[string]any, args []string)
		PrintResults(results chan map[string]any)
	}

	Calculator struct {
		PrecisionFlag    *flags.CalculatorPrecision
		TrigonometryFlag *flags.CalculatorTrigonometry
		AreaFlag         *flags.CalculatorFigureArea
		arithmetic       arithmetic.ArithmeticArea
		geometry         geometry.GeometryArea
		trigonometry     trigonometry.TrigonometryArea
	}

	MathFields struct {
		arithmetic.ArithmeticArea
		geometry.GeometryArea
		trigonometry.TrigonometryArea
	}
)

func NewCalculator(flags flags.CalculatorFlags, mathsFields MathFields) CalculatorInstance {
	c := &Calculator{
		PrecisionFlag:    flags.CalculatorPrecision,
		TrigonometryFlag: flags.CalculatorTrigonometry,
		AreaFlag:         flags.CalculatorFigureArea,
		arithmetic:       mathsFields.ArithmeticArea,
		geometry:         mathsFields.GeometryArea,
		trigonometry:     mathsFields.TrigonometryArea,
	}
	return c
}

func (c *Calculator) StartCalculation(resultOperationCh chan map[string]any, args []string) {
	go func() {
		chMap := make(map[string]any)
		defer close(resultOperationCh)

		if *c.AreaFlag != " " {
			res, err := c.geometry.CalculateArea(args, string(*c.AreaFlag))
			if err != nil {
				chMap["errorResult"] = err
			} else {
				chMap["result"] = res
			}
		} else if *c.TrigonometryFlag {
			chMap["result"] = c.trigonometry.Calculate(args)
		} else {
			args = utils.PrepareArguments(args)
			chMap["result"] = c.arithmetic.ResolveOperation(args)
		}

		resultOperationCh <- chMap
	}()
}

func (c *Calculator) PrintResults(results chan map[string]any) {
	for ch := range results {
		for k, r := range ch {

			if k == "errorResult" {
				fmt.Fprintf(os.Stderr, "\v \033[01;05;31mError: \033[01;05;36m%v\033[00m\n\v", r.(error))
			} else if k == "result" {
				fmt.Fprintf(os.Stdout, "\v \033[01;05;32mResult: \033[01;05;36m "+string(*c.PrecisionFlag)+"\033[00m\n\v", r.(float64))
			}
		}
	}
}
