// ex1.12 serves the lissajous gif on a web server
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{G: 0xff, A: 0xff}}

const (
	blackIndex = 0 // first colour in palette
	greenIndex = 1 // next color in palette
)

func main() {
	defaultCycles := 5

	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission
	rand.Seed(time.Now().UTC().UnixNano())
	handler := func(w http.ResponseWriter, r *http.Request) {
		cycleStr := r.URL.Query().Get("cycles")
		if cycleStr != "" {
			cycles, err := strconv.Atoi(cycleStr)
			if err != nil {
				fmt.Fprintf(w, "Error getting cycles. Given: %s Error: %v", cycleStr, err)
				return
			}
			lissajous(w, cycles)
		}
		lissajous(w, defaultCycles)
	}
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int) {
	const (
		res = 0.001  // angular resolution
		size = 100   // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames
		delay = 8    // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
	if err != nil {
		fmt.Println(err)
	}
}