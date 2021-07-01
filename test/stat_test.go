package test

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

func TestStat(t *testing.T) {
	fs := []float64{
		32.32, 56.98, 21.52, 44.32,
		55.63, 13.75, 43.47, 43.34,
		12.34,
	}
	// 加权平均
	mean := stat.Mean(fs, nil)
	t.Log(mean)
	// 标准差
	variance := stat.Variance(fs, nil)
	t.Log(variance)
	// 方差
	sqrt := math.Sqrt(variance)
	t.Log(sqrt)
	// 中位数
	sort.Float64s(fs)
	media := stat.Quantile(0.5, stat.Empirical, fs, nil)
	t.Log(media)
}

func TestDense(t *testing.T) {
	matrix := mat.NewDense(3, 5, nil)
	t.Log(mat.Formatted(matrix, mat.FormatPython()))
	t.Log(mat.Formatted(matrix, mat.FormatMATLAB()))

	m := mat.NewDense(3, 3, []float64{
		2.0, 9.0, 3.0,
		4.5, 6.7, 8.0,
		1.2, 3.0, 6.0,
	})

	dst := []float64{2.0, 9.0, 3.0}

	col := mat.Col(dst, 1, m)

	fmt.Printf("col = %#v", col)
	fmt.Printf("dst = %#v", dst)

	f := mat.Norm(m, math.Inf(1))
	t.Log(f)
}
