// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MDInstrumentDefinitionSpread29 struct {
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
	MaturityMonthYear       MaturityMonthYear
	Currency                [3]byte
	SecuritySubType         [5]byte
	UserDefinedInstrument   byte
	MatchAlgorithm          byte
	MinTradeVol             uint32
	MaxTradeVol             uint32
	MinPriceIncrement       PRICE
	DisplayFactor           FLOAT
	PriceDisplayFormat      uint8
	PriceRatio              PRICENULL
	TickRule                int8
	UnitOfMeasure           [30]byte
	TradingReferencePrice   PRICENULL
	SettlPriceType          SettlPriceType
	OpenInterestQty         int32
	ClearedVolume           int32
	HighLimitPrice          PRICENULL
	LowLimitPrice           PRICENULL
	MaxPriceVariation       PRICENULL
	MainFraction            uint8
	SubFraction             uint8
	TradingReferenceDate    uint16
	NoEvents                []MDInstrumentDefinitionSpread29NoEvents
	NoMDFeedTypes           []MDInstrumentDefinitionSpread29NoMDFeedTypes
	NoInstAttrib            []MDInstrumentDefinitionSpread29NoInstAttrib
	NoLotTypeRules          []MDInstrumentDefinitionSpread29NoLotTypeRules
	NoLegs                  []MDInstrumentDefinitionSpread29NoLegs
}
type MDInstrumentDefinitionSpread29NoEvents struct {
	EventType EventTypeEnum
	EventTime uint64
}
type MDInstrumentDefinitionSpread29NoMDFeedTypes struct {
	MDFeedType  [3]byte
	MarketDepth int8
}
type MDInstrumentDefinitionSpread29NoInstAttrib struct {
	InstAttribType  int8
	InstAttribValue InstAttribValue
}
type MDInstrumentDefinitionSpread29NoLotTypeRules struct {
	LotType    int8
	MinLotSize DecimalQty
}
type MDInstrumentDefinitionSpread29NoLegs struct {
	LegSecurityID       int32
	LegSecurityIDSource [1]byte
	LegSide             LegSideEnum
	LegRatioQty         int8
	LegPrice            PRICENULL
	LegOptionDelta      DecimalQty
}

