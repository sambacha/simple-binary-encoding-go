// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type SnapshotFullRefresh52 struct {
	LastMsgSeqNumProcessed  uint32
	TotNumReports           uint32
	SecurityID              int32
	RptSeq                  uint32
	TransactTime            uint64
	LastUpdateTime          uint64
	TradeDate               uint16
	MDSecurityTradingStatus SecurityTradingStatusEnum
	HighLimitPrice          PRICENULL9
	LowLimitPrice           PRICENULL9
	MaxPriceVariation       PRICENULL9
	NoMDEntries             []SnapshotFullRefresh52NoMDEntries
}
type SnapshotFullRefresh52NoMDEntries struct {
	MDEntryPx            PRICENULL9
	MDEntrySize          int32
	NumberOfOrders       int32
	MDPriceLevel         int8
	TradingReferenceDate uint16
	OpenCloseSettlFlag   OpenCloseSettlFlagEnum
	SettlPriceType       SettlPriceType
	MDEntryType          MDEntryTypeEnum
}

func (s *SnapshotFullRefresh52) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := s.RangeCheck(s.SbeSchemaVersion(), s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, s.LastMsgSeqNumProcessed); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, s.TotNumReports); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, s.RptSeq); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.LastUpdateTime); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, s.TradeDate); err != nil {
		return err
	}
	if err := s.MDSecurityTradingStatus.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.HighLimitPrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.LowLimitPrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.MaxPriceVariation.Encode(_m, _w); err != nil {
		return err
	}
	var NoMDEntriesBlockLength uint16 = 22
	if err := _m.WriteUint16(_w, NoMDEntriesBlockLength); err != nil {
		return err
	}
	var NoMDEntriesNumInGroup uint8 = uint8(len(s.NoMDEntries))
	if err := _m.WriteUint8(_w, NoMDEntriesNumInGroup); err != nil {
		return err
	}
	for _, prop := range s.NoMDEntries {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (s *SnapshotFullRefresh52) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !s.LastMsgSeqNumProcessedInActingVersion(actingVersion) {
		s.LastMsgSeqNumProcessed = s.LastMsgSeqNumProcessedNullValue()
	} else {
		if err := _m.ReadUint32(_r, &s.LastMsgSeqNumProcessed); err != nil {
			return err
		}
	}
	if !s.TotNumReportsInActingVersion(actingVersion) {
		s.TotNumReports = s.TotNumReportsNullValue()
	} else {
		if err := _m.ReadUint32(_r, &s.TotNumReports); err != nil {
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
	if !s.RptSeqInActingVersion(actingVersion) {
		s.RptSeq = s.RptSeqNullValue()
	} else {
		if err := _m.ReadUint32(_r, &s.RptSeq); err != nil {
			return err
		}
	}
	if !s.TransactTimeInActingVersion(actingVersion) {
		s.TransactTime = s.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.TransactTime); err != nil {
			return err
		}
	}
	if !s.LastUpdateTimeInActingVersion(actingVersion) {
		s.LastUpdateTime = s.LastUpdateTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.LastUpdateTime); err != nil {
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
	if s.MDSecurityTradingStatusInActingVersion(actingVersion) {
		if err := s.MDSecurityTradingStatus.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.HighLimitPriceInActingVersion(actingVersion) {
		if err := s.HighLimitPrice.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.LowLimitPriceInActingVersion(actingVersion) {
		if err := s.LowLimitPrice.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.MaxPriceVariationInActingVersion(actingVersion) {
		if err := s.MaxPriceVariation.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > s.SbeSchemaVersion() && blockLength > s.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-s.SbeBlockLength()))
	}

	if s.NoMDEntriesInActingVersion(actingVersion) {
		var NoMDEntriesBlockLength uint16
		if err := _m.ReadUint16(_r, &NoMDEntriesBlockLength); err != nil {
			return err
		}
		var NoMDEntriesNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoMDEntriesNumInGroup); err != nil {
			return err
		}
		if cap(s.NoMDEntries) < int(NoMDEntriesNumInGroup) {
			s.NoMDEntries = make([]SnapshotFullRefresh52NoMDEntries, NoMDEntriesNumInGroup)
		}
		s.NoMDEntries = s.NoMDEntries[:NoMDEntriesNumInGroup]
		for i, _ := range s.NoMDEntries {
			if err := s.NoMDEntries[i].Decode(_m, _r, actingVersion, uint(NoMDEntriesBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := s.RangeCheck(actingVersion, s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (s *SnapshotFullRefresh52) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.LastMsgSeqNumProcessedInActingVersion(actingVersion) {
		if s.LastMsgSeqNumProcessed < s.LastMsgSeqNumProcessedMinValue() || s.LastMsgSeqNumProcessed > s.LastMsgSeqNumProcessedMaxValue() {
			return fmt.Errorf("Range check failed on s.LastMsgSeqNumProcessed (%v < %v > %v)", s.LastMsgSeqNumProcessedMinValue(), s.LastMsgSeqNumProcessed, s.LastMsgSeqNumProcessedMaxValue())
		}
	}
	if s.TotNumReportsInActingVersion(actingVersion) {
		if s.TotNumReports < s.TotNumReportsMinValue() || s.TotNumReports > s.TotNumReportsMaxValue() {
			return fmt.Errorf("Range check failed on s.TotNumReports (%v < %v > %v)", s.TotNumReportsMinValue(), s.TotNumReports, s.TotNumReportsMaxValue())
		}
	}
	if s.SecurityIDInActingVersion(actingVersion) {
		if s.SecurityID < s.SecurityIDMinValue() || s.SecurityID > s.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on s.SecurityID (%v < %v > %v)", s.SecurityIDMinValue(), s.SecurityID, s.SecurityIDMaxValue())
		}
	}
	if s.RptSeqInActingVersion(actingVersion) {
		if s.RptSeq < s.RptSeqMinValue() || s.RptSeq > s.RptSeqMaxValue() {
			return fmt.Errorf("Range check failed on s.RptSeq (%v < %v > %v)", s.RptSeqMinValue(), s.RptSeq, s.RptSeqMaxValue())
		}
	}
	if s.TransactTimeInActingVersion(actingVersion) {
		if s.TransactTime < s.TransactTimeMinValue() || s.TransactTime > s.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on s.TransactTime (%v < %v > %v)", s.TransactTimeMinValue(), s.TransactTime, s.TransactTimeMaxValue())
		}
	}
	if s.LastUpdateTimeInActingVersion(actingVersion) {
		if s.LastUpdateTime < s.LastUpdateTimeMinValue() || s.LastUpdateTime > s.LastUpdateTimeMaxValue() {
			return fmt.Errorf("Range check failed on s.LastUpdateTime (%v < %v > %v)", s.LastUpdateTimeMinValue(), s.LastUpdateTime, s.LastUpdateTimeMaxValue())
		}
	}
	if s.TradeDateInActingVersion(actingVersion) {
		if s.TradeDate != s.TradeDateNullValue() && (s.TradeDate < s.TradeDateMinValue() || s.TradeDate > s.TradeDateMaxValue()) {
			return fmt.Errorf("Range check failed on s.TradeDate (%v < %v > %v)", s.TradeDateMinValue(), s.TradeDate, s.TradeDateMaxValue())
		}
	}
	if err := s.MDSecurityTradingStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	for _, prop := range s.NoMDEntries {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func SnapshotFullRefresh52Init(s *SnapshotFullRefresh52) {
	s.TradeDate = 65535
	return
}

func (s *SnapshotFullRefresh52NoMDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := s.MDEntryPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.MDEntrySize); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.NumberOfOrders); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, s.MDPriceLevel); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, s.TradingReferenceDate); err != nil {
		return err
	}
	if err := s.OpenCloseSettlFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.SettlPriceType.Encode(_m, _w); err != nil {
		return err
	}
	if err := s.MDEntryType.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (s *SnapshotFullRefresh52NoMDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if s.MDEntryPxInActingVersion(actingVersion) {
		if err := s.MDEntryPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !s.MDEntrySizeInActingVersion(actingVersion) {
		s.MDEntrySize = s.MDEntrySizeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.MDEntrySize); err != nil {
			return err
		}
	}
	if !s.NumberOfOrdersInActingVersion(actingVersion) {
		s.NumberOfOrders = s.NumberOfOrdersNullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.NumberOfOrders); err != nil {
			return err
		}
	}
	if !s.MDPriceLevelInActingVersion(actingVersion) {
		s.MDPriceLevel = s.MDPriceLevelNullValue()
	} else {
		if err := _m.ReadInt8(_r, &s.MDPriceLevel); err != nil {
			return err
		}
	}
	if !s.TradingReferenceDateInActingVersion(actingVersion) {
		s.TradingReferenceDate = s.TradingReferenceDateNullValue()
	} else {
		if err := _m.ReadUint16(_r, &s.TradingReferenceDate); err != nil {
			return err
		}
	}
	if s.OpenCloseSettlFlagInActingVersion(actingVersion) {
		if err := s.OpenCloseSettlFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.SettlPriceTypeInActingVersion(actingVersion) {
		if err := s.SettlPriceType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if s.MDEntryTypeInActingVersion(actingVersion) {
		if err := s.MDEntryType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > s.SbeSchemaVersion() && blockLength > s.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-s.SbeBlockLength()))
	}
	return nil
}

func (s *SnapshotFullRefresh52NoMDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.MDEntrySizeInActingVersion(actingVersion) {
		if s.MDEntrySize != s.MDEntrySizeNullValue() && (s.MDEntrySize < s.MDEntrySizeMinValue() || s.MDEntrySize > s.MDEntrySizeMaxValue()) {
			return fmt.Errorf("Range check failed on s.MDEntrySize (%v < %v > %v)", s.MDEntrySizeMinValue(), s.MDEntrySize, s.MDEntrySizeMaxValue())
		}
	}
	if s.NumberOfOrdersInActingVersion(actingVersion) {
		if s.NumberOfOrders != s.NumberOfOrdersNullValue() && (s.NumberOfOrders < s.NumberOfOrdersMinValue() || s.NumberOfOrders > s.NumberOfOrdersMaxValue()) {
			return fmt.Errorf("Range check failed on s.NumberOfOrders (%v < %v > %v)", s.NumberOfOrdersMinValue(), s.NumberOfOrders, s.NumberOfOrdersMaxValue())
		}
	}
	if s.MDPriceLevelInActingVersion(actingVersion) {
		if s.MDPriceLevel != s.MDPriceLevelNullValue() && (s.MDPriceLevel < s.MDPriceLevelMinValue() || s.MDPriceLevel > s.MDPriceLevelMaxValue()) {
			return fmt.Errorf("Range check failed on s.MDPriceLevel (%v < %v > %v)", s.MDPriceLevelMinValue(), s.MDPriceLevel, s.MDPriceLevelMaxValue())
		}
	}
	if s.TradingReferenceDateInActingVersion(actingVersion) {
		if s.TradingReferenceDate != s.TradingReferenceDateNullValue() && (s.TradingReferenceDate < s.TradingReferenceDateMinValue() || s.TradingReferenceDate > s.TradingReferenceDateMaxValue()) {
			return fmt.Errorf("Range check failed on s.TradingReferenceDate (%v < %v > %v)", s.TradingReferenceDateMinValue(), s.TradingReferenceDate, s.TradingReferenceDateMaxValue())
		}
	}
	if err := s.OpenCloseSettlFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := s.MDEntryType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func SnapshotFullRefresh52NoMDEntriesInit(s *SnapshotFullRefresh52NoMDEntries) {
	s.MDEntrySize = 2147483647
	s.NumberOfOrders = 2147483647
	s.MDPriceLevel = 127
	s.TradingReferenceDate = 65535
	return
}

func (*SnapshotFullRefresh52) SbeBlockLength() (blockLength uint16) {
	return 59
}

func (*SnapshotFullRefresh52) SbeTemplateId() (templateId uint16) {
	return 52
}

func (*SnapshotFullRefresh52) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*SnapshotFullRefresh52) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*SnapshotFullRefresh52) SbeSemanticType() (semanticType []byte) {
	return []byte("W")
}

func (*SnapshotFullRefresh52) LastMsgSeqNumProcessedId() uint16 {
	return 369
}

func (*SnapshotFullRefresh52) LastMsgSeqNumProcessedSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52) LastMsgSeqNumProcessedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LastMsgSeqNumProcessedSinceVersion()
}

