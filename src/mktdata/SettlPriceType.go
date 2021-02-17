// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"io"
)

type SettlPriceType [8]bool
type SettlPriceTypeChoiceValue uint8
type SettlPriceTypeChoiceValues struct {
	FinalDaily   SettlPriceTypeChoiceValue
	Actual       SettlPriceTypeChoiceValue
	Rounded      SettlPriceTypeChoiceValue
	Intraday     SettlPriceTypeChoiceValue
	ReservedBits SettlPriceTypeChoiceValue
	NullValue    SettlPriceTypeChoiceValue
}

var SettlPriceTypeChoice = SettlPriceTypeChoiceValues{0, 1, 2, 3, 4, 7}

func (s *SettlPriceType) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint8 = 0
	for k, v := range s {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint8(_w, wireval)
}

func (s *SettlPriceType) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	var wireval uint8

	if err := _m.ReadUint8(_r, &wireval); err != nil {
		return err
	}

	var idx uint
	for idx = 0; idx < 8; idx++ {
		s[idx] = (wireval & (1 << idx)) > 0
	}
	return nil
}

func (SettlPriceType) EncodedLength() int64 {
	return 1
}

func (*SettlPriceType) FinalDailySinceVersion() uint16 {
	return 0
}

func (s *SettlPriceType) FinalDailyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.FinalDailySinceVersion()
}

func (*SettlPriceType) FinalDailyDeprecated() uint16 {
	return 0
}

func (*SettlPriceType) ActualSinceVersion() uint16 {
	return 0
}

func (s *SettlPriceType) ActualInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ActualSinceVersion()
}

func (*SettlPriceType) ActualDeprecated() uint16 {
	return 0
}

func (*SettlPriceType) RoundedSinceVersion() uint16 {
	return 0
}

func (s *SettlPriceType) RoundedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.RoundedSinceVersion()
}

func (*SettlPriceType) RoundedDeprecated() uint16 {
	return 0
}

func (*SettlPriceType) IntradaySinceVersion() uint16 {
	return 4
}

func (s *SettlPriceType) IntradayInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.IntradaySinceVersion()
}

func (*SettlPriceType) IntradayDeprecated() uint16 {
	return 0
}

func (*SettlPriceType) ReservedBitsSinceVersion() uint16 {
	return 0
}

func (s *SettlPriceType) ReservedBitsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ReservedBitsSinceVersion()
}

func (*SettlPriceType) ReservedBitsDeprecated() uint16 {
	return 0
}

func (*SettlPriceType) NullValueSinceVersion() uint16 {
	return 0
}

func (s *SettlPriceType) NullValueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NullValueSinceVersion()
}

func (*SettlPriceType) NullValueDeprecated() uint16 {
	return 0
}
