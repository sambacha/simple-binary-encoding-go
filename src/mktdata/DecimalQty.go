// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"math"
)

type DecimalQty struct {
	Mantissa int32
	Exponent int8
}

func (d *DecimalQty) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt32(_w, d.Mantissa); err != nil {
		return err
	}
	return nil
}

func (d *DecimalQty) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !d.MantissaInActingVersion(actingVersion) {
		d.Mantissa = d.MantissaNullValue()
	} else {
		if err := _m.ReadInt32(_r, &d.Mantissa); err != nil {
			return err
		}
	}
	d.Exponent = -4
	return nil
}

func (d *DecimalQty) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if d.MantissaInActingVersion(actingVersion) {
		if d.Mantissa != d.MantissaNullValue() && (d.Mantissa < d.MantissaMinValue() || d.Mantissa > d.MantissaMaxValue()) {
			return fmt.Errorf("Range check failed on d.Mantissa (%v < %v > %v)", d.MantissaMinValue(), d.Mantissa, d.MantissaMaxValue())
		}
	}
	return nil
}

func DecimalQtyInit(d *DecimalQty) {
	d.Mantissa = 2147483647
	d.Exponent = -4
	return
}

func (*DecimalQty) EncodedLength() int64 {
	return 4
}

func (*DecimalQty) MantissaMinValue() int32 {
	return math.MinInt32 + 1
}

func (*DecimalQty) MantissaMaxValue() int32 {
	return math.MaxInt32
}

func (*DecimalQty) MantissaNullValue() int32 {
	return 2147483647
}

func (*DecimalQty) MantissaSinceVersion() uint16 {
	return 0
}

func (d *DecimalQty) MantissaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.MantissaSinceVersion()
}

func (*DecimalQty) MantissaDeprecated() uint16 {
	return 0
}

func (*DecimalQty) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*DecimalQty) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*DecimalQty) ExponentNullValue() int8 {
	return math.MinInt8
}

func (*DecimalQty) ExponentSinceVersion() uint16 {
	return 0
}

func (d *DecimalQty) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= d.ExponentSinceVersion()
}

func (*DecimalQty) ExponentDeprecated() uint16 {
	return 0
}
