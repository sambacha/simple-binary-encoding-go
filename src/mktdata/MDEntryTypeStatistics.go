// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type MDEntryTypeStatisticsEnum byte
type MDEntryTypeStatisticsValues struct {
	OpenPrice   MDEntryTypeStatisticsEnum
	HighTrade   MDEntryTypeStatisticsEnum
	LowTrade    MDEntryTypeStatisticsEnum
	HighestBid  MDEntryTypeStatisticsEnum
	LowestOffer MDEntryTypeStatisticsEnum
	NullValue   MDEntryTypeStatisticsEnum
}

var MDEntryTypeStatistics = MDEntryTypeStatisticsValues{52, 55, 56, 78, 79, 0}

func (m MDEntryTypeStatisticsEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(m)); err != nil {
		return err
	}
	return nil
}

func (m *MDEntryTypeStatisticsEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(m)); err != nil {
		return err
	}
	return nil
}

func (m MDEntryTypeStatisticsEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MDEntryTypeStatistics)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MDEntryTypeStatistics, unknown enumeration value %d", m)
}

func (*MDEntryTypeStatisticsEnum) EncodedLength() int64 {
	return 1
}

func (*MDEntryTypeStatisticsEnum) OpenPriceSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeStatisticsEnum) OpenPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OpenPriceSinceVersion()
}

func (*MDEntryTypeStatisticsEnum) OpenPriceDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeStatisticsEnum) HighTradeSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeStatisticsEnum) HighTradeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.HighTradeSinceVersion()
}

func (*MDEntryTypeStatisticsEnum) HighTradeDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeStatisticsEnum) LowTradeSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeStatisticsEnum) LowTradeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LowTradeSinceVersion()
}

func (*MDEntryTypeStatisticsEnum) LowTradeDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeStatisticsEnum) HighestBidSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeStatisticsEnum) HighestBidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.HighestBidSinceVersion()
}

func (*MDEntryTypeStatisticsEnum) HighestBidDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeStatisticsEnum) LowestOfferSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeStatisticsEnum) LowestOfferInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LowestOfferSinceVersion()
}

func (*MDEntryTypeStatisticsEnum) LowestOfferDeprecated() uint16 {
	return 0
}