func (*SnapshotFullRefresh52) LastMsgSeqNumProcessedDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52) LastMsgSeqNumProcessedMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "SeqNum"
	case 4:
		return "required"
	}
	return ""
}

func (*SnapshotFullRefresh52) LastMsgSeqNumProcessedMinValue() uint32 {
	return 0
}

func (*SnapshotFullRefresh52) LastMsgSeqNumProcessedMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*SnapshotFullRefresh52) LastMsgSeqNumProcessedNullValue() uint32 {
	return math.MaxUint32
}

func (*SnapshotFullRefresh52) TotNumReportsId() uint16 {
	return 911
}

func (*SnapshotFullRefresh52) TotNumReportsSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52) TotNumReportsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TotNumReportsSinceVersion()
}

func (*SnapshotFullRefresh52) TotNumReportsDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52) TotNumReportsMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh52) TotNumReportsMinValue() uint32 {
	return 0
}

func (*SnapshotFullRefresh52) TotNumReportsMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*SnapshotFullRefresh52) TotNumReportsNullValue() uint32 {
	return math.MaxUint32
}

func (*SnapshotFullRefresh52) SecurityIDId() uint16 {
	return 48
}

func (*SnapshotFullRefresh52) SecurityIDSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityIDSinceVersion()
}

func (*SnapshotFullRefresh52) SecurityIDDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52) SecurityIDMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh52) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SnapshotFullRefresh52) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*SnapshotFullRefresh52) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*SnapshotFullRefresh52) RptSeqId() uint16 {
	return 83
}

