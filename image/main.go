// package main

// import (
// 	"image"
// 	"image/color"
// 	"image/draw"
// 	"image/png"
// 	"log"
// 	"os"
// )

// func main() {
// 	data := []int{100, 33, 73, 64}
// 	w, h := len(data)*60+10, 100
// 	r := image.Rect(0, 0, w, h)
// 	img := image.NewRGBA(r)
// 	bg := image.NewUniform(color.RGBA{200, 200, 200, 255})

// 	draw.Draw(img, r, bg, image.Point{0, 0}, draw.Src)
// 	// for y := 0; y < h; y++ {
// 	// 	for x := 0; x < w; x++ {
// 	// 		img.Set(x, y, color.RGBA{255, 255, 255, 255})
// 	// 	}
// 	// }
// 	mask := image.NewNRGBA(image.Rect(0, 0, w, h))
// 	for y := 0; y < h; y++ {
// 		for x := 0; x < w; x++ {
// 			alpha := uint8(0)
// 			switch {
// 			case y < 30:
// 				alpha = 255
// 			case y < 50:
// 				alpha = 100
// 			}
// 			mask.Set(x, y, color.NRGBA{
// 				R: uint8((x + y) & 255),
// 				G: uint8((x + y) << 1 & 255),
// 				B: uint8((x + y) << 2 & 255),
// 				A: alpha,
// 			})
// 		}
// 	}

// 	for i, dp := range data {
// 		x0, y0 := (i*60 + 10), 100-dp
// 		x1, y1 := (i+1)*60-1, 100
// 		bar := image.Rect(x0, y0, x1, y1)
// 		grey := image.NewUniform(color.RGBA{180, 180, 180, 255})
// 		draw.Draw(img, bar, grey, image.Point{0, 0}, draw.Src)

// 		blue := image.NewUniform(color.RGBA{0, 0, 255, 255})
// 		draw.DrawMask(img, bar, blue, image.Point{0, 0}, mask, image.Point{x0, y0}, draw.Over)
// 		// for x := i*60 + 10; x < (i+1)*60; x++ {
// 		// 	for y := 100; y >= (100 - dp); y-- {
// 		// 		img.Set(x, y, color.RGBA{180, 180, 250, 250})
// 		// 	}
// 		// }
// 	}

// 	f, err := os.Create("image.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
// 	if err := png.Encode(f, img); err != nil {
// 		f.Close()
// 		log.Fatal(err)
// 	}
// }
package main

import (
	"math/rand"
	"os"
	"time"

	svg "github.com/ajstarks/svgo"
)

func rn(n int) int { return rand.Intn(n) }

func main() {
	canvas := svg.New(os.Stdout)
	data := []struct {
		Month string
		Usage int
	}{
		{"Jan", 171},
		{"Feb", 180},
		{"Mar", 100},
		{"Apr", 87},
		{"May", 66},
		{"Jun", 40},
		{"Jul", 32},
		{"Aug", 55},
	}
	width := len(data)*60 + 10
	height := 300
	threshold := 120
	max := 0
	for _, item := range data {
		if item.Usage > max {
			max = item.Usage
		}
	}
	rand.Seed(time.Now().Unix())
	canvas.Start(width, height)
	for i, val := range data {
		percent := val.Usage * (height - 50) / max
		canvas.Rect(i*60+10, (height-50)-percent, 50, percent, "fill:rgb(77,200,232)")
		canvas.Text(i*60+35, height-24, val.Month, "font-size:14pt;fill:rgb(150,150,150);text-anchor:middle")
	}
	threshPercent := threshold * (height - 50) / max
	canvas.Line(0, height-threshPercent, width, height-threshPercent, "stroke:rgb(250,100,100);opacity:0.8;stroke-width:2px")
	canvas.Rect(0, 0, width, height-threshPercent, "fill:rgb(255,100,100);opacity:0.3")
	canvas.Line(0, height-50, width, height-50, "stroke: rgb(150,150,150);stroke-width:2")
	canvas.End()
}
