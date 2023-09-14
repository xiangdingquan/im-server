package main

import (
	"fmt"
	"open.chat/app/service/dfs/internal/imaging"
	"open.chat/pkg/bytes2"
	"os"
)

func main() {
	img, err := imaging.Open(".//tl_card_connect.gif")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(img.Bounds())
	dst := imaging.Resize(img, 40, 40)
	fmt.Println(dst.Bounds())

	w := bytes2.NewBuffer(make([]byte, 0, 4096))
	err = imaging.Encode(w, dst, 30)
	if err != nil {
		fmt.Println(err)
		return
	}
	bb := w.Bytes()
	fmt.Println(len(bb))

	file, err := os.Create("./tl_card_connect.jpeg")
	if err != nil {
	}
	defer file.Close()
	h := imaging.JpegHeader
	h[164] = bb[1]
	h[166] = bb[2]
	file.Write(imaging.JpegHeader)
	file.Write(bb[3:])
	file.Write(imaging.JpegFooter)
}