func (*SnapshotFullRefresh52) RptSeqSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52) RptSeqInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.RptSeqSinceVersion()
}

func (*SnapshotFullRefresh52) RptSeqDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52) RptSeqMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh52) RptSeqMinValue() uint32 {
	return 0
}

func (*SnapshotFullRefresh52) RptSeqMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*SnapshotFullRefresh52) RptSeqNullValue() uint32 {
	return math.MaxUint32
}

func (*SnapshotFullRefresh52) TransactTimeId() uint16 {
	return 60
}

func (*SnapshotFullRefresh52) TransactTimeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TransactTimeSinceVersion()
}

func (*SnapshotFullRefresh52) TransactTimeDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52) TransactTimeMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh52) TransactTimeMinValue() uint64 {
	return 0
}

func (*SnapshotFullRefresh52) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SnapshotFullRefresh52) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*SnapshotFullRefresh52) LastUpdateTimeId() uint16 {
	return 779
}

func (*SnapshotFullRefresh52) LastUpdateTimeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52) LastUpdateTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LastUpdateTimeSinceVersion()
}

func (*SnapshotFullRefresh52) LastUpdateTimeDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52) LastUpdateTimeMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh52) LastUpdateTimeMinValue() uint64 {
	return 0
}

func (*SnapshotFullRefresh52) LastUpdateTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SnapshotFullRefresh52) LastUpdateTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*SnapshotFullRefresh52) TradeDateId() uint16 {
	return 75
}

