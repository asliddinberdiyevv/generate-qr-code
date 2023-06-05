package main

import (
	"image/color"
	"log"

	"github.com/skip2/go-qrcode"
)

func main() {
	if err := qrcode.WriteColorFile("https://example.org", qrcode.Medium, 256, color.Black, color.White, "qr.png"); err != nil {
		log.Println(err)
	} else {
		log.Println("Create qr code")
	}

}
