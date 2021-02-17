// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type SnapshotFullRefreshOrderBook53 struct {
	LastMsgSeqNumProcessed uint32
	TotNumReports          uint32
	SecurityID             int32
	NoChunks               uint32
	CurrentChunk           uint32
	TransactTime           uint64
	NoMDEntries            []SnapshotFullRefreshOrderBook53NoMDEntries
}
type SnapshotFullRefreshOrderBook53NoMDEntries struct {
	OrderID         uint64
	MDOrderPriority uint64
	MDEntryPx       PRICE9
	MDDisplayQty    int32
	MDEntryType     MDEntryTypeBookEnum
}

func (s *SnapshotFullRefreshOrderBook53) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteUint32(_w, s.NoChunks); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, s.CurrentChunk); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.TransactTime); err != nil {
		return err
	}
	var NoMDEntriesBlockLength uint16 = 29
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

func (s *SnapshotFullRefreshOrderBook53) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !s.NoChunksInActingVersion(actingVersion) {
		s.NoChunks = s.NoChunksNullValue()
	} else {
		if err := _m.ReadUint32(_r, &s.NoChunks); err != nil {
			return err
		}
	}
	if !s.CurrentChunkInActingVersion(actingVersion) {
		s.CurrentChunk = s.CurrentChunkNullValue()
	} else {
		if err := _m.ReadUint32(_r, &s.CurrentChunk); err != nil {
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
			s.NoMDEntries = make([]SnapshotFullRefreshOrderBook53NoMDEntries, NoMDEntriesNumInGroup)
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

func (s *SnapshotFullRefreshOrderBook53) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if s.NoChunksInActingVersion(actingVersion) {
		if s.NoChunks < s.NoChunksMinValue() || s.NoChunks > s.NoChunksMaxValue() {
			return fmt.Errorf("Range check failed on s.NoChunks (%v < %v > %v)", s.NoChunksMinValue(), s.NoChunks, s.NoChunksMaxValue())
		}
	}
	if s.CurrentChunkInActingVersion(actingVersion) {
		if s.CurrentChunk < s.CurrentChunkMinValue() || s.CurrentChunk > s.CurrentChunkMaxValue() {
			return fmt.Errorf("Range check failed on s.CurrentChunk (%v < %v > %v)", s.CurrentChunkMinValue(), s.CurrentChunk, s.CurrentChunkMaxValue())
		}
	}
	if s.TransactTimeInActingVersion(actingVersion) {
		if s.TransactTime < s.TransactTimeMinValue() || s.TransactTime > s.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on s.TransactTime (%v < %v > %v)", s.TransactTimeMinValue(), s.TransactTime, s.TransactTimeMaxValue())
		}
	}
	for _, prop := range s.NoMDEntries {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func SnapshotFullRefreshOrderBook53Init(s *SnapshotFullRefreshOrderBook53) {
	return
}

func (s *SnapshotFullRefreshOrderBook53NoMDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint64(_w, s.OrderID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.MDOrderPriority); err != nil {
		return err
	}
	if err := s.MDEntryPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.MDDisplayQty); err != nil {
		return err
	}
	if err := s.MDEntryType.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (s *SnapshotFullRefreshOrderBook53NoMDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !s.OrderIDInActingVersion(actingVersion) {
		s.OrderID = s.OrderIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.OrderID); err != nil {
			return err
		}
	}
	if !s.MDOrderPriorityInActingVersion(actingVersion) {
		s.MDOrderPriority = s.MDOrderPriorityNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.MDOrderPriority); err != nil {
			return err
		}
	}
	if s.MDEntryPxInActingVersion(actingVersion) {
		if err := s.MDEntryPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !s.MDDisplayQtyInActingVersion(actingVersion) {
		s.MDDisplayQty = s.MDDisplayQtyNullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.MDDisplayQty); err != nil {
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

func (s *SnapshotFullRefreshOrderBook53NoMDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.OrderIDInActingVersion(actingVersion) {
		if s.OrderID < s.OrderIDMinValue() || s.OrderID > s.OrderIDMaxValue() {
			return fmt.Errorf("Range check failed on s.OrderID (%v < %v > %v)", s.OrderIDMinValue(), s.OrderID, s.OrderIDMaxValue())
		}
	}
	if s.MDOrderPriorityInActingVersion(actingVersion) {
		if s.MDOrderPriority != s.MDOrderPriorityNullValue() && (s.MDOrderPriority < s.MDOrderPriorityMinValue() || s.MDOrderPriority > s.MDOrderPriorityMaxValue()) {
			return fmt.Errorf("Range check failed on s.MDOrderPriority (%v < %v > %v)", s.MDOrderPriorityMinValue(), s.MDOrderPriority, s.MDOrderPriorityMaxValue())
		}
	}
	if s.MDDisplayQtyInActingVersion(actingVersion) {
		if s.MDDisplayQty < s.MDDisplayQtyMinValue() || s.MDDisplayQty > s.MDDisplayQtyMaxValue() {
			return fmt.Errorf("Range check failed on s.MDDisplayQty (%v < %v > %v)", s.MDDisplayQtyMinValue(), s.MDDisplayQty, s.MDDisplayQtyMaxValue())
		}
	}
	if err := s.MDEntryType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func SnapshotFullRefreshOrderBook53NoMDEntriesInit(s *SnapshotFullRefreshOrderBook53NoMDEntries) {
	s.MDOrderPriority = 18446744073709551615
	return
}

func (*SnapshotFullRefreshOrderBook53) SbeBlockLength() (blockLength uint16) {
	return 28
}

func (*SnapshotFullRefreshOrderBook53) SbeTemplateId() (templateId uint16) {
	return 53
}

func (*SnapshotFullRefreshOrderBook53) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*SnapshotFullRefreshOrderBook53) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*SnapshotFullRefreshOrderBook53) SbeSemanticType() (semanticType []byte) {
	return []byte("W")
}

func (*SnapshotFullRefreshOrderBook53) LastMsgSeqNumProcessedId() uint16 {
	return 369
}

func (*SnapshotFullRefreshOrderBook53) LastMsgSeqNumProcessedSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefreshOrderBook53) LastMsgSeqNumProcessedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LastMsgSeqNumProcessedSinceVersion()
}

