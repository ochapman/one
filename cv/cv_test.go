package main

import (
	"image"
	"testing"
)

func TestSplitLargestRect(t *testing.T) {
	//index: 10, element, size 191, rec: image.Rectangle{Min:image.Point{X:646, Y:613}, Max:image.Point{X:769, Y:968}}
	rect := image.Rectangle{Min: image.Point{X: 646, Y: 613}, Max: image.Point{X: 769, Y: 968}}
	rects := SplitLargestRect(rect)
	t.Logf("origin: %v\nrects: %v", rect, rects)
}
