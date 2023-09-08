package trigonometry

import (
	"calculator/utils"
	"math"
)

const (
	Sin      = "sin"
	Sinh     = "sinh"
	Cos      = "cos"
	Cosh     = "cosh"
	Tan      = "tan"
	Cot      = "cot"
	Sec      = "sec"
	Cosec    = "cosec"
	Hypot    = "hypot"
	RootN    = "rootN"
	PotencyN = "^N"
)

func SinFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	result := math.Sin(radians)
	return result
}

func SinHFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	result := math.Sinh(radians)
	return result
}

func CosineFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	result := math.Cos(radians)
	return result
}

func CosineHFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	result := math.Cosh(radians)
	return result
}

func TangentFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	result := math.Tan(radians)
	return result
}

func CotangentFn(degrees float64) float64 {
	radians := utils.DegreesToRadians(degrees)
	result := math.Atan(radians)
	return result
}

func SecantFn(degrees float64) float64 {
	result := 1 / CosineFn(degrees)
	return result
}

func CosecantFn(degrees float64) float64 {
	result := 1 / SinFn(degrees)
	return result
}

func HipotenuseFn(catheiA, catheiB float64) float64 {
	result := math.Hypot(catheiA, catheiB)
	return result
}

func RootNFn(value float64, root float64) float64 {
	var result = value
	counter := root
	for counter != 0 {
		counter--
		result /= root
	}
	return result
}

func PotencyNFn(value, exponent float64) float64 {
	result := math.Pow(value, value)
	return result
}
