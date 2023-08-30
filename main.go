package main

import (
	calculator "calculator/calc"
	"calculator/flags"
	"flag"
)

var usage = "p (precision) specifies the precision result printed default.\nAccepts a number or a go float64 verb.\n"
var precision flags.PrecisionFlag

func init() {
	usage += "i.e:\n\t-p 3 or -p \\%3.4f\nDefaults to %g"

	precision.PrecisionFlagVar("precision", usage)
	precision.PrecisionFlagVar("p", usage+" (shorthand)")
	precision.PrecisionFlagVar("P", usage+" (shorthand)")
}

func main() {
	flag.Parse()
	var (
		// chanle that contains arguments readed from standard input (Stdin)
		entries = make(chan []string)
		// channel slice of arguments to do the necessary math operations
		operationParts = make(chan [][]string)
		result         = make(chan float64)
	)

	c := calculator.NewCalculator()
	calculator.Start(entries, operationParts)
	go calculator.Operations(operationParts, result)
	c.Results(result)
}
