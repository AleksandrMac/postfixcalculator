package stack

import (
	"fmt"
	"testing"

	values "github.com/AleksandrMac/postfixcalculator/common"
)

/*
Running exemples where fmt is needed to check output.
*/
func ExampleRPNStack() {
	fmt.Printf("%.13f", RPNStack(values.Input))
	// Output: 3.0001220703125
}

func ExampleRPNStack_sqrt() {
	fmt.Printf("%f", RPNStack(values.RPNInput))
	// Output: 30.000000
}

func TestRPNStack(t *testing.T) {
	if got := RPNStack(values.RPNInput); got != values.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", values.Input, got, values.RPNInputWant)
	}
}

func BenchmarkRPNStack(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		RPNStack(values.Input)
	}
}
