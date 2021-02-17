// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type QuoteRequest39 struct {
	TransactTime        uint64
	QuoteReqID          [23]byte
	MatchEventIndicator MatchEventIndicator
	NoRelatedSym        []QuoteRequest39NoRelatedSym
}
type QuoteRequest39NoRelatedSym struct {
	Symbol     [20]byte
	SecurityID int32
	OrderQty   int32
	QuoteType  int8
	Side       int8
}

func (q *QuoteRequest39) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := q.RangeCheck(q.SbeSchemaVersion(), q.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, q.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, q.QuoteReqID[:]); err != nil {
		return err
	}
	if err := q.MatchEventIndicator.Encode(_m, _w); err != nil {
		return err
	}

	for i := 0; i < 3; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	var NoRelatedSymBlockLength uint16 = 32
	if err := _m.WriteUint16(_w, NoRelatedSymBlockLength); err != nil {
		return err
	}
	var NoRelatedSymNumInGroup uint8 = uint8(len(q.NoRelatedSym))
	if err := _m.WriteUint8(_w, NoRelatedSymNumInGroup); err != nil {
		return err
	}
	for _, prop := range q.NoRelatedSym {
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

func (q *QuoteRequest39) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !q.TransactTimeInActingVersion(actingVersion) {
		q.TransactTime = q.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &q.TransactTime); err != nil {
			return err
		}
	}
	if !q.QuoteReqIDInActingVersion(actingVersion) {
		for idx := 0; idx < 23; idx++ {
			q.QuoteReqID[idx] = q.QuoteReqIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, q.QuoteReqID[:]); err != nil {
			return err
		}
	}
	if q.MatchEventIndicatorInActingVersion(actingVersion) {
		if err := q.MatchEventIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > q.SbeSchemaVersion() && blockLength > q.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-q.SbeBlockLength()))
	}
	io.CopyN(ioutil.Discard, _r, 3)

	if q.NoRelatedSymInActingVersion(actingVersion) {
		var NoRelatedSymBlockLength uint16
		if err := _m.ReadUint16(_r, &NoRelatedSymBlockLength); err != nil {
			return err
		}
		var NoRelatedSymNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoRelatedSymNumInGroup); err != nil {
			return err
		}
		if cap(q.NoRelatedSym) < int(NoRelatedSymNumInGroup) {
			q.NoRelatedSym = make([]QuoteRequest39NoRelatedSym, NoRelatedSymNumInGroup)
		}
		q.NoRelatedSym = q.NoRelatedSym[:NoRelatedSymNumInGroup]
		for i, _ := range q.NoRelatedSym {
			if err := q.NoRelatedSym[i].Decode(_m, _r, actingVersion, uint(NoRelatedSymBlockLength)); err != nil {
				return err
			}
		}
		io.CopyN(ioutil.Discard, _r, 2)
	}
	if doRangeCheck {
		if err := q.RangeCheck(actingVersion, q.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (q *QuoteRequest39) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if q.TransactTimeInActingVersion(actingVersion) {
		if q.TransactTime < q.TransactTimeMinValue() || q.TransactTime > q.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on q.TransactTime (%v < %v > %v)", q.TransactTimeMinValue(), q.TransactTime, q.TransactTimeMaxValue())
		}
	}
	if q.QuoteReqIDInActingVersion(actingVersion) {
		for idx := 0; idx < 23; idx++ {
			if q.QuoteReqID[idx] < q.QuoteReqIDMinValue() || q.QuoteReqID[idx] > q.QuoteReqIDMaxValue() {
				return fmt.Errorf("Range check failed on q.QuoteReqID[%d] (%v < %v > %v)", idx, q.QuoteReqIDMinValue(), q.QuoteReqID[idx], q.QuoteReqIDMaxValue())
			}
		}
	}
	for _, prop := range q.NoRelatedSym {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func QuoteRequest39Init(q *QuoteRequest39) {
	return
}

func (q *QuoteRequest39NoRelatedSym) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, q.Symbol[:]); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, q.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, q.OrderQty); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, q.QuoteType); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, q.Side); err != nil {
		return err
	}
	return nil
}

