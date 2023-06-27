package filter

import (
	"github.com/disintegration/imaging"
)

type FiltreGris struct{}

func (f *FiltreGris) Process(srcPath, dstPath string) error {
	srcImg, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	grayImg := imaging.Grayscale(srcImg)

	err = imaging.Save(grayImg, dstPath)
	if err != nil {
		return err
	}

	return nil
}

type FiltreFlou struct{}

func (f *FiltreFlou) Process(srcPath, dstPath string) error {
	srcImg, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	blurImg := imaging.Blur(srcImg, 10.0)

	err = imaging.Save(blurImg, dstPath)
	if err != nil {
		return err
	}

	return nil
}
