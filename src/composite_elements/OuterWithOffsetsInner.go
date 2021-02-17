// Generated SBE (Simple Binary Encoding) message codec

package composite_elements

import (
	"fmt"
	"io"
	"math"
)

type OuterWithOffsetsInner struct {
	First  int64
	Second int64
}

func (o *OuterWithOffsetsInner) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, o.First); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.Second); err != nil {
		return err
	}
	return nil
}

func (o *OuterWithOffsetsInner) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !o.FirstInActingVersion(actingVersion) {
		o.First = o.FirstNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.First); err != nil {
			return err
		}
	}
	if !o.SecondInActingVersion(actingVersion) {
		o.Second = o.SecondNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.Second); err != nil {
			return err
		}
	}
	return nil
}

func (o *OuterWithOffsetsInner) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.FirstInActingVersion(actingVersion) {
		if o.First < o.FirstMinValue() || o.First > o.FirstMaxValue() {
			return fmt.Errorf("Range check failed on o.First (%v < %v > %v)", o.FirstMinValue(), o.First, o.FirstMaxValue())
		}
	}
	if o.SecondInActingVersion(actingVersion) {
		if o.Second < o.SecondMinValue() || o.Second > o.SecondMaxValue() {
			return fmt.Errorf("Range check failed on o.Second (%v < %v > %v)", o.SecondMinValue(), o.Second, o.SecondMaxValue())
		}
	}
	return nil
}

func OuterWithOffsetsInnerInit(o *OuterWithOffsetsInner) {
	return
}

func (*OuterWithOffsetsInner) EncodedLength() int64 {
	return 16
}

func (*OuterWithOffsetsInner) FirstMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OuterWithOffsetsInner) FirstMaxValue() int64 {
	return math.MaxInt64
}

func (*OuterWithOffsetsInner) FirstNullValue() int64 {
	return math.MinInt64
}

func (*OuterWithOffsetsInner) FirstSinceVersion() uint16 {
	return 0
}

func (o *OuterWithOffsetsInner) FirstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.FirstSinceVersion()
}

func (*OuterWithOffsetsInner) FirstDeprecated() uint16 {
	return 0
}

func (*OuterWithOffsetsInner) SecondMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OuterWithOffsetsInner) SecondMaxValue() int64 {
	return math.MaxInt64
}

func (*OuterWithOffsetsInner) SecondNullValue() int64 {
	return math.MinInt64
}

func (*OuterWithOffsetsInner) SecondSinceVersion() uint16 {
	return 0
}

func (o *OuterWithOffsetsInner) SecondInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecondSinceVersion()
}

func (*OuterWithOffsetsInner) SecondDeprecated() uint16 {
	return 0
}
