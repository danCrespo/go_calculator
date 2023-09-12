package main

import (
	"github.com/danCrespo/go_calculator/arithmetic"
	calculator "github.com/danCrespo/go_calculator/calc"
	"github.com/danCrespo/go_calculator/flags"
	"github.com/danCrespo/go_calculator/geometry"
	"github.com/danCrespo/go_calculator/trigonometry"
)

var (
	calcPrecision     *flags.CalculatorPrecision
	percisionUsage    = "Sets the `precision` of the result\nAccepts a number or a go float64 verb.\n"
	precisionExamples = "\n./calculator -p 3 '5.5 * 5'   Output: Resutl: 27.500\n" +
		"./calculator --precision %.4f '255.99 * 35.5'   Output: Resutl: 9087.6450\n"

	calcTrigonometry  *flags.CalculatorTrigonometry
	trigonometryUsage = "Sets the trigonometry mode.\nAccepts `falsy: no, false, off; truthy: yes, true, on`.\n" +
		"The \"trigonometry\" mode allows to resolve complex maths. The follow is a list of supported functions:\n" +
		"sin, cos, tan, cot, sinh, cosh, hypot, sec, cosec, rootN, ^N (potency N)\n"
	trigonometryExamples = "\n./calculator -t y sin(60)\n" +
		"./calculator --trigonometry true cos(30)\n"

	calcFigureArea *flags.CalculatorFigureArea
	areaUsage      = "Used to calculate the area of a figure given their measures.\nAccepts the `figure name` or an alias of the name; the arguments of the measures will depend of the figure.\n" +
		"The follow is a list of accepted figures, their aliases and the arguments to be passed:\n" +
		"  Name      \tAliases                   \tArguments\n" +
		"  Square    \tsquare, sqr.              \tone argument (measure) [number]\n" +
		"  Triangle  \ttriangle, tri.            \tb=[number] h=[number] or just [number] [number], in the second form first argument will be b=num and the second will be h=num\n" +
		"  Circle    \tcircle, circ.             \tone argument (measure) this would be the radius\n" +
		"  Rectangle \trectangle, rect.          \tb=[number] h=[number] or just [number] [number], in the second form first argument will be b=num and the second will be h=num\n" +
		"  Trapezoid \ttrapezoid, trapeze, trap. \tsideA=[number] sideB=[number] h=[number] or just [number] [number] [number], in the second form first argument will be sideA=num, second will be sideB=num and third will be h=num\n" +
		"  Rhombus   \trhombus, diamond, rhom.   \td1=[number] d2=[number] or b=[number] h=[number] or side=[number] a=[number]\n" +
		"  Ellipse  \tellipse, ellip.\n"
	areaExample = "\n./calculator -a square 105   Output: Resutl: 420\n" +
		"./calculator --area rhom. side=20 a=60   Output: Resutl: 346.410161\n"
)

func init() {
	calcPrecision = &flags.Precision("p, precision", "%g", percisionUsage+" ", precisionExamples).CalculatorPrecision
	calcTrigonometry = &flags.Trigonometry("t, trigonometry", false, trigonometryUsage, trigonometryExamples).CalculatorTrigonometry
	calcFigureArea = &flags.Area("a, area", " ", areaUsage, areaExample).CalculatorFigureArea
}

func main() {
	flags.Parse()

	var operationResultCh = make(chan map[string]any)

	a := arithmetic.Arithmetic()
	g := geometry.Geometry(a)
	t := trigonometry.Trigonometry(a)

	calcFlags := flags.CalculatorFlags{
		CalculatorPrecision:    calcPrecision,
		CalculatorTrigonometry: calcTrigonometry,
		CalculatorFigureArea:   calcFigureArea,
	}

	mathFields := calculator.MathFields{ArithmeticArea: a, GeometryArea: g, TrigonometryArea: t}

	c := calculator.NewCalculator(calcFlags, mathFields)
	c.StartCalculation(operationResultCh, flags.Args())
	c.PrintResults(operationResultCh)
}
