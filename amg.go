package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"os/exec"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var (
	// ascii grayscale rank, If you find a better way to combine grayscale values than this, submit plz!
	ramp = " .,-=+CNMGB@"

	// specify Reverse to change contrast
	Reverse = false

	// specify ScaleX and ScaleY to resize the pictures
	ScaleX, ScaleY = 3, 1
)

func init() {
	if Reverse {
		ramp = reverseString(ramp)
	}
}

func reverseString(str string) string {
	s := []rune(str)
	l, r := 0, len(s)-1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return string(s)
}

// load image
func LoadImage(path string) (image.Image, error) {
	if path[0:4] == "http" {
		// http get
		res, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		image, _, err := image.Decode(res.Body)
		return image, err

	} else {
		// local path
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		image, _, err := image.Decode(f)
		return image, err
	}
}

// rgb to gray
func rgb2gray(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
}

// image to ascii character
func image2Ascii(img image.Image) []byte {
	buf := bytes.Buffer{}
	max := img.Bounds().Max

	for y := 0; y < max.Y; y += ScaleX {
		for x := 0; x < max.X; x += ScaleY {
			c := avgPixel(img, x, y, ScaleX, ScaleY)
			buf.WriteByte(ramp[len(ramp)*c/65536])
		}
		buf.WriteByte('\n')
	}
	return buf.Bytes()
}

func SaveFrameAsAscii(path string, index int) error {
	// get frame from video
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(path).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", index)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "png"}).
		WithOutput(buf).Run()
	if err != nil {
		return err
	}

	// decode
	img, _, err := image.Decode(buf)
	if err != nil {
		return err
	}

	// img to ascii
	file, err := os.OpenFile(fmt.Sprintf("frames/frame:%d", index), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes := image2Ascii(img)
	file.Write(bytes)
	file.Close()

	fmt.Printf("save video frame:%d as file success\n", index)
	return nil
}

func avgPixel(img image.Image, x, y, w, h int) int {
	cnt, sum, max := 0, 0, img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			sum += rgb2gray(img.At(i, j))
			cnt++
		}
	}
	return sum / cnt
}

// clear terminal
func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	// img, err := LoadImage("https://i0.hdslb.com/bfs/new_dyn/615c8071c1c4beba47e6c7971b8561e4470962000.jpg")
	// if err != nil {
	// 	panic(err)
	// }

	// bytes := image2Ascii(img)

	// fmt.Println(string(bytes))

	// file := "download/test1.flv"
	// for i := 0; ; i++ {
	// 	if err := SaveFrameAsAscii(file, i); err != nil {
	// 		break
	// 	}
	// }
	// fmt.Println("ascii videos complete")
	Play()
}

// play ascii video on console
func Play() {
	for i := 0; ; i++ {
		file, err := os.ReadFile(fmt.Sprintf("frames/frame:%d", i))
		if err != nil {
			break
		}
		// clear()
		fmt.Println(string(file))
	}
}
