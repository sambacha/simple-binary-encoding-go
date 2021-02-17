// Generated SBE (Simple Binary Encoding) message codec

package composite_elements

import (
	"fmt"
	"io"
	"math"
)

type Inner struct {
	First  int64
	Second int64
}

func (i *Inner) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, i.First); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.Second); err != nil {
		return err
	}
	return nil
}

func (i *Inner) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !i.FirstInActingVersion(actingVersion) {
		i.First = i.FirstNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.First); err != nil {
			return err
		}
	}
	if !i.SecondInActingVersion(actingVersion) {
		i.Second = i.SecondNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.Second); err != nil {
			return err
		}
	}
	return nil
}

func (i *Inner) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if i.FirstInActingVersion(actingVersion) {
		if i.First < i.FirstMinValue() || i.First > i.FirstMaxValue() {
			return fmt.Errorf("Range check failed on i.First (%v < %v > %v)", i.FirstMinValue(), i.First, i.FirstMaxValue())
		}
	}
	if i.SecondInActingVersion(actingVersion) {
		if i.Second < i.SecondMinValue() || i.Second > i.SecondMaxValue() {
			return fmt.Errorf("Range check failed on i.Second (%v < %v > %v)", i.SecondMinValue(), i.Second, i.SecondMaxValue())
		}
	}
	return nil
}

func InnerInit(i *Inner) {
	return
}

func (*Inner) EncodedLength() int64 {
	return 16
}

func (*Inner) FirstMinValue() int64 {
	return math.MinInt64 + 1
}

func (*Inner) FirstMaxValue() int64 {
	return math.MaxInt64
}

func (*Inner) FirstNullValue() int64 {
	return math.MinInt64
}

func (*Inner) FirstSinceVersion() uint16 {
	return 0
}

func (i *Inner) FirstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.FirstSinceVersion()
}

func (*Inner) FirstDeprecated() uint16 {
	return 0
}

func (*Inner) SecondMinValue() int64 {
	return math.MinInt64 + 1
}

func (*Inner) SecondMaxValue() int64 {
	return math.MaxInt64
}

func (*Inner) SecondNullValue() int64 {
	return math.MinInt64
}

func (*Inner) SecondSinceVersion() uint16 {
	return 0
}

func (i *Inner) SecondInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.SecondSinceVersion()
}

func (*Inner) SecondDeprecated() uint16 {
	return 0
}
