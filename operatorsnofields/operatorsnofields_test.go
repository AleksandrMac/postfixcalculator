package operatorsnofields

import (
	"fmt"
	"testing"

	values "github.com/AleksandrMac/postfixcalculator/common"
)

/* if fmt.Print has an output (debug mode) */
func ExampleRPNOperatorsNoFields() {
	fmt.Printf("%f", RPNOperatorsNoFields(values.RPNInput))
	// Output: 30.000000
}

func ExampleRPNOperatorsNoFields_exp() {
	fmt.Printf("%.13f", RPNOperatorsNoFields(values.Input))
	// Output: 3.0001220703125
}

func TestRPNOperatorsNoFields(t *testing.T) {
	if got := RPNOperatorsNoFields(values.RPNInput); got != values.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", values.RPNInput, got, values.RPNInputWant)
	}
}

func BenchmarkRPNOperatorsNoFields(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNOperatorsNoFields(values.Input)
	}
}
