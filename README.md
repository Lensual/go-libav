# DEPRECATED

根本没有必要开发这个库。不如直接在Go项目里写C模块。对于不会C和不熟悉手动内存管理的Go开发者，不如直接调用命令。

# go-libav

[![Go Reference](https://pkg.go.dev/badge/github.com/Lensual/go-libav.svg)](https://pkg.go.dev/github.com/Lensual/go-libav)
![GitHub release (with filter)](https://img.shields.io/github/v/release/Lensual/go-libav?include_prereleases&color=blue)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/Lensual/go-libav)
![Static Badge](https://img.shields.io/github/license/Lensual/go-libav)
[![Go Report Card](https://goreportcard.com/badge/github.com/Lensual/go-libav)](https://goreportcard.com/report/github.com/Lensual/go-libav)
![Static Badge](https://img.shields.io/badge/FFmpeg_version-6.1-purple)

[English](./README.md) | [中文](./README_zh.md)

Golang binding for [FFmpeg and libav](https://ffmpeg.org/).

## Things to know

1. This package is implemented using CGO and requires FFmpeg's header files and libraries for compilation.
2. Due to well-known performance issues with CGO calls, it is recommended to use C/CPP to achieve your needs.

## Getting Started

Package `go-libav/av*` are cgo bindings. Before using this package, you'd better to master the basic usage of the FFmpeg library.

Package `golibav/advance` is a secondary package to simplify use in go.

Need Golang version `1.21` and FFmpeg library version `6.1`.

Set your environment variable if your library is not in the default directory.

```env
# using pkg-config
PKG_CONFIG_PATH="/ffmpeg/lib/pkgconfig"

# or set compiler flags
CGO_LDFLAGS="-L/ffmpeg/lib -lavcodec -lavdevice -lavfilter -lavformat -lavutil -lswresample -lswscale"
CGO_CFLAGS="-I/ffmpeg/include"
```

## Examples

Check [advance/example](./advance/example) directory.

Try examples With Docker

```bash
# Run container with Debian 13 (trixie).
sudo docker run --rm -it debian:trixie
# Install development environment.
apt install golang git
# Clone project.
git clone https://github.com/Lensual/go-libav
# Install FFmpeg development libraries.
apt install libavcodec-dev libavdevice-dev libavfilter-dev libavformat-dev libavutil-dev libswresample-dev libswscale-dev
# Build example like 'mux' and then run iut.
cd go-libav/example/mux
go build .
./mux ./test.mp4
```

## License

[GNU Lesser General Public License version 2.1](./LICENSE)
