// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MDIncrementalRefreshBook46 struct {
	TransactTime        uint64
	MatchEventIndicator MatchEventIndicator
	NoMDEntries         []MDIncrementalRefreshBook46NoMDEntries
	NoOrderIDEntries    []MDIncrementalRefreshBook46NoOrderIDEntries
}
type MDIncrementalRefreshBook46NoMDEntries struct {
	MDEntryPx      PRICENULL9
	MDEntrySize    int32
	SecurityID     int32
	RptSeq         uint32
	NumberOfOrders int32
	MDPriceLevel   uint8
	MDUpdateAction MDUpdateActionEnum
	MDEntryType    MDEntryTypeBookEnum
}
type MDIncrementalRefreshBook46NoOrderIDEntries struct {
	OrderID           uint64
	MDOrderPriority   uint64
	MDDisplayQty      int32
	ReferenceID       uint8
	OrderUpdateAction OrderUpdateActionEnum
}

func (m *MDIncrementalRefreshBook46) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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

		for i := 0; i < 5; i++ {
			if err := _m.WriteUint8(_w, uint8(0)); err != nil {
				return err
			}
		}
	}
	var NoOrderIDEntriesBlockLength uint16 = 24
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

		for i := 0; i < 2; i++ {
			if err := _m.WriteUint8(_w, uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *MDIncrementalRefreshBook46) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
			m.NoMDEntries = make([]MDIncrementalRefreshBook46NoMDEntries, NoMDEntriesNumInGroup)
		}
		m.NoMDEntries = m.NoMDEntries[:NoMDEntriesNumInGroup]
		for i, _ := range m.NoMDEntries {
			if err := m.NoMDEntries[i].Decode(_m, _r, actingVersion, uint(NoMDEntriesBlockLength)); err != nil {
				return err
			}
		}
		io.CopyN(ioutil.Discard, _r, 5)
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
			m.NoOrderIDEntries = make([]MDIncrementalRefreshBook46NoOrderIDEntries, NoOrderIDEntriesNumInGroup)
		}
		m.NoOrderIDEntries = m.NoOrderIDEntries[:NoOrderIDEntriesNumInGroup]
		for i, _ := range m.NoOrderIDEntries {
			if err := m.NoOrderIDEntries[i].Decode(_m, _r, actingVersion, uint(NoOrderIDEntriesBlockLength)); err != nil {
				return err
			}
		}
		io.CopyN(ioutil.Discard, _r, 2)
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MDIncrementalRefreshBook46) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDIncrementalRefreshBook46Init(m *MDIncrementalRefreshBook46) {
	return
}

