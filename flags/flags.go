package flags

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	geometry "github.com/danCrespo/go_calculator/utils"
)

type (
	CalculatorFlags struct {
		*CalculatorPrecision
		*CalculatorTrigonometry
		*CalculatorFigureArea
	}

	CalculatorPrecision string

	CalculatorTrigonometry bool

	CalculatorFigureArea string

	PrecisionFlag struct {
		CalculatorPrecision
	}

	TrigonometryFlag struct {
		CalculatorTrigonometry
	}

	AreaFlag struct {
		CalculatorFigureArea
	}
)

func (c CalculatorPrecision) String() string {
	return string(c)
}

func (c CalculatorTrigonometry) String() string {
	return strconv.FormatBool(bool(c))
}

func (c CalculatorFigureArea) String() string {
	return string(c)
}

var cmd = newCmdLine()

func Args() []string {
	return cmd.Args()
}

func Parse() {
	if err := cmd.FlagSet.Parse(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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

func (f *TrigonometryFlag) Set(v string) error {
	var value string

	fmt.Sscanf(v, "%s", &value)
	switch value {
	case "no", "false", "off":
		f.CalculatorTrigonometry = CalculatorTrigonometry(false)
		return nil

	case "yes", "true", "on", "y":
		f.CalculatorTrigonometry = CalculatorTrigonometry(true)
		return nil
	}

	return fmt.Errorf("%s is not a valid argument", v)
}

func (f *TrigonometryFlag) Var(name, usage string, example ...string) *CalculatorTrigonometry {
	cmd.Var(f, name, usage, example...)
	return &f.CalculatorTrigonometry
}

func Trigonometry(name string, value CalculatorTrigonometry, usage string, examples ...string) *TrigonometryFlag {
	f := TrigonometryFlag{value}
	cmd.Var(&f, name, usage, examples...)
	return &f
}

func (f *AreaFlag) Set(value string) error {
	var geometricFigure CalculatorFigureArea
	fmt.Sscanf(value, "%s", &geometricFigure)

	switch value {
	case "square", "Square", "sqr.":
		fmt.Println("Set value", geometricFigure, value)
		f.CalculatorFigureArea = CalculatorFigureArea(geometry.Square)
		return nil
	case "Triangle", "triangle", "tri.":
		f.CalculatorFigureArea = CalculatorFigureArea(geometry.Triangle)
		return nil
	case "Circle", "circle", "circ.":
		f.CalculatorFigureArea = CalculatorFigureArea(geometry.Circle)
		return nil
	case "Rectangle", "rectangle", "rect.":
		f.CalculatorFigureArea = CalculatorFigureArea(geometry.Rectangle)
		return nil
	case "Trapezoid", "trapezoid", "trapeze", "trap.":
		f.CalculatorFigureArea = CalculatorFigureArea(geometry.Trapezoid)
		return nil
	case "Rhombus", "rhombus", "diamond", "rhom.":
		f.CalculatorFigureArea = CalculatorFigureArea(geometry.Rhombus)
		return nil
	case "Ellipse", "ellipse", "ellip.":
		f.CalculatorFigureArea = CalculatorFigureArea(geometry.Ellipse)
		return nil

	}

	return fmt.Errorf("%s is not a valid argument", value)
}

func (f *AreaFlag) Var(name, usage string, example ...string) *CalculatorFigureArea {
	cmd.Var(f, name, usage, example...)
	return &f.CalculatorFigureArea
}

func Area(name string, value CalculatorFigureArea, usage string, example ...string) *AreaFlag {
	f := AreaFlag{value}
	cmd.Var(&f, name, usage, example...)
	return &f
}
