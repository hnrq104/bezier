package geometry

import (
	"image"
	"testing"
)

func TestCross(t *testing.T) {
	c := Cross(image.Pt(1, 0), image.Pt(0, 1))
	if !(c == 1) {
		t.Errorf("Cross((1,0), (0,1)) = %d: want %d", c, 1)
	}
}

func TestDirection(t *testing.T) {
	c := Direction(image.Pt(0, 0), image.Pt(0, 1), image.Pt(1, 0))
	if !(c == 1) {
		t.Errorf("Direction((0,0), (1,0), (0,1)) = %d: want %d", c, 1)
	}
}

func TestGrahamScan(t *testing.T) {
	points := []image.Point{
		image.Pt(-2, 3),
		image.Pt(0, 4),
		image.Pt(3, 4),
		image.Pt(-1, 2),
		image.Pt(0, 1),
		image.Pt(3, 1),
		image.Pt(-2, 0),
		image.Pt(2, 0),
		image.Pt(0, -2),
	}

	answer := []image.Point{
		image.Pt(0, -2),
		image.Pt(3, 1),
		image.Pt(3, 4),
		image.Pt(0, 4),
		image.Pt(-2, 3),
		image.Pt(-2, 0),
	}

	var correct = true

	c := GrahamScan(points)
	for i := range c {
		if i >= len(answer) || c[i] != answer[i] {
			correct = false
			break
		}
	}

	if !correct {
		t.Errorf("GrahamScan error: got %v : want %v", c, answer)
	}

}
