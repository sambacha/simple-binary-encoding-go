// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MDIncrementalRefreshLimitsBanding34 struct {
	TransactTime        uint64
	MatchEventIndicator MatchEventIndicator
	NoMDEntries         []MDIncrementalRefreshLimitsBanding34NoMDEntries
}
type MDIncrementalRefreshLimitsBanding34NoMDEntries struct {
	HighLimitPrice    PRICENULL
	LowLimitPrice     PRICENULL
	MaxPriceVariation PRICENULL
	SecurityID        int32
	RptSeq            uint32
	MDUpdateAction    int8
	MDEntryType       [1]byte
}

func (m *MDIncrementalRefreshLimitsBanding34) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	}
	return nil
}

func (m *MDIncrementalRefreshLimitsBanding34) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
			m.NoMDEntries = make([]MDIncrementalRefreshLimitsBanding34NoMDEntries, NoMDEntriesNumInGroup)
		}
		m.NoMDEntries = m.NoMDEntries[:NoMDEntriesNumInGroup]
		for i, _ := range m.NoMDEntries {
			if err := m.NoMDEntries[i].Decode(_m, _r, actingVersion, uint(NoMDEntriesBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MDIncrementalRefreshLimitsBanding34) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDIncrementalRefreshLimitsBanding34Init(m *MDIncrementalRefreshLimitsBanding34) {
	return
}

func (m *MDIncrementalRefreshLimitsBanding34NoMDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.HighLimitPrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.LowLimitPrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MaxPriceVariation.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.RptSeq); err != nil {
		return err
	}
	return nil
}

func (m *MDIncrementalRefreshLimitsBanding34NoMDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if m.HighLimitPriceInActingVersion(actingVersion) {
		if err := m.HighLimitPrice.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.LowLimitPriceInActingVersion(actingVersion) {
		if err := m.LowLimitPrice.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.MaxPriceVariationInActingVersion(actingVersion) {
		if err := m.MaxPriceVariation.Decode(_m, _r, actingVersion); err != nil {
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
	m.MDUpdateAction = 0
	m.MDEntryType[0] = 103
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDIncrementalRefreshLimitsBanding34NoMDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	return nil
}

func MDIncrementalRefreshLimitsBanding34NoMDEntriesInit(m *MDIncrementalRefreshLimitsBanding34NoMDEntries) {
	m.MDUpdateAction = 0
	m.MDEntryType[0] = 103
	return
}

func (*MDIncrementalRefreshLimitsBanding34) SbeBlockLength() (blockLength uint16) {
	return 11
}

func (*MDIncrementalRefreshLimitsBanding34) SbeTemplateId() (templateId uint16) {
	return 34
}

func (*MDIncrementalRefreshLimitsBanding34) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MDIncrementalRefreshLimitsBanding34) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDIncrementalRefreshLimitsBanding34) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MDIncrementalRefreshLimitsBanding34) TransactTimeId() uint16 {
	return 60
}

func (*MDIncrementalRefreshLimitsBanding34) TransactTimeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding34) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TransactTimeSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding34) TransactTimeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding34) TransactTimeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding34) TransactTimeMinValue() uint64 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding34) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDIncrementalRefreshLimitsBanding34) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDIncrementalRefreshLimitsBanding34) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MDIncrementalRefreshLimitsBanding34) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding34) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding34) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding34) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) HighLimitPriceId() uint16 {
	return 1149
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) HighLimitPriceSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding34NoMDEntries) HighLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.HighLimitPriceSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) HighLimitPriceDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) HighLimitPriceMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) LowLimitPriceId() uint16 {
	return 1148
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) LowLimitPriceSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding34NoMDEntries) LowLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LowLimitPriceSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) LowLimitPriceDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) LowLimitPriceMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MaxPriceVariationId() uint16 {
	return 1143
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MaxPriceVariationSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding34NoMDEntries) MaxPriceVariationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaxPriceVariationSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MaxPriceVariationDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MaxPriceVariationMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) SecurityIDId() uint16 {
	return 48
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding34NoMDEntries) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) SecurityIDMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) RptSeqId() uint16 {
	return 83
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) RptSeqSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding34NoMDEntries) RptSeqInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RptSeqSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) RptSeqDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) RptSeqMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) RptSeqMinValue() uint32 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) RptSeqMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) RptSeqNullValue() uint32 {
	return math.MaxUint32
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDUpdateActionId() uint16 {
	return 279
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding34NoMDEntries) MDUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDUpdateActionSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDUpdateActionDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDUpdateActionMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "constant"
	}
	return ""
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDUpdateActionMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDUpdateActionMaxValue() int8 {
	return math.MaxInt8
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDUpdateActionNullValue() int8 {
	return math.MinInt8
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDEntryTypeId() uint16 {
	return 269
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDEntryTypeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding34NoMDEntries) MDEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryTypeSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDEntryTypeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDEntryTypeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDEntryTypeMinValue() byte {
	return byte(32)
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDEntryTypeMaxValue() byte {
	return byte(126)
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) MDEntryTypeNullValue() byte {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding34) NoMDEntriesId() uint16 {
	return 268
}

func (*MDIncrementalRefreshLimitsBanding34) NoMDEntriesSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding34) NoMDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoMDEntriesSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding34) NoMDEntriesDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) SbeBlockLength() (blockLength uint) {
	return 32
}

func (*MDIncrementalRefreshLimitsBanding34NoMDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