func (m *MDInstrumentDefinitionSpread29) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := m.MaturityMonthYear.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.Currency[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.SecuritySubType[:]); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.UserDefinedInstrument); err != nil {
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
	if err := m.DisplayFactor.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.PriceDisplayFormat); err != nil {
		return err
	}
	if err := m.PriceRatio.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.TickRule); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.UnitOfMeasure[:]); err != nil {
		return err
	}
	if err := m.TradingReferencePrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.SettlPriceType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.OpenInterestQty); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.ClearedVolume); err != nil {
		return err
	}
	if err := m.HighLimitPrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.LowLimitPrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MaxPriceVariation.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.MainFraction); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, m.SubFraction); err != nil {
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
	var NoLegsBlockLength uint16 = 18
	if err := _m.WriteUint16(_w, NoLegsBlockLength); err != nil {
		return err
	}
	var NoLegsNumInGroup uint8 = uint8(len(m.NoLegs))
	if err := _m.WriteUint8(_w, NoLegsNumInGroup); err != nil {
		return err
	}
	for _, prop := range m.NoLegs {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (m *MDInstrumentDefinitionSpread29) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !m.SecuritySubTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			m.SecuritySubType[idx] = m.SecuritySubTypeNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.SecuritySubType[:]); err != nil {
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
	if m.DisplayFactorInActingVersion(actingVersion) {
		if err := m.DisplayFactor.Decode(_m, _r, actingVersion); err != nil {
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
	if m.PriceRatioInActingVersion(actingVersion) {
		if err := m.PriceRatio.Decode(_m, _r, actingVersion); err != nil {
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
	if !m.UnitOfMeasureInActingVersion(actingVersion) {
		for idx := 0; idx < 30; idx++ {
			m.UnitOfMeasure[idx] = m.UnitOfMeasureNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.UnitOfMeasure[:]); err != nil {
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
	if !m.OpenInterestQtyInActingVersion(actingVersion) {
		m.OpenInterestQty = m.OpenInterestQtyNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.OpenInterestQty); err != nil {
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
			m.NoEvents = make([]MDInstrumentDefinitionSpread29NoEvents, NoEventsNumInGroup)
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
			m.NoMDFeedTypes = make([]MDInstrumentDefinitionSpread29NoMDFeedTypes, NoMDFeedTypesNumInGroup)
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
			m.NoInstAttrib = make([]MDInstrumentDefinitionSpread29NoInstAttrib, NoInstAttribNumInGroup)
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
			m.NoLotTypeRules = make([]MDInstrumentDefinitionSpread29NoLotTypeRules, NoLotTypeRulesNumInGroup)
		}
		m.NoLotTypeRules = m.NoLotTypeRules[:NoLotTypeRulesNumInGroup]
		for i, _ := range m.NoLotTypeRules {
			if err := m.NoLotTypeRules[i].Decode(_m, _r, actingVersion, uint(NoLotTypeRulesBlockLength)); err != nil {
				return err
			}
		}
	}

	if m.NoLegsInActingVersion(actingVersion) {
		var NoLegsBlockLength uint16
		if err := _m.ReadUint16(_r, &NoLegsBlockLength); err != nil {
			return err
		}
		var NoLegsNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoLegsNumInGroup); err != nil {
			return err
		}
		if cap(m.NoLegs) < int(NoLegsNumInGroup) {
			m.NoLegs = make([]MDInstrumentDefinitionSpread29NoLegs, NoLegsNumInGroup)
		}
		m.NoLegs = m.NoLegs[:NoLegsNumInGroup]
		for i, _ := range m.NoLegs {
			if err := m.NoLegs[i].Decode(_m, _r, actingVersion, uint(NoLegsBlockLength)); err != nil {
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

func (m *MDInstrumentDefinitionSpread29) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
		if m.UnderlyingProduct != m.UnderlyingProductNullValue() && (m.UnderlyingProduct < m.UnderlyingProductMinValue() || m.UnderlyingProduct > m.UnderlyingProductMaxValue()) {
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
	if m.CurrencyInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if m.Currency[idx] < m.CurrencyMinValue() || m.Currency[idx] > m.CurrencyMaxValue() {
				return fmt.Errorf("Range check failed on m.Currency[%d] (%v < %v > %v)", idx, m.CurrencyMinValue(), m.Currency[idx], m.CurrencyMaxValue())
			}
		}
	}
	if m.SecuritySubTypeInActingVersion(actingVersion) {
		for idx := 0; idx < 5; idx++ {
			if m.SecuritySubType[idx] < m.SecuritySubTypeMinValue() || m.SecuritySubType[idx] > m.SecuritySubTypeMaxValue() {
				return fmt.Errorf("Range check failed on m.SecuritySubType[%d] (%v < %v > %v)", idx, m.SecuritySubTypeMinValue(), m.SecuritySubType[idx], m.SecuritySubTypeMaxValue())
			}
		}
	}
	if m.UserDefinedInstrumentInActingVersion(actingVersion) {
		if m.UserDefinedInstrument < m.UserDefinedInstrumentMinValue() || m.UserDefinedInstrument > m.UserDefinedInstrumentMaxValue() {
			return fmt.Errorf("Range check failed on m.UserDefinedInstrument (%v < %v > %v)", m.UserDefinedInstrumentMinValue(), m.UserDefinedInstrument, m.UserDefinedInstrumentMaxValue())
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
	if m.PriceDisplayFormatInActingVersion(actingVersion) {
		if m.PriceDisplayFormat != m.PriceDisplayFormatNullValue() && (m.PriceDisplayFormat < m.PriceDisplayFormatMinValue() || m.PriceDisplayFormat > m.PriceDisplayFormatMaxValue()) {
			return fmt.Errorf("Range check failed on m.PriceDisplayFormat (%v < %v > %v)", m.PriceDisplayFormatMinValue(), m.PriceDisplayFormat, m.PriceDisplayFormatMaxValue())
		}
	}
	if m.TickRuleInActingVersion(actingVersion) {
		if m.TickRule != m.TickRuleNullValue() && (m.TickRule < m.TickRuleMinValue() || m.TickRule > m.TickRuleMaxValue()) {
			return fmt.Errorf("Range check failed on m.TickRule (%v < %v > %v)", m.TickRuleMinValue(), m.TickRule, m.TickRuleMaxValue())
		}
	}
	if m.UnitOfMeasureInActingVersion(actingVersion) {
		for idx := 0; idx < 30; idx++ {
			if m.UnitOfMeasure[idx] < m.UnitOfMeasureMinValue() || m.UnitOfMeasure[idx] > m.UnitOfMeasureMaxValue() {
				return fmt.Errorf("Range check failed on m.UnitOfMeasure[%d] (%v < %v > %v)", idx, m.UnitOfMeasureMinValue(), m.UnitOfMeasure[idx], m.UnitOfMeasureMaxValue())
			}
		}
	}
	if m.OpenInterestQtyInActingVersion(actingVersion) {
		if m.OpenInterestQty != m.OpenInterestQtyNullValue() && (m.OpenInterestQty < m.OpenInterestQtyMinValue() || m.OpenInterestQty > m.OpenInterestQtyMaxValue()) {
			return fmt.Errorf("Range check failed on m.OpenInterestQty (%v < %v > %v)", m.OpenInterestQtyMinValue(), m.OpenInterestQty, m.OpenInterestQtyMaxValue())
		}
	}
	if m.ClearedVolumeInActingVersion(actingVersion) {
		if m.ClearedVolume != m.ClearedVolumeNullValue() && (m.ClearedVolume < m.ClearedVolumeMinValue() || m.ClearedVolume > m.ClearedVolumeMaxValue()) {
			return fmt.Errorf("Range check failed on m.ClearedVolume (%v < %v > %v)", m.ClearedVolumeMinValue(), m.ClearedVolume, m.ClearedVolumeMaxValue())
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
	for _, prop := range m.NoLegs {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func MDInstrumentDefinitionSpread29Init(m *MDInstrumentDefinitionSpread29) {
	m.TotNumReports = 4294967295
	m.UnderlyingProduct = 255
	m.SecurityIDSource[0] = 56
	m.PriceDisplayFormat = 255
	m.TickRule = 127
	m.OpenInterestQty = 2147483647
	m.ClearedVolume = 2147483647
	m.MainFraction = 255
	m.SubFraction = 255
	m.TradingReferenceDate = 65535
	return
}

func (m *MDInstrumentDefinitionSpread29NoEvents) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.EventType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.EventTime); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionSpread29NoEvents) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MDInstrumentDefinitionSpread29NoEvents) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDInstrumentDefinitionSpread29NoEventsInit(m *MDInstrumentDefinitionSpread29NoEvents) {
	return
}

func (m *MDInstrumentDefinitionSpread29NoMDFeedTypes) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, m.MDFeedType[:]); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.MarketDepth); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionSpread29NoMDFeedTypes) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MDInstrumentDefinitionSpread29NoMDFeedTypes) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDInstrumentDefinitionSpread29NoMDFeedTypesInit(m *MDInstrumentDefinitionSpread29NoMDFeedTypes) {
	return
}

func (m *MDInstrumentDefinitionSpread29NoInstAttrib) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.InstAttribValue.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionSpread29NoInstAttrib) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MDInstrumentDefinitionSpread29NoInstAttrib) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func MDInstrumentDefinitionSpread29NoInstAttribInit(m *MDInstrumentDefinitionSpread29NoInstAttrib) {
	m.InstAttribType = 24
	return
}

func (m *MDInstrumentDefinitionSpread29NoLotTypeRules) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, m.LotType); err != nil {
		return err
	}
	if err := m.MinLotSize.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionSpread29NoLotTypeRules) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MDInstrumentDefinitionSpread29NoLotTypeRules) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.LotTypeInActingVersion(actingVersion) {
		if m.LotType < m.LotTypeMinValue() || m.LotType > m.LotTypeMaxValue() {
			return fmt.Errorf("Range check failed on m.LotType (%v < %v > %v)", m.LotTypeMinValue(), m.LotType, m.LotTypeMaxValue())
		}
	}
	return nil
}

