package main

import (
	"flag"
	"image"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"github.com/rasmi/personify/facefinder"
	"github.com/disintegration/imaging"
)

var haarCascade = flag.String("haar", os.Getenv("GOPATH")+"/src/github.com/rasmi/personify/haarcascade_frontalface_alt.xml", "The location of the Haar Cascade XML configuration to be provided to OpenCV.")
var facesDir = flag.String("faces", "faces/", "The directory to search for faces.")

func main() {
	flag.Parse()

	var personFaces FaceList

	var facesPath string
	var err error

	if *facesDir != "" {
		facesPath, err = filepath.Abs(*facesDir)
		if err != nil {
			panic(err)
		}
	}

	err = personFaces.Load(facesPath)
	if err != nil {
		panic(err)
	}
	if len(personFaces) == 0 {
		panic("no faces found")
	}

	file := flag.Arg(0)

	finder := facefinder.NewFinder(*haarCascade)

	baseImage := loadImage(file)

	faces := finder.Detect(baseImage)

	bounds := baseImage.Bounds()

	canvas := canvasFromImage(baseImage)

	for _, face := range faces {
		rect := rectMargin(30.0, face)

		newFace := personFaces.Random()
		if newFace == nil {
			panic("nil face")
		}
		personFace := imaging.Fit(newFace, rect.Dx(), rect.Dy(), imaging.Lanczos)

		draw.Draw(
			canvas,
			rect,
			personFace,
			bounds.Min,
			draw.Over,
		)
	}

	if len(faces) == 0 {
		face := imaging.Resize(
			personFaces[0],
			bounds.Dx()/3,
			0,
			imaging.Lanczos,
		)
		face_bounds := face.Bounds()
		draw.Draw(
			canvas,
			bounds,
			face,
			bounds.Min.Add(image.Pt(-bounds.Max.X/2+face_bounds.Max.X/2, -bounds.Max.Y+int(float64(face_bounds.Max.Y)/1.9))),
			draw.Over,
		)
	}

	jpeg.Encode(os.Stdout, canvas, &jpeg.Options{jpeg.DefaultQuality})
}