func (*SnapshotFullRefreshOrderBook53) LastMsgSeqNumProcessedDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53) LastMsgSeqNumProcessedMetaAttribute(meta int) string {
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

func (*SnapshotFullRefreshOrderBook53) LastMsgSeqNumProcessedMinValue() uint32 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53) LastMsgSeqNumProcessedMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*SnapshotFullRefreshOrderBook53) LastMsgSeqNumProcessedNullValue() uint32 {
	return math.MaxUint32
}

func (*SnapshotFullRefreshOrderBook53) TotNumReportsId() uint16 {
	return 911
}

func (*SnapshotFullRefreshOrderBook53) TotNumReportsSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefreshOrderBook53) TotNumReportsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TotNumReportsSinceVersion()
}

func (*SnapshotFullRefreshOrderBook53) TotNumReportsDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53) TotNumReportsMetaAttribute(meta int) string {
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

func (*SnapshotFullRefreshOrderBook53) TotNumReportsMinValue() uint32 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53) TotNumReportsMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*SnapshotFullRefreshOrderBook53) TotNumReportsNullValue() uint32 {
	return math.MaxUint32
}

func (*SnapshotFullRefreshOrderBook53) SecurityIDId() uint16 {
	return 48
}

func (*SnapshotFullRefreshOrderBook53) SecurityIDSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefreshOrderBook53) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SecurityIDSinceVersion()
}

func (*SnapshotFullRefreshOrderBook53) SecurityIDDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53) SecurityIDMetaAttribute(meta int) string {
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

func (*SnapshotFullRefreshOrderBook53) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SnapshotFullRefreshOrderBook53) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*SnapshotFullRefreshOrderBook53) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*SnapshotFullRefreshOrderBook53) NoChunksId() uint16 {
	return 37709
}

func (*SnapshotFullRefreshOrderBook53) NoChunksSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefreshOrderBook53) NoChunksInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NoChunksSinceVersion()
}

