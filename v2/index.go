// This package aims to help developers to sum 2 numbers.
// Please contact adrianlocurcio95@gmail.com if you need custom support for it.
package sum_numbers

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Sum 2 numbers and return the result
func Add[T Number](a, b T) T {
	return a + b
}