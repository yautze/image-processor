package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
)

const (
	imgURL      = "testjpg.jpg"
	outFileName = "out.jpg"
)

func main() {
	f, err := os.Open(imgURL)
	if err != nil {
		log.Panicln("讀取圖檔失敗", err)
		return
	}
	img, _, err := image.Decode(f)
	if err != nil {
		log.Panicln("Decode圖檔失敗", err)
		return
	}

	buf := bytes.Buffer{}
	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 90})
	if err != nil {
		log.Panicln("encode JPEG image failed", err)
		return
	}

	err = ioutil.WriteFile(outFileName, buf.Bytes(), 0664)
	if err != nil {
		log.Panicln("write out image failed", err)
	}

}
