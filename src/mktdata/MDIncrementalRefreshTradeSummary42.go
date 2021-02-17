// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MDIncrementalRefreshTradeSummary42 struct {
	TransactTime        uint64
	MatchEventIndicator MatchEventIndicator
	NoMDEntries         []MDIncrementalRefreshTradeSummary42NoMDEntries
	NoOrderIDEntries    []MDIncrementalRefreshTradeSummary42NoOrderIDEntries
}
type MDIncrementalRefreshTradeSummary42NoMDEntries struct {
	MDEntryPx      PRICE
	MDEntrySize    int32
	SecurityID     int32
	RptSeq         uint32
	NumberOfOrders int32
	AggressorSide  AggressorSideEnum
	MDUpdateAction MDUpdateActionEnum
	MDEntryType    [1]byte
	MDTradeEntryID uint32
}
type MDIncrementalRefreshTradeSummary42NoOrderIDEntries struct {
	OrderID uint64
	LastQty int32
}

func (m *MDIncrementalRefreshTradeSummary42) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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

		for i := 0; i < 2; i++ {
			if err := _m.WriteUint8(_w, uint8(0)); err != nil {
				return err
			}
		}
	}
	var NoOrderIDEntriesBlockLength uint16 = 16
	if err := _m.WriteUint16(_w, NoOrderIDEntriesBlockLength); err != nil {
		return err
	}
	var NoOrderIDEntriesNumInGroup uint8 = uint8(len(m.NoOrderIDEntries))
	if err := _m.WriteUint8(_w, NoOrderIDEntriesNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.NoOrderIDEntries {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}

		for i := 0; i < 4; i++ {
			if err := _m.WriteUint8(_w, uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *MDIncrementalRefreshTradeSummary42) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
			m.NoMDEntries = make([]MDIncrementalRefreshTradeSummary42NoMDEntries, NoMDEntriesNumInGroup)
		}
		m.NoMDEntries = m.NoMDEntries[:NoMDEntriesNumInGroup]
		for i, _ := range m.NoMDEntries {
			if err := m.NoMDEntries[i].Decode(_m, _r, actingVersion, uint(NoMDEntriesBlockLength)); err != nil {
				return err
			}
		}
		io.CopyN(ioutil.Discard, _r, 2)
	}

	if m.NoOrderIDEntriesInActingVersion(actingVersion) {
		var NoOrderIDEntriesBlockLength uint16
		if err := _m.ReadUint16(_r, &NoOrderIDEntriesBlockLength); err != nil {
			return err
		}
		var NoOrderIDEntriesNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoOrderIDEntriesNumInGroup); err != nil {
			return err
		}
		if cap(m.NoOrderIDEntries) < int(NoOrderIDEntriesNumInGroup) {
			m.NoOrderIDEntries = make([]MDIncrementalRefreshTradeSummary42NoOrderIDEntries, NoOrderIDEntriesNumInGroup)
		}
		m.NoOrderIDEntries = m.NoOrderIDEntries[:NoOrderIDEntriesNumInGroup]
		for i, _ := range m.NoOrderIDEntries {
			if err := m.NoOrderIDEntries[i].Decode(_m, _r, actingVersion, uint(NoOrderIDEntriesBlockLength)); err != nil {
				return err
			}
		}
		io.CopyN(ioutil.Discard, _r, 4)
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MDIncrementalRefreshTradeSummary42) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	for _, prop := range m.NoOrderIDEntries {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func MDIncrementalRefreshTradeSummary42Init(m *MDIncrementalRefreshTradeSummary42) {
	return
}

func (m *MDIncrementalRefreshTradeSummary42NoMDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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
	if err := _m.WriteInt32(_w, m.NumberOfOrders); err != nil {
		return err
	}
	if err := m.AggressorSide.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MDUpdateAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.MDTradeEntryID); err != nil {
		return err
	}
	return nil
}

