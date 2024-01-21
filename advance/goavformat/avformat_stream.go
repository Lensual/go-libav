package goavformat

import (
	"unsafe"

	"github.com/Lensual/go-libav/avcodec"
	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
)

type AVStream struct {
	CAVStream *avformat.CAVStream
}

//#region members

func (st *AVStream) GetAvClass() *avutil.CAVClass {
	return st.CAVStream.GetAvClass()
}

func (st *AVStream) GetIndex() int {
	return st.CAVStream.GetIndex()
}
func (st *AVStream) SetIndex(index int) {
	st.CAVStream.SetIndex(index)
}

func (st *AVStream) GetId() int {
	return st.CAVStream.GetId()
}
func (st *AVStream) SetId(id int) {
	st.CAVStream.SetId(id)
}

func (st *AVStream) GetCodecPar() *avcodec.CAVCodecParameters {
	return st.CAVStream.GetCodecPar()
}
func (st *AVStream) SetCodecPar(codecPar *avcodec.CAVCodecParameters) {
	st.CAVStream.SetCodecPar(codecPar)
}

func (st *AVStream) GetPrivData() unsafe.Pointer {
	return st.CAVStream.GetPrivData()
}
func (st *AVStream) SetPrivData(privData unsafe.Pointer) {
	st.CAVStream.SetPrivData(privData)
}

func (st *AVStream) GetTimeBase() avutil.CAVRational {
	return st.CAVStream.GetTimeBase()
}
func (st *AVStream) SetTimeBase(timeBase avutil.CAVRational) {
	st.CAVStream.SetTimeBase(timeBase)
}

//#endregion members
