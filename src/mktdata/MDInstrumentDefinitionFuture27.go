// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MDInstrumentDefinitionFuture27 struct {
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
	SettlCurrency           [3]byte
	MatchAlgorithm          byte
	MinTradeVol             uint32
	MaxTradeVol             uint32
	MinPriceIncrement       PRICE
	DisplayFactor           FLOAT
	MainFraction            uint8
	SubFraction             uint8
	PriceDisplayFormat      uint8
	UnitOfMeasure           [30]byte
	UnitOfMeasureQty        PRICENULL
	TradingReferencePrice   PRICENULL
	SettlPriceType          SettlPriceType
	OpenInterestQty         int32
	ClearedVolume           int32
	HighLimitPrice          PRICENULL
	LowLimitPrice           PRICENULL
	MaxPriceVariation       PRICENULL
	DecayQuantity           int32
	DecayStartDate          uint16
	OriginalContractSize    int32
	ContractMultiplier      int32
	ContractMultiplierUnit  int8
	FlowScheduleType        int8
	MinPriceIncrementAmount PRICENULL
	UserDefinedInstrument   byte
	TradingReferenceDate    uint16
	NoEvents                []MDInstrumentDefinitionFuture27NoEvents
	NoMDFeedTypes           []MDInstrumentDefinitionFuture27NoMDFeedTypes
	NoInstAttrib            []MDInstrumentDefinitionFuture27NoInstAttrib
	NoLotTypeRules          []MDInstrumentDefinitionFuture27NoLotTypeRules
}
type MDInstrumentDefinitionFuture27NoEvents struct {
	EventType EventTypeEnum
	EventTime uint64
}
type MDInstrumentDefinitionFuture27NoMDFeedTypes struct {
	MDFeedType  [3]byte
	MarketDepth int8
}
type MDInstrumentDefinitionFuture27NoInstAttrib struct {
	InstAttribType  int8
	InstAttribValue InstAttribValue
}
type MDInstrumentDefinitionFuture27NoLotTypeRules struct {
	LotType    int8
	MinLotSize DecimalQty
}

func (m *MDInstrumentDefinitionFuture27) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteBytes(_w, m.SettlCurrency[:]); err != nil {
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
	if err := _m.WriteInt32(_w, m.DecayQuantity); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.DecayStartDate); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.OriginalContractSize); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.ContractMultiplier); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.ContractMultiplierUnit); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.FlowScheduleType); err != nil {
		return err
	}
	if err := m.MinPriceIncrementAmount.Encode(_m, _w); err != nil {
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
	return nil
}

