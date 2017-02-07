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

// OMIGA Ω 的值，越大，模糊程度越高
const OMIGA = 50

func main() {
	num := 10

	PrintImg("source.jpg", "o50n10.jpg", GetAvgArr(num), num)

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
			var newColor color.RGBA64
			var sumR, sumG, sumB, sumA uint16

			for p := ((-1) * num); p <= num; p++ {
				for q := ((-1) * num); q <= num; q++ {
					trueX := i + p
					trueY := j + q

					// 若超出边界则使用边界值
					if trueX < 0 {
						trueX = 0
					} else if trueX > xWidth {
						trueX = xWidth
					}
					if trueY < 0 {
						trueY = 0
					} else if trueY > yHeight {
						trueY = yHeight
					}
					thisR, thisG, thisB, thisA := img.At(trueX, trueY).RGBA()
					sumR += uint16(arr[p+num][q+num] * float64(thisR))
					sumG += uint16(arr[p+num][q+num] * float64(thisG))
					sumB += uint16(arr[p+num][q+num] * float64(thisB))
					sumA += uint16(arr[p+num][q+num] * float64(thisA))
					// fmt.Println(sumA)
				}
			}
			newColor.R = sumR
			newColor.G = sumG
			newColor.B = sumB
			newColor.A = sumA
			jpg.SetRGBA64(i, j, newColor)

		}
	}
	draw.Draw(jpg, img.Bounds().Add(image.Pt(xWidth, yHeight)), img, img.Bounds().Min, draw.Src)
	jpeg.Encode(file1, jpg, nil)
}
