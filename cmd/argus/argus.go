package main

import (
	"log"
	"os"

	"github.com/asjoyner/argus/capture"
)

func main() {
	img, err := capture.GetImage(nil)
	if err != nil {
		log.Fatalf("could not read image from camera: %s", err)
	}

	fh, err := os.Create("/tmp/output.jpg")
	if err != nil {
		log.Fatalf("could not open /tmp/output.jpg: %s", err)
	}
	defer fh.Close()

	if _, err := fh.Write(img); err != nil {
		log.Fatalf("could not write image to /tmp/output.jpg: %s", err)
	}

}