func (m *MDInstrumentDefinitionFuture27) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !m.SettlCurrencyInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			m.SettlCurrency[idx] = m.SettlCurrencyNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.SettlCurrency[:]); err != nil {
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
	if !m.DecayQuantityInActingVersion(actingVersion) {
		m.DecayQuantity = m.DecayQuantityNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.DecayQuantity); err != nil {
			return err
		}
	}
	if !m.DecayStartDateInActingVersion(actingVersion) {
		m.DecayStartDate = m.DecayStartDateNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.DecayStartDate); err != nil {
			return err
		}
	}
	if !m.OriginalContractSizeInActingVersion(actingVersion) {
		m.OriginalContractSize = m.OriginalContractSizeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.OriginalContractSize); err != nil {
			return err
		}
	}
	if !m.ContractMultiplierInActingVersion(actingVersion) {
		m.ContractMultiplier = m.ContractMultiplierNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.ContractMultiplier); err != nil {
			return err
		}
	}
	if !m.ContractMultiplierUnitInActingVersion(actingVersion) {
		m.ContractMultiplierUnit = m.ContractMultiplierUnitNullValue()
	} else {
		if err := _m.ReadInt8(_r, &m.ContractMultiplierUnit); err != nil {
			return err
		}
	}
	if !m.FlowScheduleTypeInActingVersion(actingVersion) {
		m.FlowScheduleType = m.FlowScheduleTypeNullValue()
	} else {
		if err := _m.ReadInt8(_r, &m.FlowScheduleType); err != nil {
			return err
		}
	}
	if m.MinPriceIncrementAmountInActingVersion(actingVersion) {
		if err := m.MinPriceIncrementAmount.Decode(_m, _r, actingVersion); err != nil {
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
			m.NoEvents = make([]MDInstrumentDefinitionFuture27NoEvents, NoEventsNumInGroup)
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
			m.NoMDFeedTypes = make([]MDInstrumentDefinitionFuture27NoMDFeedTypes, NoMDFeedTypesNumInGroup)
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
			m.NoInstAttrib = make([]MDInstrumentDefinitionFuture27NoInstAttrib, NoInstAttribNumInGroup)
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
			m.NoLotTypeRules = make([]MDInstrumentDefinitionFuture27NoLotTypeRules, NoLotTypeRulesNumInGroup)
		}
		m.NoLotTypeRules = m.NoLotTypeRules[:NoLotTypeRulesNumInGroup]
		for i, _ := range m.NoLotTypeRules {
			if err := m.NoLotTypeRules[i].Decode(_m, _r, actingVersion, uint(NoLotTypeRulesBlockLength)); err != nil {
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

func (m *MDInstrumentDefinitionFuture27) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if m.CurrencyInActingVersion(actingVersion) {
		for idx := 0; idx < 3; idx++ {
			if m.Currency[idx] < m.CurrencyMinValue() || m.Currency[idx] > m.CurrencyMaxValue() {
				return fmt.Errorf("Range check failed on m.Currency[%d] (%v < %v > %v)", idx, m.CurrencyMinValue(), m.Currency[idx], m.CurrencyMaxValue())
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
	if m.DecayQuantityInActingVersion(actingVersion) {
		if m.DecayQuantity != m.DecayQuantityNullValue() && (m.DecayQuantity < m.DecayQuantityMinValue() || m.DecayQuantity > m.DecayQuantityMaxValue()) {
			return fmt.Errorf("Range check failed on m.DecayQuantity (%v < %v > %v)", m.DecayQuantityMinValue(), m.DecayQuantity, m.DecayQuantityMaxValue())
		}
	}
	if m.DecayStartDateInActingVersion(actingVersion) {
		if m.DecayStartDate != m.DecayStartDateNullValue() && (m.DecayStartDate < m.DecayStartDateMinValue() || m.DecayStartDate > m.DecayStartDateMaxValue()) {
			return fmt.Errorf("Range check failed on m.DecayStartDate (%v < %v > %v)", m.DecayStartDateMinValue(), m.DecayStartDate, m.DecayStartDateMaxValue())
		}
	}
	if m.OriginalContractSizeInActingVersion(actingVersion) {
		if m.OriginalContractSize != m.OriginalContractSizeNullValue() && (m.OriginalContractSize < m.OriginalContractSizeMinValue() || m.OriginalContractSize > m.OriginalContractSizeMaxValue()) {
			return fmt.Errorf("Range check failed on m.OriginalContractSize (%v < %v > %v)", m.OriginalContractSizeMinValue(), m.OriginalContractSize, m.OriginalContractSizeMaxValue())
		}
	}
	if m.ContractMultiplierInActingVersion(actingVersion) {
		if m.ContractMultiplier != m.ContractMultiplierNullValue() && (m.ContractMultiplier < m.ContractMultiplierMinValue() || m.ContractMultiplier > m.ContractMultiplierMaxValue()) {
			return fmt.Errorf("Range check failed on m.ContractMultiplier (%v < %v > %v)", m.ContractMultiplierMinValue(), m.ContractMultiplier, m.ContractMultiplierMaxValue())
		}
	}
	if m.ContractMultiplierUnitInActingVersion(actingVersion) {
		if m.ContractMultiplierUnit != m.ContractMultiplierUnitNullValue() && (m.ContractMultiplierUnit < m.ContractMultiplierUnitMinValue() || m.ContractMultiplierUnit > m.ContractMultiplierUnitMaxValue()) {
			return fmt.Errorf("Range check failed on m.ContractMultiplierUnit (%v < %v > %v)", m.ContractMultiplierUnitMinValue(), m.ContractMultiplierUnit, m.ContractMultiplierUnitMaxValue())
		}
	}
	if m.FlowScheduleTypeInActingVersion(actingVersion) {
		if m.FlowScheduleType != m.FlowScheduleTypeNullValue() && (m.FlowScheduleType < m.FlowScheduleTypeMinValue() || m.FlowScheduleType > m.FlowScheduleTypeMaxValue()) {
			return fmt.Errorf("Range check failed on m.FlowScheduleType (%v < %v > %v)", m.FlowScheduleTypeMinValue(), m.FlowScheduleType, m.FlowScheduleTypeMaxValue())
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
	return nil
}

func MDInstrumentDefinitionFuture27Init(m *MDInstrumentDefinitionFuture27) {
	m.TotNumReports = 4294967295
	m.SecurityIDSource[0] = 56
	m.MainFraction = 255
	m.SubFraction = 255
	m.PriceDisplayFormat = 255
	m.OpenInterestQty = 2147483647
	m.ClearedVolume = 2147483647
	m.DecayQuantity = 2147483647
	m.DecayStartDate = 65535
	m.OriginalContractSize = 2147483647
	m.ContractMultiplier = 2147483647
	m.ContractMultiplierUnit = 127
	m.FlowScheduleType = 127
	m.TradingReferenceDate = 65535
	return
}

func (m *MDInstrumentDefinitionFuture27NoEvents) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.EventType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.EventTime); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionFuture27NoEvents) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MDInstrumentDefinitionFuture27NoEvents) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDInstrumentDefinitionFuture27NoEventsInit(m *MDInstrumentDefinitionFuture27NoEvents) {
	return
}

func (m *MDInstrumentDefinitionFuture27NoMDFeedTypes) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, m.MDFeedType[:]); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.MarketDepth); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionFuture27NoMDFeedTypes) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MDInstrumentDefinitionFuture27NoMDFeedTypes) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDInstrumentDefinitionFuture27NoMDFeedTypesInit(m *MDInstrumentDefinitionFuture27NoMDFeedTypes) {
	return
}

