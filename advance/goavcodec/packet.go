package goavcodec

import (
	"unsafe"

	"github.com/Lensual/go-libav/avcodec"
	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/ctypes"
)

type AVPacket struct {
	CAVPacket *avcodec.CAVPacket
}

//#region member

func (pkt *AVPacket) GetBuf() *avutil.CAVBufferRef {
	return pkt.CAVPacket.GetBuf()
}

func (pkt *AVPacket) SetBuf(buf *avutil.CAVBufferRef) {
	pkt.CAVPacket.SetBuf(buf)
}

func (pkt *AVPacket) GetPts() int64 {
	return pkt.CAVPacket.GetPts()
}

func (pkt *AVPacket) SetPts(pts int64) {
	pkt.CAVPacket.SetPts(pts)
}

func (pkt *AVPacket) GetDts() int64 {
	return pkt.CAVPacket.GetDts()
}

func (pkt *AVPacket) SetDts(dts int64) {
	pkt.CAVPacket.SetDts(dts)
}

func (pkt *AVPacket) GetData() []byte {
	return unsafe.Slice((*byte)(pkt.CAVPacket.GetData()), pkt.CAVPacket.GetSize())
}

func (pkt *AVPacket) SetData(data []byte) {
	pkt.CAVPacket.SetData(unsafe.Pointer(unsafe.SliceData(data)))
}

func (pkt *AVPacket) GetSize() int {
	return pkt.CAVPacket.GetSize()
}

func (pkt *AVPacket) SetSize(size int) {
	pkt.CAVPacket.SetSize(size)
}

func (pkt *AVPacket) GetStreamIndex() int {
	return pkt.CAVPacket.GetStreamIndex()
}

func (pkt *AVPacket) SetStreamIndex(streamIndex int) {
	pkt.CAVPacket.SetStreamIndex(streamIndex)
}

func (pkt *AVPacket) GetFlags() int {
	return pkt.CAVPacket.GetFlags()
}

func (pkt *AVPacket) SetFlags(flags int) {
	pkt.CAVPacket.SetFlags(flags)
}

func (pkt *AVPacket) GetSideData() *avcodec.CAVPacketSideData {
	return pkt.CAVPacket.GetSideData()
}

func (pkt *AVPacket) SetSideData(sideData *avcodec.CAVPacketSideData) {
	pkt.CAVPacket.SetSideData(sideData)
}

func (pkt *AVPacket) GetSideDataElems() int {
	return pkt.CAVPacket.GetSideDataElems()
}

func (pkt *AVPacket) SetSideDataElems(sideDataElems int) {
	pkt.CAVPacket.SetSideDataElems(sideDataElems)
}

func (pkt *AVPacket) GetDuration() int64 {
	return pkt.CAVPacket.GetDuration()
}

func (pkt *AVPacket) SetDuration(duration int64) {
	pkt.CAVPacket.SetDuration(duration)
}

func (pkt *AVPacket) GetPos() int64 {
	return pkt.CAVPacket.GetPos()
}

func (pkt *AVPacket) SetPos(pos int64) {
	pkt.CAVPacket.SetPos(pos)
}

func (pkt *AVPacket) GetOpaque() unsafe.Pointer {
	return pkt.CAVPacket.GetOpaque()
}

func (pkt *AVPacket) SetOpaque(opaque unsafe.Pointer) {
	pkt.CAVPacket.SetOpaque(opaque)
}

func (pkt *AVPacket) GetOpaqueRef() *avutil.CAVBufferRef {
	return pkt.CAVPacket.GetOpaqueRef()
}

func (pkt *AVPacket) SetOpaqueRef(opaqueRef *avutil.CAVBufferRef) {
	pkt.CAVPacket.SetOpaqueRef(opaqueRef)
}

func (pkt *AVPacket) GetTimeBase() avutil.CAVRational {
	return pkt.CAVPacket.GetTimeBase()
}

