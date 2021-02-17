// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"math"
)

type PRICENULL struct {
	Mantissa int64
	Exponent int8
}

func (p *PRICENULL) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, p.Mantissa); err != nil {
		return err
	}
	return nil
}

func (p *PRICENULL) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !p.MantissaInActingVersion(actingVersion) {
		p.Mantissa = p.MantissaNullValue()
	} else {
		if err := _m.ReadInt64(_r, &p.Mantissa); err != nil {
			return err
		}
	}
	p.Exponent = -7
	return nil
}

func (p *PRICENULL) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.MantissaInActingVersion(actingVersion) {
		if p.Mantissa != p.MantissaNullValue() && (p.Mantissa < p.MantissaMinValue() || p.Mantissa > p.MantissaMaxValue()) {
			return fmt.Errorf("Range check failed on p.Mantissa (%v < %v > %v)", p.MantissaMinValue(), p.Mantissa, p.MantissaMaxValue())
		}
	}
	return nil
}

func PRICENULLInit(p *PRICENULL) {
	p.Mantissa = int64(9223372036854775807)
	p.Exponent = -7
	return
}

func (*PRICENULL) EncodedLength() int64 {
	return 8
}

func (*PRICENULL) MantissaMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PRICENULL) MantissaMaxValue() int64 {
	return math.MaxInt64
}

func (*PRICENULL) MantissaNullValue() int64 {
	return int64(9223372036854775807)
}

func (*PRICENULL) MantissaSinceVersion() uint16 {
	return 0
}

func (p *PRICENULL) MantissaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.MantissaSinceVersion()
}

func (*PRICENULL) MantissaDeprecated() uint16 {
	return 0
}

func (*PRICENULL) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*PRICENULL) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*PRICENULL) ExponentNullValue() int8 {
	return math.MinInt8
}

func (*PRICENULL) ExponentSinceVersion() uint16 {
	return 0
}

func (p *PRICENULL) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ExponentSinceVersion()
}

func (*PRICENULL) ExponentDeprecated() uint16 {
	return 0
}
