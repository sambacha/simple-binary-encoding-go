// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type SecurityStatus30 struct {
	TransactTime          uint64
	SecurityGroup         [6]byte
	Asset                 [6]byte
	SecurityID            int32
	TradeDate             uint16
	MatchEventIndicator   MatchEventIndicator
	SecurityTradingStatus SecurityTradingStatusEnum
	HaltReason            HaltReasonEnum
	SecurityTradingEvent  SecurityTradingEventEnum
}

func (s *SecurityStatus30) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := s.RangeCheck(s.SbeSchemaVersion(), s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, s.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.SecurityGroup[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.Asset[:]); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, s.TradeDate); err != nil {
		return err
	}
	if err := s.MatchEventIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.SecurityTradingStatus.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.HaltReason.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.SecurityTradingEvent.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (s *SecurityStatus30) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !s.TransactTimeInActingVersion(actingVersion) {
		s.TransactTime = s.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.TransactTime); err != nil {
			return err
		}
	}
	if !s.SecurityGroupInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			s.SecurityGroup[idx] = s.SecurityGroupNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.SecurityGroup[:]); err != nil {
			return err
		}
	}
	if !s.AssetInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			s.Asset[idx] = s.AssetNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.Asset[:]); err != nil {
			return err
		}
	}
	if !s.SecurityIDInActingVersion(actingVersion) {
		s.SecurityID = s.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.SecurityID); err != nil {
			return err
		}
	}
	if !s.TradeDateInActingVersion(actingVersion) {
		s.TradeDate = s.TradeDateNullValue()
	} else {
		if err := _m.ReadUint16(_r, &s.TradeDate); err != nil {
			return err
		}
	}
	if s.MatchEventIndicatorInActingVersion(actingVersion) {
		if err := s.MatchEventIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.SecurityTradingStatusInActingVersion(actingVersion) {
		if err := s.SecurityTradingStatus.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.HaltReasonInActingVersion(actingVersion) {
		if err := s.HaltReason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.SecurityTradingEventInActingVersion(actingVersion) {
		if err := s.SecurityTradingEvent.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > s.SbeSchemaVersion() && blockLength > s.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-s.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := s.RangeCheck(actingVersion, s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (s *SecurityStatus30) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.TransactTimeInActingVersion(actingVersion) {
		if s.TransactTime < s.TransactTimeMinValue() || s.TransactTime > s.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on s.TransactTime (%v < %v > %v)", s.TransactTimeMinValue(), s.TransactTime, s.TransactTimeMaxValue())
		}
	}
	if s.SecurityGroupInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if s.SecurityGroup[idx] < s.SecurityGroupMinValue() || s.SecurityGroup[idx] > s.SecurityGroupMaxValue() {
				return fmt.Errorf("Range check failed on s.SecurityGroup[%d] (%v < %v > %v)", idx, s.SecurityGroupMinValue(), s.SecurityGroup[idx], s.SecurityGroupMaxValue())
			}
		}
	}
	if s.AssetInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if s.Asset[idx] < s.AssetMinValue() || s.Asset[idx] > s.AssetMaxValue() {
				return fmt.Errorf("Range check failed on s.Asset[%d] (%v < %v > %v)", idx, s.AssetMinValue(), s.Asset[idx], s.AssetMaxValue())
			}
		}
	}
	if s.SecurityIDInActingVersion(actingVersion) {
		if s.SecurityID != s.SecurityIDNullValue() && (s.SecurityID < s.SecurityIDMinValue() || s.SecurityID > s.SecurityIDMaxValue()) {
			return fmt.Errorf("Range check failed on s.SecurityID (%v < %v > %v)", s.SecurityIDMinValue(), s.SecurityID, s.SecurityIDMaxValue())
		}
	}
	if s.TradeDateInActingVersion(actingVersion) {
		if s.TradeDate != s.TradeDateNullValue() && (s.TradeDate < s.TradeDateMinValue() || s.TradeDate > s.TradeDateMaxValue()) {
			return fmt.Errorf("Range check failed on s.TradeDate (%v < %v > %v)", s.TradeDateMinValue(), s.TradeDate, s.TradeDateMaxValue())
		}
	}
	if err := s.SecurityTradingStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := s.HaltReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := s.SecurityTradingEvent.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func SecurityStatus30Init(s *SecurityStatus30) {
	s.SecurityID = 2147483647
	s.TradeDate = 65535
	return
}

func (*SecurityStatus30) SbeBlockLength() (blockLength uint16) {
	return 30
}

func (*SecurityStatus30) SbeTemplateId() (templateId uint16) {
	return 30
}

func (*SecurityStatus30) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*SecurityStatus30) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*SecurityStatus30) SbeSemanticType() (semanticType []byte) {
	return []byte("f")
}

func (*SecurityStatus30) TransactTimeId() uint16 {
	return 60
}

