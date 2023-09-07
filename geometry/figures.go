package geometry

import (
	"calculator/utils"
	"fmt"
	"math"
	"strings"
)

const (
	Square    = "square"
	Triangle  = "triangle"
	Circle    = "circle"
	Rectangle = "rectangle"
	Trapezoid = "trapezoid"
	Rhombus   = "rhombus"
	Ellipse   = "ellipse"
)

func Area(measures []string, figure string) float64 {
	switch figure {
	case Square:
		return squareArea(measures)
	case Triangle:
		return triangleArea(measures)
	case Circle:
		return circleArea(measures)
	case Rectangle:
		return rectangleArea(measures)
	case Trapezoid:
		return trapezoidArea(measures)
	case Rhombus:
		return rhombusdArea(measures)
	case Ellipse:
	}
	return 0
}

func squareArea(measures []string) float64 {
	result := utils.StringToFloat64(measures[0]) * 4
	return result
}

func triangleArea(measures []string) float64 {
	var (
		base   float64
		height float64
	)
	b := measures[0]
	h := measures[1]

	if strings.Contains(b, "b=") {
		fmt.Sscanf(b, "b=%f", &base)
	} else {
		fmt.Sscanf(b, "%f", &base)
	}

	if strings.Contains(h, "h=") {
		fmt.Sscanf(h, "h=%f", &height)
	} else {
		fmt.Sscanf(h, "%f", &height)
	}

	result := (base * height) / 2
	return result
}

func circleArea(measures []string) float64 {
	radius := utils.StringToFloat64(measures[0])
	radiusSquared := math.Pow(radius, 2)
	result := math.Pi * radiusSquared
	return result
}

func rectangleArea(measures []string) float64 {
	var (
		base   float64
		height float64
	)
	b := measures[0]
	h := measures[1]

	if strings.Contains(b, "b=") {
		fmt.Sscanf(b, "b=%f", &base)
	} else {
		fmt.Sscanf(b, "%f", &base)
	}

	if strings.Contains(h, "h=") {
		fmt.Sscanf(h, "h=%f", &height)
	} else {
		fmt.Sscanf(h, "%f", &height)
	}

	result := base * height
	return result
}

func trapezoidArea(measures []string) float64 {
	var (
		sideA  float64
		sideB  float64
		height float64
	)

	measuresStr := strings.Join(measures, " ")
	fmt.Sscanf(measuresStr, "a=%f b=%f h=%f", &sideA, &sideB, &height)
	result := ((sideA + sideB) / 2) * height
	return result
}

func rhombusdArea(measures []string) float64 {
	var (
		diagonalA float64
		diagonalB float64
		base      float64
		height    float64
		angle     float64
		result    float64
	)

	a := measures[0]
	b := measures[1]

	if strings.Contains(a, "d1=") && strings.Contains(b, "d2=") {
		fmt.Sscanf(a, "d1=%f", &diagonalA)
		fmt.Sscanf(b, "d2=%f", &diagonalB)
		result = (diagonalA * diagonalB) / 2
	}

	if strings.Contains(a, "b=") && strings.Contains(b, "h=") {
		fmt.Sscanf(a, "b=%f", &base)
		fmt.Sscanf(b, "h=%f", &height)
		result = base * height
	}

	if strings.Contains(a, "side=") && strings.Contains(b, "a=") {
		fmt.Sscanf(a, "side=%f", &base)
		fmt.Sscanf(b, "a=%f", &angle)
		result = math.Pow(base, 2) * math.Sin(angle)
	}

	return result
}
