package avcodec

/*
#cgo LDFLAGS: -lavcodec

#include "libavcodec/avcodec.h"

#include <stdint.h>
*/
import "C"

/**
 * A channel layout is a 64-bits integer with a bit set for every channel.
 * The number of bits set must be equal to the number of channels.
 * The value 0 means that the channel layout is not known.
 * @note this data structure is not powerful enough to handle channels
 * combinations that have the same channel multiple times, such as
 * dual-mono.
 */
type AV_CH_LAYOUT C.uint64_t

const (
	AV_CH_LAYOUT_Mono              AV_CH_LAYOUT = C.AV_CH_LAYOUT_MONO
	AV_CH_LAYOUT_Stereo            AV_CH_LAYOUT = C.AV_CH_LAYOUT_STEREO
	AV_CH_LAYOUT_2Point1           AV_CH_LAYOUT = C.AV_CH_LAYOUT_2POINT1
	AV_CH_LAYOUT_2_1               AV_CH_LAYOUT = C.AV_CH_LAYOUT_2_1
	AV_CH_LAYOUT_Surround          AV_CH_LAYOUT = C.AV_CH_LAYOUT_SURROUND
	AV_CH_LAYOUT_3Point1           AV_CH_LAYOUT = C.AV_CH_LAYOUT_3POINT1
	AV_CH_LAYOUT_4Point0           AV_CH_LAYOUT = C.AV_CH_LAYOUT_4POINT0
	AV_CH_LAYOUT_4Point1           AV_CH_LAYOUT = C.AV_CH_LAYOUT_4POINT1
	AV_CH_LAYOUT_2_2               AV_CH_LAYOUT = C.AV_CH_LAYOUT_2_2
	AV_CH_LAYOUT_Quad              AV_CH_LAYOUT = C.AV_CH_LAYOUT_QUAD
	AV_CH_LAYOUT_5Point0           AV_CH_LAYOUT = C.AV_CH_LAYOUT_5POINT0
	AV_CH_LAYOUT_5Point1           AV_CH_LAYOUT = C.AV_CH_LAYOUT_5POINT1
	AV_CH_LAYOUT_5Point0_Back      AV_CH_LAYOUT = C.AV_CH_LAYOUT_5POINT0_BACK
	AV_CH_LAYOUT_5Point1_Back      AV_CH_LAYOUT = C.AV_CH_LAYOUT_5POINT1_BACK
	AV_CH_LAYOUT_6Point0           AV_CH_LAYOUT = C.AV_CH_LAYOUT_6POINT0
	AV_CH_LAYOUT_6Point0_Front     AV_CH_LAYOUT = C.AV_CH_LAYOUT_6POINT0_FRONT
	AV_CH_LAYOUT_Hexagonal         AV_CH_LAYOUT = C.AV_CH_LAYOUT_HEXAGONAL
	AV_CH_LAYOUT_6Point1           AV_CH_LAYOUT = C.AV_CH_LAYOUT_6POINT1
	AV_CH_LAYOUT_6Point1_Back      AV_CH_LAYOUT = C.AV_CH_LAYOUT_6POINT1_BACK
	AV_CH_LAYOUT_6Point1_Front     AV_CH_LAYOUT = C.AV_CH_LAYOUT_6POINT1_FRONT
	AV_CH_LAYOUT_7Point0           AV_CH_LAYOUT = C.AV_CH_LAYOUT_7POINT0
	AV_CH_LAYOUT_7Point0_Front     AV_CH_LAYOUT = C.AV_CH_LAYOUT_7POINT0_FRONT
	AV_CH_LAYOUT_7Point1           AV_CH_LAYOUT = C.AV_CH_LAYOUT_7POINT1
	AV_CH_LAYOUT_7Point1_Wide      AV_CH_LAYOUT = C.AV_CH_LAYOUT_7POINT1_WIDE
	AV_CH_LAYOUT_7Point1_Wide_Back AV_CH_LAYOUT = C.AV_CH_LAYOUT_7POINT1_WIDE_BACK
	AV_CH_LAYOUT_Octagonal         AV_CH_LAYOUT = C.AV_CH_LAYOUT_OCTAGONAL
	AV_CH_LAYOUT_Hexadecagonal     AV_CH_LAYOUT = C.AV_CH_LAYOUT_HEXADECAGONAL
	AV_CH_LAYOUT_Stero_Downmix     AV_CH_LAYOUT = C.AV_CH_LAYOUT_STEREO_DOWNMIX
	// AV_CH_LAYOUT_22POINT2          AV_CH_LAYOUT = C.AV_CH_LAYOUT_22POINT2
)
