package source

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"io"
	"log"
)

// ImgCoderMap -
var ImgCoderMap = map[string]Coder{
	"jpg":  new(jpg),
	"jpeg": new(jpg),
}

// Coder -
type Coder interface {
	// Decode -
	Decode(r io.Reader) (image.Image, error)

	// Encode -
	Encode(imgSource image.Image, quality int) ([]byte, error)
}

// NewCoder -
func NewCoder(currentImgType string) (Coder, error) {
	coder, ok := ImgCoderMap[currentImgType]
	if !ok {
		return nil, errors.New("new image coder failed")
	}

	return coder, nil
}

// ----- JPG / JPEG -----
type jpg struct{}

// Decode - jpg / jpeg decode
func (j *jpg) Decode(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		log.Println(err)
		return nil, errors.New("decode jpg/jpeg image failed")
	}

	return img, nil
}

// Encode - jpg / jpeg encode
func (j *jpg) Encode(imgSource image.Image, quality int) ([]byte, error) {
	buf := bytes.Buffer{}
	if err := jpeg.Encode(&buf, imgSource, &jpeg.Options{Quality: quality}); err != nil {
		log.Println(err)
		return nil, errors.New("encode jpg/jpeg image to buffer failed")
	}

	return buf.Bytes(), nil
}
