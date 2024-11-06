package main

import "image"

func sample(width, height, i, j int) Color {
	rgb := Color{}

	rgb.R = float64(i) / float64(width-1)
	rgb.G = float64(j) / float64(height-1)
	rgb.B = 0.0

	return rgb
}

func render(width, height int, ch chan<- int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for row := range height {
		for col := range width {
			rgb := sample(width, height, col, row)
			img.Set(col, height-row-1, rgb)
		}
		ch <- 1
	}

	return img
}
