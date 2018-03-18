package RPN

import (
	"fmt"
	"testing"
)

/* if fmt.Print has an output (debug mode) */
func ExampleRPN_operators_no_fields() {
	fmt.Printf("%f", RPN_operators_no_fields(RPNInput))
	// Output: 30.000000
}

func ExampleRPN_operators_no_fieldsExp() {
	fmt.Printf("%.13f", RPN_operators_no_fields(input))
	// Output: 3.0001220703125
}

/* */

func TestRPN_operators_no_fields(t *testing.T) {
	actual := RPN_operators_no_fields(RPNInput)
	if actual != 30 {
		t.Errorf("RPN(%s): expected %d, actual %f", RPNInput, 30, actual)
	}
}

/*
go test -bench=. allows to test the func
*/
func BenchmarkRPN_operators_no_fields(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPN_operators_no_fields(RPNInput)
	}
}