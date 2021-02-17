// Generated SBE (Simple Binary Encoding) message codec

package composite_elements

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OuterWithOffsets struct {
	EnumOne EnumOneEnum
	Zeroth  uint8
	SetOne  SetOne
	Inner   OuterWithOffsetsInner
}

func (o *OuterWithOffsets) Encode(_m *SbeGoMarshaller, _w io.Writer) error {

	for i := 0; i < 2; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	if err := o.EnumOne.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.Zeroth); err != nil {
		return err
	}

	for i := 0; i < 4; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	if err := o.SetOne.Encode(_m, _w); err != nil {
		return err
	}

	for i := 0; i < 4; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	if err := o.Inner.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (o *OuterWithOffsets) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	io.CopyN(ioutil.Discard, _r, 2)
	if o.EnumOneInActingVersion(actingVersion) {
		if err := o.EnumOne.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.ZerothInActingVersion(actingVersion) {
		o.Zeroth = o.ZerothNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.Zeroth); err != nil {
			return err
		}
	}
	io.CopyN(ioutil.Discard, _r, 4)
	if o.SetOneInActingVersion(actingVersion) {
		if err := o.SetOne.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	io.CopyN(ioutil.Discard, _r, 4)
	if o.InnerInActingVersion(actingVersion) {
		if err := o.Inner.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *OuterWithOffsets) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.ZerothInActingVersion(actingVersion) {
		if o.Zeroth < o.ZerothMinValue() || o.Zeroth > o.ZerothMaxValue() {
			return fmt.Errorf("Range check failed on o.Zeroth (%v < %v > %v)", o.ZerothMinValue(), o.Zeroth, o.ZerothMaxValue())
		}
	}
	if err := o.Inner.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func OuterWithOffsetsInit(o *OuterWithOffsets) {
	return
}

func (*OuterWithOffsets) EncodedLength() int64 {
	return 32
}

func (*OuterWithOffsets) EnumOneSinceVersion() uint16 {
	return 0
}

func (o *OuterWithOffsets) EnumOneInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.EnumOneSinceVersion()
}

func (*OuterWithOffsets) EnumOneDeprecated() uint16 {
	return 0
}

func (*OuterWithOffsets) ZerothMinValue() uint8 {
	return 0
}

func (*OuterWithOffsets) ZerothMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OuterWithOffsets) ZerothNullValue() uint8 {
	return math.MaxUint8
}

func (*OuterWithOffsets) ZerothSinceVersion() uint16 {
	return 0
}

func (o *OuterWithOffsets) ZerothInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ZerothSinceVersion()
}

func (*OuterWithOffsets) ZerothDeprecated() uint16 {
	return 0
}

func (*OuterWithOffsets) SetOneSinceVersion() uint16 {
	return 0
}

func (o *OuterWithOffsets) SetOneInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SetOneSinceVersion()
}

func (*OuterWithOffsets) SetOneDeprecated() uint16 {
	return 0
}

func (*OuterWithOffsets) InnerSinceVersion() uint16 {
	return 0
}

func (o *OuterWithOffsets) InnerInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.InnerSinceVersion()
}

func (*OuterWithOffsets) InnerDeprecated() uint16 {
	return 0
}
