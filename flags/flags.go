package flags

import (
	"os"

	calculator "calculator/calc"
)

type PrecisionFlag struct {
	calculator.CalculatorPrecision
}

var cmd = newCmdLine()

func Args() []string {
	return cmd.Args()
}

func Parse() {
	cmd.FlagSet.Parse(os.Args[1:])
}

func (f *PrecisionFlag) Var(name, usage string, example ...string) *calculator.CalculatorPrecision {
	cmd.Var(f, name, usage, example...)
	return &f.CalculatorPrecision
}

func Precision(name string, value calculator.CalculatorPrecision, usage string, examples ...string) *PrecisionFlag {
	f := PrecisionFlag{value}
	cmd.Var(&f, name, usage, examples...)
	return &f
}
