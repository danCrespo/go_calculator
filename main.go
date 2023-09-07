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

	calcFigureArea *flags.CalculatorFigureArea
	areaUsage      = "Used to calculate the area of a figure given their measures.\nAccepts the `figure name` or an alias of the name; the arguments of the measures will depend of the figure.\n" +
		"The follow is a list of accepted figures, their aliases and the arguments to be passed:\n" +
		"  Name      \tAliases                   \tArguments\n" +
		"  Square    \tsquare, sqr.              \tone argument (measure) [number]\n" +
		"  Triangle  \ttriangle, tri.            \tb=[number] h=[number] or just [number] [number]\n" +
		"  Circle    \tcircle, circ.             \tone argument (measure) this would be the radius\n" +
		"  Rectangle \trectangle, rect.          \tb=[number] h=[number] or just [number] [number]\n" +
		"  Trapezoid \ttrapezoid, trapeze, trap. \ta=[number] b=[number] h=[number]\n" +
		"  Rhombus   \trhombus, diamond, rhom.    \td1=[number] d2=[number] or b=[number] h=[number] or side=[number] a=[number]\n" +
		"  Ellipse  \tellipse, ellip.\n"
	areaExample = "\n./calculator -a square 105   Output: Resutl: 420\n" +
		"\n./calculator --area rhom side=20 a=60   Output: Resutl: 346.410161\n"
)

func init() {
	calcPrecision = &flags.Precision("p, precision", "%g", percisionUsage+" ", precisionExamples).CalculatorPrecision
	calcHierachy = &flags.Hierachy("H, hierachy", true, hierachyUsage, hierachyExamples).CalculatorHierarchy
	calcFigureArea = &flags.Area("a, area", " ", areaUsage, areaExample).CalculatorFigureArea
}

func main() {
	flags.Parse()

	var operationResultCh = make(chan float64)

	calcFlags := flags.CalculatorFlags{
		CalculatorPrecision:  calcPrecision,
		CalculatorHierarchy:  calcHierachy,
		CalculatorFigureArea: calcFigureArea,
	}

	c := calculator.NewCalculator(calcFlags)
	c.StartCalculation(operationResultCh, flags.Args())
	c.PrintResults(operationResultCh)
}
