// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type MDEntryTypeDailyStatisticsEnum byte
type MDEntryTypeDailyStatisticsValues struct {
	SettlementPrice MDEntryTypeDailyStatisticsEnum
	ClearedVolume   MDEntryTypeDailyStatisticsEnum
	OpenInterest    MDEntryTypeDailyStatisticsEnum
	FixingPrice     MDEntryTypeDailyStatisticsEnum
	NullValue       MDEntryTypeDailyStatisticsEnum
}

var MDEntryTypeDailyStatistics = MDEntryTypeDailyStatisticsValues{54, 66, 67, 87, 0}

func (m MDEntryTypeDailyStatisticsEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(m)); err != nil {
		return err
	}
	return nil
}

func (m *MDEntryTypeDailyStatisticsEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(m)); err != nil {
		return err
	}
	return nil
}

func (m MDEntryTypeDailyStatisticsEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MDEntryTypeDailyStatistics)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MDEntryTypeDailyStatistics, unknown enumeration value %d", m)
}

func (*MDEntryTypeDailyStatisticsEnum) EncodedLength() int64 {
	return 1
}

func (*MDEntryTypeDailyStatisticsEnum) SettlementPriceSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeDailyStatisticsEnum) SettlementPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SettlementPriceSinceVersion()
}

func (*MDEntryTypeDailyStatisticsEnum) SettlementPriceDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeDailyStatisticsEnum) ClearedVolumeSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeDailyStatisticsEnum) ClearedVolumeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ClearedVolumeSinceVersion()
}

func (*MDEntryTypeDailyStatisticsEnum) ClearedVolumeDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeDailyStatisticsEnum) OpenInterestSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeDailyStatisticsEnum) OpenInterestInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OpenInterestSinceVersion()
}

func (*MDEntryTypeDailyStatisticsEnum) OpenInterestDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeDailyStatisticsEnum) FixingPriceSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeDailyStatisticsEnum) FixingPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.FixingPriceSinceVersion()
}

func (*MDEntryTypeDailyStatisticsEnum) FixingPriceDeprecated() uint16 {
	return 0
}
