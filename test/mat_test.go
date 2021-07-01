package test

import (
	"fmt"
	"testing"
	"time"

	"gonum.org/v1/gonum/mat"
)

func TestMul(t *testing.T) {
	m1 := mat.NewDense(2, 3, []float64{
		2.0, 9.0, 3.0,
		4.5, 6.7, 8.0,
	})

	m2 := mat.NewDense(3, 2, []float64{
		2.0, 9.0,
		4.5, 6.7,
		1.2, 3.0,
	})

	var m mat.Dense
	start := time.Now()
	m.Mul(m1, m2)
	fc := mat.Formatted(&m, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("c = %v", fc)
	t.Log(time.Since(start).Milliseconds())
}

func TestPow(t *testing.T) {
	m1 := mat.NewDense(2, 2, []float64{
		2, 3,
		4, 5,
	})

	var m mat.Dense
	start := time.Now()
	m.Pow(m1, 2)
	fc := mat.Formatted(&m, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("c = %v", fc)
	t.Log(time.Since(start).Milliseconds())

	m2 := mat.NewDense(3, 3, []float64{0.5})
	fc = mat.Formatted(m2, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("c = %v", fc)
}

func TestCopy(t *testing.T) {
	m := mat.NewDense(3, 3, nil)
	m1 := mat.NewDense(2, 3, []float64{
		2.0, 9.0, 3.0,
		4.5, 6.7, 8.0,
	})

	m.Copy(m1)
	m.SetRow(3, []float64{1})

	fc := mat.Formatted(m, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("c = %v", fc)
}
