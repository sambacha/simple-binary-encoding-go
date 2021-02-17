// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MDIncrementalRefreshOrderBook47 struct {
	TransactTime        uint64
	MatchEventIndicator MatchEventIndicator
	NoMDEntries         []MDIncrementalRefreshOrderBook47NoMDEntries
}
type MDIncrementalRefreshOrderBook47NoMDEntries struct {
	OrderID         uint64
	MDOrderPriority uint64
	MDEntryPx       PRICENULL9
	MDDisplayQty    int32
	SecurityID      int32
	MDUpdateAction  MDUpdateActionEnum
	MDEntryType     MDEntryTypeBookEnum
}

func (m *MDIncrementalRefreshOrderBook47) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	var NoMDEntriesBlockLength uint16 = 40
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

		for i := 0; i < 6; i++ {
			if err := _m.WriteUint8(_w, uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *MDIncrementalRefreshOrderBook47) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
			m.NoMDEntries = make([]MDIncrementalRefreshOrderBook47NoMDEntries, NoMDEntriesNumInGroup)
		}
		m.NoMDEntries = m.NoMDEntries[:NoMDEntriesNumInGroup]
		for i, _ := range m.NoMDEntries {
			if err := m.NoMDEntries[i].Decode(_m, _r, actingVersion, uint(NoMDEntriesBlockLength)); err != nil {
				return err
			}
		}
		io.CopyN(ioutil.Discard, _r, 6)
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MDIncrementalRefreshOrderBook47) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDIncrementalRefreshOrderBook47Init(m *MDIncrementalRefreshOrderBook47) {
	return
}

func (m *MDIncrementalRefreshOrderBook47NoMDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint64(_w, m.OrderID); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.MDOrderPriority); err != nil {
		return err
	}
	if err := m.MDEntryPx.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.MDDisplayQty); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.SecurityID); err != nil {
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

func (m *MDIncrementalRefreshOrderBook47NoMDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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
	if m.MDEntryPxInActingVersion(actingVersion) {
		if err := m.MDEntryPx.Decode(_m, _r, actingVersion); err != nil {
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
	if !m.SecurityIDInActingVersion(actingVersion) {
		m.SecurityID = m.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.SecurityID); err != nil {
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

func (m *MDIncrementalRefreshOrderBook47NoMDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.OrderIDInActingVersion(actingVersion) {
		if m.OrderID != m.OrderIDNullValue() && (m.OrderID < m.OrderIDMinValue() || m.OrderID > m.OrderIDMaxValue()) {
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
	if m.SecurityIDInActingVersion(actingVersion) {
		if m.SecurityID < m.SecurityIDMinValue() || m.SecurityID > m.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on m.SecurityID (%v < %v > %v)", m.SecurityIDMinValue(), m.SecurityID, m.SecurityIDMaxValue())
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

func MDIncrementalRefreshOrderBook47NoMDEntriesInit(m *MDIncrementalRefreshOrderBook47NoMDEntries) {
	m.OrderID = 18446744073709551615
	m.MDOrderPriority = 18446744073709551615
	m.MDDisplayQty = 2147483647
	return
}

func (*MDIncrementalRefreshOrderBook47) SbeBlockLength() (blockLength uint16) {
	return 11
}

func (*MDIncrementalRefreshOrderBook47) SbeTemplateId() (templateId uint16) {
	return 47
}

func (*MDIncrementalRefreshOrderBook47) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MDIncrementalRefreshOrderBook47) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDIncrementalRefreshOrderBook47) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MDIncrementalRefreshOrderBook47) TransactTimeId() uint16 {
	return 60
}

func (*MDIncrementalRefreshOrderBook47) TransactTimeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshOrderBook47) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TransactTimeSinceVersion()
}

func (*MDIncrementalRefreshOrderBook47) TransactTimeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshOrderBook47) TransactTimeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshOrderBook47) TransactTimeMinValue() uint64 {
	return 0
}

func (*MDIncrementalRefreshOrderBook47) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDIncrementalRefreshOrderBook47) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDIncrementalRefreshOrderBook47) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MDIncrementalRefreshOrderBook47) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshOrderBook47) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MDIncrementalRefreshOrderBook47) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshOrderBook47) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshOrderBook47NoMDEntries) OrderIDId() uint16 {
	return 37
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) OrderIDSinceVersion() uint16 {
	return 7
}

func (m *MDIncrementalRefreshOrderBook47NoMDEntries) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OrderIDSinceVersion()
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) OrderIDDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) OrderIDMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshOrderBook47NoMDEntries) OrderIDMinValue() uint64 {
	return 0
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) OrderIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) OrderIDNullValue() uint64 {
	return 18446744073709551615
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDOrderPriorityId() uint16 {
	return 37707
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDOrderPrioritySinceVersion() uint16 {
	return 7
}

func (m *MDIncrementalRefreshOrderBook47NoMDEntries) MDOrderPriorityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDOrderPrioritySinceVersion()
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDOrderPriorityDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDOrderPriorityMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDOrderPriorityMinValue() uint64 {
	return 0
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDOrderPriorityMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDOrderPriorityNullValue() uint64 {
	return 18446744073709551615
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDEntryPxId() uint16 {
	return 270
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDEntryPxSinceVersion() uint16 {
	return 9
}

func (m *MDIncrementalRefreshOrderBook47NoMDEntries) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryPxSinceVersion()
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDEntryPxMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDDisplayQtyId() uint16 {
	return 37706
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDDisplayQtySinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshOrderBook47NoMDEntries) MDDisplayQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDDisplayQtySinceVersion()
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDDisplayQtyDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDDisplayQtyMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDDisplayQtyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDDisplayQtyMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDDisplayQtyNullValue() int32 {
	return 2147483647
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) SecurityIDId() uint16 {
	return 48
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshOrderBook47NoMDEntries) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) SecurityIDMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshOrderBook47NoMDEntries) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDUpdateActionId() uint16 {
	return 279
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshOrderBook47NoMDEntries) MDUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDUpdateActionSinceVersion()
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDUpdateActionDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDUpdateActionMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDEntryTypeId() uint16 {
	return 269
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDEntryTypeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshOrderBook47NoMDEntries) MDEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryTypeSinceVersion()
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDEntryTypeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) MDEntryTypeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshOrderBook47) NoMDEntriesId() uint16 {
	return 268
}

func (*MDIncrementalRefreshOrderBook47) NoMDEntriesSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshOrderBook47) NoMDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoMDEntriesSinceVersion()
}

func (*MDIncrementalRefreshOrderBook47) NoMDEntriesDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) SbeBlockLength() (blockLength uint) {
	return 40
}

func (*MDIncrementalRefreshOrderBook47NoMDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