func (m *MDInstrumentDefinitionFuture27NoInstAttrib) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.InstAttribValue.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionFuture27NoInstAttrib) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MDInstrumentDefinitionFuture27NoInstAttrib) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func MDInstrumentDefinitionFuture27NoInstAttribInit(m *MDInstrumentDefinitionFuture27NoInstAttrib) {
	m.InstAttribType = 24
	return
}

func (m *MDInstrumentDefinitionFuture27NoLotTypeRules) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, m.LotType); err != nil {
		return err
	}
	if err := m.MinLotSize.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionFuture27NoLotTypeRules) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MDInstrumentDefinitionFuture27NoLotTypeRules) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.LotTypeInActingVersion(actingVersion) {
		if m.LotType < m.LotTypeMinValue() || m.LotType > m.LotTypeMaxValue() {
			return fmt.Errorf("Range check failed on m.LotType (%v < %v > %v)", m.LotTypeMinValue(), m.LotType, m.LotTypeMaxValue())
		}
	}
	return nil
}

func MDInstrumentDefinitionFuture27NoLotTypeRulesInit(m *MDInstrumentDefinitionFuture27NoLotTypeRules) {
	return
}

func (*MDInstrumentDefinitionFuture27) SbeBlockLength() (blockLength uint16) {
	return 216
}

func (*MDInstrumentDefinitionFuture27) SbeTemplateId() (templateId uint16) {
	return 27
}

func (*MDInstrumentDefinitionFuture27) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MDInstrumentDefinitionFuture27) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionFuture27) SbeSemanticType() (semanticType []byte) {
	return []byte("d")
}

func (*MDInstrumentDefinitionFuture27) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MDInstrumentDefinitionFuture27) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) TotNumReportsId() uint16 {
	return 911
}

