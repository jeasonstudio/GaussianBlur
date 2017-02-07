package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math"
	"os"
)

// OMIGA Ω
const OMIGA = 1.5

func main() {
	num := 1

	// fmt.Println(GetAvgArr(1))

	// fmt.Println(GaussFunc(0, 0, OMIGA))

	PrintImg("ava.jpg", "bvb.jpg", GetAvgArr(num), num)

}

// GaussFunc 二维高斯函数
func GaussFunc(x, y int, OmiGa float64) float64 {
	return (1.0 / (2.0 * math.Pi * OmiGa * OmiGa)) * math.Pow(math.E, ((-1.0)*(float64(x*x+y*y)/(2.0*OmiGa*OmiGa))))
}

// GetAvgArr 计算权重矩阵
func GetAvgArr(len int) [][]float64 {
	sum := 0.0
	arr := make([][]float64, (2*len + 1), (2*len + 1))
	for i := 0; i < (2*len + 1); i++ {
		arr2 := make([]float64, (2*len + 1), (2*len + 1))
		for j := 0; j < (2*len + 1); j++ {
			thisGaussResult := GaussFunc(i-len, j-len, OMIGA)
			arr2[j] = thisGaussResult
			sum += thisGaussResult
		}
		arr[i] = arr2
	}

	for i := 0; i < (2*len + 1); i++ {
		thisArr := arr[i]
		for j := 0; j < (2*len + 1); j++ {
			thisArr[j] = thisArr[j] / sum
		}
		arr[i] = thisArr
	}

	return arr
}

// PrintImg 打印图片
func PrintImg(sourceImg, tagImg string, arr [][]float64, num int) {
	file, err := os.Open(sourceImg)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Create(tagImg)

	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()

	img, _ := jpeg.Decode(file)

	jpg := image.NewRGBA64(img.Bounds())

	xWidth := img.Bounds().Dx()
	yHeight := img.Bounds().Dy()

	for i := 0; i < xWidth; i++ {
		for j := 0; j < yHeight; j++ {
			for k := 0; k < ((2*num + 1) * (2*num + 1)); k++ {
			}
			thisR, thisG, thisB, thisA := img.At(i, j).RGBA()
			var newColor color.RGBA64
			newColor.R = uint16(thisR)
			newColor.G = uint16(thisG)
			newColor.B = uint16(thisB)
			newColor.A = uint16(thisA)

			jpg.SetRGBA64(i, j, newColor)
		}
	}
	// r, g, b, a := img.At(0, 0).RGBA()
	// fmt.Println(r, g, b, a)
	// jpg := image.NewGray(img.Bounds())
	draw.Draw(jpg, img.Bounds().Add(image.Pt(xWidth, yHeight)), img, img.Bounds().Min, draw.Src)
	jpeg.Encode(file1, jpg, nil)
}
