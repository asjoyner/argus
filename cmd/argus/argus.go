package main

import (
	"log"
	"time"

	"github.com/asjoyner/argus/capture"
)

var imageHistory = 5

func main() {
	images := make([][]byte, 0, imageHistory)

	for {
		img, err := capture.GetImage(nil)
		if err != nil {
			log.Fatalf("could not read image from camera: %s", err)
		}
		if len(images) > imageHistory {
			images = append(images[1:], img)
		} else {
			images = append(images, img)
		}
		time.Sleep(4 * time.Second)
		log.Println(len(images))
	}

}