func (m *MDIncrementalRefreshBook46NoMDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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
	if err := _m.WriteUint8(_w, m.MDPriceLevel); err != nil {
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

func (m *MDIncrementalRefreshBook46NoMDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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
	if !m.MDPriceLevelInActingVersion(actingVersion) {
		m.MDPriceLevel = m.MDPriceLevelNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.MDPriceLevel); err != nil {
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

func (m *MDIncrementalRefreshBook46NoMDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if m.NumberOfOrdersInActingVersion(actingVersion) {
		if m.NumberOfOrders != m.NumberOfOrdersNullValue() && (m.NumberOfOrders < m.NumberOfOrdersMinValue() || m.NumberOfOrders > m.NumberOfOrdersMaxValue()) {
			return fmt.Errorf("Range check failed on m.NumberOfOrders (%v < %v > %v)", m.NumberOfOrdersMinValue(), m.NumberOfOrders, m.NumberOfOrdersMaxValue())
		}
	}
	if m.MDPriceLevelInActingVersion(actingVersion) {
		if m.MDPriceLevel < m.MDPriceLevelMinValue() || m.MDPriceLevel > m.MDPriceLevelMaxValue() {
			return fmt.Errorf("Range check failed on m.MDPriceLevel (%v < %v > %v)", m.MDPriceLevelMinValue(), m.MDPriceLevel, m.MDPriceLevelMaxValue())
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

func MDIncrementalRefreshBook46NoMDEntriesInit(m *MDIncrementalRefreshBook46NoMDEntries) {
	m.MDEntrySize = 2147483647
	m.NumberOfOrders = 2147483647
	return
}

func (m *MDIncrementalRefreshBook46NoOrderIDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint64(_w, m.OrderID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.MDOrderPriority); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.MDDisplayQty); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.ReferenceID); err != nil {
		return err
	}
	if err := m.OrderUpdateAction.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MDIncrementalRefreshBook46NoOrderIDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !m.OrderIDInActingVersion(actingVersion) {
		m.OrderID = m.OrderIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.OrderID); err != nil {
			return err
		}
	}
	if !m.MDOrderPriorityInActingVersion(actingVersion) {
		m.MDOrderPriority = m.MDOrderPriorityNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.MDOrderPriority); err != nil {
			return err
		}
	}
	if !m.MDDisplayQtyInActingVersion(actingVersion) {
		m.MDDisplayQty = m.MDDisplayQtyNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.MDDisplayQty); err != nil {
			return err
		}
	}
	if !m.ReferenceIDInActingVersion(actingVersion) {
		m.ReferenceID = m.ReferenceIDNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.ReferenceID); err != nil {
			return err
		}
	}
	if m.OrderUpdateActionInActingVersion(actingVersion) {
		if err := m.OrderUpdateAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDIncrementalRefreshBook46NoOrderIDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.OrderIDInActingVersion(actingVersion) {
		if m.OrderID < m.OrderIDMinValue() || m.OrderID > m.OrderIDMaxValue() {
			return fmt.Errorf("Range check failed on m.OrderID (%v < %v > %v)", m.OrderIDMinValue(), m.OrderID, m.OrderIDMaxValue())
		}
	}
	if m.MDOrderPriorityInActingVersion(actingVersion) {
		if m.MDOrderPriority != m.MDOrderPriorityNullValue() && (m.MDOrderPriority < m.MDOrderPriorityMinValue() || m.MDOrderPriority > m.MDOrderPriorityMaxValue()) {
			return fmt.Errorf("Range check failed on m.MDOrderPriority (%v < %v > %v)", m.MDOrderPriorityMinValue(), m.MDOrderPriority, m.MDOrderPriorityMaxValue())
		}
	}
	if m.MDDisplayQtyInActingVersion(actingVersion) {
		if m.MDDisplayQty != m.MDDisplayQtyNullValue() && (m.MDDisplayQty < m.MDDisplayQtyMinValue() || m.MDDisplayQty > m.MDDisplayQtyMaxValue()) {
			return fmt.Errorf("Range check failed on m.MDDisplayQty (%v < %v > %v)", m.MDDisplayQtyMinValue(), m.MDDisplayQty, m.MDDisplayQtyMaxValue())
		}
	}
	if m.ReferenceIDInActingVersion(actingVersion) {
		if m.ReferenceID != m.ReferenceIDNullValue() && (m.ReferenceID < m.ReferenceIDMinValue() || m.ReferenceID > m.ReferenceIDMaxValue()) {
			return fmt.Errorf("Range check failed on m.ReferenceID (%v < %v > %v)", m.ReferenceIDMinValue(), m.ReferenceID, m.ReferenceIDMaxValue())
		}
	}
	if err := m.OrderUpdateAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func MDIncrementalRefreshBook46NoOrderIDEntriesInit(m *MDIncrementalRefreshBook46NoOrderIDEntries) {
	m.MDOrderPriority = 18446744073709551615
	m.MDDisplayQty = 2147483647
	m.ReferenceID = 255
	return
}

func (*MDIncrementalRefreshBook46) SbeBlockLength() (blockLength uint16) {
	return 11
}

func (*MDIncrementalRefreshBook46) SbeTemplateId() (templateId uint16) {
	return 46
}

func (*MDIncrementalRefreshBook46) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MDIncrementalRefreshBook46) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDIncrementalRefreshBook46) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MDIncrementalRefreshBook46) TransactTimeId() uint16 {
	return 60
}

func (*MDIncrementalRefreshBook46) TransactTimeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TransactTimeSinceVersion()
}

func (*MDIncrementalRefreshBook46) TransactTimeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46) TransactTimeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46) TransactTimeMinValue() uint64 {
	return 0
}

func (*MDIncrementalRefreshBook46) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDIncrementalRefreshBook46) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDIncrementalRefreshBook46) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MDIncrementalRefreshBook46) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MDIncrementalRefreshBook46) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntryPxId() uint16 {
	return 270
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntryPxSinceVersion() uint16 {
	return 9
}

func (m *MDIncrementalRefreshBook46NoMDEntries) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryPxSinceVersion()
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntryPxMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntrySizeId() uint16 {
	return 271
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46NoMDEntries) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntrySizeSinceVersion()
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntrySizeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntrySizeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntrySizeMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntrySizeNullValue() int32 {
	return 2147483647
}

func (*MDIncrementalRefreshBook46NoMDEntries) SecurityIDId() uint16 {
	return 48
}

func (*MDIncrementalRefreshBook46NoMDEntries) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46NoMDEntries) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MDIncrementalRefreshBook46NoMDEntries) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoMDEntries) SecurityIDMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46NoMDEntries) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshBook46NoMDEntries) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshBook46NoMDEntries) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDIncrementalRefreshBook46NoMDEntries) RptSeqId() uint16 {
	return 83
}

