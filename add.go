// Samlpe module to learn how to create a Go module.
package adder_sample

import "golang.org/x/exp/constraints"

// Add the two numbers and return the result.
func Add[T constraints.Float | constraints.Integer](a, b T) T {
	return a + b
}
