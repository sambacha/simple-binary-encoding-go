// Generated SBE (Simple Binary Encoding) message codec

package composite_elements

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type FuturesPrice struct {
	Mantissa     int64
	Exponent     int8
	IsSettlement BooleanEnumEnum
}

func (f *FuturesPrice) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, f.Mantissa); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, f.Exponent); err != nil {
		return err
	}

	for i := 0; i < 1; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	if err := f.IsSettlement.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (f *FuturesPrice) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !f.MantissaInActingVersion(actingVersion) {
		f.Mantissa = f.MantissaNullValue()
	} else {
		if err := _m.ReadInt64(_r, &f.Mantissa); err != nil {
			return err
		}
	}
	if !f.ExponentInActingVersion(actingVersion) {
		f.Exponent = f.ExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &f.Exponent); err != nil {
			return err
		}
	}
	io.CopyN(ioutil.Discard, _r, 1)
	if f.IsSettlementInActingVersion(actingVersion) {
		if err := f.IsSettlement.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	return nil
}

func (f *FuturesPrice) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if f.MantissaInActingVersion(actingVersion) {
		if f.Mantissa < f.MantissaMinValue() || f.Mantissa > f.MantissaMaxValue() {
			return fmt.Errorf("Range check failed on f.Mantissa (%v < %v > %v)", f.MantissaMinValue(), f.Mantissa, f.MantissaMaxValue())
		}
	}
	if f.ExponentInActingVersion(actingVersion) {
		if f.Exponent < f.ExponentMinValue() || f.Exponent > f.ExponentMaxValue() {
			return fmt.Errorf("Range check failed on f.Exponent (%v < %v > %v)", f.ExponentMinValue(), f.Exponent, f.ExponentMaxValue())
		}
	}
	return nil
}

func FuturesPriceInit(f *FuturesPrice) {
	return
}

func (*FuturesPrice) EncodedLength() int64 {
	return 11
}

func (*FuturesPrice) MantissaMinValue() int64 {
	return math.MinInt64 + 1
}

func (*FuturesPrice) MantissaMaxValue() int64 {
	return math.MaxInt64
}

func (*FuturesPrice) MantissaNullValue() int64 {
	return math.MinInt64
}

func (*FuturesPrice) MantissaSinceVersion() uint16 {
	return 0
}

func (f *FuturesPrice) MantissaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= f.MantissaSinceVersion()
}

func (*FuturesPrice) MantissaDeprecated() uint16 {
	return 0
}

func (*FuturesPrice) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*FuturesPrice) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*FuturesPrice) ExponentNullValue() int8 {
	return math.MinInt8
}

func (*FuturesPrice) ExponentSinceVersion() uint16 {
	return 0
}

func (f *FuturesPrice) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= f.ExponentSinceVersion()
}

func (*FuturesPrice) ExponentDeprecated() uint16 {
	return 0
}

func (*FuturesPrice) IsSettlementSinceVersion() uint16 {
	return 0
}

func (f *FuturesPrice) IsSettlementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= f.IsSettlementSinceVersion()
}

func (*FuturesPrice) IsSettlementDeprecated() uint16 {
	return 0
}