func (*MDInstrumentDefinitionFuture27) TotNumReportsSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) TotNumReportsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TotNumReportsSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) TotNumReportsDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) TotNumReportsMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) TotNumReportsMinValue() uint32 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) TotNumReportsMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDInstrumentDefinitionFuture27) TotNumReportsNullValue() uint32 {
	return 4294967295
}

func (*MDInstrumentDefinitionFuture27) SecurityUpdateActionId() uint16 {
	return 980
}

func (*MDInstrumentDefinitionFuture27) SecurityUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SecurityUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityUpdateActionSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) SecurityUpdateActionDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) SecurityUpdateActionMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) LastUpdateTimeId() uint16 {
	return 779
}

func (*MDInstrumentDefinitionFuture27) LastUpdateTimeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) LastUpdateTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastUpdateTimeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) LastUpdateTimeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) LastUpdateTimeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) LastUpdateTimeMinValue() uint64 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) LastUpdateTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDInstrumentDefinitionFuture27) LastUpdateTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDInstrumentDefinitionFuture27) MDSecurityTradingStatusId() uint16 {
	return 1682
}

func (*MDInstrumentDefinitionFuture27) MDSecurityTradingStatusSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) MDSecurityTradingStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDSecurityTradingStatusSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) MDSecurityTradingStatusDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MDSecurityTradingStatusMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) ApplIDId() uint16 {
	return 1180
}

func (*MDInstrumentDefinitionFuture27) ApplIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) ApplIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ApplIDSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) ApplIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) ApplIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) ApplIDMinValue() int16 {
	return math.MinInt16 + 1
}

func (*MDInstrumentDefinitionFuture27) ApplIDMaxValue() int16 {
	return math.MaxInt16
}

func (*MDInstrumentDefinitionFuture27) ApplIDNullValue() int16 {
	return math.MinInt16
}

func (*MDInstrumentDefinitionFuture27) MarketSegmentIDId() uint16 {
	return 1300
}

func (*MDInstrumentDefinitionFuture27) MarketSegmentIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) MarketSegmentIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MarketSegmentIDSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) MarketSegmentIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MarketSegmentIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) MarketSegmentIDMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MarketSegmentIDMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionFuture27) MarketSegmentIDNullValue() uint8 {
	return math.MaxUint8
}

func (*MDInstrumentDefinitionFuture27) UnderlyingProductId() uint16 {
	return 462
}

func (*MDInstrumentDefinitionFuture27) UnderlyingProductSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) UnderlyingProductInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnderlyingProductSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) UnderlyingProductDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) UnderlyingProductMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) UnderlyingProductMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) UnderlyingProductMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionFuture27) UnderlyingProductNullValue() uint8 {
	return math.MaxUint8
}

func (*MDInstrumentDefinitionFuture27) SecurityExchangeId() uint16 {
	return 207
}

func (*MDInstrumentDefinitionFuture27) SecurityExchangeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SecurityExchangeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityExchangeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) SecurityExchangeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) SecurityExchangeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) SecurityExchangeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture27) SecurityExchangeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture27) SecurityExchangeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SecurityExchangeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture27) SecurityGroupId() uint16 {
	return 1151
}

func (*MDInstrumentDefinitionFuture27) SecurityGroupSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SecurityGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityGroupSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) SecurityGroupDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) SecurityGroupMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) SecurityGroupMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture27) SecurityGroupMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture27) SecurityGroupNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SecurityGroupCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture27) AssetId() uint16 {
	return 6937
}

func (*MDInstrumentDefinitionFuture27) AssetSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) AssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AssetSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) AssetDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) AssetMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) AssetMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture27) AssetMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture27) AssetNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) AssetCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture27) SymbolId() uint16 {
	return 55
}

func (*MDInstrumentDefinitionFuture27) SymbolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SymbolSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) SymbolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) SymbolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) SymbolMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture27) SymbolMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture27) SymbolNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture27) SecurityIDId() uint16 {
	return 48
}

func (*MDInstrumentDefinitionFuture27) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) SecurityIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionFuture27) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionFuture27) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDInstrumentDefinitionFuture27) SecurityIDSourceId() uint16 {
	return 22
}