func MDInstrumentDefinitionSpread29NoLotTypeRulesInit(m *MDInstrumentDefinitionSpread29NoLotTypeRules) {
	return
}

func (m *MDInstrumentDefinitionSpread29NoLegs) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt32(_w, m.LegSecurityID); err != nil {
		return err
	}
	if err := m.LegSide.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.LegRatioQty); err != nil {
		return err
	}
	if err := m.LegPrice.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.LegOptionDelta.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionSpread29NoLegs) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !m.LegSecurityIDInActingVersion(actingVersion) {
		m.LegSecurityID = m.LegSecurityIDNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.LegSecurityID); err != nil {
			return err
		}
	}
	m.LegSecurityIDSource[0] = 56
	if m.LegSideInActingVersion(actingVersion) {
		if err := m.LegSide.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.LegRatioQtyInActingVersion(actingVersion) {
		m.LegRatioQty = m.LegRatioQtyNullValue()
	} else {
		if err := _m.ReadInt8(_r, &m.LegRatioQty); err != nil {
			return err
		}
	}
	if m.LegPriceInActingVersion(actingVersion) {
		if err := m.LegPrice.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.LegOptionDeltaInActingVersion(actingVersion) {
		if err := m.LegOptionDelta.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MDInstrumentDefinitionSpread29NoLegs) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.LegSecurityIDInActingVersion(actingVersion) {
		if m.LegSecurityID < m.LegSecurityIDMinValue() || m.LegSecurityID > m.LegSecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on m.LegSecurityID (%v < %v > %v)", m.LegSecurityIDMinValue(), m.LegSecurityID, m.LegSecurityIDMaxValue())
		}
	}
	if err := m.LegSide.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.LegRatioQtyInActingVersion(actingVersion) {
		if m.LegRatioQty < m.LegRatioQtyMinValue() || m.LegRatioQty > m.LegRatioQtyMaxValue() {
			return fmt.Errorf("Range check failed on m.LegRatioQty (%v < %v > %v)", m.LegRatioQtyMinValue(), m.LegRatioQty, m.LegRatioQtyMaxValue())
		}
	}
	return nil
}

