# go-libav

[![Go Reference](https://pkg.go.dev/badge/github.com/Lensual/go-libav.svg)](https://pkg.go.dev/github.com/Lensual/go-libav)
![GitHub release (with filter)](https://img.shields.io/github/v/release/Lensual/go-libav?include_prereleases&color=blue)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/Lensual/go-libav)
![Static Badge](https://img.shields.io/github/license/Lensual/go-libav)
[![Go Report Card](https://goreportcard.com/badge/github.com/Lensual/go-libav)](https://goreportcard.com/report/github.com/Lensual/go-libav)
![Static Badge](https://img.shields.io/badge/FFmpeg_version-6.1-purple)

[English](./README.md) | [中文](./README_zh.md)

[FFmpeg和libav](https://ffmpeg.org/)的Go绑定。

## 须知

1. 这个包是使用CGO实现的，需要FFmpeg的头文件和库进行编译。
2. 由于CGO调用存在众所周知的性能问题，建议使用C/CPP实现你的需求。

## 开始使用

`av***`目录下是单纯的CGO绑定。使用这个包之前，你最好掌握FFmpeg库的基本用法。

`advance`目录是为方便使用进行二次封装的包

go版本至少`1.21`，FFmpeg至少`6.1`

如果你的库不在默认目录下，需要设置好环境变量。

```env
# 使用 pkg-config
PKG_CONFIG_PATH="/ffmpeg/lib/pkgconfig"

# 或者设置编译器 flags
CGO_LDFLAGS="-L/ffmpeg/lib -lavcodec -lavdevice -lavfilter -lavformat -lavutil -lswresample -lswscale"
CGO_CFLAGS="-I/ffmpeg/include"
```

## 例子

检查 [advance/example](./advance/example) 目录。

可以使用Docker运行例子

```bash
# 运行 Debian 13 (trixie) 容器
sudo docker run --rm -it debian:trixie
# 安装开发环境
apt install golang git
# Clone 项目
git clone https://github.com/Lensual/go-libav
# 安装FFmpeg开发库
apt install libavcodec-dev libavdevice-dev libavfilter-dev libavformat-dev libavutil-dev libswresample-dev libswscale-dev
# 构建并运行例子'mux'
cd go-libav/example/mux
go build .
./mux ./test.mp4
```

## License

[GNU Lesser General Public License version 2.1](./LICENSE)
