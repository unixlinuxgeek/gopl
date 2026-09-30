// Exercise 1.6:
// Modify the Lissajous program to produce images in multiple colors by adding more values to palette and
// then displaying them by changing the third argument of SetColorIndex in some interesting way.

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
)

var palette = []color.Color{
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}, // black
	color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, // red
	color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}, // green
	color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}, // blue
}

const (
	redIndex   = 0 // first color in palette
	greenIndex = 1 // next color in palette
	blueIndex  = 2 // end color in palette
)

func main() {
	// delete old png images if exists or exit from loop
	for {
		err := os.Remove("*.png")
		if err != nil {
			break
		}
	}

	for i, _ := range palette[1:] {
		f, err := os.Create(strconv.Itoa(i+1) + ".png")
		if err != nil {
			fmt.Fprintf(os.Stderr, "create: %v\n", err)
			os.Exit(1)
		}
		lissajous(f, i+1)
		f.Close() // NOTE: ignoring Unhandled error
	}
}

func lissajous(out io.Writer, colInd int) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colInd))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
