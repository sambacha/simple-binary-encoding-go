// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type GroupSize8Byte struct {
	BlockLength uint16
	NumInGroup  uint8
}

func (g *GroupSize8Byte) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, g.BlockLength); err != nil {
		return err
	}

	for i := 0; i < 5; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, g.NumInGroup); err != nil {
		return err
	}
	return nil
}

func (g *GroupSize8Byte) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !g.BlockLengthInActingVersion(actingVersion) {
		g.BlockLength = g.BlockLengthNullValue()
	} else {
		if err := _m.ReadUint16(_r, &g.BlockLength); err != nil {
			return err
		}
	}
	io.CopyN(ioutil.Discard, _r, 5)
	if !g.NumInGroupInActingVersion(actingVersion) {
		g.NumInGroup = g.NumInGroupNullValue()
	} else {
		if err := _m.ReadUint8(_r, &g.NumInGroup); err != nil {
			return err
		}
	}
	return nil
}

func (g *GroupSize8Byte) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if g.BlockLengthInActingVersion(actingVersion) {
		if g.BlockLength < g.BlockLengthMinValue() || g.BlockLength > g.BlockLengthMaxValue() {
			return fmt.Errorf("Range check failed on g.BlockLength (%v < %v > %v)", g.BlockLengthMinValue(), g.BlockLength, g.BlockLengthMaxValue())
		}
	}
	if g.NumInGroupInActingVersion(actingVersion) {
		if g.NumInGroup < g.NumInGroupMinValue() || g.NumInGroup > g.NumInGroupMaxValue() {
			return fmt.Errorf("Range check failed on g.NumInGroup (%v < %v > %v)", g.NumInGroupMinValue(), g.NumInGroup, g.NumInGroupMaxValue())
		}
	}
	return nil
}

func GroupSize8ByteInit(g *GroupSize8Byte) {
	return
}

func (*GroupSize8Byte) EncodedLength() int64 {
	return 8
}

func (*GroupSize8Byte) BlockLengthMinValue() uint16 {
	return 0
}

func (*GroupSize8Byte) BlockLengthMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*GroupSize8Byte) BlockLengthNullValue() uint16 {
	return math.MaxUint16
}

func (*GroupSize8Byte) BlockLengthSinceVersion() uint16 {
	return 0
}

func (g *GroupSize8Byte) BlockLengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= g.BlockLengthSinceVersion()
}

func (*GroupSize8Byte) BlockLengthDeprecated() uint16 {
	return 0
}

func (*GroupSize8Byte) NumInGroupMinValue() uint8 {
	return 0
}

func (*GroupSize8Byte) NumInGroupMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*GroupSize8Byte) NumInGroupNullValue() uint8 {
	return math.MaxUint8
}

func (*GroupSize8Byte) NumInGroupSinceVersion() uint16 {
	return 0
}

func (g *GroupSize8Byte) NumInGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= g.NumInGroupSinceVersion()
}

func (*GroupSize8Byte) NumInGroupDeprecated() uint16 {
	return 0
}
