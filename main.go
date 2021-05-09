package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"path/filepath"

	"github.com/MeNoln/fractal/palette"
)

const (
	max_iter = 800
)

var (
	samples = []*complexSample{
		{
			x_lower: -2.5,
			x_upper: 1.0555555,
			y_lower: -1,
			y_upper: 1,
			fname:   "image1.png",
		},
		{
			x_lower: -0.7512082704194535,
			x_upper: -0.7189233668101299,
			y_lower: 0.171229302623112,
			y_upper: 0.1954429803301047,
			fname:   "image2.png",
		},
		{
			x_lower: -0.8430837175066829,
			x_upper: -0.7097679463836053,
			y_lower: 0.1298062156920607,
			y_upper: 0.2297930440343688,
			fname:   "image3.png",
		},
		{
			x_lower: -0.7436848654999562,
			x_upper: -0.7382828140917086,
			y_lower: 0.1820790907660431,
			y_upper: 0.1861306293222289,
			fname:   "image4.png",
		},
		{
			x_lower: -1.2578645782749025,
			x_upper: -1.2575820364511403,
			y_lower: 0.0293170055498651,
			y_upper: 0.0295289119176868,
			fname:   "image5.png",
		},
		{
			x_lower: -1.1590295228484262,
			x_upper: -1.1590294146037756,
			y_lower: 0.2725089787759562,
			y_upper: 0.272509059959444,
			fname:   "image6.png",
		},
		{
			x_lower: 0.2634939406642726,
			x_upper: 0.2645498803213035,
			y_lower: 0.0025203902052182,
			y_upper: 0.0033123449479914,
			fname:   "image7.png",
		},
	}
)

type complexSample struct {
	// min real value
	x_lower float64
	// max real value
	x_upper float64

	// min imaginary value
	y_lower float64
	// max imaginary value
	y_upper float64

	// image name to save on disk
	fname string
}

func main() {
	pn := "default"
	p := palette.Gradients[pn]

	dirName := fmt.Sprintf("img_%v", pn)
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	for _, sample := range samples {
		img := render(1920, 1080, sample, p)

		f, err := os.Create(filepath.Join(dirName, sample.fname))
		if err != nil {
			log.Fatal(err)
		}

		if err := png.Encode(f, img); err != nil {
			f.Close()
			log.Fatal(err)
		}

		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func render(width, height int, cs *complexSample, p *palette.Gradient) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b := calcRgb(x, y, width, height, cs)
			r, g, b = palette.Recolor(int(r), int(g), int(b), p)

			img.Set(x, y, color.NRGBA{
				R: r,
				G: g,
				B: b,
				A: 255,
			})
		}
	}

	return img
}

func calcRgb(x, y, maxX, maxY int, cs *complexSample) (r, g, b uint8) {
	cx := convertRange(float64(x), 0, float64(maxX), cs.x_lower, cs.x_upper)
	cy := convertRange(float64(y), 0, float64(maxY), cs.y_lower, cs.y_upper)
	iters := iter(cx, cy, max_iter)

	color := uint8(convertRange(float64(iters), 0, max_iter, 0, 255))
	r, g, b = color, color, color

	return
}

func iter(r, im float64, maxIter int) int {
	z := complex(0, 0)
	c := complex(r, im)

	i := 0
	for i < maxIter {
		z = z*z + c
		if cmplx.Abs(z) > 4 {
			break
		}

		i++
	}

	return i
}

func convertRange(v, lbound, ubound, nlbound, nubound float64) float64 {
	return float64((((v-lbound)/(ubound-lbound))*(nubound-nlbound) + nlbound))
}
