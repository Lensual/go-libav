package goavutil

import (
	"github.com/pkg/errors"

	"github.com/Lensual/go-libav/avutil"
)

func AvErr(code int) error {
	return errors.New(avutil.AvErr2str(code))
}