func MDInstrumentDefinitionSpread29NoLegsInit(m *MDInstrumentDefinitionSpread29NoLegs) {
	m.LegSecurityIDSource[0] = 56
	return
}

func (*MDInstrumentDefinitionSpread29) SbeBlockLength() (blockLength uint16) {
	return 195
}

func (*MDInstrumentDefinitionSpread29) SbeTemplateId() (templateId uint16) {
	return 29
}

func (*MDInstrumentDefinitionSpread29) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MDInstrumentDefinitionSpread29) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionSpread29) SbeSemanticType() (semanticType []byte) {
	return []byte("d")
}

func (*MDInstrumentDefinitionSpread29) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MDInstrumentDefinitionSpread29) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) TotNumReportsId() uint16 {
	return 911
}

func (*MDInstrumentDefinitionSpread29) TotNumReportsSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) TotNumReportsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TotNumReportsSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) TotNumReportsDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) TotNumReportsMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) TotNumReportsMinValue() uint32 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) TotNumReportsMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDInstrumentDefinitionSpread29) TotNumReportsNullValue() uint32 {
	return 4294967295
}

func (*MDInstrumentDefinitionSpread29) SecurityUpdateActionId() uint16 {
	return 980
}

func (*MDInstrumentDefinitionSpread29) SecurityUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SecurityUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityUpdateActionSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) SecurityUpdateActionDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) SecurityUpdateActionMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) LastUpdateTimeId() uint16 {
	return 779
}

func (*MDInstrumentDefinitionSpread29) LastUpdateTimeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) LastUpdateTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastUpdateTimeSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) LastUpdateTimeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) LastUpdateTimeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) LastUpdateTimeMinValue() uint64 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) LastUpdateTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDInstrumentDefinitionSpread29) LastUpdateTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDInstrumentDefinitionSpread29) MDSecurityTradingStatusId() uint16 {
	return 1682
}

func (*MDInstrumentDefinitionSpread29) MDSecurityTradingStatusSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) MDSecurityTradingStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDSecurityTradingStatusSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) MDSecurityTradingStatusDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MDSecurityTradingStatusMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) ApplIDId() uint16 {
	return 1180
}

func (*MDInstrumentDefinitionSpread29) ApplIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) ApplIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ApplIDSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) ApplIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) ApplIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) ApplIDMinValue() int16 {
	return math.MinInt16 + 1
}

func (*MDInstrumentDefinitionSpread29) ApplIDMaxValue() int16 {
	return math.MaxInt16
}

