package flags

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	CalculatorFlags struct {
		*CalculatorPrecision
		*CalculatorHierarchy
	}

	CalculatorPrecision string

	CalculatorHierarchy bool

	PrecisionFlag struct {
		CalculatorPrecision
	}

	HierarchyFlag struct {
		CalculatorHierarchy
	}
)

func (c CalculatorPrecision) String() string {
	return string(c)
}

func (c CalculatorHierarchy) String() string {
	return strconv.FormatBool(bool(c))
}

var cmd = newCmdLine()

func Args() []string {
	return cmd.Args()
}

func Parse() {
	cmd.FlagSet.Parse(os.Args[1:])
}

func (f *PrecisionFlag) Var(name, usage string, example ...string) *CalculatorPrecision {
	cmd.Var(f, name, usage, example...)
	return &f.CalculatorPrecision
}

func (f *PrecisionFlag) Set(p string) error {
	var verb string

	fmt.Sscanf(p, "%s", &verb)

	if !strings.Contains(verb, "%") {
		verb = strings.Replace(verb, verb, "%."+verb+"f", 1)
	}
	f.CalculatorPrecision = CalculatorPrecision(verb)
	return nil
}

func Precision(name string, value CalculatorPrecision, usage string, examples ...string) *PrecisionFlag {
	f := PrecisionFlag{value}
	cmd.Var(&f, name, usage, examples...)
	return &f
}

func (f *HierarchyFlag) Set(v string) error {
	var value bool

	fmt.Sscanf(v, "%b", &value)
	switch v {
	case "no":
	case "false":
	case "off":
		f.CalculatorHierarchy = CalculatorHierarchy(!value)
		return nil

	case "yes":
	case "true":
	case "on":
		f.CalculatorHierarchy = CalculatorHierarchy(value)
		return nil
	}

	return fmt.Errorf("%s is not a valid argument", v)
}

func (f *HierarchyFlag) Var(name, usage string, example ...string) *CalculatorHierarchy {
	cmd.Var(f, name, usage, example...)
	return &f.CalculatorHierarchy
}

func Hierachy(name string, value CalculatorHierarchy, usage string, examples ...string) *HierarchyFlag {
	f := HierarchyFlag{value}
	cmd.Var(&f, name, usage, examples...)
	return &f
}