func (*MDInstrumentDefinitionFuture27) SecurityIDSourceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SecurityIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSourceSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) SecurityIDSourceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) SecurityIDSourceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) SecurityIDSourceMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture27) SecurityIDSourceMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture27) SecurityIDSourceNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionFuture27) SecurityTypeId() uint16 {
	return 167
}

func (*MDInstrumentDefinitionFuture27) SecurityTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SecurityTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) SecurityTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) SecurityTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) SecurityTypeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture27) SecurityTypeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture27) SecurityTypeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SecurityTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture27) CFICodeId() uint16 {
	return 461
}

func (*MDInstrumentDefinitionFuture27) CFICodeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) CFICodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CFICodeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) CFICodeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) CFICodeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) CFICodeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture27) CFICodeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture27) CFICodeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) CFICodeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture27) MaturityMonthYearId() uint16 {
	return 200
}

func (*MDInstrumentDefinitionFuture27) MaturityMonthYearSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) MaturityMonthYearInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaturityMonthYearSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) MaturityMonthYearDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MaturityMonthYearMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) CurrencyId() uint16 {
	return 15
}

func (*MDInstrumentDefinitionFuture27) CurrencySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) CurrencyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CurrencySinceVersion()
}

func (*MDInstrumentDefinitionFuture27) CurrencyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) CurrencyMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) CurrencyMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture27) CurrencyMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture27) CurrencyNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) CurrencyCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture27) SettlCurrencyId() uint16 {
	return 120
}

func (*MDInstrumentDefinitionFuture27) SettlCurrencySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SettlCurrencyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SettlCurrencySinceVersion()
}

func (*MDInstrumentDefinitionFuture27) SettlCurrencyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) SettlCurrencyMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) SettlCurrencyMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture27) SettlCurrencyMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture27) SettlCurrencyNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SettlCurrencyCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture27) MatchAlgorithmId() uint16 {
	return 1142
}

func (*MDInstrumentDefinitionFuture27) MatchAlgorithmSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) MatchAlgorithmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchAlgorithmSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) MatchAlgorithmDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MatchAlgorithmMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) MatchAlgorithmMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture27) MatchAlgorithmMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture27) MatchAlgorithmNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MinTradeVolId() uint16 {
	return 562
}

func (*MDInstrumentDefinitionFuture27) MinTradeVolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) MinTradeVolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinTradeVolSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) MinTradeVolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MinTradeVolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) MinTradeVolMinValue() uint32 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MinTradeVolMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDInstrumentDefinitionFuture27) MinTradeVolNullValue() uint32 {
	return math.MaxUint32
}

func (*MDInstrumentDefinitionFuture27) MaxTradeVolId() uint16 {
	return 1140
}

func (*MDInstrumentDefinitionFuture27) MaxTradeVolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) MaxTradeVolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaxTradeVolSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) MaxTradeVolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MaxTradeVolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) MaxTradeVolMinValue() uint32 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MaxTradeVolMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDInstrumentDefinitionFuture27) MaxTradeVolNullValue() uint32 {
	return math.MaxUint32
}

func (*MDInstrumentDefinitionFuture27) MinPriceIncrementId() uint16 {
	return 969
}

func (*MDInstrumentDefinitionFuture27) MinPriceIncrementSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) MinPriceIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinPriceIncrementSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) MinPriceIncrementDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MinPriceIncrementMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) DisplayFactorId() uint16 {
	return 9787
}

func (*MDInstrumentDefinitionFuture27) DisplayFactorSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) DisplayFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DisplayFactorSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) DisplayFactorDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) DisplayFactorMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) MainFractionId() uint16 {
	return 37702
}

func (*MDInstrumentDefinitionFuture27) MainFractionSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) MainFractionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MainFractionSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) MainFractionDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MainFractionMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) MainFractionMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MainFractionMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionFuture27) MainFractionNullValue() uint8 {
	return 255
}

func (*MDInstrumentDefinitionFuture27) SubFractionId() uint16 {
	return 37703
}