func (*MDInstrumentDefinitionSpread29) ApplIDNullValue() int16 {
	return math.MinInt16
}

func (*MDInstrumentDefinitionSpread29) MarketSegmentIDId() uint16 {
	return 1300
}

func (*MDInstrumentDefinitionSpread29) MarketSegmentIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) MarketSegmentIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MarketSegmentIDSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) MarketSegmentIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MarketSegmentIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) MarketSegmentIDMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MarketSegmentIDMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionSpread29) MarketSegmentIDNullValue() uint8 {
	return math.MaxUint8
}

func (*MDInstrumentDefinitionSpread29) UnderlyingProductId() uint16 {
	return 462
}

func (*MDInstrumentDefinitionSpread29) UnderlyingProductSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) UnderlyingProductInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnderlyingProductSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) UnderlyingProductDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) UnderlyingProductMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) UnderlyingProductMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) UnderlyingProductMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionSpread29) UnderlyingProductNullValue() uint8 {
	return 255
}

func (*MDInstrumentDefinitionSpread29) SecurityExchangeId() uint16 {
	return 207
}

func (*MDInstrumentDefinitionSpread29) SecurityExchangeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SecurityExchangeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityExchangeSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) SecurityExchangeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) SecurityExchangeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) SecurityExchangeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29) SecurityExchangeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29) SecurityExchangeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SecurityExchangeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionSpread29) SecurityGroupId() uint16 {
	return 1151
}

func (*MDInstrumentDefinitionSpread29) SecurityGroupSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SecurityGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityGroupSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) SecurityGroupDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) SecurityGroupMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) SecurityGroupMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29) SecurityGroupMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29) SecurityGroupNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SecurityGroupCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionSpread29) AssetId() uint16 {
	return 6937
}

func (*MDInstrumentDefinitionSpread29) AssetSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) AssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AssetSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) AssetDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) AssetMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) AssetMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29) AssetMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29) AssetNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) AssetCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionSpread29) SymbolId() uint16 {
	return 55
}

func (*MDInstrumentDefinitionSpread29) SymbolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SymbolSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) SymbolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) SymbolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) SymbolMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29) SymbolMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29) SymbolNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionSpread29) SecurityIDId() uint16 {
	return 48
}

func (*MDInstrumentDefinitionSpread29) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) SecurityIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionSpread29) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionSpread29) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDInstrumentDefinitionSpread29) SecurityIDSourceId() uint16 {
	return 22
}

func (*MDInstrumentDefinitionSpread29) SecurityIDSourceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SecurityIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSourceSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) SecurityIDSourceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) SecurityIDSourceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) SecurityIDSourceMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29) SecurityIDSourceMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29) SecurityIDSourceNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionSpread29) SecurityTypeId() uint16 {
	return 167
}

func (*MDInstrumentDefinitionSpread29) SecurityTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SecurityTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityTypeSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) SecurityTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) SecurityTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) SecurityTypeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29) SecurityTypeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29) SecurityTypeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SecurityTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionSpread29) CFICodeId() uint16 {
	return 461
}

func (*MDInstrumentDefinitionSpread29) CFICodeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) CFICodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CFICodeSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) CFICodeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) CFICodeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) CFICodeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29) CFICodeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29) CFICodeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) CFICodeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionSpread29) MaturityMonthYearId() uint16 {
	return 200
}

func (*MDInstrumentDefinitionSpread29) MaturityMonthYearSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) MaturityMonthYearInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaturityMonthYearSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) MaturityMonthYearDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MaturityMonthYearMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) CurrencyId() uint16 {
	return 15
}

func (*MDInstrumentDefinitionSpread29) CurrencySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) CurrencyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CurrencySinceVersion()
}

func (*MDInstrumentDefinitionSpread29) CurrencyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) CurrencyMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) CurrencyMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29) CurrencyMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29) CurrencyNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) CurrencyCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionSpread29) SecuritySubTypeId() uint16 {
	return 762
}

func (*MDInstrumentDefinitionSpread29) SecuritySubTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SecuritySubTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecuritySubTypeSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) SecuritySubTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) SecuritySubTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) SecuritySubTypeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29) SecuritySubTypeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29) SecuritySubTypeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SecuritySubTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionSpread29) UserDefinedInstrumentId() uint16 {
	return 9779
}

