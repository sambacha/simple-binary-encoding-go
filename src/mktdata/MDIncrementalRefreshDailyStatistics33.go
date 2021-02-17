// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MDIncrementalRefreshDailyStatistics33 struct {
	TransactTime        uint64
	MatchEventIndicator MatchEventIndicator
	NoMDEntries         []MDIncrementalRefreshDailyStatistics33NoMDEntries
}
type MDIncrementalRefreshDailyStatistics33NoMDEntries struct {
	MDEntryPx            PRICENULL
	MDEntrySize          int32
	SecurityID           int32
	RptSeq               uint32
	TradingReferenceDate uint16
	SettlPriceType       SettlPriceType
	MDUpdateAction       MDUpdateActionEnum
	MDEntryType          MDEntryTypeDailyStatisticsEnum
}

func (m *MDIncrementalRefreshDailyStatistics33) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, m.TransactTime); err != nil {
		return err
	}
	if err := m.MatchEventIndicator.Encode(_m, _w); err != nil {
		return err
	}

	for i := 0; i < 2; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	var NoMDEntriesBlockLength uint16 = 32
	if err := _m.WriteUint16(_w, NoMDEntriesBlockLength); err != nil {
		return err
	}
	var NoMDEntriesNumInGroup uint8 = uint8(len(m.NoMDEntries))
	if err := _m.WriteUint8(_w, NoMDEntriesNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.NoMDEntries {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}

		for i := 0; i < 7; i++ {
			if err := _m.WriteUint8(_w, uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *MDIncrementalRefreshDailyStatistics33) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.TransactTimeInActingVersion(actingVersion) {
		m.TransactTime = m.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.TransactTime); err != nil {
			return err
		}
	}
	if m.MatchEventIndicatorInActingVersion(actingVersion) {
		if err := m.MatchEventIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	io.CopyN(ioutil.Discard, _r, 2)

	if m.NoMDEntriesInActingVersion(actingVersion) {
		var NoMDEntriesBlockLength uint16
		if err := _m.ReadUint16(_r, &NoMDEntriesBlockLength); err != nil {
			return err
		}
		var NoMDEntriesNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoMDEntriesNumInGroup); err != nil {
			return err
		}
		if cap(m.NoMDEntries) < int(NoMDEntriesNumInGroup) {
			m.NoMDEntries = make([]MDIncrementalRefreshDailyStatistics33NoMDEntries, NoMDEntriesNumInGroup)
		}
		m.NoMDEntries = m.NoMDEntries[:NoMDEntriesNumInGroup]
		for i, _ := range m.NoMDEntries {
			if err := m.NoMDEntries[i].Decode(_m, _r, actingVersion, uint(NoMDEntriesBlockLength)); err != nil {
				return err
			}
		}
		io.CopyN(ioutil.Discard, _r, 7)
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MDIncrementalRefreshDailyStatistics33) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.TransactTimeInActingVersion(actingVersion) {
		if m.TransactTime < m.TransactTimeMinValue() || m.TransactTime > m.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on m.TransactTime (%v < %v > %v)", m.TransactTimeMinValue(), m.TransactTime, m.TransactTimeMaxValue())
		}
	}
	for _, prop := range m.NoMDEntries {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func MDIncrementalRefreshDailyStatistics33Init(m *MDIncrementalRefreshDailyStatistics33) {
	return
}

func (m *MDIncrementalRefreshDailyStatistics33NoMDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.MDEntryPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.MDEntrySize); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.RptSeq); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.TradingReferenceDate); err != nil {
		return err
	}
	if err := m.SettlPriceType.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MDUpdateAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MDEntryType.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MDIncrementalRefreshDailyStatistics33NoMDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if m.MDEntryPxInActingVersion(actingVersion) {
		if err := m.MDEntryPx.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.MDEntrySizeInActingVersion(actingVersion) {
		m.MDEntrySize = m.MDEntrySizeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.MDEntrySize); err != nil {
			return err
		}
	}
	if !m.SecurityIDInActingVersion(actingVersion) {
		m.SecurityID = m.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.SecurityID); err != nil {
			return err
		}
	}
	if !m.RptSeqInActingVersion(actingVersion) {
		m.RptSeq = m.RptSeqNullValue()
	} else {
		if err := _m.ReadUint32(_r, &m.RptSeq); err != nil {
			return err
		}
	}
	if !m.TradingReferenceDateInActingVersion(actingVersion) {
		m.TradingReferenceDate = m.TradingReferenceDateNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.TradingReferenceDate); err != nil {
			return err
		}
	}
	if m.SettlPriceTypeInActingVersion(actingVersion) {
		if err := m.SettlPriceType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.MDUpdateActionInActingVersion(actingVersion) {
		if err := m.MDUpdateAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.MDEntryTypeInActingVersion(actingVersion) {
		if err := m.MDEntryType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDIncrementalRefreshDailyStatistics33NoMDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.MDEntrySizeInActingVersion(actingVersion) {
		if m.MDEntrySize != m.MDEntrySizeNullValue() && (m.MDEntrySize < m.MDEntrySizeMinValue() || m.MDEntrySize > m.MDEntrySizeMaxValue()) {
			return fmt.Errorf("Range check failed on m.MDEntrySize (%v < %v > %v)", m.MDEntrySizeMinValue(), m.MDEntrySize, m.MDEntrySizeMaxValue())
		}
	}
	if m.SecurityIDInActingVersion(actingVersion) {
		if m.SecurityID < m.SecurityIDMinValue() || m.SecurityID > m.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on m.SecurityID (%v < %v > %v)", m.SecurityIDMinValue(), m.SecurityID, m.SecurityIDMaxValue())
		}
	}
	if m.RptSeqInActingVersion(actingVersion) {
		if m.RptSeq < m.RptSeqMinValue() || m.RptSeq > m.RptSeqMaxValue() {
			return fmt.Errorf("Range check failed on m.RptSeq (%v < %v > %v)", m.RptSeqMinValue(), m.RptSeq, m.RptSeqMaxValue())
		}
	}
	if m.TradingReferenceDateInActingVersion(actingVersion) {
		if m.TradingReferenceDate != m.TradingReferenceDateNullValue() && (m.TradingReferenceDate < m.TradingReferenceDateMinValue() || m.TradingReferenceDate > m.TradingReferenceDateMaxValue()) {
			return fmt.Errorf("Range check failed on m.TradingReferenceDate (%v < %v > %v)", m.TradingReferenceDateMinValue(), m.TradingReferenceDate, m.TradingReferenceDateMaxValue())
		}
	}
	if err := m.MDUpdateAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.MDEntryType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func MDIncrementalRefreshDailyStatistics33NoMDEntriesInit(m *MDIncrementalRefreshDailyStatistics33NoMDEntries) {
	m.MDEntrySize = 2147483647
	m.TradingReferenceDate = 65535
	return
}

func (*MDIncrementalRefreshDailyStatistics33) SbeBlockLength() (blockLength uint16) {
	return 11
}

func (*MDIncrementalRefreshDailyStatistics33) SbeTemplateId() (templateId uint16) {
	return 33
}

func (*MDIncrementalRefreshDailyStatistics33) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MDIncrementalRefreshDailyStatistics33) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDIncrementalRefreshDailyStatistics33) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MDIncrementalRefreshDailyStatistics33) TransactTimeId() uint16 {
	return 60
}