func (*SnapshotFullRefresh52) TradeDateSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52) TradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TradeDateSinceVersion()
}

func (*SnapshotFullRefresh52) TradeDateDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52) TradeDateMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh52) TradeDateMinValue() uint16 {
	return 0
}

func (*SnapshotFullRefresh52) TradeDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*SnapshotFullRefresh52) TradeDateNullValue() uint16 {
	return 65535
}

func (*SnapshotFullRefresh52) MDSecurityTradingStatusId() uint16 {
	return 1682
}

func (*SnapshotFullRefresh52) MDSecurityTradingStatusSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52) MDSecurityTradingStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDSecurityTradingStatusSinceVersion()
}

func (*SnapshotFullRefresh52) MDSecurityTradingStatusDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52) MDSecurityTradingStatusMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh52) HighLimitPriceId() uint16 {
	return 1149
}

func (*SnapshotFullRefresh52) HighLimitPriceSinceVersion() uint16 {
	return 9
}

func (s *SnapshotFullRefresh52) HighLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.HighLimitPriceSinceVersion()
}

func (*SnapshotFullRefresh52) HighLimitPriceDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52) HighLimitPriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Price"
	case 4:
		return "required"
	}
	return ""
}

func (*SnapshotFullRefresh52) LowLimitPriceId() uint16 {
	return 1148
}

func (*SnapshotFullRefresh52) LowLimitPriceSinceVersion() uint16 {
	return 9
}

func (s *SnapshotFullRefresh52) LowLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LowLimitPriceSinceVersion()
}

func (*SnapshotFullRefresh52) LowLimitPriceDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52) LowLimitPriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Price"
	case 4:
		return "required"
	}
	return ""
}

func (*SnapshotFullRefresh52) MaxPriceVariationId() uint16 {
	return 1143
}

func (*SnapshotFullRefresh52) MaxPriceVariationSinceVersion() uint16 {
	return 9
}

func (s *SnapshotFullRefresh52) MaxPriceVariationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MaxPriceVariationSinceVersion()
}

