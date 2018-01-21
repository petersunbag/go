package dynamic

import (
	"unsafe"
	"github.com/thrift-iterator/go/spi"
)

type boolEncoder struct {
}

func (encoder *boolEncoder) encode(ptr unsafe.Pointer, iter spi.Stream) {
	iter.WriteBool(*(*bool)(ptr))
}

type int32Encoder struct {
}

func (encoder *int32Encoder) encode(ptr unsafe.Pointer, iter spi.Stream) {
	iter.WriteInt32(*(*int32)(ptr))
}