package interpolation

import (
	"math"
	"splines/matrix"
)

type Curve struct {
	n       int //number of functions (one less than number of points)
	Splines []float64
}

func NewCurve(Y, D []float64) Curve {
	nfunc := len(Y) - 1

	spline := make([]float64, 4*nfunc) // each function has 4 parameter a,b,c,d

	//as described in 28-2 a these can be calculated as follow

	for i := 0; i < nfunc; i++ {
		//ai
		spline[4*i+0] = Y[i]
		//bi
		spline[4*i+1] = D[i]
		//ci
		spline[4*i+2] = 3*(Y[i+1]-Y[i]) - D[i+1] - 2*D[i]
		//di
		spline[4*i+3] = 2*(Y[i]-Y[i+1]) + D[i+1] + D[i]
	}

	return Curve{nfunc, spline}
}

func NewInterpolationMatrix(n int) matrix.Tridiagonal {
	l := make([]float64, n-1)
	u := make([]float64, n-1)
	m := make([]float64, n)

	for i := 0; i < n-1; i++ {
		l[i], u[i] = 1, 1
		m[i+1] = 4
	}
	m[0], m[n-1] = 2, 2
	return matrix.NewTridiagonal(l, m, u)
}

func deltaY(Y []float64) []float64 {
	n := len(Y)
	delta := make([]float64, n)

	delta[0] = 3 * (Y[1] - Y[0])
	delta[n-1] = 3 * (Y[n-1] - Y[n-2])

	for i := 1; i <= n-2; i++ {
		delta[i] = 3 * (Y[i+1] - Y[i-1])
	}
	return delta
}

// Find's D0,D1.. Dn (the derivatives at the each yi point)
func FindDerivatives(Y []float64) []float64 {

	n := len(Y)
	delta := make([]float64, n)

	delta[0] = 3 * (Y[1] - Y[0])
	delta[n-1] = 3 * (Y[n-1] - Y[n-2])

	for i := 1; i <= n-2; i++ {
		delta[i] = 3 * (Y[i+1] - Y[i-1])
	}

	// Trigonal interpolation matrix
	M := NewInterpolationMatrix(n)

	matrix.LU(M)
	return matrix.LUSolveFast(M, delta)
}

func Interpolate(X []float64) Curve {
	D := FindDerivatives(X)
	return NewCurve(X, D)
}

// becareful how you use this function, x will only have sensible data
// between [0,1]
func (c Curve) At(i int, x float64) float64 {
	if i < 0 || i >= c.n {
		return math.NaN()
	}

	// really ugly way to return a + bx + cx^2 + dx^3
	return c.Splines[4*i] + x*(c.Splines[4*i+1]+x*(c.Splines[4*i+2]+x*c.Splines[4*i+3]))

}

func (c Curve) Len() int {
	return c.n
}
