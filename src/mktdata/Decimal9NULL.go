// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"math"
)

type Decimal9NULL struct {
	Mantissa int64
	Exponent int8
}

func (d *Decimal9NULL) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, d.Mantissa); err != nil {
		return err
	}
	return nil
}

func (d *Decimal9NULL) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
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

func (d *Decimal9NULL) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if d.MantissaInActingVersion(actingVersion) {
		if d.Mantissa != d.MantissaNullValue() && (d.Mantissa < d.MantissaMinValue() || d.Mantissa > d.MantissaMaxValue()) {
			return fmt.Errorf("Range check failed on d.Mantissa (%v < %v > %v)", d.MantissaMinValue(), d.Mantissa, d.MantissaMaxValue())
		}
	}
	return nil
}

func Decimal9NULLInit(d *Decimal9NULL) {
	d.Mantissa = int64(9223372036854775807)
	d.Exponent = -9
	return
}

func (*Decimal9NULL) EncodedLength() int64 {
	return 8
}

func (*Decimal9NULL) MantissaMinValue() int64 {
	return math.MinInt64 + 1
}

func (*Decimal9NULL) MantissaMaxValue() int64 {
	return math.MaxInt64
}

func (*Decimal9NULL) MantissaNullValue() int64 {
	return int64(9223372036854775807)
}

func (*Decimal9NULL) MantissaSinceVersion() uint16 {
	return 0
}

func (d *Decimal9NULL) MantissaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.MantissaSinceVersion()
}

func (*Decimal9NULL) MantissaDeprecated() uint16 {
	return 0
}

func (*Decimal9NULL) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*Decimal9NULL) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*Decimal9NULL) ExponentNullValue() int8 {
	return math.MinInt8
}

func (*Decimal9NULL) ExponentSinceVersion() uint16 {
	return 0
}

func (d *Decimal9NULL) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.ExponentSinceVersion()
}

func (*Decimal9NULL) ExponentDeprecated() uint16 {
	return 0
}
