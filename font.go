package persiancaptcha

import (
	"embed"
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed fonts/IranNastaliq.ttf
var fontData embed.FS

var (
	fontFace font.Face
	loadOnce sync.Once
	fontErr  error
)

func loadFontOnce() {
	loadOnce.Do(func() {
		data, err := fontData.ReadFile("fonts/IranNastaliq.ttf")
		if err != nil {
			fontErr = err
			return
		}
		ft, err := opentype.Parse(data)
		if err != nil {
			fontErr = err
			return
		}
		fontFace, fontErr = opentype.NewFace(ft, &opentype.FaceOptions{
			Size:    fontSize,
			DPI:     72,
			Hinting: font.HintingFull,
		})
	})
}

func getFontFace(size float64) (font.Face, error) {
	loadFontOnce()
	return fontFace, fontErr
}
