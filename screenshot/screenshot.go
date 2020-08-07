package screenshot

import (
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func savesreenshot(img *image.RGBA, filename string) {
	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	png.Encode(file, img)
}

func Screenshot(filename string) (bool, error) {
	a := image.Rect(0, 0, 0, 0)
	bounds := screenshot.GetDisplayBounds(0)
	a = bounds.Union(a)

	tk, err := screenshot.CaptureRect(bounds)

	if err != nil {
		return false, err
	}

	savesreenshot(tk, filename)
	return true, nil
}
