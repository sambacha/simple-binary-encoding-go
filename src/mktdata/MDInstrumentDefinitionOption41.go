// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MDInstrumentDefinitionOption41 struct {
	MatchEventIndicator     MatchEventIndicator
	TotNumReports           uint32
	SecurityUpdateAction    SecurityUpdateActionEnum
	LastUpdateTime          uint64
	MDSecurityTradingStatus SecurityTradingStatusEnum
	ApplID                  int16
	MarketSegmentID         uint8
	UnderlyingProduct       uint8
	SecurityExchange        [4]byte
	SecurityGroup           [6]byte
	Asset                   [6]byte
	Symbol                  [20]byte
	SecurityID              int32
	SecurityIDSource        [1]byte
	SecurityType            [6]byte
	CFICode                 [6]byte
	PutOrCall               PutOrCallEnum
	MaturityMonthYear       MaturityMonthYear
	Currency                [3]byte
	StrikePrice             PRICENULL
	StrikeCurrency          [3]byte
	SettlCurrency           [3]byte
	MinCabPrice             PRICENULL
	MatchAlgorithm          byte
	MinTradeVol             uint32
	MaxTradeVol             uint32
	MinPriceIncrement       PRICENULL
	MinPriceIncrementAmount PRICENULL
	DisplayFactor           FLOAT
	TickRule                int8
	MainFraction            uint8
	SubFraction             uint8
	PriceDisplayFormat      uint8
	UnitOfMeasure           [30]byte
	UnitOfMeasureQty        PRICENULL
	TradingReferencePrice   PRICENULL
	SettlPriceType          SettlPriceType
	ClearedVolume           int32
	OpenInterestQty         int32
	LowLimitPrice           PRICENULL
	HighLimitPrice          PRICENULL
	UserDefinedInstrument   byte
	TradingReferenceDate    uint16
	NoEvents                []MDInstrumentDefinitionOption41NoEvents
	NoMDFeedTypes           []MDInstrumentDefinitionOption41NoMDFeedTypes
	NoInstAttrib            []MDInstrumentDefinitionOption41NoInstAttrib
	NoLotTypeRules          []MDInstrumentDefinitionOption41NoLotTypeRules
	NoUnderlyings           []MDInstrumentDefinitionOption41NoUnderlyings
	NoRelatedInstruments    []MDInstrumentDefinitionOption41NoRelatedInstruments
}
type MDInstrumentDefinitionOption41NoEvents struct {
	EventType EventTypeEnum
	EventTime uint64
}
type MDInstrumentDefinitionOption41NoMDFeedTypes struct {
	MDFeedType  [3]byte
	MarketDepth int8
}
type MDInstrumentDefinitionOption41NoInstAttrib struct {
	InstAttribType  int8
	InstAttribValue InstAttribValue
}
type MDInstrumentDefinitionOption41NoLotTypeRules struct {
	LotType    int8
	MinLotSize DecimalQty
}
type MDInstrumentDefinitionOption41NoUnderlyings struct {
	UnderlyingSecurityID       int32
	UnderlyingSecurityIDSource [1]byte
	UnderlyingSymbol           [20]byte
}
type MDInstrumentDefinitionOption41NoRelatedInstruments struct {
	RelatedSecurityID       int32
	RelatedSecurityIDSource [1]byte
	RelatedSymbol           [20]byte
}

