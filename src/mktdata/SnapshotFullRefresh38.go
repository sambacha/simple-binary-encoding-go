// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type SnapshotFullRefresh38 struct {
	LastMsgSeqNumProcessed  uint32
	TotNumReports           uint32
	SecurityID              int32
	RptSeq                  uint32
	TransactTime            uint64
	LastUpdateTime          uint64
	TradeDate               uint16
	MDSecurityTradingStatus SecurityTradingStatusEnum
	HighLimitPrice          PRICENULL
	LowLimitPrice           PRICENULL
	MaxPriceVariation       PRICENULL
	NoMDEntries             []SnapshotFullRefresh38NoMDEntries
}
type SnapshotFullRefresh38NoMDEntries struct {
	MDEntryPx            PRICENULL
	MDEntrySize          int32
	NumberOfOrders       int32
	MDPriceLevel         int8
	TradingReferenceDate uint16
	OpenCloseSettlFlag   OpenCloseSettlFlagEnum
	SettlPriceType       SettlPriceType
	MDEntryType          MDEntryTypeEnum
}

func (s *SnapshotFullRefresh38) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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

func (s *SnapshotFullRefresh38) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
			s.NoMDEntries = make([]SnapshotFullRefresh38NoMDEntries, NoMDEntriesNumInGroup)
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

func (s *SnapshotFullRefresh38) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func SnapshotFullRefresh38Init(s *SnapshotFullRefresh38) {
	s.TradeDate = 65535
	return
}

func (s *SnapshotFullRefresh38NoMDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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

func (s *SnapshotFullRefresh38NoMDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (s *SnapshotFullRefresh38NoMDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func SnapshotFullRefresh38NoMDEntriesInit(s *SnapshotFullRefresh38NoMDEntries) {
	s.MDEntrySize = 2147483647
	s.NumberOfOrders = 2147483647
	s.MDPriceLevel = 127
	s.TradingReferenceDate = 65535
	return
}

func (*SnapshotFullRefresh38) SbeBlockLength() (blockLength uint16) {
	return 59
}

func (*SnapshotFullRefresh38) SbeTemplateId() (templateId uint16) {
	return 38
}

func (*SnapshotFullRefresh38) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*SnapshotFullRefresh38) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*SnapshotFullRefresh38) SbeSemanticType() (semanticType []byte) {
	return []byte("W")
}

func (*SnapshotFullRefresh38) LastMsgSeqNumProcessedId() uint16 {
	return 369
}

func (*SnapshotFullRefresh38) LastMsgSeqNumProcessedSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38) LastMsgSeqNumProcessedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LastMsgSeqNumProcessedSinceVersion()
}

func (*SnapshotFullRefresh38) LastMsgSeqNumProcessedDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38) LastMsgSeqNumProcessedMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38) LastMsgSeqNumProcessedMinValue() uint32 {
	return 0
}

func (*SnapshotFullRefresh38) LastMsgSeqNumProcessedMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*SnapshotFullRefresh38) LastMsgSeqNumProcessedNullValue() uint32 {
	return math.MaxUint32
}

func (*SnapshotFullRefresh38) TotNumReportsId() uint16 {
	return 911
}

func (*SnapshotFullRefresh38) TotNumReportsSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38) TotNumReportsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TotNumReportsSinceVersion()
}

func (*SnapshotFullRefresh38) TotNumReportsDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38) TotNumReportsMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38) TotNumReportsMinValue() uint32 {
	return 0
}

func (*SnapshotFullRefresh38) TotNumReportsMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*SnapshotFullRefresh38) TotNumReportsNullValue() uint32 {
	return math.MaxUint32
}

func (*SnapshotFullRefresh38) SecurityIDId() uint16 {
	return 48
}

func (*SnapshotFullRefresh38) SecurityIDSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityIDSinceVersion()
}

func (*SnapshotFullRefresh38) SecurityIDDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38) SecurityIDMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SnapshotFullRefresh38) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*SnapshotFullRefresh38) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*SnapshotFullRefresh38) RptSeqId() uint16 {
	return 83
}

