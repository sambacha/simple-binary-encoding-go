// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MDIncrementalRefreshSessionStatistics51 struct {
	TransactTime        uint64
	MatchEventIndicator MatchEventIndicator
	NoMDEntries         []MDIncrementalRefreshSessionStatistics51NoMDEntries
}
type MDIncrementalRefreshSessionStatistics51NoMDEntries struct {
	MDEntryPx          PRICE9
	SecurityID         int32
	RptSeq             uint32
	OpenCloseSettlFlag OpenCloseSettlFlagEnum
	MDUpdateAction     MDUpdateActionEnum
	MDEntryType        MDEntryTypeStatisticsEnum
	MDEntrySize        int32
}

func (m *MDIncrementalRefreshSessionStatistics51) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	var NoMDEntriesBlockLength uint16 = 24
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

		for i := 0; i < 1; i++ {
			if err := _m.WriteUint8(_w, uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *MDIncrementalRefreshSessionStatistics51) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
			m.NoMDEntries = make([]MDIncrementalRefreshSessionStatistics51NoMDEntries, NoMDEntriesNumInGroup)
		}
		m.NoMDEntries = m.NoMDEntries[:NoMDEntriesNumInGroup]
		for i, _ := range m.NoMDEntries {
			if err := m.NoMDEntries[i].Decode(_m, _r, actingVersion, uint(NoMDEntriesBlockLength)); err != nil {
				return err
			}
		}
		io.CopyN(ioutil.Discard, _r, 1)
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MDIncrementalRefreshSessionStatistics51) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDIncrementalRefreshSessionStatistics51Init(m *MDIncrementalRefreshSessionStatistics51) {
	return
}

func (m *MDIncrementalRefreshSessionStatistics51NoMDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.MDEntryPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.RptSeq); err != nil {
		return err
	}
	if err := m.OpenCloseSettlFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MDUpdateAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MDEntryType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.MDEntrySize); err != nil {
		return err
	}
	return nil
}

func (m *MDIncrementalRefreshSessionStatistics51NoMDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if m.MDEntryPxInActingVersion(actingVersion) {
		if err := m.MDEntryPx.Decode(_m, _r, actingVersion); err != nil {
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
	if m.OpenCloseSettlFlagInActingVersion(actingVersion) {
		if err := m.OpenCloseSettlFlag.Decode(_m, _r, actingVersion); err != nil {
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
	if !m.MDEntrySizeInActingVersion(actingVersion) {
		m.MDEntrySize = m.MDEntrySizeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.MDEntrySize); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDIncrementalRefreshSessionStatistics51NoMDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if err := m.OpenCloseSettlFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.MDUpdateAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.MDEntryType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.MDEntrySizeInActingVersion(actingVersion) {
		if m.MDEntrySize != m.MDEntrySizeNullValue() && (m.MDEntrySize < m.MDEntrySizeMinValue() || m.MDEntrySize > m.MDEntrySizeMaxValue()) {
			return fmt.Errorf("Range check failed on m.MDEntrySize (%v < %v > %v)", m.MDEntrySizeMinValue(), m.MDEntrySize, m.MDEntrySizeMaxValue())
		}
	}
	return nil
}

func MDIncrementalRefreshSessionStatistics51NoMDEntriesInit(m *MDIncrementalRefreshSessionStatistics51NoMDEntries) {
	m.MDEntrySize = 2147483647
	return
}

func (*MDIncrementalRefreshSessionStatistics51) SbeBlockLength() (blockLength uint16) {
	return 11
}

func (*MDIncrementalRefreshSessionStatistics51) SbeTemplateId() (templateId uint16) {
	return 51
}

func (*MDIncrementalRefreshSessionStatistics51) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MDIncrementalRefreshSessionStatistics51) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDIncrementalRefreshSessionStatistics51) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MDIncrementalRefreshSessionStatistics51) TransactTimeId() uint16 {
	return 60
}

func (*MDIncrementalRefreshSessionStatistics51) TransactTimeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshSessionStatistics51) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TransactTimeSinceVersion()
}

func (*MDIncrementalRefreshSessionStatistics51) TransactTimeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshSessionStatistics51) TransactTimeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshSessionStatistics51) TransactTimeMinValue() uint64 {
	return 0
}

func (*MDIncrementalRefreshSessionStatistics51) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDIncrementalRefreshSessionStatistics51) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDIncrementalRefreshSessionStatistics51) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MDIncrementalRefreshSessionStatistics51) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshSessionStatistics51) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MDIncrementalRefreshSessionStatistics51) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshSessionStatistics51) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntryPxId() uint16 {
	return 270
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntryPxSinceVersion() uint16 {
	return 9
}

func (m *MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryPxSinceVersion()
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntryPxMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) SecurityIDId() uint16 {
	return 48
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshSessionStatistics51NoMDEntries) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) SecurityIDMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) RptSeqId() uint16 {
	return 83
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) RptSeqSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshSessionStatistics51NoMDEntries) RptSeqInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RptSeqSinceVersion()
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) RptSeqDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) RptSeqMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) RptSeqMinValue() uint32 {
	return 0
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) RptSeqMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) RptSeqNullValue() uint32 {
	return math.MaxUint32
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) OpenCloseSettlFlagId() uint16 {
	return 286
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) OpenCloseSettlFlagSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshSessionStatistics51NoMDEntries) OpenCloseSettlFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OpenCloseSettlFlagSinceVersion()
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) OpenCloseSettlFlagDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) OpenCloseSettlFlagMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDUpdateActionId() uint16 {
	return 279
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshSessionStatistics51NoMDEntries) MDUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDUpdateActionSinceVersion()
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDUpdateActionDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDUpdateActionMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntryTypeId() uint16 {
	return 269
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntryTypeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryTypeSinceVersion()
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntryTypeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntryTypeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntrySizeId() uint16 {
	return 271
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntrySizeSinceVersion()
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntrySizeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntrySizeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntrySizeMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) MDEntrySizeNullValue() int32 {
	return 2147483647
}

func (*MDIncrementalRefreshSessionStatistics51) NoMDEntriesId() uint16 {
	return 268
}

func (*MDIncrementalRefreshSessionStatistics51) NoMDEntriesSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshSessionStatistics51) NoMDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoMDEntriesSinceVersion()
}

func (*MDIncrementalRefreshSessionStatistics51) NoMDEntriesDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) SbeBlockLength() (blockLength uint) {
	return 24
}

func (*MDIncrementalRefreshSessionStatistics51NoMDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
