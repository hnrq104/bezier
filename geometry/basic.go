package geometry

import "image"

//Diferente do que foi feito em outro pacote, (goalg)
//Aqui os nossos vetores terão partes inteiras, (para facilitar o plot)

// Retorna o determinante (ou cross product) dos dois vetores
func Cross(p1, p2 image.Point) int {
	return p1.X*p2.Y - p2.X*p1.Y
}

// Retorna > 0 se pipk está a "direita" (sentido anti-horario) de pipj
func Direction(pi, pj, pk image.Point) int {
	return Cross(pk.Sub(pi), pj.Sub(pi))
}

// Dado que pk é colinear com pi, pj, diz se ele está no segmento (pi,pj)
func onSegment(pi, pj, pk image.Point) bool {
	return (min(pi.X, pj.X) <= pk.X) && (pk.X <= max(pi.X, pj.X)) &&
		(min(pi.Y, pj.Y) <= pk.Y) && (pk.Y <= max(pi.Y, pj.Y))
}

// Retorna verdadeiro se o segmento (p1,p2) intersecta com (p3,p4)
// Isso provavelmente não será usado no projeto mas é divertido de ter

func SegmentIntersect(p1, p2, p3, p4 image.Point) bool {
	d1 := Direction(p3, p4, p1)
	d2 := Direction(p3, p4, p2)
	d3 := Direction(p1, p2, p3)
	d4 := Direction(p1, p2, p4)

	if ((d1 < 0 && d2 > 0) || (d1 > 0 && d2 < 0)) &&
		((d3 < 0 && d4 > 0) || (d3 > 0 && d4 < 0)) {
		return true
	} else if d1 == 0 && onSegment(p3, p4, p1) {
		return true
	} else if d2 == 0 && onSegment(p3, p4, p2) {
		return true
	} else if d3 == 0 && onSegment(p1, p2, p3) {
		return true
	} else if d4 == 0 && onSegment(p1, p2, p4) {
		return true
	}
	return false
}