func (m *MDInstrumentDefinitionOption41) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := m.MatchEventIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.TotNumReports); err != nil {
		return err
	}
	if err := m.SecurityUpdateAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.LastUpdateTime); err != nil {
		return err
	}
	if err := m.MDSecurityTradingStatus.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt16(_w, m.ApplID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.MarketSegmentID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.UnderlyingProduct); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.SecurityExchange[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.SecurityGroup[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.Asset[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.Symbol[:]); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.SecurityID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.SecurityType[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.CFICode[:]); err != nil {
		return err
	}
	if err := m.PutOrCall.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MaturityMonthYear.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.Currency[:]); err != nil {
		return err
	}
	if err := m.StrikePrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.StrikeCurrency[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.SettlCurrency[:]); err != nil {
		return err
	}
	if err := m.MinCabPrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.MatchAlgorithm); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.MinTradeVol); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.MaxTradeVol); err != nil {
		return err
	}
	if err := m.MinPriceIncrement.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MinPriceIncrementAmount.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.DisplayFactor.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.TickRule); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.MainFraction); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.SubFraction); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.PriceDisplayFormat); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.UnitOfMeasure[:]); err != nil {
		return err
	}
	if err := m.UnitOfMeasureQty.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.TradingReferencePrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.SettlPriceType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.ClearedVolume); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.OpenInterestQty); err != nil {
		return err
	}
	if err := m.LowLimitPrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.HighLimitPrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.UserDefinedInstrument); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.TradingReferenceDate); err != nil {
		return err
	}
	var NoEventsBlockLength uint16 = 9
	if err := _m.WriteUint16(_w, NoEventsBlockLength); err != nil {
		return err
	}
	var NoEventsNumInGroup uint8 = uint8(len(m.NoEvents))
	if err := _m.WriteUint8(_w, NoEventsNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.NoEvents {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	var NoMDFeedTypesBlockLength uint16 = 4
	if err := _m.WriteUint16(_w, NoMDFeedTypesBlockLength); err != nil {
		return err
	}
	var NoMDFeedTypesNumInGroup uint8 = uint8(len(m.NoMDFeedTypes))
	if err := _m.WriteUint8(_w, NoMDFeedTypesNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.NoMDFeedTypes {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	var NoInstAttribBlockLength uint16 = 4
	if err := _m.WriteUint16(_w, NoInstAttribBlockLength); err != nil {
		return err
	}
	var NoInstAttribNumInGroup uint8 = uint8(len(m.NoInstAttrib))
	if err := _m.WriteUint8(_w, NoInstAttribNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.NoInstAttrib {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	var NoLotTypeRulesBlockLength uint16 = 5
	if err := _m.WriteUint16(_w, NoLotTypeRulesBlockLength); err != nil {
		return err
	}
	var NoLotTypeRulesNumInGroup uint8 = uint8(len(m.NoLotTypeRules))
	if err := _m.WriteUint8(_w, NoLotTypeRulesNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.NoLotTypeRules {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	var NoUnderlyingsBlockLength uint16 = 24
	if err := _m.WriteUint16(_w, NoUnderlyingsBlockLength); err != nil {
		return err
	}
	var NoUnderlyingsNumInGroup uint8 = uint8(len(m.NoUnderlyings))
	if err := _m.WriteUint8(_w, NoUnderlyingsNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.NoUnderlyings {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	var NoRelatedInstrumentsBlockLength uint16 = 24
	if err := _m.WriteUint16(_w, NoRelatedInstrumentsBlockLength); err != nil {
		return err
	}
	var NoRelatedInstrumentsNumInGroup uint8 = uint8(len(m.NoRelatedInstruments))
	if err := _m.WriteUint8(_w, NoRelatedInstrumentsNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.NoRelatedInstruments {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (m *MDInstrumentDefinitionOption41) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if m.MatchEventIndicatorInActingVersion(actingVersion) {
		if err := m.MatchEventIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.TotNumReportsInActingVersion(actingVersion) {
		m.TotNumReports = m.TotNumReportsNullValue()
	} else {
		if err := _m.ReadUint32(_r, &m.TotNumReports); err != nil {
			return err
		}
	}
	if m.SecurityUpdateActionInActingVersion(actingVersion) {
		if err := m.SecurityUpdateAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.LastUpdateTimeInActingVersion(actingVersion) {
		m.LastUpdateTime = m.LastUpdateTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.LastUpdateTime); err != nil {
			return err
		}
	}
	if m.MDSecurityTradingStatusInActingVersion(actingVersion) {
		if err := m.MDSecurityTradingStatus.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.ApplIDInActingVersion(actingVersion) {
		m.ApplID = m.ApplIDNullValue()
	} else {
		if err := _m.ReadInt16(_r, &m.ApplID); err != nil {
			return err
		}
	}
	if !m.MarketSegmentIDInActingVersion(actingVersion) {
		m.MarketSegmentID = m.MarketSegmentIDNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.MarketSegmentID); err != nil {
			return err
		}
	}
	if !m.UnderlyingProductInActingVersion(actingVersion) {
		m.UnderlyingProduct = m.UnderlyingProductNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.UnderlyingProduct); err != nil {
			return err
		}
	}
	if !m.SecurityExchangeInActingVersion(actingVersion) {
		for idx := 0; idx < 4; idx++ {
			m.SecurityExchange[idx] = m.SecurityExchangeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.SecurityExchange[:]); err != nil {
			return err
		}
	}
	if !m.SecurityGroupInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			m.SecurityGroup[idx] = m.SecurityGroupNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.SecurityGroup[:]); err != nil {
			return err
		}
	}
	if !m.AssetInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			m.Asset[idx] = m.AssetNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.Asset[:]); err != nil {
			return err
		}
	}
	if !m.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			m.Symbol[idx] = m.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.Symbol[:]); err != nil {
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
	m.SecurityIDSource[0] = 56
	if !m.SecurityTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			m.SecurityType[idx] = m.SecurityTypeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.SecurityType[:]); err != nil {
			return err
		}
	}
	if !m.CFICodeInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			m.CFICode[idx] = m.CFICodeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.CFICode[:]); err != nil {
			return err
		}
	}
	if m.PutOrCallInActingVersion(actingVersion) {
		if err := m.PutOrCall.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.MaturityMonthYearInActingVersion(actingVersion) {
		if err := m.MaturityMonthYear.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.CurrencyInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			m.Currency[idx] = m.CurrencyNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.Currency[:]); err != nil {
			return err
		}
	}
	if m.StrikePriceInActingVersion(actingVersion) {
		if err := m.StrikePrice.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.StrikeCurrencyInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			m.StrikeCurrency[idx] = m.StrikeCurrencyNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.StrikeCurrency[:]); err != nil {
			return err
		}
	}
	if !m.SettlCurrencyInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			m.SettlCurrency[idx] = m.SettlCurrencyNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.SettlCurrency[:]); err != nil {
			return err
		}
	}
	if m.MinCabPriceInActingVersion(actingVersion) {
		if err := m.MinCabPrice.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.MatchAlgorithmInActingVersion(actingVersion) {
		m.MatchAlgorithm = m.MatchAlgorithmNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.MatchAlgorithm); err != nil {
			return err
		}
	}
	if !m.MinTradeVolInActingVersion(actingVersion) {
		m.MinTradeVol = m.MinTradeVolNullValue()
	} else {
		if err := _m.ReadUint32(_r, &m.MinTradeVol); err != nil {
			return err
		}
	}
	if !m.MaxTradeVolInActingVersion(actingVersion) {
		m.MaxTradeVol = m.MaxTradeVolNullValue()
	} else {
		if err := _m.ReadUint32(_r, &m.MaxTradeVol); err != nil {
			return err
		}
	}
	if m.MinPriceIncrementInActingVersion(actingVersion) {
		if err := m.MinPriceIncrement.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.MinPriceIncrementAmountInActingVersion(actingVersion) {
		if err := m.MinPriceIncrementAmount.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.DisplayFactorInActingVersion(actingVersion) {
		if err := m.DisplayFactor.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.TickRuleInActingVersion(actingVersion) {
		m.TickRule = m.TickRuleNullValue()
	} else {
		if err := _m.ReadInt8(_r, &m.TickRule); err != nil {
			return err
		}
	}
	if !m.MainFractionInActingVersion(actingVersion) {
		m.MainFraction = m.MainFractionNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.MainFraction); err != nil {
			return err
		}
	}
	if !m.SubFractionInActingVersion(actingVersion) {
		m.SubFraction = m.SubFractionNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.SubFraction); err != nil {
			return err
		}
	}
	if !m.PriceDisplayFormatInActingVersion(actingVersion) {
		m.PriceDisplayFormat = m.PriceDisplayFormatNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.PriceDisplayFormat); err != nil {
			return err
		}
	}
	if !m.UnitOfMeasureInActingVersion(actingVersion) {
		for idx := 0; idx < 30; idx++ {
			m.UnitOfMeasure[idx] = m.UnitOfMeasureNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.UnitOfMeasure[:]); err != nil {
			return err
		}
	}
	if m.UnitOfMeasureQtyInActingVersion(actingVersion) {
		if err := m.UnitOfMeasureQty.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.TradingReferencePriceInActingVersion(actingVersion) {
		if err := m.TradingReferencePrice.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.SettlPriceTypeInActingVersion(actingVersion) {
		if err := m.SettlPriceType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.ClearedVolumeInActingVersion(actingVersion) {
		m.ClearedVolume = m.ClearedVolumeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.ClearedVolume); err != nil {
			return err
		}
	}
	if !m.OpenInterestQtyInActingVersion(actingVersion) {
		m.OpenInterestQty = m.OpenInterestQtyNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.OpenInterestQty); err != nil {
			return err
		}
	}
	if m.LowLimitPriceInActingVersion(actingVersion) {
		if err := m.LowLimitPrice.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.HighLimitPriceInActingVersion(actingVersion) {
		if err := m.HighLimitPrice.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.UserDefinedInstrumentInActingVersion(actingVersion) {
		m.UserDefinedInstrument = m.UserDefinedInstrumentNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.UserDefinedInstrument); err != nil {
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
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.NoEventsInActingVersion(actingVersion) {
		var NoEventsBlockLength uint16
		if err := _m.ReadUint16(_r, &NoEventsBlockLength); err != nil {
			return err
		}
		var NoEventsNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoEventsNumInGroup); err != nil {
			return err
		}
		if cap(m.NoEvents) < int(NoEventsNumInGroup) {
			m.NoEvents = make([]MDInstrumentDefinitionOption41NoEvents, NoEventsNumInGroup)
		}
		m.NoEvents = m.NoEvents[:NoEventsNumInGroup]
		for i, _ := range m.NoEvents {
			if err := m.NoEvents[i].Decode(_m, _r, actingVersion, uint(NoEventsBlockLength)); err != nil {
				return err
			}
		}
	}

	if m.NoMDFeedTypesInActingVersion(actingVersion) {
		var NoMDFeedTypesBlockLength uint16
		if err := _m.ReadUint16(_r, &NoMDFeedTypesBlockLength); err != nil {
			return err
		}
		var NoMDFeedTypesNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoMDFeedTypesNumInGroup); err != nil {
			return err
		}
		if cap(m.NoMDFeedTypes) < int(NoMDFeedTypesNumInGroup) {
			m.NoMDFeedTypes = make([]MDInstrumentDefinitionOption41NoMDFeedTypes, NoMDFeedTypesNumInGroup)
		}
		m.NoMDFeedTypes = m.NoMDFeedTypes[:NoMDFeedTypesNumInGroup]
		for i, _ := range m.NoMDFeedTypes {
			if err := m.NoMDFeedTypes[i].Decode(_m, _r, actingVersion, uint(NoMDFeedTypesBlockLength)); err != nil {
				return err
			}
		}
	}

	if m.NoInstAttribInActingVersion(actingVersion) {
		var NoInstAttribBlockLength uint16
		if err := _m.ReadUint16(_r, &NoInstAttribBlockLength); err != nil {
			return err
		}
		var NoInstAttribNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoInstAttribNumInGroup); err != nil {
			return err
		}
		if cap(m.NoInstAttrib) < int(NoInstAttribNumInGroup) {
			m.NoInstAttrib = make([]MDInstrumentDefinitionOption41NoInstAttrib, NoInstAttribNumInGroup)
		}
		m.NoInstAttrib = m.NoInstAttrib[:NoInstAttribNumInGroup]
		for i, _ := range m.NoInstAttrib {
			if err := m.NoInstAttrib[i].Decode(_m, _r, actingVersion, uint(NoInstAttribBlockLength)); err != nil {
				return err
			}
		}
	}

	if m.NoLotTypeRulesInActingVersion(actingVersion) {
		var NoLotTypeRulesBlockLength uint16
		if err := _m.ReadUint16(_r, &NoLotTypeRulesBlockLength); err != nil {
			return err
		}
		var NoLotTypeRulesNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoLotTypeRulesNumInGroup); err != nil {
			return err
		}
		if cap(m.NoLotTypeRules) < int(NoLotTypeRulesNumInGroup) {
			m.NoLotTypeRules = make([]MDInstrumentDefinitionOption41NoLotTypeRules, NoLotTypeRulesNumInGroup)
		}
		m.NoLotTypeRules = m.NoLotTypeRules[:NoLotTypeRulesNumInGroup]
		for i, _ := range m.NoLotTypeRules {
			if err := m.NoLotTypeRules[i].Decode(_m, _r, actingVersion, uint(NoLotTypeRulesBlockLength)); err != nil {
				return err
			}
		}
	}

	if m.NoUnderlyingsInActingVersion(actingVersion) {
		var NoUnderlyingsBlockLength uint16
		if err := _m.ReadUint16(_r, &NoUnderlyingsBlockLength); err != nil {
			return err
		}
		var NoUnderlyingsNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoUnderlyingsNumInGroup); err != nil {
			return err
		}
		if cap(m.NoUnderlyings) < int(NoUnderlyingsNumInGroup) {
			m.NoUnderlyings = make([]MDInstrumentDefinitionOption41NoUnderlyings, NoUnderlyingsNumInGroup)
		}
		m.NoUnderlyings = m.NoUnderlyings[:NoUnderlyingsNumInGroup]
		for i, _ := range m.NoUnderlyings {
			if err := m.NoUnderlyings[i].Decode(_m, _r, actingVersion, uint(NoUnderlyingsBlockLength)); err != nil {
				return err
			}
		}
	}

	if m.NoRelatedInstrumentsInActingVersion(actingVersion) {
		var NoRelatedInstrumentsBlockLength uint16
		if err := _m.ReadUint16(_r, &NoRelatedInstrumentsBlockLength); err != nil {
			return err
		}
		var NoRelatedInstrumentsNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoRelatedInstrumentsNumInGroup); err != nil {
			return err
		}
		if cap(m.NoRelatedInstruments) < int(NoRelatedInstrumentsNumInGroup) {
			m.NoRelatedInstruments = make([]MDInstrumentDefinitionOption41NoRelatedInstruments, NoRelatedInstrumentsNumInGroup)
		}
		m.NoRelatedInstruments = m.NoRelatedInstruments[:NoRelatedInstrumentsNumInGroup]
		for i, _ := range m.NoRelatedInstruments {
			if err := m.NoRelatedInstruments[i].Decode(_m, _r, actingVersion, uint(NoRelatedInstrumentsBlockLength)); err != nil {
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

func (m *MDInstrumentDefinitionOption41) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.TotNumReportsInActingVersion(actingVersion) {
		if m.TotNumReports != m.TotNumReportsNullValue() && (m.TotNumReports < m.TotNumReportsMinValue() || m.TotNumReports > m.TotNumReportsMaxValue()) {
			return fmt.Errorf("Range check failed on m.TotNumReports (%v < %v > %v)", m.TotNumReportsMinValue(), m.TotNumReports, m.TotNumReportsMaxValue())
		}
	}
	if err := m.SecurityUpdateAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.LastUpdateTimeInActingVersion(actingVersion) {
		if m.LastUpdateTime < m.LastUpdateTimeMinValue() || m.LastUpdateTime > m.LastUpdateTimeMaxValue() {
			return fmt.Errorf("Range check failed on m.LastUpdateTime (%v < %v > %v)", m.LastUpdateTimeMinValue(), m.LastUpdateTime, m.LastUpdateTimeMaxValue())
		}
	}
	if err := m.MDSecurityTradingStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.ApplIDInActingVersion(actingVersion) {
		if m.ApplID < m.ApplIDMinValue() || m.ApplID > m.ApplIDMaxValue() {
			return fmt.Errorf("Range check failed on m.ApplID (%v < %v > %v)", m.ApplIDMinValue(), m.ApplID, m.ApplIDMaxValue())
		}
	}
	if m.MarketSegmentIDInActingVersion(actingVersion) {
		if m.MarketSegmentID < m.MarketSegmentIDMinValue() || m.MarketSegmentID > m.MarketSegmentIDMaxValue() {
			return fmt.Errorf("Range check failed on m.MarketSegmentID (%v < %v > %v)", m.MarketSegmentIDMinValue(), m.MarketSegmentID, m.MarketSegmentIDMaxValue())
		}
	}
	if m.UnderlyingProductInActingVersion(actingVersion) {
		if m.UnderlyingProduct < m.UnderlyingProductMinValue() || m.UnderlyingProduct > m.UnderlyingProductMaxValue() {
			return fmt.Errorf("Range check failed on m.UnderlyingProduct (%v < %v > %v)", m.UnderlyingProductMinValue(), m.UnderlyingProduct, m.UnderlyingProductMaxValue())
		}
	}
	if m.SecurityExchangeInActingVersion(actingVersion) {
		for idx := 0; idx < 4; idx++ {
			if m.SecurityExchange[idx] < m.SecurityExchangeMinValue() || m.SecurityExchange[idx] > m.SecurityExchangeMaxValue() {
				return fmt.Errorf("Range check failed on m.SecurityExchange[%d] (%v < %v > %v)", idx, m.SecurityExchangeMinValue(), m.SecurityExchange[idx], m.SecurityExchangeMaxValue())
			}
		}
	}
	if m.SecurityGroupInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if m.SecurityGroup[idx] < m.SecurityGroupMinValue() || m.SecurityGroup[idx] > m.SecurityGroupMaxValue() {
				return fmt.Errorf("Range check failed on m.SecurityGroup[%d] (%v < %v > %v)", idx, m.SecurityGroupMinValue(), m.SecurityGroup[idx], m.SecurityGroupMaxValue())
			}
		}
	}
	if m.AssetInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if m.Asset[idx] < m.AssetMinValue() || m.Asset[idx] > m.AssetMaxValue() {
				return fmt.Errorf("Range check failed on m.Asset[%d] (%v < %v > %v)", idx, m.AssetMinValue(), m.Asset[idx], m.AssetMaxValue())
			}
		}
	}
	if m.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if m.Symbol[idx] < m.SymbolMinValue() || m.Symbol[idx] > m.SymbolMaxValue() {
				return fmt.Errorf("Range check failed on m.Symbol[%d] (%v < %v > %v)", idx, m.SymbolMinValue(), m.Symbol[idx], m.SymbolMaxValue())
			}
		}
	}
	if m.SecurityIDInActingVersion(actingVersion) {
		if m.SecurityID < m.SecurityIDMinValue() || m.SecurityID > m.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on m.SecurityID (%v < %v > %v)", m.SecurityIDMinValue(), m.SecurityID, m.SecurityIDMaxValue())
		}
	}
	if m.SecurityTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if m.SecurityType[idx] < m.SecurityTypeMinValue() || m.SecurityType[idx] > m.SecurityTypeMaxValue() {
				return fmt.Errorf("Range check failed on m.SecurityType[%d] (%v < %v > %v)", idx, m.SecurityTypeMinValue(), m.SecurityType[idx], m.SecurityTypeMaxValue())
			}
		}
	}
	if m.CFICodeInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if m.CFICode[idx] < m.CFICodeMinValue() || m.CFICode[idx] > m.CFICodeMaxValue() {
				return fmt.Errorf("Range check failed on m.CFICode[%d] (%v < %v > %v)", idx, m.CFICodeMinValue(), m.CFICode[idx], m.CFICodeMaxValue())
			}
		}
	}
	if err := m.PutOrCall.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.CurrencyInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if m.Currency[idx] < m.CurrencyMinValue() || m.Currency[idx] > m.CurrencyMaxValue() {
				return fmt.Errorf("Range check failed on m.Currency[%d] (%v < %v > %v)", idx, m.CurrencyMinValue(), m.Currency[idx], m.CurrencyMaxValue())
			}
		}
	}
	if m.StrikeCurrencyInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if m.StrikeCurrency[idx] < m.StrikeCurrencyMinValue() || m.StrikeCurrency[idx] > m.StrikeCurrencyMaxValue() {
				return fmt.Errorf("Range check failed on m.StrikeCurrency[%d] (%v < %v > %v)", idx, m.StrikeCurrencyMinValue(), m.StrikeCurrency[idx], m.StrikeCurrencyMaxValue())
			}
		}
	}
	if m.SettlCurrencyInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if m.SettlCurrency[idx] < m.SettlCurrencyMinValue() || m.SettlCurrency[idx] > m.SettlCurrencyMaxValue() {
				return fmt.Errorf("Range check failed on m.SettlCurrency[%d] (%v < %v > %v)", idx, m.SettlCurrencyMinValue(), m.SettlCurrency[idx], m.SettlCurrencyMaxValue())
			}
		}
	}
	if m.MatchAlgorithmInActingVersion(actingVersion) {
		if m.MatchAlgorithm < m.MatchAlgorithmMinValue() || m.MatchAlgorithm > m.MatchAlgorithmMaxValue() {
			return fmt.Errorf("Range check failed on m.MatchAlgorithm (%v < %v > %v)", m.MatchAlgorithmMinValue(), m.MatchAlgorithm, m.MatchAlgorithmMaxValue())
		}
	}
	if m.MinTradeVolInActingVersion(actingVersion) {
		if m.MinTradeVol < m.MinTradeVolMinValue() || m.MinTradeVol > m.MinTradeVolMaxValue() {
			return fmt.Errorf("Range check failed on m.MinTradeVol (%v < %v > %v)", m.MinTradeVolMinValue(), m.MinTradeVol, m.MinTradeVolMaxValue())
		}
	}
	if m.MaxTradeVolInActingVersion(actingVersion) {
		if m.MaxTradeVol < m.MaxTradeVolMinValue() || m.MaxTradeVol > m.MaxTradeVolMaxValue() {
			return fmt.Errorf("Range check failed on m.MaxTradeVol (%v < %v > %v)", m.MaxTradeVolMinValue(), m.MaxTradeVol, m.MaxTradeVolMaxValue())
		}
	}
	if m.TickRuleInActingVersion(actingVersion) {
		if m.TickRule != m.TickRuleNullValue() && (m.TickRule < m.TickRuleMinValue() || m.TickRule > m.TickRuleMaxValue()) {
			return fmt.Errorf("Range check failed on m.TickRule (%v < %v > %v)", m.TickRuleMinValue(), m.TickRule, m.TickRuleMaxValue())
		}
	}
	if m.MainFractionInActingVersion(actingVersion) {
		if m.MainFraction != m.MainFractionNullValue() && (m.MainFraction < m.MainFractionMinValue() || m.MainFraction > m.MainFractionMaxValue()) {
			return fmt.Errorf("Range check failed on m.MainFraction (%v < %v > %v)", m.MainFractionMinValue(), m.MainFraction, m.MainFractionMaxValue())
		}
	}
	if m.SubFractionInActingVersion(actingVersion) {
		if m.SubFraction != m.SubFractionNullValue() && (m.SubFraction < m.SubFractionMinValue() || m.SubFraction > m.SubFractionMaxValue()) {
			return fmt.Errorf("Range check failed on m.SubFraction (%v < %v > %v)", m.SubFractionMinValue(), m.SubFraction, m.SubFractionMaxValue())
		}
	}
	if m.PriceDisplayFormatInActingVersion(actingVersion) {
		if m.PriceDisplayFormat != m.PriceDisplayFormatNullValue() && (m.PriceDisplayFormat < m.PriceDisplayFormatMinValue() || m.PriceDisplayFormat > m.PriceDisplayFormatMaxValue()) {
			return fmt.Errorf("Range check failed on m.PriceDisplayFormat (%v < %v > %v)", m.PriceDisplayFormatMinValue(), m.PriceDisplayFormat, m.PriceDisplayFormatMaxValue())
		}
	}
	if m.UnitOfMeasureInActingVersion(actingVersion) {
		for idx := 0; idx < 30; idx++ {
			if m.UnitOfMeasure[idx] < m.UnitOfMeasureMinValue() || m.UnitOfMeasure[idx] > m.UnitOfMeasureMaxValue() {
				return fmt.Errorf("Range check failed on m.UnitOfMeasure[%d] (%v < %v > %v)", idx, m.UnitOfMeasureMinValue(), m.UnitOfMeasure[idx], m.UnitOfMeasureMaxValue())
			}
		}
	}
	if m.ClearedVolumeInActingVersion(actingVersion) {
		if m.ClearedVolume != m.ClearedVolumeNullValue() && (m.ClearedVolume < m.ClearedVolumeMinValue() || m.ClearedVolume > m.ClearedVolumeMaxValue()) {
			return fmt.Errorf("Range check failed on m.ClearedVolume (%v < %v > %v)", m.ClearedVolumeMinValue(), m.ClearedVolume, m.ClearedVolumeMaxValue())
		}
	}
	if m.OpenInterestQtyInActingVersion(actingVersion) {
		if m.OpenInterestQty != m.OpenInterestQtyNullValue() && (m.OpenInterestQty < m.OpenInterestQtyMinValue() || m.OpenInterestQty > m.OpenInterestQtyMaxValue()) {
			return fmt.Errorf("Range check failed on m.OpenInterestQty (%v < %v > %v)", m.OpenInterestQtyMinValue(), m.OpenInterestQty, m.OpenInterestQtyMaxValue())
		}
	}
	if m.UserDefinedInstrumentInActingVersion(actingVersion) {
		if m.UserDefinedInstrument < m.UserDefinedInstrumentMinValue() || m.UserDefinedInstrument > m.UserDefinedInstrumentMaxValue() {
			return fmt.Errorf("Range check failed on m.UserDefinedInstrument (%v < %v > %v)", m.UserDefinedInstrumentMinValue(), m.UserDefinedInstrument, m.UserDefinedInstrumentMaxValue())
		}
	}
	if m.TradingReferenceDateInActingVersion(actingVersion) {
		if m.TradingReferenceDate != m.TradingReferenceDateNullValue() && (m.TradingReferenceDate < m.TradingReferenceDateMinValue() || m.TradingReferenceDate > m.TradingReferenceDateMaxValue()) {
			return fmt.Errorf("Range check failed on m.TradingReferenceDate (%v < %v > %v)", m.TradingReferenceDateMinValue(), m.TradingReferenceDate, m.TradingReferenceDateMaxValue())
		}
	}
	for _, prop := range m.NoEvents {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range m.NoMDFeedTypes {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range m.NoInstAttrib {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range m.NoLotTypeRules {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range m.NoUnderlyings {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for _, prop := range m.NoRelatedInstruments {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func MDInstrumentDefinitionOption41Init(m *MDInstrumentDefinitionOption41) {
	m.TotNumReports = 4294967295
	m.SecurityIDSource[0] = 56
	m.TickRule = 127
	m.MainFraction = 255
	m.SubFraction = 255
	m.PriceDisplayFormat = 255
	m.ClearedVolume = 2147483647
	m.OpenInterestQty = 2147483647
	m.TradingReferenceDate = 65535
	return
}

func (m *MDInstrumentDefinitionOption41NoEvents) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.EventType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.EventTime); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionOption41NoEvents) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if m.EventTypeInActingVersion(actingVersion) {
		if err := m.EventType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.EventTimeInActingVersion(actingVersion) {
		m.EventTime = m.EventTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.EventTime); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDInstrumentDefinitionOption41NoEvents) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := m.EventType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.EventTimeInActingVersion(actingVersion) {
		if m.EventTime < m.EventTimeMinValue() || m.EventTime > m.EventTimeMaxValue() {
			return fmt.Errorf("Range check failed on m.EventTime (%v < %v > %v)", m.EventTimeMinValue(), m.EventTime, m.EventTimeMaxValue())
		}
	}
	return nil
}

func MDInstrumentDefinitionOption41NoEventsInit(m *MDInstrumentDefinitionOption41NoEvents) {
	return
}

func (m *MDInstrumentDefinitionOption41NoMDFeedTypes) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, m.MDFeedType[:]); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.MarketDepth); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionOption41NoMDFeedTypes) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !m.MDFeedTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			m.MDFeedType[idx] = m.MDFeedTypeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.MDFeedType[:]); err != nil {
			return err
		}
	}
	if !m.MarketDepthInActingVersion(actingVersion) {
		m.MarketDepth = m.MarketDepthNullValue()
	} else {
		if err := _m.ReadInt8(_r, &m.MarketDepth); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDInstrumentDefinitionOption41NoMDFeedTypes) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.MDFeedTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if m.MDFeedType[idx] < m.MDFeedTypeMinValue() || m.MDFeedType[idx] > m.MDFeedTypeMaxValue() {
				return fmt.Errorf("Range check failed on m.MDFeedType[%d] (%v < %v > %v)", idx, m.MDFeedTypeMinValue(), m.MDFeedType[idx], m.MDFeedTypeMaxValue())
			}
		}
	}
	if m.MarketDepthInActingVersion(actingVersion) {
		if m.MarketDepth < m.MarketDepthMinValue() || m.MarketDepth > m.MarketDepthMaxValue() {
			return fmt.Errorf("Range check failed on m.MarketDepth (%v < %v > %v)", m.MarketDepthMinValue(), m.MarketDepth, m.MarketDepthMaxValue())
		}
	}
	return nil
}

