// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"math"
)

type Decimal9 struct {
	Mantissa int64
	Exponent int8
}

func (d *Decimal9) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, d.Mantissa); err != nil {
		return err
	}
	return nil
}

func (d *Decimal9) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !d.MantissaInActingVersion(actingVersion) {
		d.Mantissa = d.MantissaNullValue()
	} else {
		if err := _m.ReadInt64(_r, &d.Mantissa); err != nil {
			return err
		}
	}
	d.Exponent = -9
	return nil
}

func (d *Decimal9) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if d.MantissaInActingVersion(actingVersion) {
		if d.Mantissa < d.MantissaMinValue() || d.Mantissa > d.MantissaMaxValue() {
			return fmt.Errorf("Range check failed on d.Mantissa (%v < %v > %v)", d.MantissaMinValue(), d.Mantissa, d.MantissaMaxValue())
		}
	}
	return nil
}

func Decimal9Init(d *Decimal9) {
	d.Exponent = -9
	return
}

func (*Decimal9) EncodedLength() int64 {
	return 8
}

func (*Decimal9) MantissaMinValue() int64 {
	return math.MinInt64 + 1
}

func (*Decimal9) MantissaMaxValue() int64 {
	return math.MaxInt64
}

func (*Decimal9) MantissaNullValue() int64 {
	return math.MinInt64
}

func (*Decimal9) MantissaSinceVersion() uint16 {
	return 0
}

func (d *Decimal9) MantissaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.MantissaSinceVersion()
}

func (*Decimal9) MantissaDeprecated() uint16 {
	return 0
}

func (*Decimal9) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*Decimal9) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*Decimal9) ExponentNullValue() int8 {
	return math.MinInt8
}

func (*Decimal9) ExponentSinceVersion() uint16 {
	return 0
}

func (d *Decimal9) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.ExponentSinceVersion()
}

func (*Decimal9) ExponentDeprecated() uint16 {
	return 0
}
