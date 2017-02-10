package main

import (
	"fmt"
)



func main() {
    fmt.Println("start GaussianBlur")
	GaussianBlur.GBlurInit("source.jpg", "tag.jpg", 5, 100.0)
}