func (*MDIncrementalRefreshBook46NoMDEntries) RptSeqSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46NoMDEntries) RptSeqInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RptSeqSinceVersion()
}

func (*MDIncrementalRefreshBook46NoMDEntries) RptSeqDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoMDEntries) RptSeqMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46NoMDEntries) RptSeqMinValue() uint32 {
	return 0
}

func (*MDIncrementalRefreshBook46NoMDEntries) RptSeqMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDIncrementalRefreshBook46NoMDEntries) RptSeqNullValue() uint32 {
	return math.MaxUint32
}

func (*MDIncrementalRefreshBook46NoMDEntries) NumberOfOrdersId() uint16 {
	return 346
}

func (*MDIncrementalRefreshBook46NoMDEntries) NumberOfOrdersSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46NoMDEntries) NumberOfOrdersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NumberOfOrdersSinceVersion()
}

func (*MDIncrementalRefreshBook46NoMDEntries) NumberOfOrdersDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoMDEntries) NumberOfOrdersMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46NoMDEntries) NumberOfOrdersMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshBook46NoMDEntries) NumberOfOrdersMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshBook46NoMDEntries) NumberOfOrdersNullValue() int32 {
	return 2147483647
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDPriceLevelId() uint16 {
	return 1023
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDPriceLevelSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46NoMDEntries) MDPriceLevelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDPriceLevelSinceVersion()
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDPriceLevelDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDPriceLevelMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46NoMDEntries) MDPriceLevelMinValue() uint8 {
	return 0
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDPriceLevelMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDPriceLevelNullValue() uint8 {
	return math.MaxUint8
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDUpdateActionId() uint16 {
	return 279
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46NoMDEntries) MDUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDUpdateActionSinceVersion()
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDUpdateActionDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDUpdateActionMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntryTypeId() uint16 {
	return 269
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntryTypeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46NoMDEntries) MDEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryTypeSinceVersion()
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntryTypeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoMDEntries) MDEntryTypeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46NoOrderIDEntries) OrderIDId() uint16 {
	return 37
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) OrderIDSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46NoOrderIDEntries) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OrderIDSinceVersion()
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) OrderIDDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) OrderIDMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46NoOrderIDEntries) OrderIDMinValue() uint64 {
	return 0
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) OrderIDNullValue() uint64 {
	return math.MaxUint64
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDOrderPriorityId() uint16 {
	return 37707
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDOrderPrioritySinceVersion() uint16 {
	return 7
}

func (m *MDIncrementalRefreshBook46NoOrderIDEntries) MDOrderPriorityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDOrderPrioritySinceVersion()
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDOrderPriorityDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDOrderPriorityMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDOrderPriorityMinValue() uint64 {
	return 0
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDOrderPriorityMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDOrderPriorityNullValue() uint64 {
	return 18446744073709551615
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDDisplayQtyId() uint16 {
	return 37706
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDDisplayQtySinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46NoOrderIDEntries) MDDisplayQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDDisplayQtySinceVersion()
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDDisplayQtyDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDDisplayQtyMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDDisplayQtyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDDisplayQtyMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) MDDisplayQtyNullValue() int32 {
	return 2147483647
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) ReferenceIDId() uint16 {
	return 9633
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) ReferenceIDSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46NoOrderIDEntries) ReferenceIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ReferenceIDSinceVersion()
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) ReferenceIDDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) ReferenceIDMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46NoOrderIDEntries) ReferenceIDMinValue() uint8 {
	return 0
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) ReferenceIDMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) ReferenceIDNullValue() uint8 {
	return 255
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) OrderUpdateActionId() uint16 {
	return 37708
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) OrderUpdateActionSinceVersion() uint16 {
	return 7
}

func (m *MDIncrementalRefreshBook46NoOrderIDEntries) OrderUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OrderUpdateActionSinceVersion()
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) OrderUpdateActionDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) OrderUpdateActionMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshBook46) NoMDEntriesId() uint16 {
	return 268
}

func (*MDIncrementalRefreshBook46) NoMDEntriesSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46) NoMDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoMDEntriesSinceVersion()
}

func (*MDIncrementalRefreshBook46) NoMDEntriesDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoMDEntries) SbeBlockLength() (blockLength uint) {
	return 32
}

func (*MDIncrementalRefreshBook46NoMDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDIncrementalRefreshBook46) NoOrderIDEntriesId() uint16 {
	return 37705
}

func (*MDIncrementalRefreshBook46) NoOrderIDEntriesSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshBook46) NoOrderIDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoOrderIDEntriesSinceVersion()
}

func (*MDIncrementalRefreshBook46) NoOrderIDEntriesDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) SbeBlockLength() (blockLength uint) {
	return 24
}

func (*MDIncrementalRefreshBook46NoOrderIDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
