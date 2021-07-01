package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

func main() {
	fs := 1
	sm := 20
	w := mat.NewDense(1, fs+1, nil)
	X := mat.NewDense(fs+1, sm, []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	})
	yf := []float64{
		3, 4, 5, 5, 2, 4, 7, 8, 11, 8, 12,
		11, 13, 13, 16, 17, 18, 17, 19, 21,
	}
	// 加权平均
	mean := stat.Mean(yf, nil)
	// 标准差
	variance := stat.Variance(yf, nil)
	floats.AddConst(-1*mean, yf)
	floats.Scale(1.0/variance, yf)
	y := mat.NewDense(1, sm, yf)

	fw := mat.Formatted(X, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("X = %v \n", fw)

	fw = mat.Formatted(y, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("Y = %v \n", fw)
	for i := 1; i <= 1000; i++ {
		los := dg(w, X, y)
		fw := mat.Formatted(w, mat.Prefix("    "), mat.Squeeze())
		fmt.Printf("iterator=%v : c = %v ------los=%v\n", i, fw, los)
		if los <= 0 {
			break
		}
	}
}

func dg(w *mat.Dense, X *mat.Dense, y *mat.Dense) (los float64) {
	fc, sm := X.Dims()

	c := mat.NewDense(1, sm, nil)
	c.Mul(w, X)

	c.Sub(c, y)
	c.Scale(1.0/float64(sm), c)

	dw := mat.NewDense(1, fc, nil)
	dw.Mul(c, X.T())

	a := 0.001
	dw.Scale(a, dw)
	w.Sub(w, dw)

	cs := mat.NewDense(1, sm, nil)
	cs.MulElem(c, c)
	los = mat.Sum(cs) / float64(sm*2)
	return
}
