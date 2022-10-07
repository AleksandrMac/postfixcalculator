package operatorslist

import (
	"fmt"
	"testing"

	values "github.com/AleksandrMac/postfixcalculator/common"
)

/* if fmt.Print has an output (debug mode) */
func ExampleRPNOperatorsList() {
	fmt.Printf("%f", RPNOperatorsList(values.RPNInput))
	// Output: 30.000000
}

func ExampleRPNOperatorsList_exp() {
	fmt.Printf("%.13f", RPNOperatorsList(values.Input))
	// Output: 3.0001220703125
}

func TestRPNOperatorsList(t *testing.T) {
	if got := RPNOperatorsList(values.RPNInput); got != values.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", values.RPNInput, got, values.RPNInputWant)
	}
}

func BenchmarkRPNOperatorsList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNOperatorsList(values.Input)
	}
}
