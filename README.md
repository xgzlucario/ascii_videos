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



## Pictures

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



## Videos

Use project [The fastest Bilibili video downloader](https://github.com/sodaling/FastestBilibiliDownloader) to download bilibili videos. clone project and use **go build** to make a binary file in `./cmd` path. download videos in `./download` path.

```bash
$ ./cmd/download_bili_videos
欢迎使用B站视频下载器 v1.0.1
支持以下几种格式的输入：
·  https://www.bilibili.com/video/旧版的av号/ | av号 是以`av`开头的一串数字
·  https://www.bilibili.com/video/新版的BV号/ | BV号 是以`BV`开头的一串字符
·  https://space.bilibili.com/UP主的ID/       | UP主的ID 是一串数字

请输入想要下载的视频网址/up主个人主页网址: https://www.bilibili.com/video/BV1pS4y1x7RR
2022/08/16 22:42:27 开始下载...
即将开始下载： 【4K60FPS】Fool's Garden《Lemon Tree》经典现场！英语老师推荐曲目！
```

