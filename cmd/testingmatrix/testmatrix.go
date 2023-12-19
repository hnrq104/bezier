package main

import (
	"fmt"
	"splines/interpolation"
	"splines/matrix"
)

func main() {
	// a := []float64{
	// 	2, 0, 2, 0.6,
	// 	3, 3, 4, -2,
	// 	5, 5, 4, 2,
	// 	-1, -2, 3.4, -1,
	// }
	// A := matrices.NewMatrix(4, 4, a)

	// p, _ := matrices.LUPDecomposition(A)
	// // fmt.Println(a, p, err)
	// for i := 0; i < 4; i++ {
	// 	fmt.Println(a[4*i : 4*(i+1)])
	// }
	// fmt.Println(p)

	// a := []float64{
	// 	1, -1, 0, 0, 0,
	// 	-1, 2, -1, 0, 0,
	// 	0, -1, 2, -1, 0,
	// 	0, 0, -1, 2, -1,
	// 	0, 0, 0, -1, 2,
	// }

	// A := matrices.NewMatrix(5, 5, a)
	// fmt.Println(A.VecMul([]float64{1, 1, 1, 1, 1}))

	l := []float64{
		-1, -1, -1, -1,
	}
	m := []float64{
		1, 2, 2, 2, 2,
	}
	u := []float64{
		-1, -1, -1, -1,
	}

	t := matrix.NewTridiagonal(l, m, u)
	// matrices.LUP(t)
	// fmt.Println(matrices.LUPSolveReady(t, pi, []float64{1, 1, 1, 1, 1}))

	matrix.LU(t)

	x := matrix.LUSolveFast(t, []float64{1, 1, 1, 1, 1})
	fmt.Println(x)

	// fmt.Println(interpolation.FindDerivatives([]float64{3, -5, 40, 6, 7}))
	C := interpolation.Interpolate([]float64{3, -5, 40, 6, 7})
	for i := 0; i < len(C.Splines); i += 4 {
		var sum float64
		for j := i; j < i+4; j++ {
			sum += C.Splines[j]
			fmt.Printf("%.2f ", C.Splines[j])
		}
		fmt.Println("sum = ", sum)

	}

}
