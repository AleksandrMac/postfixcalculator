package slowturingmachine

import (
	"fmt"
	"testing"

	values "github.com/AleksandrMac/postfixcalculator/common"
)

func ExampleRPNSlowTuringMachine() {
	fmt.Printf("%f", RPNSlowTuringMachine(values.RPNInput))
	// Output: 30.000000
}

func ExampleRPNSlowTuringMachine_exp() {
	fmt.Printf("%.13f", RPNSlowTuringMachine(values.Input))
	// Output: 3.0001220703125
}

func TestRPNSlowTuringMachine(t *testing.T) {
	if got := RPNSlowTuringMachine(values.RPNInput); got != values.RPNInputWant {
		t.Errorf("RPN(%s): got %f, want %f", values.RPNInput, got, values.RPNInputWant)
	}
}

func BenchmarkRPNSlowTuringMachine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNSlowTuringMachine(values.Input)
	}
}
