# ascii_videos

A project that converts pictures or video to ascii characters and output in terminal.

## Environment

Before start, make sure you have **ffmpeg** installed.

```bash
$ ffmpeg
ffmpeg version 4.2.7-0ubuntu0.1 Copyright (c) 2000-2022 the FFmpeg developers
  built with gcc 9 (Ubuntu 9.4.0-1ubuntu1~20.04.1)
...
```

If not, use apt install on Ubuntu:

```bash
$ sudo apt install ffmpeg
```

Or compile by source code:

```bash
$ wget http://www.ffmpeg.org/releases/ffmpeg-5.1.tar.gz
$ tar -zxvf ffmpeg-5.1.tar.gz
$ cd ffmpeg-5.1
$ ./configure --prefix=/usr/local/ffmpeg
$ make
$ make install
```

## Example

Load a image from remote or local path, convert rgb to grayscale, and calculate average pixel to ascii characters, then output or save as a file.

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

![image-20220815210749832](https://s3.bmp.ovh/imgs/2022/08/16/d188bb8672d2d43a.png)