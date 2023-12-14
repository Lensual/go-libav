package goavfilter

import "github.com/Lensual/go-libav/avfilter"

type AVFilter struct {
	CAVFilter *avfilter.CAVFilter
}

func GetByName(name string) *AVFilter {
	f := avfilter.AvfilterGetByName(name)
	if f == nil {
		return nil
	}
	return &AVFilter{
		CAVFilter: f,
	}
}
