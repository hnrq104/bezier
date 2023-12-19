package drawing

import (
	"image"
	"image/color"
	"image/draw"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func DrawLine(x0, y0, x1, y1 int) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)

	sx := 1
	sy := 1
	if x0 > x1 {
		sx = -1
	}
	if y0 > y1 {
		sy = -1
	}

	err := dx - dy

	for {
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func DrawLineP(p0, p1 image.Point, c color.Color, img draw.Image) {
	dx := abs(p0.X - p1.X)
	dy := abs(p0.Y - p1.Y)

	sx := 1
	sy := 1
	if p0.X > p1.X {
		sx = -1
	}
	if p0.Y > p1.Y {
		sy = -1
	}

	err := dx - dy
	for {
		img.Set(p0.X, p0.Y, c)
		if p0.X == p1.X && p0.Y == p1.Y {
			break
		}
		e2 := 2 * err
		if e2 >= -dy {
			err -= dy
			p0.X += sx
		}
		if e2 <= dx {
			err += dx
			p0.Y += sy
		}
	}
}