func (*SnapshotFullRefreshOrderBook53) NoChunksDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53) NoChunksMetaAttribute(meta int) string {
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

func (*SnapshotFullRefreshOrderBook53) NoChunksMinValue() uint32 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53) NoChunksMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*SnapshotFullRefreshOrderBook53) NoChunksNullValue() uint32 {
	return math.MaxUint32
}

func (*SnapshotFullRefreshOrderBook53) CurrentChunkId() uint16 {
	return 37710
}

func (*SnapshotFullRefreshOrderBook53) CurrentChunkSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefreshOrderBook53) CurrentChunkInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.CurrentChunkSinceVersion()
}

func (*SnapshotFullRefreshOrderBook53) CurrentChunkDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53) CurrentChunkMetaAttribute(meta int) string {
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

func (*SnapshotFullRefreshOrderBook53) CurrentChunkMinValue() uint32 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53) CurrentChunkMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*SnapshotFullRefreshOrderBook53) CurrentChunkNullValue() uint32 {
	return math.MaxUint32
}

func (*SnapshotFullRefreshOrderBook53) TransactTimeId() uint16 {
	return 60
}

func (*SnapshotFullRefreshOrderBook53) TransactTimeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefreshOrderBook53) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TransactTimeSinceVersion()
}

func (*SnapshotFullRefreshOrderBook53) TransactTimeDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53) TransactTimeMetaAttribute(meta int) string {
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

func (*SnapshotFullRefreshOrderBook53) TransactTimeMinValue() uint64 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SnapshotFullRefreshOrderBook53) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) OrderIDId() uint16 {
	return 37
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) OrderIDSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefreshOrderBook53NoMDEntries) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.OrderIDSinceVersion()
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) OrderIDDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) OrderIDMetaAttribute(meta int) string {
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

func (*SnapshotFullRefreshOrderBook53NoMDEntries) OrderIDMinValue() uint64 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDOrderPriorityId() uint16 {
	return 37707
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDOrderPrioritySinceVersion() uint16 {
	return 7
}

func (s *SnapshotFullRefreshOrderBook53NoMDEntries) MDOrderPriorityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDOrderPrioritySinceVersion()
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDOrderPriorityDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDOrderPriorityMetaAttribute(meta int) string {
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

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDOrderPriorityMinValue() uint64 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDOrderPriorityMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDOrderPriorityNullValue() uint64 {
	return 18446744073709551615
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDEntryPxId() uint16 {
	return 270
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDEntryPxSinceVersion() uint16 {
	return 9
}

func (s *SnapshotFullRefreshOrderBook53NoMDEntries) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDEntryPxSinceVersion()
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDEntryPxMetaAttribute(meta int) string {
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

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDDisplayQtyId() uint16 {
	return 37706
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDDisplayQtySinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefreshOrderBook53NoMDEntries) MDDisplayQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDDisplayQtySinceVersion()
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDDisplayQtyDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDDisplayQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Qty"
	case 4:
		return "required"
	}
	return ""
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDDisplayQtyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDDisplayQtyMaxValue() int32 {
	return math.MaxInt32
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDDisplayQtyNullValue() int32 {
	return math.MinInt32
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDEntryTypeId() uint16 {
	return 269
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDEntryTypeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefreshOrderBook53NoMDEntries) MDEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MDEntryTypeSinceVersion()
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDEntryTypeDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) MDEntryTypeMetaAttribute(meta int) string {
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

func (*SnapshotFullRefreshOrderBook53) NoMDEntriesId() uint16 {
	return 268
}

func (*SnapshotFullRefreshOrderBook53) NoMDEntriesSinceVersion() uint16 {
	return 0
}

func (s *SnapshotFullRefreshOrderBook53) NoMDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NoMDEntriesSinceVersion()
}

func (*SnapshotFullRefreshOrderBook53) NoMDEntriesDeprecated() uint16 {
	return 0
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) SbeBlockLength() (blockLength uint) {
	return 29
}

func (*SnapshotFullRefreshOrderBook53NoMDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