func (*SnapshotFullRefresh38) RptSeqSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38) RptSeqInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.RptSeqSinceVersion()
}

func (*SnapshotFullRefresh38) RptSeqDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38) RptSeqMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38) RptSeqMinValue() uint32 {
	return 0
}

func (*SnapshotFullRefresh38) RptSeqMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*SnapshotFullRefresh38) RptSeqNullValue() uint32 {
	return math.MaxUint32
}

func (*SnapshotFullRefresh38) TransactTimeId() uint16 {
	return 60
}

func (*SnapshotFullRefresh38) TransactTimeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TransactTimeSinceVersion()
}

func (*SnapshotFullRefresh38) TransactTimeDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38) TransactTimeMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38) TransactTimeMinValue() uint64 {
	return 0
}

func (*SnapshotFullRefresh38) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SnapshotFullRefresh38) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*SnapshotFullRefresh38) LastUpdateTimeId() uint16 {
	return 779
}

func (*SnapshotFullRefresh38) LastUpdateTimeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38) LastUpdateTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LastUpdateTimeSinceVersion()
}

func (*SnapshotFullRefresh38) LastUpdateTimeDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38) LastUpdateTimeMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38) LastUpdateTimeMinValue() uint64 {
	return 0
}

func (*SnapshotFullRefresh38) LastUpdateTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SnapshotFullRefresh38) LastUpdateTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*SnapshotFullRefresh38) TradeDateId() uint16 {
	return 75
}

func (*SnapshotFullRefresh38) TradeDateSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38) TradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TradeDateSinceVersion()
}

func (*SnapshotFullRefresh38) TradeDateDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38) TradeDateMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38) TradeDateMinValue() uint16 {
	return 0
}

func (*SnapshotFullRefresh38) TradeDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*SnapshotFullRefresh38) TradeDateNullValue() uint16 {
	return 65535
}

func (*SnapshotFullRefresh38) MDSecurityTradingStatusId() uint16 {
	return 1682
}

func (*SnapshotFullRefresh38) MDSecurityTradingStatusSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38) MDSecurityTradingStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDSecurityTradingStatusSinceVersion()
}

func (*SnapshotFullRefresh38) MDSecurityTradingStatusDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38) MDSecurityTradingStatusMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38) HighLimitPriceId() uint16 {
	return 1149
}

func (*SnapshotFullRefresh38) HighLimitPriceSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38) HighLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.HighLimitPriceSinceVersion()
}

func (*SnapshotFullRefresh38) HighLimitPriceDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38) HighLimitPriceMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38) LowLimitPriceId() uint16 {
	return 1148
}

func (*SnapshotFullRefresh38) LowLimitPriceSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38) LowLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LowLimitPriceSinceVersion()
}

func (*SnapshotFullRefresh38) LowLimitPriceDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38) LowLimitPriceMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38) MaxPriceVariationId() uint16 {
	return 1143
}

func (*SnapshotFullRefresh38) MaxPriceVariationSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38) MaxPriceVariationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MaxPriceVariationSinceVersion()
}

func (*SnapshotFullRefresh38) MaxPriceVariationDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38) MaxPriceVariationMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38NoMDEntries) MDEntryPxId() uint16 {
	return 270
}

func (*SnapshotFullRefresh38NoMDEntries) MDEntryPxSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38NoMDEntries) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDEntryPxSinceVersion()
}

func (*SnapshotFullRefresh38NoMDEntries) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38NoMDEntries) MDEntryPxMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38NoMDEntries) MDEntrySizeId() uint16 {
	return 271
}

func (*SnapshotFullRefresh38NoMDEntries) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38NoMDEntries) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDEntrySizeSinceVersion()
}

func (*SnapshotFullRefresh38NoMDEntries) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38NoMDEntries) MDEntrySizeMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38NoMDEntries) MDEntrySizeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SnapshotFullRefresh38NoMDEntries) MDEntrySizeMaxValue() int32 {
	return math.MaxInt32
}

