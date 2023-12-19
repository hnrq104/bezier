package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"splines/drawing"
	"splines/geometry"
)

const (
	pointRectWidth  = 6
	pointRectHeight = 6
)

func main() {

	const (
		width     = 512
		height    = 512
		numPoints = 100
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	colors := []color.Color{
		color.RGBA{0xff, 0, 0, 0xff},
		color.RGBA{0, 0xff, 0, 0xff},
		color.RGBA{0, 0, 0xff, 0xff},
	}

	// genereate 100 random points
	points := make([]image.Point, 0, numPoints)
	for i := 0; i < numPoints; i++ {
		points = append(points, image.Pt(width/4+rand.Intn(width/2), height/4+rand.Intn(height/2)))
	}

	// draw little rectangles
	// will do all concurrently later
	draw.Draw(img, img.Bounds(), image.NewUniform(color.White), image.Point{}, draw.Src)

	for _, p := range points {
		littleRect := image.Rect(p.X-pointRectWidth/2, p.Y-pointRectHeight/2,
			p.X+pointRectWidth/2, p.Y+pointRectHeight/2)
		draw.Draw(img, littleRect, image.NewUniform(colors[rand.Intn(len(colors))]),
			image.Pt(0, 0), draw.Src)
	}

	// for each point draw a small rectangle

	polygon := geometry.GrahamScan(points)

	for i := 0; i < len(polygon); i++ {
		j := (i + 1) % len(polygon)
		drawing.DrawLineP(polygon[i], polygon[j], color.Black, img)
	}

	file, err := os.Create("img.png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
}