func (*MDInstrumentDefinitionSpread29) UserDefinedInstrumentSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) UserDefinedInstrumentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UserDefinedInstrumentSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) UserDefinedInstrumentDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) UserDefinedInstrumentMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) UserDefinedInstrumentMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29) UserDefinedInstrumentMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29) UserDefinedInstrumentNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MatchAlgorithmId() uint16 {
	return 1142
}

func (*MDInstrumentDefinitionSpread29) MatchAlgorithmSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) MatchAlgorithmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchAlgorithmSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) MatchAlgorithmDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MatchAlgorithmMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) MatchAlgorithmMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29) MatchAlgorithmMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29) MatchAlgorithmNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MinTradeVolId() uint16 {
	return 562
}

func (*MDInstrumentDefinitionSpread29) MinTradeVolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) MinTradeVolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinTradeVolSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) MinTradeVolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MinTradeVolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) MinTradeVolMinValue() uint32 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MinTradeVolMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDInstrumentDefinitionSpread29) MinTradeVolNullValue() uint32 {
	return math.MaxUint32
}

func (*MDInstrumentDefinitionSpread29) MaxTradeVolId() uint16 {
	return 1140
}

func (*MDInstrumentDefinitionSpread29) MaxTradeVolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) MaxTradeVolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaxTradeVolSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) MaxTradeVolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MaxTradeVolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) MaxTradeVolMinValue() uint32 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MaxTradeVolMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDInstrumentDefinitionSpread29) MaxTradeVolNullValue() uint32 {
	return math.MaxUint32
}

func (*MDInstrumentDefinitionSpread29) MinPriceIncrementId() uint16 {
	return 969
}

func (*MDInstrumentDefinitionSpread29) MinPriceIncrementSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) MinPriceIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinPriceIncrementSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) MinPriceIncrementDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MinPriceIncrementMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) DisplayFactorId() uint16 {
	return 9787
}

func (*MDInstrumentDefinitionSpread29) DisplayFactorSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) DisplayFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DisplayFactorSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) DisplayFactorDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) DisplayFactorMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) PriceDisplayFormatId() uint16 {
	return 9800
}

func (*MDInstrumentDefinitionSpread29) PriceDisplayFormatSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) PriceDisplayFormatInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceDisplayFormatSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) PriceDisplayFormatDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) PriceDisplayFormatMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) PriceDisplayFormatMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) PriceDisplayFormatMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionSpread29) PriceDisplayFormatNullValue() uint8 {
	return 255
}

func (*MDInstrumentDefinitionSpread29) PriceRatioId() uint16 {
	return 5770
}

func (*MDInstrumentDefinitionSpread29) PriceRatioSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) PriceRatioInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceRatioSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) PriceRatioDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) PriceRatioMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) TickRuleId() uint16 {
	return 6350
}

func (*MDInstrumentDefinitionSpread29) TickRuleSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) TickRuleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TickRuleSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) TickRuleDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) TickRuleMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) TickRuleMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionSpread29) TickRuleMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionSpread29) TickRuleNullValue() int8 {
	return 127
}

func (*MDInstrumentDefinitionSpread29) UnitOfMeasureId() uint16 {
	return 996
}

func (*MDInstrumentDefinitionSpread29) UnitOfMeasureSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) UnitOfMeasureInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnitOfMeasureSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) UnitOfMeasureDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) UnitOfMeasureMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) UnitOfMeasureMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29) UnitOfMeasureMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29) UnitOfMeasureNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) UnitOfMeasureCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionSpread29) TradingReferencePriceId() uint16 {
	return 1150
}

func (*MDInstrumentDefinitionSpread29) TradingReferencePriceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) TradingReferencePriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradingReferencePriceSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) TradingReferencePriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) TradingReferencePriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) SettlPriceTypeId() uint16 {
	return 731
}

func (*MDInstrumentDefinitionSpread29) SettlPriceTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SettlPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SettlPriceTypeSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) SettlPriceTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) SettlPriceTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) OpenInterestQtyId() uint16 {
	return 5792
}

func (*MDInstrumentDefinitionSpread29) OpenInterestQtySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) OpenInterestQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OpenInterestQtySinceVersion()
}