func (m *MDIncrementalRefreshTradeSummary42NoMDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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
	if !m.NumberOfOrdersInActingVersion(actingVersion) {
		m.NumberOfOrders = m.NumberOfOrdersNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.NumberOfOrders); err != nil {
			return err
		}
	}
	if m.AggressorSideInActingVersion(actingVersion) {
		if err := m.AggressorSide.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.MDUpdateActionInActingVersion(actingVersion) {
		if err := m.MDUpdateAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	m.MDEntryType[0] = 50
	if !m.MDTradeEntryIDInActingVersion(actingVersion) {
		m.MDTradeEntryID = m.MDTradeEntryIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &m.MDTradeEntryID); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDIncrementalRefreshTradeSummary42NoMDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if m.NumberOfOrdersInActingVersion(actingVersion) {
		if m.NumberOfOrders != m.NumberOfOrdersNullValue() && (m.NumberOfOrders < m.NumberOfOrdersMinValue() || m.NumberOfOrders > m.NumberOfOrdersMaxValue()) {
			return fmt.Errorf("Range check failed on m.NumberOfOrders (%v < %v > %v)", m.NumberOfOrdersMinValue(), m.NumberOfOrders, m.NumberOfOrdersMaxValue())
		}
	}
	if err := m.AggressorSide.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.MDUpdateAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.MDTradeEntryIDInActingVersion(actingVersion) {
		if m.MDTradeEntryID != m.MDTradeEntryIDNullValue() && (m.MDTradeEntryID < m.MDTradeEntryIDMinValue() || m.MDTradeEntryID > m.MDTradeEntryIDMaxValue()) {
			return fmt.Errorf("Range check failed on m.MDTradeEntryID (%v < %v > %v)", m.MDTradeEntryIDMinValue(), m.MDTradeEntryID, m.MDTradeEntryIDMaxValue())
		}
	}
	return nil
}

func MDIncrementalRefreshTradeSummary42NoMDEntriesInit(m *MDIncrementalRefreshTradeSummary42NoMDEntries) {
	m.NumberOfOrders = 2147483647
	m.MDEntryType[0] = 50
	m.MDTradeEntryID = 4294967295
	return
}

func (m *MDIncrementalRefreshTradeSummary42NoOrderIDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint64(_w, m.OrderID); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.LastQty); err != nil {
		return err
	}
	return nil
}

func (m *MDIncrementalRefreshTradeSummary42NoOrderIDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !m.OrderIDInActingVersion(actingVersion) {
		m.OrderID = m.OrderIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.OrderID); err != nil {
			return err
		}
	}
	if !m.LastQtyInActingVersion(actingVersion) {
		m.LastQty = m.LastQtyNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.LastQty); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDIncrementalRefreshTradeSummary42NoOrderIDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.OrderIDInActingVersion(actingVersion) {
		if m.OrderID < m.OrderIDMinValue() || m.OrderID > m.OrderIDMaxValue() {
			return fmt.Errorf("Range check failed on m.OrderID (%v < %v > %v)", m.OrderIDMinValue(), m.OrderID, m.OrderIDMaxValue())
		}
	}
	if m.LastQtyInActingVersion(actingVersion) {
		if m.LastQty < m.LastQtyMinValue() || m.LastQty > m.LastQtyMaxValue() {
			return fmt.Errorf("Range check failed on m.LastQty (%v < %v > %v)", m.LastQtyMinValue(), m.LastQty, m.LastQtyMaxValue())
		}
	}
	return nil
}

func MDIncrementalRefreshTradeSummary42NoOrderIDEntriesInit(m *MDIncrementalRefreshTradeSummary42NoOrderIDEntries) {
	return
}

func (*MDIncrementalRefreshTradeSummary42) SbeBlockLength() (blockLength uint16) {
	return 11
}

func (*MDIncrementalRefreshTradeSummary42) SbeTemplateId() (templateId uint16) {
	return 42
}

func (*MDIncrementalRefreshTradeSummary42) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MDIncrementalRefreshTradeSummary42) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDIncrementalRefreshTradeSummary42) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MDIncrementalRefreshTradeSummary42) TransactTimeId() uint16 {
	return 60
}

