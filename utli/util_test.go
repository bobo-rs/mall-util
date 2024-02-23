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

func Test_1688Sign(t *testing.T) {
	// 223f39ef64a35582a1b48ed2f1aecace_1708680687662
	fmt.Println(`1688原加密值：3d592bc663ac75436b5ab5547731f3a3`)
	var sign = `223f39ef64a35582a1b48ed2f1aecace&1708672769659&12574478&{"dataType":"moduleData","argString":"{\"appName\":\"offerlistAlisite\",\"memberId\":\"b2b-3930917912e7bdf\",\"appdata\":{\"sortType\":\"ninetytradenumdowndesc\",\"count\":2,\"showMode\":\"single\"}}"}`
	fmt.Println(Sign1688(sign))
}