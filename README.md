# GaussianBlur
GaussianBlur for golang go 语言图像处理库——高斯模糊

### Result

info | source | result |
---|---|---
Ω = 5; n = 5 | ![](source.jpg) | ![](o5n5.jpg)
Ω = 10; n = 10 | ![](source.jpg) | ![](o10n10.jpg)
Ω = 50; n = 10 | ![](source.jpg) | ![](o50n10.jpg)

### Usage

```go
go get github.com/jeasonstudio/GaussianBlur
```

```go
package main
import "github.com/jeasonstudio/GaussianBlur"
func main()  {
    GaussianBlur.GaussianBlur("source.jpg","tag.jpg",5,5.0)
}
```

```go
// GaussianBlur 高斯模糊处理
// sourceImg \ tagImg 处理前 \ 后图片相对路径地址
// num 高斯模糊像素，单位 px，注意，此项过高将直接影响时间
// OMIGA 欧米伽，周围像素权重
func GaussianBlur(sourceImg, tagImg string, num int, OMIGA float64)
```

### Info

> num OMIGA 都与模糊程度成正比，但 num 尽量为 5px 左右，不要超过 10。OMIGA 可以超过 50。

> num 过高影响处理时间，OMIGA 过高影响图片质量。

### Todo

 - 算法时间、空间复杂度有很大优化空间。
 - 添加对 *.png 格式文件的支持。
 - 处理后图片边缘有很大失真，尤其右边和下边。