func (q *QuoteRequest39NoRelatedSym) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !q.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			q.Symbol[idx] = q.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, q.Symbol[:]); err != nil {
			return err
		}
	}
	if !q.SecurityIDInActingVersion(actingVersion) {
		q.SecurityID = q.SecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &q.SecurityID); err != nil {
			return err
		}
	}
	if !q.OrderQtyInActingVersion(actingVersion) {
		q.OrderQty = q.OrderQtyNullValue()
	} else {
		if err := _m.ReadInt32(_r, &q.OrderQty); err != nil {
			return err
		}
	}
	if !q.QuoteTypeInActingVersion(actingVersion) {
		q.QuoteType = q.QuoteTypeNullValue()
	} else {
		if err := _m.ReadInt8(_r, &q.QuoteType); err != nil {
			return err
		}
	}
	if !q.SideInActingVersion(actingVersion) {
		q.Side = q.SideNullValue()
	} else {
		if err := _m.ReadInt8(_r, &q.Side); err != nil {
			return err
		}
	}
	if actingVersion > q.SbeSchemaVersion() && blockLength > q.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-q.SbeBlockLength()))
	}
	return nil
}

func (q *QuoteRequest39NoRelatedSym) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if q.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if q.Symbol[idx] < q.SymbolMinValue() || q.Symbol[idx] > q.SymbolMaxValue() {
				return fmt.Errorf("Range check failed on q.Symbol[%d] (%v < %v > %v)", idx, q.SymbolMinValue(), q.Symbol[idx], q.SymbolMaxValue())
			}
		}
	}
	if q.SecurityIDInActingVersion(actingVersion) {
		if q.SecurityID < q.SecurityIDMinValue() || q.SecurityID > q.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on q.SecurityID (%v < %v > %v)", q.SecurityIDMinValue(), q.SecurityID, q.SecurityIDMaxValue())
		}
	}
	if q.OrderQtyInActingVersion(actingVersion) {
		if q.OrderQty != q.OrderQtyNullValue() && (q.OrderQty < q.OrderQtyMinValue() || q.OrderQty > q.OrderQtyMaxValue()) {
			return fmt.Errorf("Range check failed on q.OrderQty (%v < %v > %v)", q.OrderQtyMinValue(), q.OrderQty, q.OrderQtyMaxValue())
		}
	}
	if q.QuoteTypeInActingVersion(actingVersion) {
		if q.QuoteType < q.QuoteTypeMinValue() || q.QuoteType > q.QuoteTypeMaxValue() {
			return fmt.Errorf("Range check failed on q.QuoteType (%v < %v > %v)", q.QuoteTypeMinValue(), q.QuoteType, q.QuoteTypeMaxValue())
		}
	}
	if q.SideInActingVersion(actingVersion) {
		if q.Side != q.SideNullValue() && (q.Side < q.SideMinValue() || q.Side > q.SideMaxValue()) {
			return fmt.Errorf("Range check failed on q.Side (%v < %v > %v)", q.SideMinValue(), q.Side, q.SideMaxValue())
		}
	}
	return nil
}

func QuoteRequest39NoRelatedSymInit(q *QuoteRequest39NoRelatedSym) {
	q.OrderQty = 2147483647
	q.Side = 127
	return
}

func (*QuoteRequest39) SbeBlockLength() (blockLength uint16) {
	return 35
}

func (*QuoteRequest39) SbeTemplateId() (templateId uint16) {
	return 39
}

func (*QuoteRequest39) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*QuoteRequest39) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*QuoteRequest39) SbeSemanticType() (semanticType []byte) {
	return []byte("R")
}

func (*QuoteRequest39) TransactTimeId() uint16 {
	return 60
}

func (*QuoteRequest39) TransactTimeSinceVersion() uint16 {
	return 0
}

func (q *QuoteRequest39) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.TransactTimeSinceVersion()
}

func (*QuoteRequest39) TransactTimeDeprecated() uint16 {
	return 0
}

func (*QuoteRequest39) TransactTimeMetaAttribute(meta int) string {
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

func (*QuoteRequest39) TransactTimeMinValue() uint64 {
	return 0
}

func (*QuoteRequest39) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*QuoteRequest39) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*QuoteRequest39) QuoteReqIDId() uint16 {
	return 131
}

func (*QuoteRequest39) QuoteReqIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteRequest39) QuoteReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteReqIDSinceVersion()
}

func (*QuoteRequest39) QuoteReqIDDeprecated() uint16 {
	return 0
}

func (*QuoteRequest39) QuoteReqIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*QuoteRequest39) QuoteReqIDMinValue() byte {
	return byte(32)
}

func (*QuoteRequest39) QuoteReqIDMaxValue() byte {
	return byte(126)
}

func (*QuoteRequest39) QuoteReqIDNullValue() byte {
	return 0
}

func (q *QuoteRequest39) QuoteReqIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*QuoteRequest39) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*QuoteRequest39) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (q *QuoteRequest39) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.MatchEventIndicatorSinceVersion()
}

func (*QuoteRequest39) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*QuoteRequest39) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*QuoteRequest39NoRelatedSym) SymbolId() uint16 {
	return 55
}

