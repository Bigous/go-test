package main

import (
	"fmt";
	"bytes";
	"encoding/base64";
	"image";
	"image/png";
	_ "math";
)

func pic(dx, dy int) *image.NRGBA {
	var (
		img *image.NRGBA;
		x, y, p int;
	)
	img = image.NewNRGBA(image.Rect(0,0, dx, dy));
	for y = 0; y < dy; y++ {
		for x = 0; x < dx; x++ {
			p = img.Stride * y + x * 4; // *4 pq é RGBA
			img.Pix[p] = uint8(x * 255 / (y + 1));
			img.Pix[p + 1] = uint8(x * y);
			img.Pix[p + 2] = uint8(y * 255 / (x + 1));
			img.Pix[p + 3] = uint8(255);
		}
	}
	return img;
}

func main() {
	var (
		img *image.NRGBA;
		buf bytes.Buffer;
		err error;
		enc string;
	)
	img = pic(256, 256);
	err = png.Encode(&buf, img);
	if (err != nil) {
		panic(err);
	}
	enc = base64.StdEncoding.EncodeToString(buf.Bytes());
	fmt.Println("data:image/png;base64," + enc);
}