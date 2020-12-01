package main

import (
	"fmt"
	"math"
)

func main() {
	s := GetValues()
	fn:= GenDisplaceFn(s)

	fmt.Println(fn(3))
}

func GetValues() []float64 {
	var acc, vel, dis float64
	sli := make([]float64, 3)

	fmt.Println("Enter a acceleration:")
	fmt.Scanln(&acc)
	sli[0] = acc

	fmt.Println("Enter a velocity:")
	fmt.Scanln(&vel)
	sli[1] = vel

	fmt.Println("Enter a displacement:")
	fmt.Scanln(&dis)
	sli[2] = dis

	return sli
}

func GenDisplaceFn(sli []float64) func(float64) float64 {
	a := sli[0]
	v := sli[1]
	d := sli[2]
	fn := func(t float64) float64 {
		return (0.5 * a*math.Pow(t, 2) + v * t) + d
	}
	return fn
}
