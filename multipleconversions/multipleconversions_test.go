package multipleconversions

import (
	"fmt"
	"testing"

	values "github.com/AleksandrMac/postfixcalculator/common"
)

func ExampleRPN() {
	fmt.Printf("%f", RPN(values.RPNInput))
	// Output: 30.000000
}

func ExampleRPN_exp() {
	fmt.Printf("%.13f", RPN(values.Input))
	// Output: 3.0001220703125
}

func TestRPN(t *testing.T) {
	if got := RPN(values.RPNInput); got != values.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", values.Input, got, values.RPNInputWant)
	}

}

func BenchmarkRPN(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPN(values.Input)
	}
}
