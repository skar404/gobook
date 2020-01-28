package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	out := os.Stdout

	const (
		cycles  = 5    // Количество полных колебаний x
		res     = .001 // Угловое разрешение
		size    = 100  // Канва изображения охват [size.. + size]
		nframes = 64   // Количество кадров
		delay   = 8    // Задержка между кадрами
	)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3. // Относительная цастота колебаний
	anim := gif.GIF{LoopCount: nframes}
	phase := 0. // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+.5), size+int(y*size+.5), blackIndex)
		}
		phase += .1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)

	if err != nil {
		fmt.Println("Error save file err=", err)
	}
}
