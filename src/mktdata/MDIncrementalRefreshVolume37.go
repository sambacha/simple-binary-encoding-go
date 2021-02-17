// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MDIncrementalRefreshVolume37 struct {
	TransactTime        uint64
	MatchEventIndicator MatchEventIndicator
	NoMDEntries         []MDIncrementalRefreshVolume37NoMDEntries
}
type MDIncrementalRefreshVolume37NoMDEntries struct {
	MDEntrySize    int32
	SecurityID     int32
	RptSeq         uint32
	MDUpdateAction MDUpdateActionEnum
	MDEntryType    [1]byte
}

func (m *MDIncrementalRefreshVolume37) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	var NoMDEntriesBlockLength uint16 = 16
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

		for i := 0; i < 3; i++ {
			if err := _m.WriteUint8(_w, uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *MDIncrementalRefreshVolume37) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
			m.NoMDEntries = make([]MDIncrementalRefreshVolume37NoMDEntries, NoMDEntriesNumInGroup)
		}
		m.NoMDEntries = m.NoMDEntries[:NoMDEntriesNumInGroup]
		for i, _ := range m.NoMDEntries {
			if err := m.NoMDEntries[i].Decode(_m, _r, actingVersion, uint(NoMDEntriesBlockLength)); err != nil {
				return err
			}
		}
		io.CopyN(ioutil.Discard, _r, 3)
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MDIncrementalRefreshVolume37) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDIncrementalRefreshVolume37Init(m *MDIncrementalRefreshVolume37) {
	return
}

func (m *MDIncrementalRefreshVolume37NoMDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt32(_w, m.MDEntrySize); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.RptSeq); err != nil {
		return err
	}
	if err := m.MDUpdateAction.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MDIncrementalRefreshVolume37NoMDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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
	if m.MDUpdateActionInActingVersion(actingVersion) {
		if err := m.MDUpdateAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	m.MDEntryType[0] = 101
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDIncrementalRefreshVolume37NoMDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.MDEntrySizeInActingVersion(actingVersion) {
		if m.MDEntrySize < m.MDEntrySizeMinValue() || m.MDEntrySize > m.MDEntrySizeMaxValue() {
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
	if err := m.MDUpdateAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func MDIncrementalRefreshVolume37NoMDEntriesInit(m *MDIncrementalRefreshVolume37NoMDEntries) {
	m.MDEntryType[0] = 101
	return
}

func (*MDIncrementalRefreshVolume37) SbeBlockLength() (blockLength uint16) {
	return 11
}

func (*MDIncrementalRefreshVolume37) SbeTemplateId() (templateId uint16) {
	return 37
}

func (*MDIncrementalRefreshVolume37) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MDIncrementalRefreshVolume37) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDIncrementalRefreshVolume37) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MDIncrementalRefreshVolume37) TransactTimeId() uint16 {
	return 60
}

func (*MDIncrementalRefreshVolume37) TransactTimeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshVolume37) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TransactTimeSinceVersion()
}

func (*MDIncrementalRefreshVolume37) TransactTimeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshVolume37) TransactTimeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshVolume37) TransactTimeMinValue() uint64 {
	return 0
}

func (*MDIncrementalRefreshVolume37) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDIncrementalRefreshVolume37) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDIncrementalRefreshVolume37) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MDIncrementalRefreshVolume37) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshVolume37) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MDIncrementalRefreshVolume37) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshVolume37) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntrySizeId() uint16 {
	return 271
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshVolume37NoMDEntries) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntrySizeSinceVersion()
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntrySizeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntrySizeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntrySizeMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntrySizeNullValue() int32 {
	return math.MinInt32
}

func (*MDIncrementalRefreshVolume37NoMDEntries) SecurityIDId() uint16 {
	return 48
}

func (*MDIncrementalRefreshVolume37NoMDEntries) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshVolume37NoMDEntries) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MDIncrementalRefreshVolume37NoMDEntries) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshVolume37NoMDEntries) SecurityIDMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshVolume37NoMDEntries) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshVolume37NoMDEntries) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshVolume37NoMDEntries) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDIncrementalRefreshVolume37NoMDEntries) RptSeqId() uint16 {
	return 83
}

func (*MDIncrementalRefreshVolume37NoMDEntries) RptSeqSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshVolume37NoMDEntries) RptSeqInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RptSeqSinceVersion()
}

func (*MDIncrementalRefreshVolume37NoMDEntries) RptSeqDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshVolume37NoMDEntries) RptSeqMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshVolume37NoMDEntries) RptSeqMinValue() uint32 {
	return 0
}

func (*MDIncrementalRefreshVolume37NoMDEntries) RptSeqMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDIncrementalRefreshVolume37NoMDEntries) RptSeqNullValue() uint32 {
	return math.MaxUint32
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDUpdateActionId() uint16 {
	return 279
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshVolume37NoMDEntries) MDUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDUpdateActionSinceVersion()
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDUpdateActionDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDUpdateActionMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntryTypeId() uint16 {
	return 269
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntryTypeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshVolume37NoMDEntries) MDEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryTypeSinceVersion()
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntryTypeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntryTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "char"
	case 4:
		return "constant"
	}
	return ""
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntryTypeMinValue() byte {
	return byte(32)
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntryTypeMaxValue() byte {
	return byte(126)
}

func (*MDIncrementalRefreshVolume37NoMDEntries) MDEntryTypeNullValue() byte {
	return 0
}

func (*MDIncrementalRefreshVolume37) NoMDEntriesId() uint16 {
	return 268
}

func (*MDIncrementalRefreshVolume37) NoMDEntriesSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshVolume37) NoMDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoMDEntriesSinceVersion()
}

func (*MDIncrementalRefreshVolume37) NoMDEntriesDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshVolume37NoMDEntries) SbeBlockLength() (blockLength uint) {
	return 16
}

func (*MDIncrementalRefreshVolume37NoMDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