func (*MDInstrumentDefinitionFuture27) SubFractionSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SubFractionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SubFractionSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) SubFractionDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) SubFractionMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) SubFractionMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) SubFractionMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionFuture27) SubFractionNullValue() uint8 {
	return 255
}

func (*MDInstrumentDefinitionFuture27) PriceDisplayFormatId() uint16 {
	return 9800
}

func (*MDInstrumentDefinitionFuture27) PriceDisplayFormatSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) PriceDisplayFormatInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceDisplayFormatSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) PriceDisplayFormatDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) PriceDisplayFormatMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) PriceDisplayFormatMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) PriceDisplayFormatMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionFuture27) PriceDisplayFormatNullValue() uint8 {
	return 255
}

func (*MDInstrumentDefinitionFuture27) UnitOfMeasureId() uint16 {
	return 996
}

func (*MDInstrumentDefinitionFuture27) UnitOfMeasureSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) UnitOfMeasureInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnitOfMeasureSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) UnitOfMeasureDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) UnitOfMeasureMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) UnitOfMeasureMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture27) UnitOfMeasureMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture27) UnitOfMeasureNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) UnitOfMeasureCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture27) UnitOfMeasureQtyId() uint16 {
	return 1147
}

func (*MDInstrumentDefinitionFuture27) UnitOfMeasureQtySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) UnitOfMeasureQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnitOfMeasureQtySinceVersion()
}

func (*MDInstrumentDefinitionFuture27) UnitOfMeasureQtyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) UnitOfMeasureQtyMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) TradingReferencePriceId() uint16 {
	return 1150
}

func (*MDInstrumentDefinitionFuture27) TradingReferencePriceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) TradingReferencePriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradingReferencePriceSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) TradingReferencePriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) TradingReferencePriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) SettlPriceTypeId() uint16 {
	return 731
}

func (*MDInstrumentDefinitionFuture27) SettlPriceTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) SettlPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SettlPriceTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) SettlPriceTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) SettlPriceTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) OpenInterestQtyId() uint16 {
	return 5792
}

func (*MDInstrumentDefinitionFuture27) OpenInterestQtySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) OpenInterestQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OpenInterestQtySinceVersion()
}

func (*MDInstrumentDefinitionFuture27) OpenInterestQtyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) OpenInterestQtyMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) OpenInterestQtyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionFuture27) OpenInterestQtyMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionFuture27) OpenInterestQtyNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionFuture27) ClearedVolumeId() uint16 {
	return 5791
}

func (*MDInstrumentDefinitionFuture27) ClearedVolumeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) ClearedVolumeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ClearedVolumeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) ClearedVolumeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) ClearedVolumeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) ClearedVolumeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionFuture27) ClearedVolumeMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionFuture27) ClearedVolumeNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionFuture27) HighLimitPriceId() uint16 {
	return 1149
}

func (*MDInstrumentDefinitionFuture27) HighLimitPriceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) HighLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.HighLimitPriceSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) HighLimitPriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) HighLimitPriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) LowLimitPriceId() uint16 {
	return 1148
}

func (*MDInstrumentDefinitionFuture27) LowLimitPriceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) LowLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LowLimitPriceSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) LowLimitPriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) LowLimitPriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) MaxPriceVariationId() uint16 {
	return 1143
}

func (*MDInstrumentDefinitionFuture27) MaxPriceVariationSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) MaxPriceVariationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaxPriceVariationSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) MaxPriceVariationDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MaxPriceVariationMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) DecayQuantityId() uint16 {
	return 5818
}

func (*MDInstrumentDefinitionFuture27) DecayQuantitySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) DecayQuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DecayQuantitySinceVersion()
}

func (*MDInstrumentDefinitionFuture27) DecayQuantityDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) DecayQuantityMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) DecayQuantityMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionFuture27) DecayQuantityMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionFuture27) DecayQuantityNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionFuture27) DecayStartDateId() uint16 {
	return 5819
}

func (*MDInstrumentDefinitionFuture27) DecayStartDateSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) DecayStartDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DecayStartDateSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) DecayStartDateDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) DecayStartDateMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) DecayStartDateMinValue() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) DecayStartDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MDInstrumentDefinitionFuture27) DecayStartDateNullValue() uint16 {
	return 65535
}

