package adder_sample

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	var value int = Add(1, 2)

	if value != 3 {
		t.Errorf("1 + 2 for int should be 3")
	}

	var fValue float32 = Add(float32(1.2), 3)
	if fValue != 4.2 {
		t.Errorf("1.2 + 3 for float32 should be 4.2")
	}

	t.Run("meu teste maroto", func(t *testing.T) {
		var f64Value float64 = Add(float64(math.Pi), 1)
		if f64Value != 4.141592653589793116 {
			t.Errorf("Pi + 1 for float64 should be %.20f", math.Pi+1)
		}
	})
}