func (*SnapshotFullRefresh52) MaxPriceVariationDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52) MaxPriceVariationMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Price"
	case 4:
		return "required"
	}
	return ""
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntryPxId() uint16 {
	return 270
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntryPxSinceVersion() uint16 {
	return 9
}

func (s *SnapshotFullRefresh52NoMDEntries) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDEntryPxSinceVersion()
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntryPxMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Price"
	case 4:
		return "required"
	}
	return ""
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntrySizeId() uint16 {
	return 271
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52NoMDEntries) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDEntrySizeSinceVersion()
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntrySizeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Qty"
	case 4:
		return "optional"
	}
	return ""
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntrySizeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntrySizeMaxValue() int32 {
	return math.MaxInt32
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntrySizeNullValue() int32 {
	return 2147483647
}

func (*SnapshotFullRefresh52NoMDEntries) NumberOfOrdersId() uint16 {
	return 346
}

func (*SnapshotFullRefresh52NoMDEntries) NumberOfOrdersSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52NoMDEntries) NumberOfOrdersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NumberOfOrdersSinceVersion()
}

func (*SnapshotFullRefresh52NoMDEntries) NumberOfOrdersDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52NoMDEntries) NumberOfOrdersMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh52NoMDEntries) NumberOfOrdersMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SnapshotFullRefresh52NoMDEntries) NumberOfOrdersMaxValue() int32 {
	return math.MaxInt32
}

func (*SnapshotFullRefresh52NoMDEntries) NumberOfOrdersNullValue() int32 {
	return 2147483647
}

func (*SnapshotFullRefresh52NoMDEntries) MDPriceLevelId() uint16 {
	return 1023
}

func (*SnapshotFullRefresh52NoMDEntries) MDPriceLevelSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52NoMDEntries) MDPriceLevelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDPriceLevelSinceVersion()
}

func (*SnapshotFullRefresh52NoMDEntries) MDPriceLevelDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52NoMDEntries) MDPriceLevelMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh52NoMDEntries) MDPriceLevelMinValue() int8 {
	return math.MinInt8 + 1
}

func (*SnapshotFullRefresh52NoMDEntries) MDPriceLevelMaxValue() int8 {
	return math.MaxInt8
}

func (*SnapshotFullRefresh52NoMDEntries) MDPriceLevelNullValue() int8 {
	return 127
}

func (*SnapshotFullRefresh52NoMDEntries) TradingReferenceDateId() uint16 {
	return 5796
}

func (*SnapshotFullRefresh52NoMDEntries) TradingReferenceDateSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52NoMDEntries) TradingReferenceDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TradingReferenceDateSinceVersion()
}

func (*SnapshotFullRefresh52NoMDEntries) TradingReferenceDateDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52NoMDEntries) TradingReferenceDateMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh52NoMDEntries) TradingReferenceDateMinValue() uint16 {
	return 0
}

func (*SnapshotFullRefresh52NoMDEntries) TradingReferenceDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*SnapshotFullRefresh52NoMDEntries) TradingReferenceDateNullValue() uint16 {
	return 65535
}

func (*SnapshotFullRefresh52NoMDEntries) OpenCloseSettlFlagId() uint16 {
	return 286
}

func (*SnapshotFullRefresh52NoMDEntries) OpenCloseSettlFlagSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52NoMDEntries) OpenCloseSettlFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.OpenCloseSettlFlagSinceVersion()
}

func (*SnapshotFullRefresh52NoMDEntries) OpenCloseSettlFlagDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52NoMDEntries) OpenCloseSettlFlagMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh52NoMDEntries) SettlPriceTypeId() uint16 {
	return 731
}

func (*SnapshotFullRefresh52NoMDEntries) SettlPriceTypeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52NoMDEntries) SettlPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SettlPriceTypeSinceVersion()
}

func (*SnapshotFullRefresh52NoMDEntries) SettlPriceTypeDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52NoMDEntries) SettlPriceTypeMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh52NoMDEntries) MDEntryTypeId() uint16 {
	return 269
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntryTypeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52NoMDEntries) MDEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDEntryTypeSinceVersion()
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntryTypeDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52NoMDEntries) MDEntryTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "char"
	case 4:
		return "required"
	}
	return ""
}

func (*SnapshotFullRefresh52) NoMDEntriesId() uint16 {
	return 268
}

func (*SnapshotFullRefresh52) NoMDEntriesSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh52) NoMDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NoMDEntriesSinceVersion()
}

func (*SnapshotFullRefresh52) NoMDEntriesDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh52NoMDEntries) SbeBlockLength() (blockLength uint) {
	return 22
}

func (*SnapshotFullRefresh52NoMDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