func (*MDIncrementalRefreshTradeSummary42) TransactTimeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TransactTimeSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42) TransactTimeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42) TransactTimeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshTradeSummary42) TransactTimeMinValue() uint64 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDIncrementalRefreshTradeSummary42) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDIncrementalRefreshTradeSummary42) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MDIncrementalRefreshTradeSummary42) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntryPxId() uint16 {
	return 270
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntryPxSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryPxSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntryPxMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntrySizeId() uint16 {
	return 271
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntrySizeSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntrySizeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntrySizeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntrySizeMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntrySizeNullValue() int32 {
	return math.MinInt32
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) SecurityIDId() uint16 {
	return 48
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42NoMDEntries) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) SecurityIDMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) RptSeqId() uint16 {
	return 83
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) RptSeqSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42NoMDEntries) RptSeqInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RptSeqSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) RptSeqDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) RptSeqMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) RptSeqMinValue() uint32 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) RptSeqMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) RptSeqNullValue() uint32 {
	return math.MaxUint32
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) NumberOfOrdersId() uint16 {
	return 346
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) NumberOfOrdersSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42NoMDEntries) NumberOfOrdersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NumberOfOrdersSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) NumberOfOrdersDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) NumberOfOrdersMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) NumberOfOrdersMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) NumberOfOrdersMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) NumberOfOrdersNullValue() int32 {
	return 2147483647
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) AggressorSideId() uint16 {
	return 5797
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) AggressorSideSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42NoMDEntries) AggressorSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AggressorSideSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) AggressorSideDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) AggressorSideMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDUpdateActionId() uint16 {
	return 279
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42NoMDEntries) MDUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDUpdateActionSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDUpdateActionDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDUpdateActionMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntryTypeId() uint16 {
	return 269
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntryTypeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryTypeSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntryTypeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntryTypeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntryTypeMinValue() byte {
	return byte(32)
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntryTypeMaxValue() byte {
	return byte(126)
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDEntryTypeNullValue() byte {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDTradeEntryIDId() uint16 {
	return 37711
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDTradeEntryIDSinceVersion() uint16 {
	return 7
}

func (m *MDIncrementalRefreshTradeSummary42NoMDEntries) MDTradeEntryIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDTradeEntryIDSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDTradeEntryIDDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDTradeEntryIDMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDTradeEntryIDMinValue() uint32 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDTradeEntryIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) MDTradeEntryIDNullValue() uint32 {
	return 4294967295
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) OrderIDId() uint16 {
	return 37
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) OrderIDSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42NoOrderIDEntries) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OrderIDSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) OrderIDDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) OrderIDMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) OrderIDMinValue() uint64 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) LastQtyId() uint16 {
	return 32
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) LastQtySinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42NoOrderIDEntries) LastQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastQtySinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) LastQtyDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) LastQtyMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) LastQtyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) LastQtyMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) LastQtyNullValue() int32 {
	return math.MinInt32
}

func (*MDIncrementalRefreshTradeSummary42) NoMDEntriesId() uint16 {
	return 268
}

func (*MDIncrementalRefreshTradeSummary42) NoMDEntriesSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42) NoMDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoMDEntriesSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42) NoMDEntriesDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) SbeBlockLength() (blockLength uint) {
	return 32
}

func (*MDIncrementalRefreshTradeSummary42NoMDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDIncrementalRefreshTradeSummary42) NoOrderIDEntriesId() uint16 {
	return 37705
}

func (*MDIncrementalRefreshTradeSummary42) NoOrderIDEntriesSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshTradeSummary42) NoOrderIDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoOrderIDEntriesSinceVersion()
}

func (*MDIncrementalRefreshTradeSummary42) NoOrderIDEntriesDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) SbeBlockLength() (blockLength uint) {
	return 16
}

func (*MDIncrementalRefreshTradeSummary42NoOrderIDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
