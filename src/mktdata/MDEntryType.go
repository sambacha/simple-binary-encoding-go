// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type MDEntryTypeEnum byte
type MDEntryTypeValues struct {
	Bid                                  MDEntryTypeEnum
	Offer                                MDEntryTypeEnum
	Trade                                MDEntryTypeEnum
	OpenPrice                            MDEntryTypeEnum
	SettlementPrice                      MDEntryTypeEnum
	TradingSessionHighPrice              MDEntryTypeEnum
	TradingSessionLowPrice               MDEntryTypeEnum
	ClearedVolume                        MDEntryTypeEnum
	OpenInterest                         MDEntryTypeEnum
	ImpliedBid                           MDEntryTypeEnum
	ImpliedOffer                         MDEntryTypeEnum
	BookReset                            MDEntryTypeEnum
	SessionHighBid                       MDEntryTypeEnum
	SessionLowOffer                      MDEntryTypeEnum
	FixingPrice                          MDEntryTypeEnum
	ElectronicVolume                     MDEntryTypeEnum
	ThresholdLimitsandPriceBandVariation MDEntryTypeEnum
	NullValue                            MDEntryTypeEnum
}

var MDEntryType = MDEntryTypeValues{48, 49, 50, 52, 54, 55, 56, 66, 67, 69, 70, 74, 78, 79, 87, 101, 103, 0}

func (m MDEntryTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(m)); err != nil {
		return err
	}
	return nil
}

func (m *MDEntryTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(m)); err != nil {
		return err
	}
	return nil
}

func (m MDEntryTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MDEntryType)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MDEntryType, unknown enumeration value %d", m)
}

func (*MDEntryTypeEnum) EncodedLength() int64 {
	return 1
}

func (*MDEntryTypeEnum) BidSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) BidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BidSinceVersion()
}

func (*MDEntryTypeEnum) BidDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) OfferSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) OfferInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OfferSinceVersion()
}

func (*MDEntryTypeEnum) OfferDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) TradeSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) TradeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradeSinceVersion()
}

func (*MDEntryTypeEnum) TradeDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) OpenPriceSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) OpenPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OpenPriceSinceVersion()
}

func (*MDEntryTypeEnum) OpenPriceDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) SettlementPriceSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) SettlementPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SettlementPriceSinceVersion()
}

func (*MDEntryTypeEnum) SettlementPriceDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) TradingSessionHighPriceSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) TradingSessionHighPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradingSessionHighPriceSinceVersion()
}

func (*MDEntryTypeEnum) TradingSessionHighPriceDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) TradingSessionLowPriceSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) TradingSessionLowPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradingSessionLowPriceSinceVersion()
}

func (*MDEntryTypeEnum) TradingSessionLowPriceDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) ClearedVolumeSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) ClearedVolumeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ClearedVolumeSinceVersion()
}

func (*MDEntryTypeEnum) ClearedVolumeDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) OpenInterestSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) OpenInterestInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OpenInterestSinceVersion()
}

func (*MDEntryTypeEnum) OpenInterestDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) ImpliedBidSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) ImpliedBidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ImpliedBidSinceVersion()
}

func (*MDEntryTypeEnum) ImpliedBidDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) ImpliedOfferSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) ImpliedOfferInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ImpliedOfferSinceVersion()
}

func (*MDEntryTypeEnum) ImpliedOfferDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) BookResetSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) BookResetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BookResetSinceVersion()
}

func (*MDEntryTypeEnum) BookResetDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) SessionHighBidSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) SessionHighBidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SessionHighBidSinceVersion()
}

func (*MDEntryTypeEnum) SessionHighBidDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) SessionLowOfferSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) SessionLowOfferInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SessionLowOfferSinceVersion()
}

func (*MDEntryTypeEnum) SessionLowOfferDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) FixingPriceSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) FixingPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.FixingPriceSinceVersion()
}

func (*MDEntryTypeEnum) FixingPriceDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) ElectronicVolumeSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) ElectronicVolumeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ElectronicVolumeSinceVersion()
}

func (*MDEntryTypeEnum) ElectronicVolumeDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeEnum) ThresholdLimitsandPriceBandVariationSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeEnum) ThresholdLimitsandPriceBandVariationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ThresholdLimitsandPriceBandVariationSinceVersion()
}

func (*MDEntryTypeEnum) ThresholdLimitsandPriceBandVariationDeprecated() uint16 {
	return 0
}
