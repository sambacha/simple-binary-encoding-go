// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MDIncrementalRefreshLimitsBanding50 struct {
	TransactTime        uint64
	MatchEventIndicator MatchEventIndicator
	NoMDEntries         []MDIncrementalRefreshLimitsBanding50NoMDEntries
}
type MDIncrementalRefreshLimitsBanding50NoMDEntries struct {
	HighLimitPrice    PRICENULL9
	LowLimitPrice     PRICENULL9
	MaxPriceVariation PRICENULL9
	SecurityID        int32
	RptSeq            uint32
	MDUpdateAction    int8
	MDEntryType       [1]byte
}

func (m *MDIncrementalRefreshLimitsBanding50) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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

func (m *MDIncrementalRefreshLimitsBanding50) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
			m.NoMDEntries = make([]MDIncrementalRefreshLimitsBanding50NoMDEntries, NoMDEntriesNumInGroup)
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

func (m *MDIncrementalRefreshLimitsBanding50) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDIncrementalRefreshLimitsBanding50Init(m *MDIncrementalRefreshLimitsBanding50) {
	return
}

func (m *MDIncrementalRefreshLimitsBanding50NoMDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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

func (m *MDIncrementalRefreshLimitsBanding50NoMDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MDIncrementalRefreshLimitsBanding50NoMDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDIncrementalRefreshLimitsBanding50NoMDEntriesInit(m *MDIncrementalRefreshLimitsBanding50NoMDEntries) {
	m.MDUpdateAction = 0
	m.MDEntryType[0] = 103
	return
}

func (*MDIncrementalRefreshLimitsBanding50) SbeBlockLength() (blockLength uint16) {
	return 11
}

func (*MDIncrementalRefreshLimitsBanding50) SbeTemplateId() (templateId uint16) {
	return 50
}

func (*MDIncrementalRefreshLimitsBanding50) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MDIncrementalRefreshLimitsBanding50) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDIncrementalRefreshLimitsBanding50) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MDIncrementalRefreshLimitsBanding50) TransactTimeId() uint16 {
	return 60
}

func (*MDIncrementalRefreshLimitsBanding50) TransactTimeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding50) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TransactTimeSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding50) TransactTimeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding50) TransactTimeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding50) TransactTimeMinValue() uint64 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding50) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDIncrementalRefreshLimitsBanding50) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDIncrementalRefreshLimitsBanding50) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MDIncrementalRefreshLimitsBanding50) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding50) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding50) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding50) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) HighLimitPriceId() uint16 {
	return 1149
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) HighLimitPriceSinceVersion() uint16 {
	return 9
}

func (m *MDIncrementalRefreshLimitsBanding50NoMDEntries) HighLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.HighLimitPriceSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) HighLimitPriceDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) HighLimitPriceMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) LowLimitPriceId() uint16 {
	return 1148
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) LowLimitPriceSinceVersion() uint16 {
	return 9
}

func (m *MDIncrementalRefreshLimitsBanding50NoMDEntries) LowLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LowLimitPriceSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) LowLimitPriceDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) LowLimitPriceMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MaxPriceVariationId() uint16 {
	return 1143
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MaxPriceVariationSinceVersion() uint16 {
	return 9
}

func (m *MDIncrementalRefreshLimitsBanding50NoMDEntries) MaxPriceVariationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaxPriceVariationSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MaxPriceVariationDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MaxPriceVariationMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) SecurityIDId() uint16 {
	return 48
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding50NoMDEntries) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) SecurityIDMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) RptSeqId() uint16 {
	return 83
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) RptSeqSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding50NoMDEntries) RptSeqInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RptSeqSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) RptSeqDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) RptSeqMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) RptSeqMinValue() uint32 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) RptSeqMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) RptSeqNullValue() uint32 {
	return math.MaxUint32
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDUpdateActionId() uint16 {
	return 279
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding50NoMDEntries) MDUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDUpdateActionSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDUpdateActionDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDUpdateActionMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDUpdateActionMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDUpdateActionMaxValue() int8 {
	return math.MaxInt8
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDUpdateActionNullValue() int8 {
	return math.MinInt8
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDEntryTypeId() uint16 {
	return 269
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDEntryTypeSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding50NoMDEntries) MDEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryTypeSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDEntryTypeDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDEntryTypeMetaAttribute(meta int) string {
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

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDEntryTypeMinValue() byte {
	return byte(32)
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDEntryTypeMaxValue() byte {
	return byte(126)
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) MDEntryTypeNullValue() byte {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding50) NoMDEntriesId() uint16 {
	return 268
}

func (*MDIncrementalRefreshLimitsBanding50) NoMDEntriesSinceVersion() uint16 {
	return 0
}

func (m *MDIncrementalRefreshLimitsBanding50) NoMDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoMDEntriesSinceVersion()
}

func (*MDIncrementalRefreshLimitsBanding50) NoMDEntriesDeprecated() uint16 {
	return 0
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) SbeBlockLength() (blockLength uint) {
	return 32
}

func (*MDIncrementalRefreshLimitsBanding50NoMDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