func (*QuoteRequest39NoRelatedSym) SymbolSinceVersion() uint16 {
	return 0
}

func (q *QuoteRequest39NoRelatedSym) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.SymbolSinceVersion()
}

func (*QuoteRequest39NoRelatedSym) SymbolDeprecated() uint16 {
	return 0
}

func (*QuoteRequest39NoRelatedSym) SymbolMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*QuoteRequest39NoRelatedSym) SymbolMinValue() byte {
	return byte(32)
}

func (*QuoteRequest39NoRelatedSym) SymbolMaxValue() byte {
	return byte(126)
}

func (*QuoteRequest39NoRelatedSym) SymbolNullValue() byte {
	return 0
}

func (q *QuoteRequest39NoRelatedSym) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*QuoteRequest39NoRelatedSym) SecurityIDId() uint16 {
	return 48
}

func (*QuoteRequest39NoRelatedSym) SecurityIDSinceVersion() uint16 {
	return 0
}

func (q *QuoteRequest39NoRelatedSym) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.SecurityIDSinceVersion()
}

func (*QuoteRequest39NoRelatedSym) SecurityIDDeprecated() uint16 {
	return 0
}

func (*QuoteRequest39NoRelatedSym) SecurityIDMetaAttribute(meta int) string {
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

func (*QuoteRequest39NoRelatedSym) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*QuoteRequest39NoRelatedSym) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*QuoteRequest39NoRelatedSym) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*QuoteRequest39NoRelatedSym) OrderQtyId() uint16 {
	return 38
}

func (*QuoteRequest39NoRelatedSym) OrderQtySinceVersion() uint16 {
	return 0
}

func (q *QuoteRequest39NoRelatedSym) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.OrderQtySinceVersion()
}

func (*QuoteRequest39NoRelatedSym) OrderQtyDeprecated() uint16 {
	return 0
}

func (*QuoteRequest39NoRelatedSym) OrderQtyMetaAttribute(meta int) string {
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

func (*QuoteRequest39NoRelatedSym) OrderQtyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*QuoteRequest39NoRelatedSym) OrderQtyMaxValue() int32 {
	return math.MaxInt32
}

func (*QuoteRequest39NoRelatedSym) OrderQtyNullValue() int32 {
	return 2147483647
}

func (*QuoteRequest39NoRelatedSym) QuoteTypeId() uint16 {
	return 537
}

func (*QuoteRequest39NoRelatedSym) QuoteTypeSinceVersion() uint16 {
	return 0
}

func (q *QuoteRequest39NoRelatedSym) QuoteTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteTypeSinceVersion()
}

func (*QuoteRequest39NoRelatedSym) QuoteTypeDeprecated() uint16 {
	return 0
}

func (*QuoteRequest39NoRelatedSym) QuoteTypeMetaAttribute(meta int) string {
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

func (*QuoteRequest39NoRelatedSym) QuoteTypeMinValue() int8 {
	return math.MinInt8 + 1
}

func (*QuoteRequest39NoRelatedSym) QuoteTypeMaxValue() int8 {
	return math.MaxInt8
}

func (*QuoteRequest39NoRelatedSym) QuoteTypeNullValue() int8 {
	return math.MinInt8
}

func (*QuoteRequest39NoRelatedSym) SideId() uint16 {
	return 54
}

func (*QuoteRequest39NoRelatedSym) SideSinceVersion() uint16 {
	return 0
}

func (q *QuoteRequest39NoRelatedSym) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.SideSinceVersion()
}

func (*QuoteRequest39NoRelatedSym) SideDeprecated() uint16 {
	return 0
}

func (*QuoteRequest39NoRelatedSym) SideMetaAttribute(meta int) string {
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

func (*QuoteRequest39NoRelatedSym) SideMinValue() int8 {
	return math.MinInt8 + 1
}

func (*QuoteRequest39NoRelatedSym) SideMaxValue() int8 {
	return math.MaxInt8
}

func (*QuoteRequest39NoRelatedSym) SideNullValue() int8 {
	return 127
}

func (*QuoteRequest39) NoRelatedSymId() uint16 {
	return 146
}

func (*QuoteRequest39) NoRelatedSymSinceVersion() uint16 {
	return 0
}

func (q *QuoteRequest39) NoRelatedSymInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.NoRelatedSymSinceVersion()
}

func (*QuoteRequest39) NoRelatedSymDeprecated() uint16 {
	return 0
}

func (*QuoteRequest39NoRelatedSym) SbeBlockLength() (blockLength uint) {
	return 32
}

func (*QuoteRequest39NoRelatedSym) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