func (*MDInstrumentDefinitionSpread29) OpenInterestQtyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) OpenInterestQtyMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) OpenInterestQtyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionSpread29) OpenInterestQtyMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionSpread29) OpenInterestQtyNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionSpread29) ClearedVolumeId() uint16 {
	return 5791
}

func (*MDInstrumentDefinitionSpread29) ClearedVolumeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) ClearedVolumeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ClearedVolumeSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) ClearedVolumeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) ClearedVolumeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) ClearedVolumeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionSpread29) ClearedVolumeMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionSpread29) ClearedVolumeNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionSpread29) HighLimitPriceId() uint16 {
	return 1149
}

func (*MDInstrumentDefinitionSpread29) HighLimitPriceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) HighLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.HighLimitPriceSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) HighLimitPriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) HighLimitPriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) LowLimitPriceId() uint16 {
	return 1148
}

func (*MDInstrumentDefinitionSpread29) LowLimitPriceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) LowLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LowLimitPriceSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) LowLimitPriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) LowLimitPriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) MaxPriceVariationId() uint16 {
	return 1143
}

func (*MDInstrumentDefinitionSpread29) MaxPriceVariationSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) MaxPriceVariationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaxPriceVariationSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) MaxPriceVariationDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MaxPriceVariationMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) MainFractionId() uint16 {
	return 37702
}

func (*MDInstrumentDefinitionSpread29) MainFractionSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) MainFractionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MainFractionSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) MainFractionDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MainFractionMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) MainFractionMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) MainFractionMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionSpread29) MainFractionNullValue() uint8 {
	return 255
}

func (*MDInstrumentDefinitionSpread29) SubFractionId() uint16 {
	return 37703
}

func (*MDInstrumentDefinitionSpread29) SubFractionSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) SubFractionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SubFractionSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) SubFractionDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) SubFractionMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) SubFractionMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) SubFractionMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionSpread29) SubFractionNullValue() uint8 {
	return 255
}

func (*MDInstrumentDefinitionSpread29) TradingReferenceDateId() uint16 {
	return 5796
}

func (*MDInstrumentDefinitionSpread29) TradingReferenceDateSinceVersion() uint16 {
	return 6
}

func (m *MDInstrumentDefinitionSpread29) TradingReferenceDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradingReferenceDateSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) TradingReferenceDateDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) TradingReferenceDateMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) TradingReferenceDateMinValue() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29) TradingReferenceDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MDInstrumentDefinitionSpread29) TradingReferenceDateNullValue() uint16 {
	return 65535
}

func (*MDInstrumentDefinitionSpread29NoEvents) EventTypeId() uint16 {
	return 865
}

func (*MDInstrumentDefinitionSpread29NoEvents) EventTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoEvents) EventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EventTypeSinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoEvents) EventTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoEvents) EventTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29NoEvents) EventTimeId() uint16 {
	return 1145
}

func (*MDInstrumentDefinitionSpread29NoEvents) EventTimeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoEvents) EventTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EventTimeSinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoEvents) EventTimeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoEvents) EventTimeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29NoEvents) EventTimeMinValue() uint64 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoEvents) EventTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDInstrumentDefinitionSpread29NoEvents) EventTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MDFeedTypeId() uint16 {
	return 1022
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MDFeedTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoMDFeedTypes) MDFeedTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDFeedTypeSinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MDFeedTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MDFeedTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MDFeedTypeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MDFeedTypeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MDFeedTypeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoMDFeedTypes) MDFeedTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MarketDepthId() uint16 {
	return 264
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MarketDepthSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoMDFeedTypes) MarketDepthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MarketDepthSinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MarketDepthDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MarketDepthMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MarketDepthMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MarketDepthMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) MarketDepthNullValue() int8 {
	return math.MinInt8
}

func (*MDInstrumentDefinitionSpread29NoInstAttrib) InstAttribTypeId() uint16 {
	return 871
}

func (*MDInstrumentDefinitionSpread29NoInstAttrib) InstAttribTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoInstAttrib) InstAttribTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstAttribTypeSinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoInstAttrib) InstAttribTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoInstAttrib) InstAttribTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29NoInstAttrib) InstAttribTypeMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionSpread29NoInstAttrib) InstAttribTypeMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionSpread29NoInstAttrib) InstAttribTypeNullValue() int8 {
	return math.MinInt8
}

