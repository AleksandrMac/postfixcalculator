package emulatingstack

import (
	"fmt"
	"testing"

	values "github.com/AleksandrMac/postfixcalculator/common"
)

func ExampleRPNEmulatingStack() {
	fmt.Printf("%f", RPNEmulatingStack(values.RPNInput))
	// Output: 30.000000
}

func ExampleRPNEmulatingStack_exp() {
	fmt.Printf("%.13f", RPNEmulatingStack(values.Input))
	// Output: 3.0001220703125
}

func TestRPNEmulatingStack(t *testing.T) {
	if got := RPNEmulatingStack(values.Input); got != values.InputWant {
		t.Errorf("RPN(%s): got %f, want %f", values.Input, got, values.InputWant)
	}
}

// To check panic if expression is invalid.
func TestRPNEmulatingStack_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	RPNEmulatingStack(values.InvalidInput)
}

func BenchmarkRPNEmulatingStack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RPNEmulatingStack(values.Input)
	}
}