func MDInstrumentDefinitionOption41NoMDFeedTypesInit(m *MDInstrumentDefinitionOption41NoMDFeedTypes) {
	return
}

func (m *MDInstrumentDefinitionOption41NoInstAttrib) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.InstAttribValue.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionOption41NoInstAttrib) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	m.InstAttribType = 24
	if m.InstAttribValueInActingVersion(actingVersion) {
		if err := m.InstAttribValue.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDInstrumentDefinitionOption41NoInstAttrib) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func MDInstrumentDefinitionOption41NoInstAttribInit(m *MDInstrumentDefinitionOption41NoInstAttrib) {
	m.InstAttribType = 24
	return
}

func (m *MDInstrumentDefinitionOption41NoLotTypeRules) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, m.LotType); err != nil {
		return err
	}
	if err := m.MinLotSize.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionOption41NoLotTypeRules) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !m.LotTypeInActingVersion(actingVersion) {
		m.LotType = m.LotTypeNullValue()
	} else {
		if err := _m.ReadInt8(_r, &m.LotType); err != nil {
			return err
		}
	}
	if m.MinLotSizeInActingVersion(actingVersion) {
		if err := m.MinLotSize.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDInstrumentDefinitionOption41NoLotTypeRules) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.LotTypeInActingVersion(actingVersion) {
		if m.LotType < m.LotTypeMinValue() || m.LotType > m.LotTypeMaxValue() {
			return fmt.Errorf("Range check failed on m.LotType (%v < %v > %v)", m.LotTypeMinValue(), m.LotType, m.LotTypeMaxValue())
		}
	}
	return nil
}

