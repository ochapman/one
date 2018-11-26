/*
* File Name:	cv.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2018-10-29
 */

// What it does:
//
// This example uses the Window class to open an image file, and then display
// the image in a Window class.
//
// How to run:
//
// go run ./cmd/showimage/main.go /home/ron/Pictures/mcp23017.jpg
//
//

package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"sort"

	"gocv.io/x/gocv"
)

func SplitLargestRect(rect image.Rectangle) []image.Rectangle {
	dx := rect.Max.X - rect.Min.X
	dy := rect.Max.Y - rect.Min.Y
	var n, l int
	var isXBig bool
	if dx > dy {
		n = dx /dy
		l = dy
		isXBig = true
	} else {
		n = dy / dx
		l = dx
		isXBig = false
	}
	rects := make([]image.Rectangle, 0)
	for i := 0; i < n; i++ {
		if isXBig {
			rects = append(rects, image.Rectangle{Min:image.Point{X:rect.Min.X+n*l, Y:rect.Min.Y}, Max:image.Point{X:rect.Max.X+n*l, Y:rect.Max.Y}})
		} else {
			rects = append(rects, image.Rectangle{Min:image.Point{X:rect.Min.X, Y:rect.Min.Y+n*l}, Max:image.Point{X:rect.Max.X, Y: rect.Max.Y+n*l}})
		}
	}
	return rects
}

func RectToGBlock(rects []image.Rectangle) {
	vx, vy := make([]int, 0), make([]int, 0)
	for _, r := range rects {
		vx = append(vx, r.Min.X)
		vy = append(vy, r.Min.Y)
	}
	sort.Ints(vx)
	sort.Ints(vy)
	fmt.Printf("vx: %v\n", vx)
	fmt.Printf("vy: %v\n", vy)

}



func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:ntshowimage [imgfile]")
		return
	}

	filename := os.Args[1]
	fmt.Printf("filename: %s\n", filename)
	window := gocv.NewWindow("Hello")
	img := gocv.IMRead(filename, gocv.IMReadColor)
	grayImage := gocv.NewMat()
	defer grayImage.Close()

	gocv.CvtColor(img, &grayImage, gocv.ColorBGRToGray)
	destImage := gocv.NewMat()
	gocv.Threshold(grayImage, &destImage, 230, 255, gocv.ThresholdBinaryInv)
	resultImage := gocv.NewMatWithSize(800, 1000, gocv.MatTypeCV8U)

	gocv.Resize(destImage, &resultImage, image.Pt(resultImage.Rows(), resultImage.Cols()), 0, 0, gocv.InterpolationCubic)
	gocv.Dilate(resultImage, &resultImage, gocv.NewMat())
	gocv.GaussianBlur(resultImage, &resultImage, image.Pt(5, 5), 0, 0, gocv.BorderWrap)
	results := gocv.FindContours(resultImage, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	imageForShowing := gocv.NewMatWithSize(resultImage.Rows(), resultImage.Cols(), gocv.MatChannels4)
	//fmt.Printf("result: %#v", results)
	rects := make([]image.Rectangle, 0)
	//sizeRects := make(map[int]image.Rectangle)
	var largestRect image.Rectangle
	elementSize := 0
	for index, element := range results {
		//gocv.DrawContours(&imageForShowing, results, index, color.RGBA{R: 0, G: 0, B: 255, A: 255}, 1)
		rec := gocv.BoundingRect(element)
		fmt.Printf("index: %d, element, size %d, rec: %#v\n", index, len(element), rec)
		if len(element) > elementSize {
			elementSize = len(element)
			largestRect = rec
		}
		gocv.Rectangle(&imageForShowing,
			rec,
			color.RGBA{R: 0, G: 255, B: 0, A: 100}, 1)
		rects = append(rects, rec)
	}
	RectToGBlock(rects)
	_ = largestRect
	if img.Empty() {
		fmt.Printf("Error reading image from: %v", filename)
		return
	}
	//fmt.Printf("imageForShowing : %#v\n", imageForShowing)

	for {
		window.IMShow(imageForShowing)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
