package drawing

import (
	"image"
	"image/color"
	"image/draw"
	"splines/interpolation"
)

//Essa é a etapa que eu considero mais interessante do projeto
// depois de pensar bastante, percebemos que podemos sim interpolar curvas
// fechadas.

// Basta abri-las em mais dimensões.

// i think 1/32 is good enough

func PolygonalSpline(pontos []image.Point, c color.Color, img draw.Image) {
	X := make([]float64, len(pontos))
	Y := make([]float64, len(pontos))

	for i := range pontos {
		X[i] = float64(pontos[i].X)
		Y[i] = float64(pontos[i].Y)
	}

	CX := interpolation.Interpolate(X)
	CY := interpolation.Interpolate(Y)

	// log.Print("finished interpolating X AND Y")
	n := CX.Len()
	for i := 0; i < n; i++ {
		old := pontos[i]
		for j := 0.; j <= float64(1); j += 1.0 / 32 {
			new := image.Pt(int(CX.At(i, j)), int(CY.At(i, j)))
			DrawLineP(old, new, c, 1, img)
			old = new
		}
	}
}