func (pkt *AVPacket) SetTimeBase(timeBase avutil.CAVRational) {
	pkt.CAVPacket.SetTimeBase(timeBase)
}

//#endregion member

func (pkt *AVPacket) Clone() *AVPacket {
	cPkt := avcodec.AvPacketClone(pkt.CAVPacket)
	if cPkt == nil {
		return nil
	}
	return &AVPacket{
		CAVPacket: cPkt,
	}
}

func (pkt *AVPacket) Free() {
	if pkt.CAVPacket != nil {
		avcodec.AvPacketFree(&pkt.CAVPacket)
		pkt.CAVPacket = nil
	}
}

func (pkt *AVPacket) Shrink(size int) {
	avcodec.AvShrinkPacket(pkt.CAVPacket, size)
}

func (pkt *AVPacket) Grow(growBy int) int {
	return avcodec.AvGrowPacket(pkt.CAVPacket, growBy)
}

func (pkt *AVPacket) RefTo(dst *AVPacket) int {
	return avcodec.AvPacketRef(dst.CAVPacket, pkt.CAVPacket)
}

func (pkt *AVPacket) Unref() {
	avcodec.AvPacketUnref(pkt.CAVPacket)
}

func (pkt *AVPacket) MoveRefTo(dst *AVPacket) {
	avcodec.AvPacketMoveRef(dst.CAVPacket, pkt.CAVPacket)
}

func (pkt *AVPacket) CopyPropsTo(dst *AVPacket) int {
	return avcodec.AvPacketCopyProps(dst.CAVPacket, pkt.CAVPacket)
}

func (pkt *AVPacket) MakeRefcounted() int {
	return avcodec.AvPacketMakeRefcounted(pkt.CAVPacket)
}

func (pkt *AVPacket) MakeWritable() int {
	return avcodec.AvPacketMakeWritable(pkt.CAVPacket)
}

func (pkt *AVPacket) RescaleTs(tbSrc avutil.CAVRational, tbDst avutil.CAVRational) {
	avcodec.AvPacketRescaleTs(pkt.CAVPacket, tbSrc, tbDst)
}

func NewAvPacketFromUnsafeData(data unsafe.Pointer, size int) (*AVPacket, int) {
	var cPkt avcodec.CAVPacket
	code := avcodec.AvPacketFromData(&cPkt, data, size)
	if code != 0 {
		return nil, code
	}
	return &AVPacket{
		CAVPacket: &cPkt,
	}, code
}

func NewAvPacketFromData(data []byte) (*AVPacket, int) {
	dataSize := len(data)
	cData := avutil.AvMallocz(ctypes.SizeT(dataSize + avcodec.AV_INPUT_BUFFER_PADDING_SIZE))
	copy(unsafe.Slice((*byte)(cData), dataSize), data)
	pkt, code := NewAvPacketFromUnsafeData(cData, dataSize)
	if code != 0 {
		avutil.AvFree(cData)
		return nil, code
	}
	return pkt, code
}

func NewAvPacketFromDataWithoutPadding(data []byte) (*AVPacket, int) {
	dataSize := len(data)
	cData := avutil.AvMalloc(ctypes.SizeT(dataSize))
	copy(unsafe.Slice((*byte)(cData), dataSize), data)
	pkt, code := NewAvPacketFromUnsafeData(cData, dataSize)
	if code != 0 {
		avutil.AvFree(cData)
		return nil, code
	}
	return pkt, code
}

func AllocAvPacket() *AVPacket {
	cPkt := avcodec.AvPacketAlloc()
	if cPkt == nil {
		return nil
	}
	return &AVPacket{
		CAVPacket: cPkt,
	}
}

func NewAvPacket(size int) (*AVPacket, int) {
	var cPkt avcodec.CAVPacket
	code := avcodec.AvNewPacket(&cPkt, size)
	if code != 0 {
		return nil, code
	}

	return &AVPacket{
		CAVPacket: &cPkt,
	}, code
}