func (*MDInstrumentDefinitionFuture27) OriginalContractSizeId() uint16 {
	return 5849
}

func (*MDInstrumentDefinitionFuture27) OriginalContractSizeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) OriginalContractSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OriginalContractSizeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) OriginalContractSizeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) OriginalContractSizeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) OriginalContractSizeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionFuture27) OriginalContractSizeMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionFuture27) OriginalContractSizeNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionFuture27) ContractMultiplierId() uint16 {
	return 231
}

func (*MDInstrumentDefinitionFuture27) ContractMultiplierSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) ContractMultiplierInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ContractMultiplierSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) ContractMultiplierDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) ContractMultiplierMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) ContractMultiplierMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionFuture27) ContractMultiplierMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionFuture27) ContractMultiplierNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionFuture27) ContractMultiplierUnitId() uint16 {
	return 1435
}

func (*MDInstrumentDefinitionFuture27) ContractMultiplierUnitSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) ContractMultiplierUnitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ContractMultiplierUnitSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) ContractMultiplierUnitDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) ContractMultiplierUnitMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) ContractMultiplierUnitMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionFuture27) ContractMultiplierUnitMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionFuture27) ContractMultiplierUnitNullValue() int8 {
	return 127
}

func (*MDInstrumentDefinitionFuture27) FlowScheduleTypeId() uint16 {
	return 1439
}

func (*MDInstrumentDefinitionFuture27) FlowScheduleTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) FlowScheduleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.FlowScheduleTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) FlowScheduleTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) FlowScheduleTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) FlowScheduleTypeMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionFuture27) FlowScheduleTypeMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionFuture27) FlowScheduleTypeNullValue() int8 {
	return 127
}

func (*MDInstrumentDefinitionFuture27) MinPriceIncrementAmountId() uint16 {
	return 1146
}

func (*MDInstrumentDefinitionFuture27) MinPriceIncrementAmountSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) MinPriceIncrementAmountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinPriceIncrementAmountSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) MinPriceIncrementAmountDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) MinPriceIncrementAmountMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) UserDefinedInstrumentId() uint16 {
	return 9779
}

func (*MDInstrumentDefinitionFuture27) UserDefinedInstrumentSinceVersion() uint16 {
	return 3
}

func (m *MDInstrumentDefinitionFuture27) UserDefinedInstrumentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UserDefinedInstrumentSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) UserDefinedInstrumentDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) UserDefinedInstrumentMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) UserDefinedInstrumentMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture27) UserDefinedInstrumentMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture27) UserDefinedInstrumentNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionFuture27) TradingReferenceDateId() uint16 {
	return 5796
}

func (*MDInstrumentDefinitionFuture27) TradingReferenceDateSinceVersion() uint16 {
	return 6
}

func (m *MDInstrumentDefinitionFuture27) TradingReferenceDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradingReferenceDateSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) TradingReferenceDateDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) TradingReferenceDateMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) TradingReferenceDateMinValue() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27) TradingReferenceDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MDInstrumentDefinitionFuture27) TradingReferenceDateNullValue() uint16 {
	return 65535
}

func (*MDInstrumentDefinitionFuture27NoEvents) EventTypeId() uint16 {
	return 865
}

func (*MDInstrumentDefinitionFuture27NoEvents) EventTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27NoEvents) EventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EventTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27NoEvents) EventTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27NoEvents) EventTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27NoEvents) EventTimeId() uint16 {
	return 1145
}

func (*MDInstrumentDefinitionFuture27NoEvents) EventTimeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27NoEvents) EventTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EventTimeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27NoEvents) EventTimeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27NoEvents) EventTimeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27NoEvents) EventTimeMinValue() uint64 {
	return 0
}

