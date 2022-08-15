# ascii_videos

A project that converts video to ascii characters and print them in the terminal

## Example

load a image from remote or local path, convert rgb to gray matrix,  and calculate ascii characters to instead.

```go
// load image
img, err := LoadImage("https://i0.hdslb.com/bfs/new_dyn/cd115a1ab8a69f121ac1ab740f45b12a173469252.png")

// image to ascii
bytes := image2Ascii(img)

// clear
cmd := exec.Command("clear")
cmd.Stdout = os.Stdout
cmd.Run()

// print
fmt.Println(string(bytes))
```

## Result

![img](https://i0.hdslb.com/bfs/new_dyn/cd115a1ab8a69f121ac1ab740f45b12a173469252.png@1709w.webp)

![image-20220815211000843](https://s3.bmp.ovh/imgs/2022/08/15/87351e1bf19fb84e.png)

![img](https://i0.hdslb.com/bfs/new_dyn/615c8071c1c4beba47e6c7971b8561e4470962000.jpg@1709w.webp)

![image-20220815210749832](https://s3.bmp.ovh/imgs/2022/08/15/0ea58718057c4a2a.png)