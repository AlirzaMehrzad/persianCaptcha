package persiancaptcha

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

const (
	width    = 250
	height   = 150
	fontSize = 55
)

func init() {
	rand.Seed(time.Now().UnixNano())
	loadFontOnce() // from font.go
}

func Generate() (*gg.Context, string, error) {
	num := rand.Intn(9000) + 1000
	persianText := toPersianDigits(num)

	dc := gg.NewContext(width, height)
	dc.SetColor(color.Opaque)
	dc.Clear()

	// Draw noisy lines
	for i := 0; i < 5; i++ {
		dc.SetRGBA(rand.Float64(), rand.Float64(), rand.Float64(), 0.5)
		x1 := rand.Float64() * width
		y1 := rand.Float64() * height
		x2 := rand.Float64() * width
		y2 := rand.Float64() * height
		dc.SetLineWidth(rand.Float64()*2 + 0.5)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}

	face, err := getFontFace(fontSize)
	if err != nil {
		return nil, "", err
	}
	dc.SetFontFace(face)

	dc.SetRGB(0, 0, 0)
	textWidth, textHeight := dc.MeasureString(persianText)
	dc.DrawStringAnchored(persianText, width-textWidth/2-20, height/2+textHeight/4, 1.5, 0)

	return dc, persianText, nil
}

func toPersianDigits(n int) string {
	persianDigits := []rune{'۰', '۱', '۲', '۳', '۴', '۵', '۶', '۷', '۸', '۹'}
	result := []rune{}
	for _, d := range []byte(fmt.Sprintf("%d", n)) {
		result = append(result, persianDigits[d-'0'])
	}
	return string(result)
}