func (*MDInstrumentDefinitionFuture27NoEvents) EventTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDInstrumentDefinitionFuture27NoEvents) EventTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MDFeedTypeId() uint16 {
	return 1022
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MDFeedTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27NoMDFeedTypes) MDFeedTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDFeedTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MDFeedTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MDFeedTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MDFeedTypeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MDFeedTypeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MDFeedTypeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture27NoMDFeedTypes) MDFeedTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MarketDepthId() uint16 {
	return 264
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MarketDepthSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27NoMDFeedTypes) MarketDepthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MarketDepthSinceVersion()
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MarketDepthDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MarketDepthMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MarketDepthMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MarketDepthMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) MarketDepthNullValue() int8 {
	return math.MinInt8
}

func (*MDInstrumentDefinitionFuture27NoInstAttrib) InstAttribTypeId() uint16 {
	return 871
}

func (*MDInstrumentDefinitionFuture27NoInstAttrib) InstAttribTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27NoInstAttrib) InstAttribTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstAttribTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27NoInstAttrib) InstAttribTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27NoInstAttrib) InstAttribTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27NoInstAttrib) InstAttribTypeMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionFuture27NoInstAttrib) InstAttribTypeMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionFuture27NoInstAttrib) InstAttribTypeNullValue() int8 {
	return math.MinInt8
}

func (*MDInstrumentDefinitionFuture27NoInstAttrib) InstAttribValueId() uint16 {
	return 872
}

func (*MDInstrumentDefinitionFuture27NoInstAttrib) InstAttribValueSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27NoInstAttrib) InstAttribValueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstAttribValueSinceVersion()
}

func (*MDInstrumentDefinitionFuture27NoInstAttrib) InstAttribValueDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27NoInstAttrib) InstAttribValueMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27NoLotTypeRules) LotTypeId() uint16 {
	return 1093
}

func (*MDInstrumentDefinitionFuture27NoLotTypeRules) LotTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27NoLotTypeRules) LotTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LotTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27NoLotTypeRules) LotTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27NoLotTypeRules) LotTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27NoLotTypeRules) LotTypeMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionFuture27NoLotTypeRules) LotTypeMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionFuture27NoLotTypeRules) LotTypeNullValue() int8 {
	return math.MinInt8
}

func (*MDInstrumentDefinitionFuture27NoLotTypeRules) MinLotSizeId() uint16 {
	return 1231
}

func (*MDInstrumentDefinitionFuture27NoLotTypeRules) MinLotSizeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27NoLotTypeRules) MinLotSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinLotSizeSinceVersion()
}

func (*MDInstrumentDefinitionFuture27NoLotTypeRules) MinLotSizeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27NoLotTypeRules) MinLotSizeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture27) NoEventsId() uint16 {
	return 864
}

func (*MDInstrumentDefinitionFuture27) NoEventsSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) NoEventsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoEventsSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) NoEventsDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27NoEvents) SbeBlockLength() (blockLength uint) {
	return 9
}

func (*MDInstrumentDefinitionFuture27NoEvents) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionFuture27) NoMDFeedTypesId() uint16 {
	return 1141
}

func (*MDInstrumentDefinitionFuture27) NoMDFeedTypesSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) NoMDFeedTypesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoMDFeedTypesSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) NoMDFeedTypesDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) SbeBlockLength() (blockLength uint) {
	return 4
}

func (*MDInstrumentDefinitionFuture27NoMDFeedTypes) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionFuture27) NoInstAttribId() uint16 {
	return 870
}

func (*MDInstrumentDefinitionFuture27) NoInstAttribSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) NoInstAttribInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoInstAttribSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) NoInstAttribDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27NoInstAttrib) SbeBlockLength() (blockLength uint) {
	return 4
}

func (*MDInstrumentDefinitionFuture27NoInstAttrib) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionFuture27) NoLotTypeRulesId() uint16 {
	return 1234
}

func (*MDInstrumentDefinitionFuture27) NoLotTypeRulesSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture27) NoLotTypeRulesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoLotTypeRulesSinceVersion()
}

func (*MDInstrumentDefinitionFuture27) NoLotTypeRulesDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture27NoLotTypeRules) SbeBlockLength() (blockLength uint) {
	return 5
}

func (*MDInstrumentDefinitionFuture27NoLotTypeRules) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
