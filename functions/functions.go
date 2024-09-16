package functions

import (
	"fmt"
	// "gonum.org/v1/gonum/blas/blas32"
	"math"
)

func NegativeAlpine(x [2]float64) float64 {
	var res float64 = 0
	for _, v := range x {
		res += math.Abs(float64(v*math.Sin(v) + 0.1*v))
	}
	return (-1) * res
}

func Rosenbrock(x [2]float64) float64 {
	var res float64 = 0
	return res
}

func Ackley(x [2]float64) float64 {
	var res float64 = 0
	return res
}

func Easom(x [2]float64) float64 {
	var res float64 = 0
	return res
}

func FourPeak(x [2]float64) float64 {
	var res float64 = 0
	return res
}

func EggCrate(x [2]float64) float64 {
	var res float64 = 0
	return res
}

func Test() {
	test := [2]float64{0, 0}
	res := NegativeAlpine(test)
	fmt.Println(res)
}
