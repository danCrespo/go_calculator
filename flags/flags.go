package flags

import (
	"flag"
	"fmt"
	"strings"

	calculator "calculator/calc"
)

type PrecisionFlag struct{ calculator.Calculator }

func (f *PrecisionFlag) Set(p string) error {
	var verb string

	fmt.Sscanf(p, "%s", &verb)

	if !strings.Contains(verb, "%") {
		verb = strings.Replace(verb, verb, "%"+verb, 1)
	}

	fmt.Println("verb", verb)
	f.Precision = verb
	return nil
}

func Precision(name string, value string, usage string) *calculator.Calculator {
	f := PrecisionFlag{Calculator: calculator.Calculator{Precision: value}}

	f.PrecisionFlagVar(name, usage)
	return &f.Calculator
}

func (f *PrecisionFlag) PrecisionFlagVar(name string, usage string) {
	flag.CommandLine.Var(f, name, usage)
}
