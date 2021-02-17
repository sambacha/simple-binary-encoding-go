// Generated SBE (Simple Binary Encoding) message codec

package composite_elements

import (
	"fmt"
	"io"
	"math"
)

type Outer struct {
	EnumOne EnumOneEnum
	Zeroth  uint8
	SetOne  SetOne
	Inner   OuterInner
}

func (o *Outer) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := o.EnumOne.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.Zeroth); err != nil {
		return err
	}
	if err := o.SetOne.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.Inner.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (o *Outer) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
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
	if o.SetOneInActingVersion(actingVersion) {
		if err := o.SetOne.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.InnerInActingVersion(actingVersion) {
		if err := o.Inner.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *Outer) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func OuterInit(o *Outer) {
	return
}

func (*Outer) EncodedLength() int64 {
	return 22
}

func (*Outer) EnumOneSinceVersion() uint16 {
	return 0
}

func (o *Outer) EnumOneInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.EnumOneSinceVersion()
}

func (*Outer) EnumOneDeprecated() uint16 {
	return 0
}

func (*Outer) ZerothMinValue() uint8 {
	return 0
}

func (*Outer) ZerothMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*Outer) ZerothNullValue() uint8 {
	return math.MaxUint8
}

func (*Outer) ZerothSinceVersion() uint16 {
	return 0
}

func (o *Outer) ZerothInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ZerothSinceVersion()
}

func (*Outer) ZerothDeprecated() uint16 {
	return 0
}

func (*Outer) SetOneSinceVersion() uint16 {
	return 0
}

func (o *Outer) SetOneInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SetOneSinceVersion()
}

func (*Outer) SetOneDeprecated() uint16 {
	return 0
}

func (*Outer) InnerSinceVersion() uint16 {
	return 0
}

func (o *Outer) InnerInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.InnerSinceVersion()
}

func (*Outer) InnerDeprecated() uint16 {
	return 0
}