func MDInstrumentDefinitionOption41NoLotTypeRulesInit(m *MDInstrumentDefinitionOption41NoLotTypeRules) {
	return
}

func (m *MDInstrumentDefinitionOption41NoUnderlyings) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt32(_w, m.UnderlyingSecurityID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.UnderlyingSymbol[:]); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionOption41NoUnderlyings) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !m.UnderlyingSecurityIDInActingVersion(actingVersion) {
		m.UnderlyingSecurityID = m.UnderlyingSecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.UnderlyingSecurityID); err != nil {
			return err
		}
	}
	m.UnderlyingSecurityIDSource[0] = 56
	if !m.UnderlyingSymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			m.UnderlyingSymbol[idx] = m.UnderlyingSymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.UnderlyingSymbol[:]); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDInstrumentDefinitionOption41NoUnderlyings) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.UnderlyingSecurityIDInActingVersion(actingVersion) {
		if m.UnderlyingSecurityID < m.UnderlyingSecurityIDMinValue() || m.UnderlyingSecurityID > m.UnderlyingSecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on m.UnderlyingSecurityID (%v < %v > %v)", m.UnderlyingSecurityIDMinValue(), m.UnderlyingSecurityID, m.UnderlyingSecurityIDMaxValue())
		}
	}
	if m.UnderlyingSymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if m.UnderlyingSymbol[idx] < m.UnderlyingSymbolMinValue() || m.UnderlyingSymbol[idx] > m.UnderlyingSymbolMaxValue() {
				return fmt.Errorf("Range check failed on m.UnderlyingSymbol[%d] (%v < %v > %v)", idx, m.UnderlyingSymbolMinValue(), m.UnderlyingSymbol[idx], m.UnderlyingSymbolMaxValue())
			}
		}
	}
	return nil
}

