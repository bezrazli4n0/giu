package giu

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/AllenDang/giu/imgui"
)

func LoadImage(imgPath string) (*image.RGBA, error) {
	imgFile, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	img, err := png.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	switch trueImg := img.(type) {
	case *image.RGBA:
		return trueImg, nil
	default:
		rgba := image.NewRGBA(trueImg.Bounds())
		draw.Draw(rgba, trueImg.Bounds(), trueImg, image.Pt(0, 0), draw.Src)
		return rgba, nil
	}
}

func ToVec4Color(col color.RGBA) imgui.Vec4 {
	return imgui.Vec4{
		X: float32(col.R) / 255,
		Y: float32(col.G) / 255,
		Z: float32(col.B) / 255,
		W: float32(col.A) / 255,
	}
}
