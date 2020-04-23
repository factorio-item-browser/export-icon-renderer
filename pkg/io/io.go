package io

import (
	"bytes"
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/env"
	"image"
	"regexp"
)

var (
	encoder       = imgio.PNGEncoder()
	regexFileName = regexp.MustCompile("^__(.*)__/(.*)$")
)

// Create will create a new transparent image with the specified size.
func Create(size int) image.Image {
	return image.NewRGBA(image.Rect(0, 0, size, size))
}

// Load loads the image with the specified fileName.
func Load(fileName string) (image.Image, error) {
	match := regexFileName.FindStringSubmatch(fileName)
	if match == nil || len(match) < 3 {
		return nil, fmt.Errorf("unable to understand filename: %s", fileName)
	}

	modName := match[1]
	imagePath := match[2]

	directory := env.FactorioModsDirectory
	if modName == "base" || modName == "core" {
		directory = env.FactorioDataDirectory
	}
	path := fmt.Sprintf("%s/%s/%s", directory, modName, imagePath)

	img, err := imgio.Open(path)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// Encode will encode the specified image to a string.
func Encode(img image.Image) (string, error) {
	buf := new(bytes.Buffer)
	err := encoder(buf, img)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