func MDInstrumentDefinitionOption41NoUnderlyingsInit(m *MDInstrumentDefinitionOption41NoUnderlyings) {
	m.UnderlyingSecurityIDSource[0] = 56
	return
}

func (m *MDInstrumentDefinitionOption41NoRelatedInstruments) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt32(_w, m.RelatedSecurityID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.RelatedSymbol[:]); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionOption41NoRelatedInstruments) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !m.RelatedSecurityIDInActingVersion(actingVersion) {
		m.RelatedSecurityID = m.RelatedSecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.RelatedSecurityID); err != nil {
			return err
		}
	}
	m.RelatedSecurityIDSource[0] = 56
	if !m.RelatedSymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			m.RelatedSymbol[idx] = m.RelatedSymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.RelatedSymbol[:]); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDInstrumentDefinitionOption41NoRelatedInstruments) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.RelatedSecurityIDInActingVersion(actingVersion) {
		if m.RelatedSecurityID < m.RelatedSecurityIDMinValue() || m.RelatedSecurityID > m.RelatedSecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on m.RelatedSecurityID (%v < %v > %v)", m.RelatedSecurityIDMinValue(), m.RelatedSecurityID, m.RelatedSecurityIDMaxValue())
		}
	}
	if m.RelatedSymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if m.RelatedSymbol[idx] < m.RelatedSymbolMinValue() || m.RelatedSymbol[idx] > m.RelatedSymbolMaxValue() {
				return fmt.Errorf("Range check failed on m.RelatedSymbol[%d] (%v < %v > %v)", idx, m.RelatedSymbolMinValue(), m.RelatedSymbol[idx], m.RelatedSymbolMaxValue())
			}
		}
	}
	return nil
}

func MDInstrumentDefinitionOption41NoRelatedInstrumentsInit(m *MDInstrumentDefinitionOption41NoRelatedInstruments) {
	m.RelatedSecurityIDSource[0] = 56
	return
}

func (*MDInstrumentDefinitionOption41) SbeBlockLength() (blockLength uint16) {
	return 213
}

func (*MDInstrumentDefinitionOption41) SbeTemplateId() (templateId uint16) {
	return 41
}

func (*MDInstrumentDefinitionOption41) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MDInstrumentDefinitionOption41) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionOption41) SbeSemanticType() (semanticType []byte) {
	return []byte("d")
}

