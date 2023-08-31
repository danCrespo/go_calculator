package main

import (
	calculator "calculator/calc"
	"calculator/flags"
)

var calcPrecision *calculator.CalculatorPrecision

func init() {
	var usage = "Sets the `precision` of the result\nAccepts a number or a go float64 verb.\v"
	example := "\n./calculator -p 3 '5.5 * 5'   Output: Operation Resutl: 27.500\n\n"
	example += "./calculator -p %.4f '255.99 * 35.5'   Output: Operation Resutl: 9087.6450\n"

	calcPrecision = &flags.Precision("p,precision", "%g", usage+" ", example).CalculatorPrecision
}

func main() {
	flags.Parse()
	var (
		// chanle that contains arguments readed from standard input (Stdin)
		entries = make(chan []string)
		// channel slice of arguments to do the necessary math operations
		operationParts = make(chan [][]string)
		result         = make(chan float64)
	)

	c := calculator.NewCalculator(calcPrecision)
	calculator.Start(entries, operationParts, flags.Args())
	go calculator.Operations(operationParts, result)
	c.Results(result)
}
