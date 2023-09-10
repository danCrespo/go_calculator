package geometry

import (
	"calculator/utils"
	"errors"
	"fmt"
	"math"
	"strings"

	"golang.org/x/exp/slices"
)

func area(measures []string, figure string) (float64, error) {
	var result float64

	if err := checkRequiredMeasures(measures, figure); err != nil {
		return result, err
	}

	switch figure {
	case utils.Square:
		result = squareArea(measures)
	case utils.Triangle:
		// if err := checkRequiredMeasures(measures, figure); err != nil {
		// 	return result, err
		// }
		result = triangleArea(measures)
	case utils.Circle:
		result = circleArea(measures)
	case utils.Rectangle:
		// if err := checkRequiredMeasures(measures, figure); err != nil {
		// 	return result, err
		// }
		result = rectangleArea(measures)
	case utils.Trapezoid:
		// if err := checkRequiredMeasures(measures, figure); err != nil {
		// 	return result, err
		// }
		result = trapezoidArea(measures)
	case utils.Rhombus:
		// if err := checkRequiredMeasures(measures, figure); err != nil {
		// 	return result, err
		// }
		result = rhombusdArea(measures)
	case utils.Ellipse:
		result = ellipseArea(measures)
	}

	return result, nil
}

func squareArea(measures []string) float64 {
	result := utils.StringToFloat64(measures[0]) * 4
	return result
}

func triangleArea(measures []string) float64 {
	var (
		base        float64
		height      float64
		measuresStr = strings.Join(measures, " ")
	)

	if utils.StringContains(measuresStr, "b=") && utils.StringContains(measuresStr, "h=") {
		slices.Sort[[]string](measures)
		measuresStr = strings.Join(measures, " ")

		fmt.Sscanf(measuresStr, "b=%f h=%f", &base, &height)
	} else {
		fmt.Sscanf(measuresStr, "%f %f", &base, &height)
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
		base        float64
		height      float64
		measuresStr = strings.Join(measures, " ")
	)

	if utils.StringContains(measuresStr, "b=") && utils.StringContains(measuresStr, "h=") {
		slices.Sort[[]string](measures)
		measuresStr = strings.Join(measures, " ")

		fmt.Sscanf(measuresStr, "b=%f h=%f", &base, &height)
	} else {
		fmt.Sscanf(measuresStr, "%f %f", &base, &height)
	}

	result := (base * height) / 2
	return result
}

func trapezoidArea(measures []string) float64 {
	var (
		sideA       float64
		sideB       float64
		height      float64
		measuresStr = strings.Join(measures, " ")
	)

	if utils.StringContains(measuresStr, "sideA=") && utils.StringContains(measuresStr, "sideB=") && utils.StringContains(measuresStr, "h=") {
		slices.Sort[[]string](measures)
		measuresStr = strings.Join(measures, " ")

		fmt.Sscanf(measuresStr, "h=%f sideA=%f sideB=%f", &height, &sideA, &sideB)
	} else {
		fmt.Sscanf(measuresStr, "%f %f %f", &sideA, &sideB, &height)

	}
	result := ((sideA + sideB) / 2) * height
	return result
}

func rhombusdArea(measures []string) float64 {
	var (
		diagonalA   float64
		diagonalB   float64
		base        float64
		height      float64
		angle       float64
		side        float64
		result      float64
		measuresStr = strings.Join(measures, " ")
	)

	if utils.StringContains(measuresStr, "d1=") && utils.StringContains(measuresStr, "d2=") {
		slices.Sort[[]string](measures)
		measuresStr = strings.Join(measures, " ")

		fmt.Sscanf(measuresStr, "d1=%f d2=%f", &diagonalA, &diagonalB)
		result = (diagonalA * diagonalB) / 2
	}

	if utils.StringContains(measuresStr, "b=") && utils.StringContains(measuresStr, "h=") {
		slices.Sort[[]string](measures)
		measuresStr = strings.Join(measures, " ")

		fmt.Sscanf(measuresStr, "b=%f h=%f", &base, &height)
		result = base * height
	}

	if utils.StringContains(measuresStr, "side=") && utils.StringContains(measuresStr, "angle=") {
		slices.Sort[[]string](measures)
		measuresStr = strings.Join(measures, " ")

		fmt.Sscanf(measuresStr, "angle=%g side=%f", &angle, &side)
		x2 := math.Pow(side, 2)
		sin := math.Sin(utils.DegreesToRadians(angle))
		result = x2 * sin
	}

	return result
}

func ellipseArea(measures []string) float64 {

	return 0
}

func checkRequiredMeasures(measures []string, figure string) error {

	measuresLen := len(measures)
	errStr := new(strings.Builder)
	errStr.WriteString("Invalid pattern or number of arguments. ")

	switch figure {
	case utils.Triangle, utils.Rectangle:
		if measuresLen == 2 {
			m1, m2 := measures[0], measures[1]
			if utils.StringContains(m1, "=") && utils.StringContains(m2, "=") {
				if !utils.StringContains(m1, "b=", "h=") || !utils.StringContains(m2, "h=", "b=") {
					errStr.WriteString(fmt.Sprintf("Arguments to the %s must be of the following form: b=[number] h=[number]\n", figure))
					return errors.New(errStr.String())
				}
			}
		} else {
			errStr.WriteString(fmt.Sprintf("For %s the number of arguments must be of just 2", figure))
			return errors.New(errStr.String())
		}

	case utils.Trapezoid:
		if measuresLen == 3 {
			m1, m2, m3 := measures[0], measures[1], measures[2]
			if utils.StringContains(m1, "=") && utils.StringContains(m2, "=") && utils.StringContains(m3, "=") {
				if !utils.StringContains(m1, "sideA=") || !utils.StringContains(m2, "sideB=") || !utils.StringContains(m3, "h=") {
					errStr.WriteString(fmt.Sprintf("Arguments to the %s must be of the following form: sideA=[number] sideB=[number] h=[number]\n", figure))
					return errors.New(errStr.String())
				}
			}
		} else {
			errStr.WriteString(fmt.Sprintf("For %s the number of arguments must be of just 3", figure))
			return errors.New(errStr.String())
		}

	case utils.Rhombus:
		if measuresLen == 2 {
			m1, m2 := measures[0], measures[1]
			if utils.StringContains(m1, "=") && utils.StringContains(m2, "=") {
				if !utils.StringContains(m1, "d1=", "d2=", "b=", "h=", "side=", "angle=") || !utils.StringContains(m2, "d1=", "d2=", "b=", "h=", "side=", "angle=") {
					errStr.WriteString(fmt.Sprintf("Arguments to the %s must be of the following form: [d1|b|side]=number [d2|h|angle]=[number]\n", figure))
					return errors.New(errStr.String())
				}
			} else {
				errStr.WriteString(fmt.Sprintf("Arguments to the %s must be of the following form: [d1|b|side]=number [d2|h|angle]=[number]\n", figure))
				return errors.New(errStr.String())
			}
		} else {
			errStr.WriteString(fmt.Sprintf("For %s the number of arguments must be of just 2", figure))
			return errors.New(errStr.String())
		}
	}

	return nil
}
