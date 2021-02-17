// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"math"
)

type FLOAT struct {
	Mantissa int64
	Exponent int8
}

func (f *FLOAT) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, f.Mantissa); err != nil {
		return err
	}
	return nil
}

func (f *FLOAT) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !f.MantissaInActingVersion(actingVersion) {
		f.Mantissa = f.MantissaNullValue()
	} else {
		if err := _m.ReadInt64(_r, &f.Mantissa); err != nil {
			return err
		}
	}
	f.Exponent = -7
	return nil
}

func (f *FLOAT) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if f.MantissaInActingVersion(actingVersion) {
		if f.Mantissa < f.MantissaMinValue() || f.Mantissa > f.MantissaMaxValue() {
			return fmt.Errorf("Range check failed on f.Mantissa (%v < %v > %v)", f.MantissaMinValue(), f.Mantissa, f.MantissaMaxValue())
		}
	}
	return nil
}

func FLOATInit(f *FLOAT) {
	f.Exponent = -7
	return
}

func (*FLOAT) EncodedLength() int64 {
	return 8
}

func (*FLOAT) MantissaMinValue() int64 {
	return math.MinInt64 + 1
}

func (*FLOAT) MantissaMaxValue() int64 {
	return math.MaxInt64
}

func (*FLOAT) MantissaNullValue() int64 {
	return math.MinInt64
}

func (*FLOAT) MantissaSinceVersion() uint16 {
	return 0
}

func (f *FLOAT) MantissaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= f.MantissaSinceVersion()
}

func (*FLOAT) MantissaDeprecated() uint16 {
	return 0
}

func (*FLOAT) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*FLOAT) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*FLOAT) ExponentNullValue() int8 {
	return math.MinInt8
}

func (*FLOAT) ExponentSinceVersion() uint16 {
	return 0
}

func (f *FLOAT) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= f.ExponentSinceVersion()
}

func (*FLOAT) ExponentDeprecated() uint16 {
	return 0
}
