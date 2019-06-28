package gif

import (
	"github.com/samcsf/daily-go/pkg/util"
	"image"
	"image/gif"
	"log"
	"os"
)

func ReverseFrames(filePath string, outPath string) {
	file, err := os.Open(filePath)
	util.ChkErr(err)
	file2, err := os.OpenFile(outPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	util.ChkErr(err)

	pic, err := gif.DecodeAll(file)
	util.ChkErr(err)

	log.Printf("Totally %d frames", len(pic.Image))
	var tmp []*image.Paletted
	for i := len(pic.Image) - 1; i >= 0; i-- {
		tmp = append(tmp, pic.Image[i])
	}
	pic.Image = tmp

	err = gif.EncodeAll(file2, pic)
	util.ChkErr(err)
}