func (*SecurityStatus30) TransactTimeSinceVersion() uint16 {
	return 0
}

func (s *SecurityStatus30) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TransactTimeSinceVersion()
}

func (*SecurityStatus30) TransactTimeDeprecated() uint16 {
	return 0
}

func (*SecurityStatus30) TransactTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "UTCTimestamp"
	case 4:
		return "required"
	}
	return ""
}

func (*SecurityStatus30) TransactTimeMinValue() uint64 {
	return 0
}

func (*SecurityStatus30) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SecurityStatus30) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*SecurityStatus30) SecurityGroupId() uint16 {
	return 1151
}

func (*SecurityStatus30) SecurityGroupSinceVersion() uint16 {
	return 0
}

func (s *SecurityStatus30) SecurityGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityGroupSinceVersion()
}

func (*SecurityStatus30) SecurityGroupDeprecated() uint16 {
	return 0
}

func (*SecurityStatus30) SecurityGroupMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*SecurityStatus30) SecurityGroupMinValue() byte {
	return byte(32)
}

func (*SecurityStatus30) SecurityGroupMaxValue() byte {
	return byte(126)
}

func (*SecurityStatus30) SecurityGroupNullValue() byte {
	return 0
}

func (s *SecurityStatus30) SecurityGroupCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityStatus30) AssetId() uint16 {
	return 6937
}

func (*SecurityStatus30) AssetSinceVersion() uint16 {
	return 0
}

func (s *SecurityStatus30) AssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.AssetSinceVersion()
}

func (*SecurityStatus30) AssetDeprecated() uint16 {
	return 0
}

func (*SecurityStatus30) AssetMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*SecurityStatus30) AssetMinValue() byte {
	return byte(32)
}

func (*SecurityStatus30) AssetMaxValue() byte {
	return byte(126)
}

func (*SecurityStatus30) AssetNullValue() byte {
	return 0
}

func (s *SecurityStatus30) AssetCharacterEncoding() string {
	return "US-ASCII"
}

func (*SecurityStatus30) SecurityIDId() uint16 {
	return 48
}

func (*SecurityStatus30) SecurityIDSinceVersion() uint16 {
	return 0
}

func (s *SecurityStatus30) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityIDSinceVersion()
}

func (*SecurityStatus30) SecurityIDDeprecated() uint16 {
	return 0
}

func (*SecurityStatus30) SecurityIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "optional"
	}
	return ""
}

func (*SecurityStatus30) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SecurityStatus30) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*SecurityStatus30) SecurityIDNullValue() int32 {
	return 2147483647
}

func (*SecurityStatus30) TradeDateId() uint16 {
	return 75
}

func (*SecurityStatus30) TradeDateSinceVersion() uint16 {
	return 0
}

func (s *SecurityStatus30) TradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TradeDateSinceVersion()
}

func (*SecurityStatus30) TradeDateDeprecated() uint16 {
	return 0
}

func (*SecurityStatus30) TradeDateMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "LocalMktDate"
	case 4:
		return "optional"
	}
	return ""
}

func (*SecurityStatus30) TradeDateMinValue() uint16 {
	return 0
}

func (*SecurityStatus30) TradeDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*SecurityStatus30) TradeDateNullValue() uint16 {
	return 65535
}

func (*SecurityStatus30) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*SecurityStatus30) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (s *SecurityStatus30) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MatchEventIndicatorSinceVersion()
}

func (*SecurityStatus30) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*SecurityStatus30) MatchEventIndicatorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "MultipleCharValue"
	case 4:
		return "required"
	}
	return ""
}

func (*SecurityStatus30) SecurityTradingStatusId() uint16 {
	return 326
}

func (*SecurityStatus30) SecurityTradingStatusSinceVersion() uint16 {
	return 0
}

func (s *SecurityStatus30) SecurityTradingStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityTradingStatusSinceVersion()
}

func (*SecurityStatus30) SecurityTradingStatusDeprecated() uint16 {
	return 0
}

func (*SecurityStatus30) SecurityTradingStatusMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*SecurityStatus30) HaltReasonId() uint16 {
	return 327
}

func (*SecurityStatus30) HaltReasonSinceVersion() uint16 {
	return 0
}

func (s *SecurityStatus30) HaltReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.HaltReasonSinceVersion()
}

func (*SecurityStatus30) HaltReasonDeprecated() uint16 {
	return 0
}

func (*SecurityStatus30) HaltReasonMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*SecurityStatus30) SecurityTradingEventId() uint16 {
	return 1174
}

func (*SecurityStatus30) SecurityTradingEventSinceVersion() uint16 {
	return 0
}

func (s *SecurityStatus30) SecurityTradingEventInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityTradingEventSinceVersion()
}

func (*SecurityStatus30) SecurityTradingEventDeprecated() uint16 {
	return 0
}

func (*SecurityStatus30) SecurityTradingEventMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}
