package shot

import (
	"bufio"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"net/http"
)

var (
	size      float64      = 26
	smallSize float64      = 16
	dpi       float64      = 144
	spacing   float64      = 1.5
	hinting   font.Hinting = font.HintingNone
	sas                    = "screenshot will be available soon"
)

func (s *Service) Draw(siteName string, w http.ResponseWriter) {
	// Read the font data.
	fontBytes := MustAsset("assets/luxisr.ttf")
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	// Draw the background and the guidelines.
	fg, bg := image.Black, image.White
	const imgW, imgH = 1200, 700
	rgba := image.NewRGBA(image.Rect(0, 0, imgW, imgH))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	d := &font.Drawer{
		Dst: rgba,
		Src: fg,
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: hinting,
		}),
	}
	y := (imgH / 2) + (int(math.Ceil(size*dpi/72)) / 2)
	//dy := int(math.Ceil(size * spacing * dpi / 72))
	d.Dot = fixed.Point26_6{
		X: (fixed.I(imgW) - d.MeasureString(siteName)) / 2,
		Y: fixed.I(y),
	}
	d.DrawString(siteName)

	d = &font.Drawer{
		Dst: rgba,
		Src: image.NewUniform(color.Gray16{0xaaaa}),
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    smallSize,
			DPI:     dpi,
			Hinting: hinting,
		}),
	}

	dy := int(math.Ceil(smallSize * spacing * dpi / 72))
	d.Dot = fixed.Point26_6{
		X: (fixed.I(imgW) - d.MeasureString(sas)) / 2,
		Y: fixed.I(y + dy),
	}
	d.DrawString(sas)

	b := bufio.NewWriter(w)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		return
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		return
	}
}
