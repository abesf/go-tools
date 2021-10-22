package image

import (
	"image"
	"net/http"
	"github.com/disintegration/imageorient"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)
//image width= img.Bounds().Dx()
//image height= img.Bounds().Dy()
func GetImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	img, _, err := imageorient.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}