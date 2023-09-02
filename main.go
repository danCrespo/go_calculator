package main

import (
	calculator "calculator/calc"
	"calculator/flags"
)

var (
	calcPrecision     *flags.CalculatorPrecision
	percisionUsage    = "Sets the `precision` of the result\nAccepts a number or a go float64 verb.\n"
	precisionExamples = "\n./calculator -p 3 '5.5 * 5'   Output: Resutl: 27.500\n" +
		"./calculator --precision %.4f '255.99 * 35.5'   Output: Resutl: 9087.6450\n"

	calcHierachy  *flags.CalculatorHierarchy
	hierachyUsage = "Controls if PEMDAS order is respect or no.\nAccepts `falsy: no, false, off; truthy: yes, true, on`.\n" +
		"The \"PEMDAS\" (Parentheses, Exponents, Multiplication, Division, Addition, Subtraction) order refers to the hierachy of the math operators\n" +
		"this order totally affects the result of the operations and if this order isn't respected, the result probably will be wrong.\n" +
		"So, for that reason it is not very logical not to respect the order, unless you have a good reason to do so.\n" +
		"\nThe order is: [()] > [x^] > [*] > [/] > [+] > [-].\n"
	hierachyExamples = "\n./calculator -H yes '(1 + 6 \u00f7 3) \u00f7 (6-5)'   Output: Result: 3\n" +
		"./calculator --hierachy false '(1 + 6 \u00f7 3) \u00f7 (6-5)'   Output: Result: 2.33\n"
)

func init() {
	calcPrecision = &flags.Precision("p, precision", "%g", percisionUsage+" ", precisionExamples).CalculatorPrecision
	calcHierachy = &flags.Hierachy("H, hierachy", true, hierachyUsage, hierachyExamples).CalculatorHierarchy
}

func main() {
	flags.Parse()

	var operationResultCh = make(chan float64)

	calcFlags := flags.CalculatorFlags{
		CalculatorPrecision: calcPrecision,
		CalculatorHierarchy: calcHierachy,
	}

	c := calculator.NewCalculator(calcFlags)
	c.StartCalculation(operationResultCh, flags.Args())
	c.PrintResults(operationResultCh)
}