func (*MDIncrementalRefreshDailyStatistics33) TransactTimeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshDailyStatistics33) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TransactTimeSinceVersion()
}

func (*MDIncrementalRefreshDailyStatistics33) TransactTimeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33) TransactTimeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshDailyStatistics33) TransactTimeMinValue() uint64 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDIncrementalRefreshDailyStatistics33) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDIncrementalRefreshDailyStatistics33) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MDIncrementalRefreshDailyStatistics33) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshDailyStatistics33) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MDIncrementalRefreshDailyStatistics33) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntryPxId() uint16 {
	return 270
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntryPxSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryPxSinceVersion()
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntryPxMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntrySizeId() uint16 {
	return 271
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntrySizeSinceVersion()
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntrySizeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntrySizeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntrySizeMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntrySizeNullValue() int32 {
	return 2147483647
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) SecurityIDId() uint16 {
	return 48
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshDailyStatistics33NoMDEntries) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) SecurityIDMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) RptSeqId() uint16 {
	return 83
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) RptSeqSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshDailyStatistics33NoMDEntries) RptSeqInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RptSeqSinceVersion()
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) RptSeqDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) RptSeqMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) RptSeqMinValue() uint32 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) RptSeqMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) RptSeqNullValue() uint32 {
	return math.MaxUint32
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) TradingReferenceDateId() uint16 {
	return 5796
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) TradingReferenceDateSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshDailyStatistics33NoMDEntries) TradingReferenceDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradingReferenceDateSinceVersion()
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) TradingReferenceDateDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) TradingReferenceDateMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) TradingReferenceDateMinValue() uint16 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) TradingReferenceDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) TradingReferenceDateNullValue() uint16 {
	return 65535
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) SettlPriceTypeId() uint16 {
	return 731
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) SettlPriceTypeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshDailyStatistics33NoMDEntries) SettlPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SettlPriceTypeSinceVersion()
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) SettlPriceTypeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) SettlPriceTypeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDUpdateActionId() uint16 {
	return 279
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshDailyStatistics33NoMDEntries) MDUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDUpdateActionSinceVersion()
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDUpdateActionDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDUpdateActionMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntryTypeId() uint16 {
	return 269
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntryTypeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryTypeSinceVersion()
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntryTypeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) MDEntryTypeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshDailyStatistics33) NoMDEntriesId() uint16 {
	return 268
}

func (*MDIncrementalRefreshDailyStatistics33) NoMDEntriesSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshDailyStatistics33) NoMDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoMDEntriesSinceVersion()
}

func (*MDIncrementalRefreshDailyStatistics33) NoMDEntriesDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) SbeBlockLength() (blockLength uint) {
	return 32
}

func (*MDIncrementalRefreshDailyStatistics33NoMDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
