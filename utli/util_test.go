package util

import (
	"fmt"
	"testing"
)

func Test_SplitWords(t *testing.T) {
	s := "428347923749\r" +
		"247823749748234\n" +
		"742947924273489,74289742974923\t" +
		"423742949734729734\t " +
		"4789237492374928\r\n" +
		"492874 29492，482648'72'93%74298 "
	fmt.Println(SplitWords(s))
}

func Test_Calculate(t *testing.T) {
	var (
		x         = 999
		y         = 8989.5645
		d float64 = 2
		//nums         = []int{888, 999, 1000}
		s = "100000hh"
	)

	fmt.Println(Add(x, y, d, s))
	fmt.Println(Sub(100, 2000, 2, 13.888))
	fmt.Printf("%.6e \n", Mul(100, 2000, 2, 100.1111))
	fmt.Println(Add(1, 1, 2))
	fmt.Println(Mul(100, 100, 2))

	fmt.Println(Div(x, y, d))
	fmt.Println(Float64("8888.99"), Float64("hh999999"))

}
