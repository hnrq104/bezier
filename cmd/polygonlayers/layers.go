package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"splines/drawing"
	"splines/geometry"
	"time"
)

const (
	pointRectWidth  = 6
	pointRectHeight = 6
)

var width = flag.Int("w", 1024, "determines the width of the created image - standart 1024")
var height = flag.Int("h", 1024, "determines the height of the created image - standart 1024")
var numPoints = flag.Int("n", 100, "determines the number of randomly generated points in space")
var dst = flag.String("file", "polygon", "determines the name of the output file")

func main() {
	flag.Parse()

	img := image.NewRGBA(image.Rect(0, 0, *width, *height))

	colorsRect := []color.Color{
		color.RGBA{0xff, 0, 0, 0xff},
		color.RGBA{0, 0xff, 0, 0xff},
		// color.RGBA{0, 0, 0xff, 0xff},
		color.White,
	}

	colors := []color.Color{
		color.RGBA{0xff, 0xf0, 0, 0xff}, // ORANGE
		color.RGBA{0xab, 0xab, 0, 0xff}, // YELLOW
		color.RGBA{0, 0xff, 0, 0xff},    // GREEN
		color.RGBA{0x60, 0, 0xf0, 0xff}, // INDIGO
		color.RGBA{0xff, 0, 0xf0, 0xff}, // VIOLET
		color.RGBA{0xff, 0, 0, 0xff},    // RED

	}

	generate := time.Now()

	points := make([]image.Point, 0, *numPoints)
	mapP := make(map[image.Point]bool, *numPoints)
	for i := 0; i < *numPoints; i++ {
		pt := image.Pt(*width/8+rand.Intn(3*(*width)/4), *height/8+rand.Intn(3*(*height)/4))
		points = append(points, pt)
		mapP[pt] = true
	}

	draw.Draw(img, img.Bounds(), image.NewUniform(color.Black), image.Point{}, draw.Src)

	for _, p := range points {
		littleRect := image.Rect(p.X-pointRectWidth/2, p.Y-pointRectHeight/2,
			p.X+pointRectWidth/2, p.Y+pointRectHeight/2)
		draw.Draw(img, littleRect, image.NewUniform(colorsRect[rand.Intn(len(colorsRect))]),
			image.Pt(0, 0), draw.Src)
	}

	log.Printf("Took %dms to generate and draw %d rectangles\n", time.Since(generate).Milliseconds(), *numPoints)

	work := time.Now()

	for len(points) >= 3 {
		polygon := geometry.GrahamScan(points)

		for _, p := range polygon {
			mapP[p] = false
		}

		points = points[:0]
		for k, v := range mapP {
			if v {
				points = append(points, k)
			}
		}

		polygon = append(polygon, polygon[0])

		for i := 0; i < len(polygon)-1; i++ {
			drawing.DrawLineP(polygon[i], polygon[i+1], color.White, 1, img)
		}

		drawing.PolygonalSpline(polygon, colors[rand.Intn(len(colors))], img)

	}

	log.Printf("Took %dms to determine layers and interpolate all of them!", time.Since(work).Milliseconds())

	file, err := os.Create(*dst)
	if err != nil {
		log.Fatalf("Could not create file %q: %v", *dst, err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
}
