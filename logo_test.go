package LogoGen_test

import (
	"image/png"
	"os"
	"testing"

	"github.com/Vexera/LogoGen"
)

func TestMeme(t *testing.T) {

	img, err := LogoGen.CreateLogo("V", "vexera")
	if err != nil {
		t.Error(err)
	}

	f, err := os.Create("img.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)

}
