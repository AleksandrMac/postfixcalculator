// Package slowturingmachine implements postfix calculator using a Turing machine using
// a regular for loop.
//
// Code is Shorter that the fast version of the Turing machine as it is using a "while" format
// but it remains slower than the usual for structure.
package slowturingmachine

import (
	"math"
	"strconv"
	"strings"

	values "github.com/AleksandrMac/postfixcalculator/common"
)

// RPNSlowTuringMachine returns the result of a string in reverse polish notation (postfix)
// using a Turing machine.
//
// The original band in the words exploded in a slice and results are held on the band but in num form.
// The band has two copies one in string and one in float. This is mandatory to avoid costly conversions.
// A processed operation or value is erased by replacing it with ? which is a reserved sign.
// A calculated value is marked as num which is also a reserved word.
// An invalid sign is interpreted as a value and the next operation panics.
func RPNSlowTuringMachine(RPNInput string) float64 {
	words := strings.Fields(RPNInput)
	numbers := make([]float64, len(words))
	i, index := 0, 0
	leftOp, rightOp := 0.0, 0.0
	var err error
	for index < len(words) {
		w := words[index]
		if strings.Contains(values.OperatorsList, w) { //"?" is always skipped
			//at least one operand exists
			//no number can be used for detection so "num" is checked
			//if word is ?, no number there, move before
			//if word is num, already converted, read numbers
			//otherwise attempt conversion
			i = index - 1
			for words[i] != "num" && i >= 0 {
				if words[i] == "?" {
					i--
				} else if words[i] != "num" {
					numbers[i], err = strconv.ParseFloat(words[i], 64)
					if err != nil {
						panic("Invalid right operand")
					}
					words[i] = "num" //conversion done
				}
			}
			//w is an operator and empty cells are skipped
			if w == "sqrt" {
				//unary operator
				numbers[i] = math.Sqrt(numbers[i])
			} else {
				//binary operator
				rightOp = numbers[i]
				words[i] = "?" //erasing value
				//looking for the left operand which is to the left...
				i--
				for words[i] != "num" && i >= 0 {
					if words[i] == "?" {
						i--
					} else if words[i] != "num" {
						numbers[i], err = strconv.ParseFloat(words[i], 64)
						if err != nil {
							panic("Invalid left operand")
						}
						words[i] = "num" //conversion done
					}
				}
				leftOp = numbers[i] //cell not erased as it keeps result
				switch w {
				case "+":
					numbers[i] = leftOp + rightOp
				case "-":
					numbers[i] = leftOp - rightOp
				case "*":
					numbers[i] = leftOp * rightOp
				case "/":
					numbers[i] = leftOp / rightOp
				case "^":
					numbers[i] = math.Pow(leftOp, rightOp)
				default:
					panic("Invalid operator : " + w)
				}
			}
			//restarting from the beginning of the expression by breaking operators for
			words[index] = "?" //erasing operator as operation is completed
			index = 0          //re-starting without re-init for. No the cleanest
		} else {
			index++
		}

	}
	/*counting remaining ops is not necessary. If you index goes above the size, you reached
	the end of the band and the value is always in the first position for a valid
	expression.
	*/
	return numbers[0]
}
