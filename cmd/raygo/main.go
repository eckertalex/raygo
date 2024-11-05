package main

import (
	"fmt"
	"os"
)

const (
	imageWidth  uint64 = 400
	imageHeight uint64 = 300
)

func main() {
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := range imageHeight {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", imageHeight-j)
		for i := range imageWidth {
			r := float64(i) / float64(imageWidth-1)
			g := float64(j) / float64(imageHeight-1)
			b := 0.0

			ir := uint64(255.999 * r)
			ig := uint64(255.999 * g)
			ib := uint64(255.999 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}

	fmt.Fprint(os.Stderr, "\rDone.                 \n")
}
