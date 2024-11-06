package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"
	"time"
)

var Version = "v0.0.0"

const (
	defaultWidth     = 800
	defaultHeight    = 600
	progressBarWidth = 80
)

var (
	width    int
	height   int
	filename string
)

func outputProgress(ch <-chan int, rows int) {
	fmt.Println()
	for i := 1; i <= rows; i++ {
		<-ch
		pct := 100 * float64(i) / float64(rows)
		filled := (progressBarWidth * i) / rows
		bar := strings.Repeat("=", filled) + strings.Repeat("-", progressBarWidth-filled)
		fmt.Printf("\r[%s] %.2f%%", bar, pct)
	}
	fmt.Println()
}

func writeFile(filename string, img image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)

	return err
}

func main() {
	flag.IntVar(&width, "w", defaultWidth, "width of image in pixels")
	flag.IntVar(&height, "h", defaultHeight, "height of image in pixels")
	flag.StringVar(&filename, "o", "out", "output filename")

	displayVersion := flag.Bool("version", false, "Display version and exit")

	flag.Parse()

	if *displayVersion {
		fmt.Printf("raygo version %s\n", Version)
		os.Exit(0)
	}

	start := time.Now()

	fmt.Printf("\nRendering %d x %d pixel", width, height)

	ch := make(chan int, height)
	defer close(ch)
	go outputProgress(ch, height)

	filename = fmt.Sprintf("%s.png", filename)

	img := render(width, height, ch)
	err := writeFile(filename, img)
	if err != nil {
		log.Fatalf("Could not write file %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nDone. Elapsed: %v", time.Since(start))
	fmt.Printf("\nOutput to: %s.png\n", filename)
}
