package geometry

import (
	"image"
	"sort"
)

type sortPolarPoints struct {
	origin image.Point
	points []image.Point
}

func (s *sortPolarPoints) Len() int { return len(s.points) }
func (s *sortPolarPoints) Swap(i, j int) {
	s.points[i], s.points[j] = s.points[j], s.points[i]
}

// Nos ordenamos em relação ao angulo polar (0,pi) da origem
func (s *sortPolarPoints) Less(i, j int) bool {
	return Direction(s.origin, s.points[j], s.points[i]) > 0
}

// (AINDA NÃO IMPLEMENTADO) Diferente do implementado anteriormente, eu vou permitir pontos colineares (mas vou ordena-los corretamente)
// se não ao tirar a proxima layer ficará ruim

func GrahamScan(points []image.Point) []image.Point {
	if len(points) < 3 {
		return nil
	}

	swap := lowestLeftmost(points)
	points[0], points[swap] = points[swap], points[0]

	sort.Sort(&sortPolarPoints{points[0], points[1:]})

	//remove pontos colineares

	sorted := make([]image.Point, 0, len(points))

	for i := 1; i < len(points); i++ {
		farthest := points[i]
		for i+1 < len(points) && Direction(points[0], farthest, points[i+1]) == 0 {
			if moduleSquare(farthest.Sub(points[0])) < moduleSquare(points[i+1].Sub(points[0])) {
				farthest = points[i+1]
			}
			i++
		}
		sorted = append(sorted, farthest)
	}

	stack := make([]image.Point, 0, len(points)/4)
	stack = append(stack, points[0], sorted[0], sorted[1])

	for i := 2; i < len(sorted); i++ {
		for Direction(stack[len(stack)-2], stack[len(stack)-1], sorted[i]) >= 0 {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, sorted[i])
	}
	return stack

}

func lowestLeftmost(points []image.Point) int {
	min := 0
	for i := 1; i < len(points); i++ {
		if points[i].Y < points[min].Y {
			min = i
		} else if points[i].Y == points[min].Y && points[i].X < points[min].X {
			min = i
		}
	}

	return min
}

func moduleSquare(img image.Point) int {
	return img.X*img.X + img.Y*img.Y
}
