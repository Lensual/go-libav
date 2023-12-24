package goavutil

import (
	"errors"

	"github.com/Lensual/go-libav/avutil"
)

func AvErr(code int) error {
	return errors.New(avutil.AvErr2str(code))
}
