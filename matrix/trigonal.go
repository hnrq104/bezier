package matrix

import (
	"fmt"
)

// tridiagonal matrix (useful for splines)
// does it have to be square? i think so
type Tridiagonal struct {
	rows, columns        int
	lower, middle, upper []float64
}

// will fragment memory, but it's a risk i'm willing to make

func NewTridiagonal(l, m, u []float64) Tridiagonal {
	t := Tridiagonal{rows: len(m), columns: len(m)}
	if len(l) != len(u) || len(l) != len(m)-1 {
		panic("Can't form tridiagonal matrix with given input")
	}
	t.lower, t.middle, t.upper = l, m, u
	return t
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (t Tridiagonal) At(i, j int) float64 {
	if i < 0 || j < 0 || i >= t.rows || j >= t.columns {
		panic("bad access at tridiagonal matrix")
	}

	if j == i+1 {
		return t.upper[i]
	}

	if j == i {
		return t.middle[i]
	}

	if j == i-1 {
		return t.lower[j]
	}

	return 0
}

func (t Tridiagonal) Set(i, j int, n float64) {
	if i < 0 || j < 0 || i >= t.rows || j >= t.columns {
		panic("bad access at tridiagonal matrix")
	}

	if j == i+1 {
		t.upper[i] = n
		return
	}

	if j == i {
		t.middle[i] = n
		return
	}

	if j == i-1 {
		t.lower[j] = n
		return
	}

	panic("bad setting location")
}

// LU DECOMPOSE TRIDIAGONAL IN PLACE
func LU(t Tridiagonal) error {
	n := t.rows
	for i := 0; i < n-1; i++ {
		if t.At(i, i) == 0 {
			return fmt.Errorf("Non positive Singular Matrix")
		}
		t.Set(i+1, i, t.At(i+1, i)/t.At(i, i))
		t.Set(i+1, i+1, t.At(i+1, i+1)-t.At(i+1, i)*t.At(i, i+1))
	}
	return nil
}

func (t Tridiagonal) Dim() (int, int) {
	return t.rows, t.columns
}

//PARA FAZER DIA 19/12 (RESOLVER LUP DA TRIGONAL - FAZER RODAR EM O(N))

// O sistema que será resolvido é simétrico e positivo definido

// não precisamos da LUP exatamente

// L and U are only bidiagonal matrices
func LUSolveFast(t Tridiagonal, b []float64) []float64 {
	if t.rows != len(b) {
		panic("wrong dimensions in LUSolve")
	}
	n, _ := t.Dim()
	x := make([]float64, n)
	x[0] = b[0]
	for i := 1; i < n; i++ {
		x[i] = b[i] - x[i-1]*t.At(i, i-1)
	}

	x[n-1] /= t.At(n-1, n-1)
	for i := n - 2; i >= 0; i-- {
		x[i] -= t.At(i, i+1) * x[i+1]
		x[i] /= t.At(i, i)
	}

	return x
}
