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
	x1 := x[0]
	x2 := x[0]
	res = math.Pow((x1-1), 2) + 100*math.Pow(x2-math.Pow(x1, 2), 2)
	return res
}

func Ackley(x [2]float64) float64 {
	var res float64 = 0
	x1 := x[0]
	x2 := x[0]
	p1 := math.Exp((-0.02) * math.Sqrt(0.5) * (math.Pow(x1, 2) + math.Pow(x2, 2)))
	p2 := math.Exp((0.05) * (math.Cos(2*math.Pi*x1) + math.Cos(2*math.Pi*x2)))
	res = -20*p1 - p2 + 20 + math.E
	return res
}

func Easom(x [2]float64) float64 {
	var res float64 = 0
	x1 := x[0]
	x2 := x[0]
	res = -(math.Cos(x1) * math.Cos(x2) * math.Exp(-(math.Pow(x1-math.Pi, 2)))) + (-math.Pow(x2-math.Pi, 2))
	return res
}

func FourPeak(x [2]float64) float64 {
	var res float64 = 0
	x1 := x[0]
	x2 := x[0]
	p1 := -(math.Pow(x1-4, 2)) - (math.Pow(x2-4, 2))
	p2 := -(math.Pow(x1+4, 2)) - (math.Pow(x2-4, 2))
	p3 := -(math.Pow(x1, 2)) - math.Pow(x2, 2)
	p4 := -(math.Pow(x1, 2)) - (math.Pow(x2+4, 2))
	res = math.Exp(p1) + math.Exp(p2) + math.Exp(p3) + math.Exp(p4)
	return res
}

func EggCrate(x [2]float64) float64 {
	var res float64 = 0
	x1 := x[0]
	x2 := x[0]
	res = math.Pow(x1, 2) + math.Pow(x2, 2) + 25*(math.Pow(math.Sin(x1), 2)+math.Pow(math.Sin(x2), 2))
	return res
}

func Test() {
	test := [2]float64{0, 0}
	res := NegativeAlpine(test)
	fmt.Println(res)
}