func (*MDInstrumentDefinitionOption41) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MDInstrumentDefinitionOption41) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MDInstrumentDefinitionOption41) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) TotNumReportsId() uint16 {
	return 911
}

func (*MDInstrumentDefinitionOption41) TotNumReportsSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) TotNumReportsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TotNumReportsSinceVersion()
}

func (*MDInstrumentDefinitionOption41) TotNumReportsDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) TotNumReportsMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) TotNumReportsMinValue() uint32 {
	return 0
}

func (*MDInstrumentDefinitionOption41) TotNumReportsMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDInstrumentDefinitionOption41) TotNumReportsNullValue() uint32 {
	return 4294967295
}

func (*MDInstrumentDefinitionOption41) SecurityUpdateActionId() uint16 {
	return 980
}

func (*MDInstrumentDefinitionOption41) SecurityUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SecurityUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityUpdateActionSinceVersion()
}

func (*MDInstrumentDefinitionOption41) SecurityUpdateActionDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) SecurityUpdateActionMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) LastUpdateTimeId() uint16 {
	return 779
}

func (*MDInstrumentDefinitionOption41) LastUpdateTimeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) LastUpdateTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastUpdateTimeSinceVersion()
}

func (*MDInstrumentDefinitionOption41) LastUpdateTimeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) LastUpdateTimeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) LastUpdateTimeMinValue() uint64 {
	return 0
}

func (*MDInstrumentDefinitionOption41) LastUpdateTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDInstrumentDefinitionOption41) LastUpdateTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDInstrumentDefinitionOption41) MDSecurityTradingStatusId() uint16 {
	return 1682
}

func (*MDInstrumentDefinitionOption41) MDSecurityTradingStatusSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) MDSecurityTradingStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDSecurityTradingStatusSinceVersion()
}

func (*MDInstrumentDefinitionOption41) MDSecurityTradingStatusDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MDSecurityTradingStatusMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) ApplIDId() uint16 {
	return 1180
}

func (*MDInstrumentDefinitionOption41) ApplIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) ApplIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ApplIDSinceVersion()
}

func (*MDInstrumentDefinitionOption41) ApplIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) ApplIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) ApplIDMinValue() int16 {
	return math.MinInt16 + 1
}

func (*MDInstrumentDefinitionOption41) ApplIDMaxValue() int16 {
	return math.MaxInt16
}

func (*MDInstrumentDefinitionOption41) ApplIDNullValue() int16 {
	return math.MinInt16
}

func (*MDInstrumentDefinitionOption41) MarketSegmentIDId() uint16 {
	return 1300
}

func (*MDInstrumentDefinitionOption41) MarketSegmentIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) MarketSegmentIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MarketSegmentIDSinceVersion()
}

func (*MDInstrumentDefinitionOption41) MarketSegmentIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MarketSegmentIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) MarketSegmentIDMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MarketSegmentIDMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionOption41) MarketSegmentIDNullValue() uint8 {
	return math.MaxUint8
}

func (*MDInstrumentDefinitionOption41) UnderlyingProductId() uint16 {
	return 462
}

func (*MDInstrumentDefinitionOption41) UnderlyingProductSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) UnderlyingProductInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnderlyingProductSinceVersion()
}

func (*MDInstrumentDefinitionOption41) UnderlyingProductDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) UnderlyingProductMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) UnderlyingProductMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionOption41) UnderlyingProductMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionOption41) UnderlyingProductNullValue() uint8 {
	return math.MaxUint8
}

func (*MDInstrumentDefinitionOption41) SecurityExchangeId() uint16 {
	return 207
}

func (*MDInstrumentDefinitionOption41) SecurityExchangeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SecurityExchangeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityExchangeSinceVersion()
}

func (*MDInstrumentDefinitionOption41) SecurityExchangeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) SecurityExchangeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Exchange"
	case 4:
		return "required"
	}
	return ""
}

func (*MDInstrumentDefinitionOption41) SecurityExchangeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41) SecurityExchangeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41) SecurityExchangeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SecurityExchangeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionOption41) SecurityGroupId() uint16 {
	return 1151
}

func (*MDInstrumentDefinitionOption41) SecurityGroupSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SecurityGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityGroupSinceVersion()
}

func (*MDInstrumentDefinitionOption41) SecurityGroupDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) SecurityGroupMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) SecurityGroupMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41) SecurityGroupMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41) SecurityGroupNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SecurityGroupCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionOption41) AssetId() uint16 {
	return 6937
}

func (*MDInstrumentDefinitionOption41) AssetSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) AssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AssetSinceVersion()
}

func (*MDInstrumentDefinitionOption41) AssetDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) AssetMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) AssetMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41) AssetMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41) AssetNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionOption41) AssetCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionOption41) SymbolId() uint16 {
	return 55
}

func (*MDInstrumentDefinitionOption41) SymbolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SymbolSinceVersion()
}

func (*MDInstrumentDefinitionOption41) SymbolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) SymbolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) SymbolMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41) SymbolMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41) SymbolNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionOption41) SecurityIDId() uint16 {
	return 48
}

func (*MDInstrumentDefinitionOption41) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MDInstrumentDefinitionOption41) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) SecurityIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionOption41) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionOption41) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDInstrumentDefinitionOption41) SecurityIDSourceId() uint16 {
	return 22
}

func (*MDInstrumentDefinitionOption41) SecurityIDSourceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SecurityIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSourceSinceVersion()
}

func (*MDInstrumentDefinitionOption41) SecurityIDSourceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) SecurityIDSourceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) SecurityIDSourceMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41) SecurityIDSourceMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41) SecurityIDSourceNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionOption41) SecurityTypeId() uint16 {
	return 167
}

func (*MDInstrumentDefinitionOption41) SecurityTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SecurityTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityTypeSinceVersion()
}

func (*MDInstrumentDefinitionOption41) SecurityTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) SecurityTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) SecurityTypeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41) SecurityTypeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41) SecurityTypeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SecurityTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionOption41) CFICodeId() uint16 {
	return 461
}

func (*MDInstrumentDefinitionOption41) CFICodeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) CFICodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CFICodeSinceVersion()
}

func (*MDInstrumentDefinitionOption41) CFICodeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) CFICodeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) CFICodeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41) CFICodeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41) CFICodeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionOption41) CFICodeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionOption41) PutOrCallId() uint16 {
	return 201
}

func (*MDInstrumentDefinitionOption41) PutOrCallSinceVersion() uint16 {
	return 3
}

func (m *MDInstrumentDefinitionOption41) PutOrCallInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PutOrCallSinceVersion()
}

func (*MDInstrumentDefinitionOption41) PutOrCallDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) PutOrCallMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) MaturityMonthYearId() uint16 {
	return 200
}

func (*MDInstrumentDefinitionOption41) MaturityMonthYearSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) MaturityMonthYearInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaturityMonthYearSinceVersion()
}

func (*MDInstrumentDefinitionOption41) MaturityMonthYearDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MaturityMonthYearMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "MonthYear"
	case 4:
		return "required"
	}
	return ""
}

func (*MDInstrumentDefinitionOption41) CurrencyId() uint16 {
	return 15
}

func (*MDInstrumentDefinitionOption41) CurrencySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) CurrencyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CurrencySinceVersion()
}

func (*MDInstrumentDefinitionOption41) CurrencyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) CurrencyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Currency"
	case 4:
		return "required"
	}
	return ""
}

func (*MDInstrumentDefinitionOption41) CurrencyMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41) CurrencyMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41) CurrencyNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionOption41) CurrencyCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionOption41) StrikePriceId() uint16 {
	return 202
}

func (*MDInstrumentDefinitionOption41) StrikePriceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) StrikePriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.StrikePriceSinceVersion()
}

func (*MDInstrumentDefinitionOption41) StrikePriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) StrikePriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) StrikeCurrencyId() uint16 {
	return 947
}

func (*MDInstrumentDefinitionOption41) StrikeCurrencySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) StrikeCurrencyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.StrikeCurrencySinceVersion()
}

