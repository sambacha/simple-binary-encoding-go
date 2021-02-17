// Generated SBE (Simple Binary Encoding) message codec

package composite

import (
	"fmt"
	"io"
	"math"
)

type Point struct {
	Name      [5]byte
	D         float64
	I         int32
	U         [2]uint8
	Truthval1 BooleanEnumEnum
	Truthval2 BooleanEnumEnum
}

func (p *Point) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, p.Name[:]); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, p.D); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, p.I); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, p.U[:]); err != nil {
		return err
	}
	if err := p.Truthval1.Encode(_m, _w); err != nil {
		return err
	}
	if err := p.Truthval2.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (p *Point) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !p.NameInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			p.Name[idx] = p.NameNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, p.Name[:]); err != nil {
			return err
		}
	}
	if !p.DInActingVersion(actingVersion) {
		p.D = p.DNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &p.D); err != nil {
			return err
		}
	}
	if !p.IInActingVersion(actingVersion) {
		p.I = p.INullValue()
	} else {
		if err := _m.ReadInt32(_r, &p.I); err != nil {
			return err
		}
	}
	if !p.UInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			p.U[idx] = p.UNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, p.U[:]); err != nil {
			return err
		}
	}
	if p.Truthval1InActingVersion(actingVersion) {
		if err := p.Truthval1.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if p.Truthval2InActingVersion(actingVersion) {
		if err := p.Truthval2.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	return nil
}

func (p *Point) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.NameInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if p.Name[idx] < p.NameMinValue() || p.Name[idx] > p.NameMaxValue() {
				return fmt.Errorf("Range check failed on p.Name[%d] (%v < %v > %v)", idx, p.NameMinValue(), p.Name[idx], p.NameMaxValue())
			}
		}
	}
	if p.DInActingVersion(actingVersion) {
		if p.D < p.DMinValue() || p.D > p.DMaxValue() {
			return fmt.Errorf("Range check failed on p.D (%v < %v > %v)", p.DMinValue(), p.D, p.DMaxValue())
		}
	}
	if p.IInActingVersion(actingVersion) {
		if p.I < p.IMinValue() || p.I > p.IMaxValue() {
			return fmt.Errorf("Range check failed on p.I (%v < %v > %v)", p.IMinValue(), p.I, p.IMaxValue())
		}
	}
	if p.UInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if p.U[idx] < p.UMinValue() || p.U[idx] > p.UMaxValue() {
				return fmt.Errorf("Range check failed on p.U[%d] (%v < %v > %v)", idx, p.UMinValue(), p.U[idx], p.UMaxValue())
			}
		}
	}
	return nil
}

func PointInit(p *Point) {
	return
}

func (*Point) EncodedLength() int64 {
	return 21
}

func (*Point) NameMinValue() byte {
	return byte(32)
}

func (*Point) NameMaxValue() byte {
	return byte(126)
}

func (*Point) NameNullValue() byte {
	return 0
}

func (p *Point) NameCharacterEncoding() string {
	return "US-ASCII"
}

func (*Point) NameSinceVersion() uint16 {
	return 0
}

func (p *Point) NameInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.NameSinceVersion()
}

func (*Point) NameDeprecated() uint16 {
	return 0
}

func (*Point) DMinValue() float64 {
	return -math.MaxFloat64
}

func (*Point) DMaxValue() float64 {
	return math.MaxFloat64
}

func (*Point) DNullValue() float64 {
	return math.NaN()
}

func (*Point) DSinceVersion() uint16 {
	return 0
}

func (p *Point) DInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.DSinceVersion()
}

func (*Point) DDeprecated() uint16 {
	return 0
}

func (*Point) IMinValue() int32 {
	return math.MinInt32 + 1
}

func (*Point) IMaxValue() int32 {
	return math.MaxInt32
}

func (*Point) INullValue() int32 {
	return math.MinInt32
}

func (*Point) ISinceVersion() uint16 {
	return 0
}

func (p *Point) IInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ISinceVersion()
}

func (*Point) IDeprecated() uint16 {
	return 0
}

func (*Point) UMinValue() uint8 {
	return 0
}

func (*Point) UMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*Point) UNullValue() uint8 {
	return math.MaxUint8
}

func (*Point) USinceVersion() uint16 {
	return 0
}

func (p *Point) UInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.USinceVersion()
}

func (*Point) UDeprecated() uint16 {
	return 0
}

func (*Point) Truthval1SinceVersion() uint16 {
	return 0
}

func (p *Point) Truthval1InActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.Truthval1SinceVersion()
}

func (*Point) Truthval1Deprecated() uint16 {
	return 0
}

func (*Point) Truthval2SinceVersion() uint16 {
	return 0
}

func (p *Point) Truthval2InActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.Truthval2SinceVersion()
}

func (*Point) Truthval2Deprecated() uint16 {
	return 0
}
