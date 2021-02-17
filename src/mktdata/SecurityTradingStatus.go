// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type SecurityTradingStatusEnum uint8
type SecurityTradingStatusValues struct {
	TradingHalt            SecurityTradingStatusEnum
	Close                  SecurityTradingStatusEnum
	NewPriceIndication     SecurityTradingStatusEnum
	ReadyToTrade           SecurityTradingStatusEnum
	NotAvailableForTrading SecurityTradingStatusEnum
	UnknownorInvalid       SecurityTradingStatusEnum
	PreOpen                SecurityTradingStatusEnum
	PreCross               SecurityTradingStatusEnum
	Cross                  SecurityTradingStatusEnum
	PostClose              SecurityTradingStatusEnum
	NoChange               SecurityTradingStatusEnum
	NullValue              SecurityTradingStatusEnum
}

var SecurityTradingStatus = SecurityTradingStatusValues{2, 4, 15, 17, 18, 20, 21, 24, 25, 26, 103, 255}

func (s SecurityTradingStatusEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(s)); err != nil {
		return err
	}
	return nil
}

func (s *SecurityTradingStatusEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(s)); err != nil {
		return err
	}
	return nil
}

func (s SecurityTradingStatusEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SecurityTradingStatus)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SecurityTradingStatus, unknown enumeration value %d", s)
}

func (*SecurityTradingStatusEnum) EncodedLength() int64 {
	return 1
}

func (*SecurityTradingStatusEnum) TradingHaltSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingStatusEnum) TradingHaltInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TradingHaltSinceVersion()
}

func (*SecurityTradingStatusEnum) TradingHaltDeprecated() uint16 {
	return 0
}

func (*SecurityTradingStatusEnum) CloseSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingStatusEnum) CloseInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.CloseSinceVersion()
}

func (*SecurityTradingStatusEnum) CloseDeprecated() uint16 {
	return 0
}

func (*SecurityTradingStatusEnum) NewPriceIndicationSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingStatusEnum) NewPriceIndicationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NewPriceIndicationSinceVersion()
}

func (*SecurityTradingStatusEnum) NewPriceIndicationDeprecated() uint16 {
	return 0
}

func (*SecurityTradingStatusEnum) ReadyToTradeSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingStatusEnum) ReadyToTradeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ReadyToTradeSinceVersion()
}

func (*SecurityTradingStatusEnum) ReadyToTradeDeprecated() uint16 {
	return 0
}

func (*SecurityTradingStatusEnum) NotAvailableForTradingSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingStatusEnum) NotAvailableForTradingInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NotAvailableForTradingSinceVersion()
}

func (*SecurityTradingStatusEnum) NotAvailableForTradingDeprecated() uint16 {
	return 0
}

func (*SecurityTradingStatusEnum) UnknownorInvalidSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingStatusEnum) UnknownorInvalidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.UnknownorInvalidSinceVersion()
}

func (*SecurityTradingStatusEnum) UnknownorInvalidDeprecated() uint16 {
	return 0
}

func (*SecurityTradingStatusEnum) PreOpenSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingStatusEnum) PreOpenInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.PreOpenSinceVersion()
}

func (*SecurityTradingStatusEnum) PreOpenDeprecated() uint16 {
	return 0
}

func (*SecurityTradingStatusEnum) PreCrossSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingStatusEnum) PreCrossInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.PreCrossSinceVersion()
}

func (*SecurityTradingStatusEnum) PreCrossDeprecated() uint16 {
	return 0
}

func (*SecurityTradingStatusEnum) CrossSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingStatusEnum) CrossInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.CrossSinceVersion()
}

func (*SecurityTradingStatusEnum) CrossDeprecated() uint16 {
	return 0
}

func (*SecurityTradingStatusEnum) PostCloseSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingStatusEnum) PostCloseInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.PostCloseSinceVersion()
}

func (*SecurityTradingStatusEnum) PostCloseDeprecated() uint16 {
	return 0
}

func (*SecurityTradingStatusEnum) NoChangeSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingStatusEnum) NoChangeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NoChangeSinceVersion()
}

func (*SecurityTradingStatusEnum) NoChangeDeprecated() uint16 {
	return 0
}