func (*MDInstrumentDefinitionOption41) StrikeCurrencyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) StrikeCurrencyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Currency"
	case 4:
		return "required"
	}
	return ""
}

func (*MDInstrumentDefinitionOption41) StrikeCurrencyMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41) StrikeCurrencyMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41) StrikeCurrencyNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionOption41) StrikeCurrencyCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionOption41) SettlCurrencyId() uint16 {
	return 120
}

func (*MDInstrumentDefinitionOption41) SettlCurrencySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SettlCurrencyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SettlCurrencySinceVersion()
}

func (*MDInstrumentDefinitionOption41) SettlCurrencyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) SettlCurrencyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "Currency"
	case 4:
		return "required"
	}
	return ""
}

func (*MDInstrumentDefinitionOption41) SettlCurrencyMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41) SettlCurrencyMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41) SettlCurrencyNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SettlCurrencyCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionOption41) MinCabPriceId() uint16 {
	return 9850
}

func (*MDInstrumentDefinitionOption41) MinCabPriceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) MinCabPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinCabPriceSinceVersion()
}

func (*MDInstrumentDefinitionOption41) MinCabPriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MinCabPriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) MatchAlgorithmId() uint16 {
	return 1142
}

func (*MDInstrumentDefinitionOption41) MatchAlgorithmSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) MatchAlgorithmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchAlgorithmSinceVersion()
}

func (*MDInstrumentDefinitionOption41) MatchAlgorithmDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MatchAlgorithmMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) MatchAlgorithmMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41) MatchAlgorithmMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41) MatchAlgorithmNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionOption41) MinTradeVolId() uint16 {
	return 562
}

func (*MDInstrumentDefinitionOption41) MinTradeVolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) MinTradeVolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinTradeVolSinceVersion()
}

func (*MDInstrumentDefinitionOption41) MinTradeVolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MinTradeVolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) MinTradeVolMinValue() uint32 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MinTradeVolMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDInstrumentDefinitionOption41) MinTradeVolNullValue() uint32 {
	return math.MaxUint32
}

func (*MDInstrumentDefinitionOption41) MaxTradeVolId() uint16 {
	return 1140
}

func (*MDInstrumentDefinitionOption41) MaxTradeVolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) MaxTradeVolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaxTradeVolSinceVersion()
}

func (*MDInstrumentDefinitionOption41) MaxTradeVolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MaxTradeVolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) MaxTradeVolMinValue() uint32 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MaxTradeVolMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDInstrumentDefinitionOption41) MaxTradeVolNullValue() uint32 {
	return math.MaxUint32
}

func (*MDInstrumentDefinitionOption41) MinPriceIncrementId() uint16 {
	return 969
}

func (*MDInstrumentDefinitionOption41) MinPriceIncrementSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) MinPriceIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinPriceIncrementSinceVersion()
}

func (*MDInstrumentDefinitionOption41) MinPriceIncrementDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MinPriceIncrementMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) MinPriceIncrementAmountId() uint16 {
	return 1146
}

func (*MDInstrumentDefinitionOption41) MinPriceIncrementAmountSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) MinPriceIncrementAmountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinPriceIncrementAmountSinceVersion()
}

func (*MDInstrumentDefinitionOption41) MinPriceIncrementAmountDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MinPriceIncrementAmountMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) DisplayFactorId() uint16 {
	return 9787
}

func (*MDInstrumentDefinitionOption41) DisplayFactorSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) DisplayFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DisplayFactorSinceVersion()
}

func (*MDInstrumentDefinitionOption41) DisplayFactorDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) DisplayFactorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "float"
	case 4:
		return "required"
	}
	return ""
}

func (*MDInstrumentDefinitionOption41) TickRuleId() uint16 {
	return 6350
}

func (*MDInstrumentDefinitionOption41) TickRuleSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) TickRuleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TickRuleSinceVersion()
}

func (*MDInstrumentDefinitionOption41) TickRuleDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) TickRuleMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) TickRuleMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionOption41) TickRuleMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionOption41) TickRuleNullValue() int8 {
	return 127
}

func (*MDInstrumentDefinitionOption41) MainFractionId() uint16 {
	return 37702
}

func (*MDInstrumentDefinitionOption41) MainFractionSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) MainFractionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MainFractionSinceVersion()
}

func (*MDInstrumentDefinitionOption41) MainFractionDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MainFractionMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) MainFractionMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionOption41) MainFractionMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionOption41) MainFractionNullValue() uint8 {
	return 255
}

func (*MDInstrumentDefinitionOption41) SubFractionId() uint16 {
	return 37703
}

func (*MDInstrumentDefinitionOption41) SubFractionSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SubFractionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SubFractionSinceVersion()
}

func (*MDInstrumentDefinitionOption41) SubFractionDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) SubFractionMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) SubFractionMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionOption41) SubFractionMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionOption41) SubFractionNullValue() uint8 {
	return 255
}

func (*MDInstrumentDefinitionOption41) PriceDisplayFormatId() uint16 {
	return 9800
}

func (*MDInstrumentDefinitionOption41) PriceDisplayFormatSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) PriceDisplayFormatInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceDisplayFormatSinceVersion()
}

func (*MDInstrumentDefinitionOption41) PriceDisplayFormatDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) PriceDisplayFormatMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) PriceDisplayFormatMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionOption41) PriceDisplayFormatMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionOption41) PriceDisplayFormatNullValue() uint8 {
	return 255
}

func (*MDInstrumentDefinitionOption41) UnitOfMeasureId() uint16 {
	return 996
}

func (*MDInstrumentDefinitionOption41) UnitOfMeasureSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) UnitOfMeasureInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnitOfMeasureSinceVersion()
}

func (*MDInstrumentDefinitionOption41) UnitOfMeasureDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) UnitOfMeasureMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) UnitOfMeasureMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41) UnitOfMeasureMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41) UnitOfMeasureNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionOption41) UnitOfMeasureCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionOption41) UnitOfMeasureQtyId() uint16 {
	return 1147
}

func (*MDInstrumentDefinitionOption41) UnitOfMeasureQtySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) UnitOfMeasureQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnitOfMeasureQtySinceVersion()
}

func (*MDInstrumentDefinitionOption41) UnitOfMeasureQtyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) UnitOfMeasureQtyMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) TradingReferencePriceId() uint16 {
	return 1150
}

func (*MDInstrumentDefinitionOption41) TradingReferencePriceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) TradingReferencePriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradingReferencePriceSinceVersion()
}

func (*MDInstrumentDefinitionOption41) TradingReferencePriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) TradingReferencePriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) SettlPriceTypeId() uint16 {
	return 731
}

func (*MDInstrumentDefinitionOption41) SettlPriceTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) SettlPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SettlPriceTypeSinceVersion()
}

func (*MDInstrumentDefinitionOption41) SettlPriceTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) SettlPriceTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) ClearedVolumeId() uint16 {
	return 5791
}

func (*MDInstrumentDefinitionOption41) ClearedVolumeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) ClearedVolumeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ClearedVolumeSinceVersion()
}

func (*MDInstrumentDefinitionOption41) ClearedVolumeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) ClearedVolumeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) ClearedVolumeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionOption41) ClearedVolumeMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionOption41) ClearedVolumeNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionOption41) OpenInterestQtyId() uint16 {
	return 5792
}

func (*MDInstrumentDefinitionOption41) OpenInterestQtySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) OpenInterestQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OpenInterestQtySinceVersion()
}

func (*MDInstrumentDefinitionOption41) OpenInterestQtyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) OpenInterestQtyMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) OpenInterestQtyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionOption41) OpenInterestQtyMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionOption41) OpenInterestQtyNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionOption41) LowLimitPriceId() uint16 {
	return 1148
}

func (*MDInstrumentDefinitionOption41) LowLimitPriceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) LowLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LowLimitPriceSinceVersion()
}

