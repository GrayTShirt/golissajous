package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xdd, 0x00, 0xFF}}
const (
	whiteIndex = 0
	blackIndex = 1
)

type Settings struct {
	Cycles  float64
	Res     float64
	Size    int
	Nframes int
	Delay   int
}


func init() {
	rand.Seed(time.Now().UnixNano())
}

func New() *Settings {
	return &Settings{
		Cycles:  5,
		Res:     0.001,
		Size:    100,
		Nframes: 64,
		Delay:   8,
	}
}
func RenderGif(out io.Writer, params *Settings) {
	freq  := rand.Float64() * 3.0
	anim  := gif.GIF{LoopCount: params.Nframes}
	phase := 0.0
	for i := 0; i < params.Nframes; i++ {
		rect := image.Rect(0, 0, 2*params.Size+1, 2*params.Size+1)
		img  := image.NewPaletted(rect, palette)
		for t := 0.0; t < params.Cycles*2*math.Pi; t += params.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(params.Size + int(x * float64(params.Size) + 0.5), params.Size + int(y * float64(params.Size) + 0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, params.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