func (*SnapshotFullRefresh38NoMDEntries) MDEntrySizeNullValue() int32 {
	return 2147483647
}

func (*SnapshotFullRefresh38NoMDEntries) NumberOfOrdersId() uint16 {
	return 346
}

func (*SnapshotFullRefresh38NoMDEntries) NumberOfOrdersSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38NoMDEntries) NumberOfOrdersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NumberOfOrdersSinceVersion()
}

func (*SnapshotFullRefresh38NoMDEntries) NumberOfOrdersDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38NoMDEntries) NumberOfOrdersMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38NoMDEntries) NumberOfOrdersMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SnapshotFullRefresh38NoMDEntries) NumberOfOrdersMaxValue() int32 {
	return math.MaxInt32
}

func (*SnapshotFullRefresh38NoMDEntries) NumberOfOrdersNullValue() int32 {
	return 2147483647
}

func (*SnapshotFullRefresh38NoMDEntries) MDPriceLevelId() uint16 {
	return 1023
}

func (*SnapshotFullRefresh38NoMDEntries) MDPriceLevelSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38NoMDEntries) MDPriceLevelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDPriceLevelSinceVersion()
}

func (*SnapshotFullRefresh38NoMDEntries) MDPriceLevelDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38NoMDEntries) MDPriceLevelMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38NoMDEntries) MDPriceLevelMinValue() int8 {
	return math.MinInt8 + 1
}

func (*SnapshotFullRefresh38NoMDEntries) MDPriceLevelMaxValue() int8 {
	return math.MaxInt8
}

func (*SnapshotFullRefresh38NoMDEntries) MDPriceLevelNullValue() int8 {
	return 127
}

func (*SnapshotFullRefresh38NoMDEntries) TradingReferenceDateId() uint16 {
	return 5796
}

func (*SnapshotFullRefresh38NoMDEntries) TradingReferenceDateSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38NoMDEntries) TradingReferenceDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TradingReferenceDateSinceVersion()
}

func (*SnapshotFullRefresh38NoMDEntries) TradingReferenceDateDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38NoMDEntries) TradingReferenceDateMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38NoMDEntries) TradingReferenceDateMinValue() uint16 {
	return 0
}

func (*SnapshotFullRefresh38NoMDEntries) TradingReferenceDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*SnapshotFullRefresh38NoMDEntries) TradingReferenceDateNullValue() uint16 {
	return 65535
}

func (*SnapshotFullRefresh38NoMDEntries) OpenCloseSettlFlagId() uint16 {
	return 286
}

func (*SnapshotFullRefresh38NoMDEntries) OpenCloseSettlFlagSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38NoMDEntries) OpenCloseSettlFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.OpenCloseSettlFlagSinceVersion()
}

func (*SnapshotFullRefresh38NoMDEntries) OpenCloseSettlFlagDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38NoMDEntries) OpenCloseSettlFlagMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38NoMDEntries) SettlPriceTypeId() uint16 {
	return 731
}

func (*SnapshotFullRefresh38NoMDEntries) SettlPriceTypeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38NoMDEntries) SettlPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SettlPriceTypeSinceVersion()
}

func (*SnapshotFullRefresh38NoMDEntries) SettlPriceTypeDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38NoMDEntries) SettlPriceTypeMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38NoMDEntries) MDEntryTypeId() uint16 {
	return 269
}

func (*SnapshotFullRefresh38NoMDEntries) MDEntryTypeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38NoMDEntries) MDEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDEntryTypeSinceVersion()
}

func (*SnapshotFullRefresh38NoMDEntries) MDEntryTypeDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38NoMDEntries) MDEntryTypeMetaAttribute(meta int) string {
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

func (*SnapshotFullRefresh38) NoMDEntriesId() uint16 {
	return 268
}

func (*SnapshotFullRefresh38) NoMDEntriesSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefresh38) NoMDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NoMDEntriesSinceVersion()
}

func (*SnapshotFullRefresh38) NoMDEntriesDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefresh38NoMDEntries) SbeBlockLength() (blockLength uint) {
	return 22
}

func (*SnapshotFullRefresh38NoMDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
