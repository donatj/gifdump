package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"log"
	"os"
)

var (
	uncomposed = flag.Bool("u", false, "Dump the raw gif frames without composing the interframe compression.")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [-u] <gif>:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	file, err := os.Open(flag.Arg(0))
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	g, err := gif.DecodeAll(file)
	if err != nil {
		log.Fatal(err)
	}

	b := g.Image[0].Bounds()
	visible := image.NewRGBA(b)

	for i, img := range g.Image {
		sb := img.Bounds()
		f, err := os.Create(fmt.Sprintf("frame_%d.png", i))
		if err != nil {
			log.Fatal(err)
		}

		if *uncomposed {
			visible = image.NewRGBA(b)
		}

		// if false {
		for y := sb.Min.Y; y < sb.Max.Y; y++ {
			for x := sb.Min.X; x < sb.Max.X; x++ {
				c := img.At(x, y)
				_, _, _, a := c.RGBA()
				if a != 0 {
					visible.Set(x, y, c)
				}
			}
		}

		png.Encode(f, visible)
	}
}