func (*MDInstrumentDefinitionSpread29NoInstAttrib) InstAttribValueId() uint16 {
	return 872
}

func (*MDInstrumentDefinitionSpread29NoInstAttrib) InstAttribValueSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoInstAttrib) InstAttribValueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstAttribValueSinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoInstAttrib) InstAttribValueDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoInstAttrib) InstAttribValueMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29NoLotTypeRules) LotTypeId() uint16 {
	return 1093
}

func (*MDInstrumentDefinitionSpread29NoLotTypeRules) LotTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoLotTypeRules) LotTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LotTypeSinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoLotTypeRules) LotTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoLotTypeRules) LotTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29NoLotTypeRules) LotTypeMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionSpread29NoLotTypeRules) LotTypeMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionSpread29NoLotTypeRules) LotTypeNullValue() int8 {
	return math.MinInt8
}

func (*MDInstrumentDefinitionSpread29NoLotTypeRules) MinLotSizeId() uint16 {
	return 1231
}

func (*MDInstrumentDefinitionSpread29NoLotTypeRules) MinLotSizeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoLotTypeRules) MinLotSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinLotSizeSinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoLotTypeRules) MinLotSizeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoLotTypeRules) MinLotSizeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDId() uint16 {
	return 602
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LegSecurityIDSinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDSourceId() uint16 {
	return 603
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDSourceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LegSecurityIDSourceSinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDSourceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDSourceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDSourceMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDSourceMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSecurityIDSourceNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSideId() uint16 {
	return 624
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSideSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoLegs) LegSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LegSideSinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSideDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegSideMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29NoLegs) LegRatioQtyId() uint16 {
	return 623
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegRatioQtySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoLegs) LegRatioQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LegRatioQtySinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegRatioQtyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegRatioQtyMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29NoLegs) LegRatioQtyMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegRatioQtyMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegRatioQtyNullValue() int8 {
	return math.MinInt8
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegPriceId() uint16 {
	return 566
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegPriceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoLegs) LegPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LegPriceSinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegPriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegPriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29NoLegs) LegOptionDeltaId() uint16 {
	return 1017
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegOptionDeltaSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29NoLegs) LegOptionDeltaInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LegOptionDeltaSinceVersion()
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegOptionDeltaDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoLegs) LegOptionDeltaMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionSpread29) NoEventsId() uint16 {
	return 864
}

func (*MDInstrumentDefinitionSpread29) NoEventsSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) NoEventsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoEventsSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) NoEventsDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoEvents) SbeBlockLength() (blockLength uint) {
	return 9
}

func (*MDInstrumentDefinitionSpread29NoEvents) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionSpread29) NoMDFeedTypesId() uint16 {
	return 1141
}

func (*MDInstrumentDefinitionSpread29) NoMDFeedTypesSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) NoMDFeedTypesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoMDFeedTypesSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) NoMDFeedTypesDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) SbeBlockLength() (blockLength uint) {
	return 4
}

func (*MDInstrumentDefinitionSpread29NoMDFeedTypes) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionSpread29) NoInstAttribId() uint16 {
	return 870
}

func (*MDInstrumentDefinitionSpread29) NoInstAttribSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) NoInstAttribInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoInstAttribSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) NoInstAttribDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoInstAttrib) SbeBlockLength() (blockLength uint) {
	return 4
}

func (*MDInstrumentDefinitionSpread29NoInstAttrib) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionSpread29) NoLotTypeRulesId() uint16 {
	return 1234
}

func (*MDInstrumentDefinitionSpread29) NoLotTypeRulesSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) NoLotTypeRulesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoLotTypeRulesSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) NoLotTypeRulesDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoLotTypeRules) SbeBlockLength() (blockLength uint) {
	return 5
}

func (*MDInstrumentDefinitionSpread29NoLotTypeRules) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionSpread29) NoLegsId() uint16 {
	return 555
}

func (*MDInstrumentDefinitionSpread29) NoLegsSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionSpread29) NoLegsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoLegsSinceVersion()
}

func (*MDInstrumentDefinitionSpread29) NoLegsDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionSpread29NoLegs) SbeBlockLength() (blockLength uint) {
	return 18
}

func (*MDInstrumentDefinitionSpread29NoLegs) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
