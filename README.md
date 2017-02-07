# GaussianBlur
GaussianBlur for golang go 语言图像处理库——高斯模糊

```bash
(i-1,j-1) (i,j-1) (i+1,j-1)

(i-1,j) (i,j) (i+1,j)

(i-1,j+1) (i,j+1) (i+1,j+1)
```

```javascript
i-num => i+num
j-num => j+num
```

info | source | show |
---|---|---
Ω = 5; n = 5 | ![](source.jpg) | ![](o5n5.jpg)
Ω = 10; n = 10 | ![](source.jpg) | ![](o10n10.jpg)
Ω = 50; n = 10 | ![](source.jpg) | ![](o50n10.jpg)
