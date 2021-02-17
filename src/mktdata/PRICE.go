// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"math"
)

type PRICE struct {
	Mantissa int64
	Exponent int8
}

func (p *PRICE) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, p.Mantissa); err != nil {
		return err
	}
	return nil
}

func (p *PRICE) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
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

func (p *PRICE) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.MantissaInActingVersion(actingVersion) {
		if p.Mantissa < p.MantissaMinValue() || p.Mantissa > p.MantissaMaxValue() {
			return fmt.Errorf("Range check failed on p.Mantissa (%v < %v > %v)", p.MantissaMinValue(), p.Mantissa, p.MantissaMaxValue())
		}
	}
	return nil
}

func PRICEInit(p *PRICE) {
	p.Exponent = -7
	return
}

func (*PRICE) EncodedLength() int64 {
	return 8
}

func (*PRICE) MantissaMinValue() int64 {
	return math.MinInt64 + 1
}

func (*PRICE) MantissaMaxValue() int64 {
	return math.MaxInt64
}

func (*PRICE) MantissaNullValue() int64 {
	return math.MinInt64
}

func (*PRICE) MantissaSinceVersion() uint16 {
	return 0
}

func (p *PRICE) MantissaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.MantissaSinceVersion()
}

func (*PRICE) MantissaDeprecated() uint16 {
	return 0
}

func (*PRICE) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*PRICE) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*PRICE) ExponentNullValue() int8 {
	return math.MinInt8
}

func (*PRICE) ExponentSinceVersion() uint16 {
	return 0
}

func (p *PRICE) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ExponentSinceVersion()
}

func (*PRICE) ExponentDeprecated() uint16 {
	return 0
}