func (*MDInstrumentDefinitionOption41) LowLimitPriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) LowLimitPriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) HighLimitPriceId() uint16 {
	return 1149
}

func (*MDInstrumentDefinitionOption41) HighLimitPriceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) HighLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.HighLimitPriceSinceVersion()
}

func (*MDInstrumentDefinitionOption41) HighLimitPriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) HighLimitPriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) UserDefinedInstrumentId() uint16 {
	return 9779
}

func (*MDInstrumentDefinitionOption41) UserDefinedInstrumentSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) UserDefinedInstrumentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UserDefinedInstrumentSinceVersion()
}

func (*MDInstrumentDefinitionOption41) UserDefinedInstrumentDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) UserDefinedInstrumentMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) UserDefinedInstrumentMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41) UserDefinedInstrumentMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41) UserDefinedInstrumentNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionOption41) TradingReferenceDateId() uint16 {
	return 5796
}

func (*MDInstrumentDefinitionOption41) TradingReferenceDateSinceVersion() uint16 {
	return 6
}

func (m *MDInstrumentDefinitionOption41) TradingReferenceDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradingReferenceDateSinceVersion()
}

func (*MDInstrumentDefinitionOption41) TradingReferenceDateDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) TradingReferenceDateMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41) TradingReferenceDateMinValue() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41) TradingReferenceDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MDInstrumentDefinitionOption41) TradingReferenceDateNullValue() uint16 {
	return 65535
}

func (*MDInstrumentDefinitionOption41NoEvents) EventTypeId() uint16 {
	return 865
}

func (*MDInstrumentDefinitionOption41NoEvents) EventTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoEvents) EventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EventTypeSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoEvents) EventTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoEvents) EventTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoEvents) EventTimeId() uint16 {
	return 1145
}

func (*MDInstrumentDefinitionOption41NoEvents) EventTimeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoEvents) EventTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EventTimeSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoEvents) EventTimeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoEvents) EventTimeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoEvents) EventTimeMinValue() uint64 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoEvents) EventTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDInstrumentDefinitionOption41NoEvents) EventTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MDFeedTypeId() uint16 {
	return 1022
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MDFeedTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoMDFeedTypes) MDFeedTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDFeedTypeSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MDFeedTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MDFeedTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MDFeedTypeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MDFeedTypeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MDFeedTypeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoMDFeedTypes) MDFeedTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MarketDepthId() uint16 {
	return 264
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MarketDepthSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoMDFeedTypes) MarketDepthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MarketDepthSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MarketDepthDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MarketDepthMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MarketDepthMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MarketDepthMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) MarketDepthNullValue() int8 {
	return math.MinInt8
}

func (*MDInstrumentDefinitionOption41NoInstAttrib) InstAttribTypeId() uint16 {
	return 871
}

func (*MDInstrumentDefinitionOption41NoInstAttrib) InstAttribTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoInstAttrib) InstAttribTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstAttribTypeSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoInstAttrib) InstAttribTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoInstAttrib) InstAttribTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoInstAttrib) InstAttribTypeMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionOption41NoInstAttrib) InstAttribTypeMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionOption41NoInstAttrib) InstAttribTypeNullValue() int8 {
	return math.MinInt8
}

func (*MDInstrumentDefinitionOption41NoInstAttrib) InstAttribValueId() uint16 {
	return 872
}

func (*MDInstrumentDefinitionOption41NoInstAttrib) InstAttribValueSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoInstAttrib) InstAttribValueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstAttribValueSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoInstAttrib) InstAttribValueDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoInstAttrib) InstAttribValueMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoLotTypeRules) LotTypeId() uint16 {
	return 1093
}

func (*MDInstrumentDefinitionOption41NoLotTypeRules) LotTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoLotTypeRules) LotTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LotTypeSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoLotTypeRules) LotTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoLotTypeRules) LotTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoLotTypeRules) LotTypeMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionOption41NoLotTypeRules) LotTypeMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionOption41NoLotTypeRules) LotTypeNullValue() int8 {
	return math.MinInt8
}

func (*MDInstrumentDefinitionOption41NoLotTypeRules) MinLotSizeId() uint16 {
	return 1231
}

func (*MDInstrumentDefinitionOption41NoLotTypeRules) MinLotSizeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoLotTypeRules) MinLotSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinLotSizeSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoLotTypeRules) MinLotSizeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoLotTypeRules) MinLotSizeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDId() uint16 {
	return 309
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnderlyingSecurityIDSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDSourceId() uint16 {
	return 305
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDSourceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnderlyingSecurityIDSourceSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDSourceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDSourceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDSourceMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDSourceMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSecurityIDSourceNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSymbolId() uint16 {
	return 311
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSymbolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnderlyingSymbolSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSymbolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSymbolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSymbolMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSymbolMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSymbolNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoUnderlyings) UnderlyingSymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDId() uint16 {
	return 1650
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RelatedSecurityIDSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDSourceId() uint16 {
	return 1651
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDSourceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RelatedSecurityIDSourceSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDSourceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDSourceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDSourceMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDSourceMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSecurityIDSourceNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSymbolId() uint16 {
	return 1649
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSymbolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RelatedSymbolSinceVersion()
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSymbolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSymbolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSymbolMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSymbolMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSymbolNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionOption41NoRelatedInstruments) RelatedSymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionOption41) NoEventsId() uint16 {
	return 864
}

func (*MDInstrumentDefinitionOption41) NoEventsSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) NoEventsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoEventsSinceVersion()
}

func (*MDInstrumentDefinitionOption41) NoEventsDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoEvents) SbeBlockLength() (blockLength uint) {
	return 9
}

func (*MDInstrumentDefinitionOption41NoEvents) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionOption41) NoMDFeedTypesId() uint16 {
	return 1141
}

func (*MDInstrumentDefinitionOption41) NoMDFeedTypesSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) NoMDFeedTypesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoMDFeedTypesSinceVersion()
}

func (*MDInstrumentDefinitionOption41) NoMDFeedTypesDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) SbeBlockLength() (blockLength uint) {
	return 4
}

func (*MDInstrumentDefinitionOption41NoMDFeedTypes) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionOption41) NoInstAttribId() uint16 {
	return 870
}

func (*MDInstrumentDefinitionOption41) NoInstAttribSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) NoInstAttribInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoInstAttribSinceVersion()
}

func (*MDInstrumentDefinitionOption41) NoInstAttribDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoInstAttrib) SbeBlockLength() (blockLength uint) {
	return 4
}

func (*MDInstrumentDefinitionOption41NoInstAttrib) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionOption41) NoLotTypeRulesId() uint16 {
	return 1234
}

func (*MDInstrumentDefinitionOption41) NoLotTypeRulesSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) NoLotTypeRulesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoLotTypeRulesSinceVersion()
}

func (*MDInstrumentDefinitionOption41) NoLotTypeRulesDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoLotTypeRules) SbeBlockLength() (blockLength uint) {
	return 5
}

func (*MDInstrumentDefinitionOption41NoLotTypeRules) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionOption41) NoUnderlyingsId() uint16 {
	return 711
}

func (*MDInstrumentDefinitionOption41) NoUnderlyingsSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionOption41) NoUnderlyingsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoUnderlyingsSinceVersion()
}

func (*MDInstrumentDefinitionOption41) NoUnderlyingsDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) SbeBlockLength() (blockLength uint) {
	return 24
}

func (*MDInstrumentDefinitionOption41NoUnderlyings) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionOption41) NoRelatedInstrumentsId() uint16 {
	return 1647
}

func (*MDInstrumentDefinitionOption41) NoRelatedInstrumentsSinceVersion() uint16 {
	return 7
}

func (m *MDInstrumentDefinitionOption41) NoRelatedInstrumentsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoRelatedInstrumentsSinceVersion()
}

func (*MDInstrumentDefinitionOption41) NoRelatedInstrumentsDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) SbeBlockLength() (blockLength uint) {
	return 24
}

func (*MDInstrumentDefinitionOption41NoRelatedInstruments) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
