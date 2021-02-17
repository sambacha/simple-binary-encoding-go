// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MDInstrumentDefinitionFuture54 struct {
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
	MinPriceIncrement       PRICE9
	DisplayFactor           Decimal9
	MainFraction            uint8
	SubFraction             uint8
	PriceDisplayFormat      uint8
	UnitOfMeasure           [30]byte
	UnitOfMeasureQty        Decimal9NULL
	TradingReferencePrice   PRICENULL9
	SettlPriceType          SettlPriceType
	OpenInterestQty         int32
	ClearedVolume           int32
	HighLimitPrice          PRICENULL9
	LowLimitPrice           PRICENULL9
	MaxPriceVariation       PRICENULL9
	DecayQuantity           int32
	DecayStartDate          uint16
	OriginalContractSize    int32
	ContractMultiplier      int32
	ContractMultiplierUnit  int8
	FlowScheduleType        int8
	MinPriceIncrementAmount PRICENULL9
	UserDefinedInstrument   byte
	TradingReferenceDate    uint16
	NoEvents                []MDInstrumentDefinitionFuture54NoEvents
	NoMDFeedTypes           []MDInstrumentDefinitionFuture54NoMDFeedTypes
	NoInstAttrib            []MDInstrumentDefinitionFuture54NoInstAttrib
	NoLotTypeRules          []MDInstrumentDefinitionFuture54NoLotTypeRules
}
type MDInstrumentDefinitionFuture54NoEvents struct {
	EventType EventTypeEnum
	EventTime uint64
}
type MDInstrumentDefinitionFuture54NoMDFeedTypes struct {
	MDFeedType  [3]byte
	MarketDepth int8
}
type MDInstrumentDefinitionFuture54NoInstAttrib struct {
	InstAttribType  int8
	InstAttribValue InstAttribValue
}
type MDInstrumentDefinitionFuture54NoLotTypeRules struct {
	LotType    int8
	MinLotSize DecimalQty
}

func (m *MDInstrumentDefinitionFuture54) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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

func (m *MDInstrumentDefinitionFuture54) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
			m.NoEvents = make([]MDInstrumentDefinitionFuture54NoEvents, NoEventsNumInGroup)
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
			m.NoMDFeedTypes = make([]MDInstrumentDefinitionFuture54NoMDFeedTypes, NoMDFeedTypesNumInGroup)
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
			m.NoInstAttrib = make([]MDInstrumentDefinitionFuture54NoInstAttrib, NoInstAttribNumInGroup)
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
			m.NoLotTypeRules = make([]MDInstrumentDefinitionFuture54NoLotTypeRules, NoLotTypeRulesNumInGroup)
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

func (m *MDInstrumentDefinitionFuture54) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDInstrumentDefinitionFuture54Init(m *MDInstrumentDefinitionFuture54) {
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

func (m *MDInstrumentDefinitionFuture54NoEvents) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.EventType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.EventTime); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionFuture54NoEvents) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MDInstrumentDefinitionFuture54NoEvents) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDInstrumentDefinitionFuture54NoEventsInit(m *MDInstrumentDefinitionFuture54NoEvents) {
	return
}

func (m *MDInstrumentDefinitionFuture54NoMDFeedTypes) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, m.MDFeedType[:]); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.MarketDepth); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionFuture54NoMDFeedTypes) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MDInstrumentDefinitionFuture54NoMDFeedTypes) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MDInstrumentDefinitionFuture54NoMDFeedTypesInit(m *MDInstrumentDefinitionFuture54NoMDFeedTypes) {
	return
}

func (m *MDInstrumentDefinitionFuture54NoInstAttrib) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.InstAttribValue.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionFuture54NoInstAttrib) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MDInstrumentDefinitionFuture54NoInstAttrib) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func MDInstrumentDefinitionFuture54NoInstAttribInit(m *MDInstrumentDefinitionFuture54NoInstAttrib) {
	m.InstAttribType = 24
	return
}

func (m *MDInstrumentDefinitionFuture54NoLotTypeRules) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, m.LotType); err != nil {
		return err
	}
	if err := m.MinLotSize.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MDInstrumentDefinitionFuture54NoLotTypeRules) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MDInstrumentDefinitionFuture54NoLotTypeRules) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.LotTypeInActingVersion(actingVersion) {
		if m.LotType < m.LotTypeMinValue() || m.LotType > m.LotTypeMaxValue() {
			return fmt.Errorf("Range check failed on m.LotType (%v < %v > %v)", m.LotTypeMinValue(), m.LotType, m.LotTypeMaxValue())
		}
	}
	return nil
}

func MDInstrumentDefinitionFuture54NoLotTypeRulesInit(m *MDInstrumentDefinitionFuture54NoLotTypeRules) {
	return
}

func (*MDInstrumentDefinitionFuture54) SbeBlockLength() (blockLength uint16) {
	return 216
}

func (*MDInstrumentDefinitionFuture54) SbeTemplateId() (templateId uint16) {
	return 54
}

func (*MDInstrumentDefinitionFuture54) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MDInstrumentDefinitionFuture54) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionFuture54) SbeSemanticType() (semanticType []byte) {
	return []byte("d")
}

func (*MDInstrumentDefinitionFuture54) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*MDInstrumentDefinitionFuture54) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchEventIndicatorSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MatchEventIndicatorMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) TotNumReportsId() uint16 {
	return 911
}

func (*MDInstrumentDefinitionFuture54) TotNumReportsSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) TotNumReportsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TotNumReportsSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) TotNumReportsDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) TotNumReportsMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) TotNumReportsMinValue() uint32 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) TotNumReportsMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDInstrumentDefinitionFuture54) TotNumReportsNullValue() uint32 {
	return 4294967295
}

func (*MDInstrumentDefinitionFuture54) SecurityUpdateActionId() uint16 {
	return 980
}

func (*MDInstrumentDefinitionFuture54) SecurityUpdateActionSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SecurityUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityUpdateActionSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) SecurityUpdateActionDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) SecurityUpdateActionMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) LastUpdateTimeId() uint16 {
	return 779
}

func (*MDInstrumentDefinitionFuture54) LastUpdateTimeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) LastUpdateTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastUpdateTimeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) LastUpdateTimeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) LastUpdateTimeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) LastUpdateTimeMinValue() uint64 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) LastUpdateTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDInstrumentDefinitionFuture54) LastUpdateTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDInstrumentDefinitionFuture54) MDSecurityTradingStatusId() uint16 {
	return 1682
}

func (*MDInstrumentDefinitionFuture54) MDSecurityTradingStatusSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) MDSecurityTradingStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDSecurityTradingStatusSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) MDSecurityTradingStatusDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MDSecurityTradingStatusMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) ApplIDId() uint16 {
	return 1180
}

func (*MDInstrumentDefinitionFuture54) ApplIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) ApplIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ApplIDSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) ApplIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) ApplIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) ApplIDMinValue() int16 {
	return math.MinInt16 + 1
}

func (*MDInstrumentDefinitionFuture54) ApplIDMaxValue() int16 {
	return math.MaxInt16
}

func (*MDInstrumentDefinitionFuture54) ApplIDNullValue() int16 {
	return math.MinInt16
}

func (*MDInstrumentDefinitionFuture54) MarketSegmentIDId() uint16 {
	return 1300
}

func (*MDInstrumentDefinitionFuture54) MarketSegmentIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) MarketSegmentIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MarketSegmentIDSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) MarketSegmentIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MarketSegmentIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) MarketSegmentIDMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MarketSegmentIDMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionFuture54) MarketSegmentIDNullValue() uint8 {
	return math.MaxUint8
}

func (*MDInstrumentDefinitionFuture54) UnderlyingProductId() uint16 {
	return 462
}

func (*MDInstrumentDefinitionFuture54) UnderlyingProductSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) UnderlyingProductInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnderlyingProductSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) UnderlyingProductDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) UnderlyingProductMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) UnderlyingProductMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) UnderlyingProductMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionFuture54) UnderlyingProductNullValue() uint8 {
	return math.MaxUint8
}

func (*MDInstrumentDefinitionFuture54) SecurityExchangeId() uint16 {
	return 207
}

func (*MDInstrumentDefinitionFuture54) SecurityExchangeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SecurityExchangeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityExchangeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) SecurityExchangeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) SecurityExchangeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) SecurityExchangeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture54) SecurityExchangeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture54) SecurityExchangeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SecurityExchangeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture54) SecurityGroupId() uint16 {
	return 1151
}

func (*MDInstrumentDefinitionFuture54) SecurityGroupSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SecurityGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityGroupSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) SecurityGroupDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) SecurityGroupMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) SecurityGroupMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture54) SecurityGroupMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture54) SecurityGroupNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SecurityGroupCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture54) AssetId() uint16 {
	return 6937
}

func (*MDInstrumentDefinitionFuture54) AssetSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) AssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AssetSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) AssetDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) AssetMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) AssetMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture54) AssetMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture54) AssetNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) AssetCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture54) SymbolId() uint16 {
	return 55
}

func (*MDInstrumentDefinitionFuture54) SymbolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SymbolSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) SymbolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) SymbolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) SymbolMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture54) SymbolMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture54) SymbolNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture54) SecurityIDId() uint16 {
	return 48
}

func (*MDInstrumentDefinitionFuture54) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) SecurityIDMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) SecurityIDMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionFuture54) SecurityIDMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionFuture54) SecurityIDNullValue() int32 {
	return math.MinInt32
}

func (*MDInstrumentDefinitionFuture54) SecurityIDSourceId() uint16 {
	return 22
}

func (*MDInstrumentDefinitionFuture54) SecurityIDSourceSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SecurityIDSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSourceSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) SecurityIDSourceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) SecurityIDSourceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) SecurityIDSourceMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture54) SecurityIDSourceMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture54) SecurityIDSourceNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionFuture54) SecurityTypeId() uint16 {
	return 167
}

func (*MDInstrumentDefinitionFuture54) SecurityTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SecurityTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) SecurityTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) SecurityTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) SecurityTypeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture54) SecurityTypeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture54) SecurityTypeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SecurityTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture54) CFICodeId() uint16 {
	return 461
}

func (*MDInstrumentDefinitionFuture54) CFICodeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) CFICodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CFICodeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) CFICodeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) CFICodeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) CFICodeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture54) CFICodeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture54) CFICodeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) CFICodeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture54) MaturityMonthYearId() uint16 {
	return 200
}

func (*MDInstrumentDefinitionFuture54) MaturityMonthYearSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) MaturityMonthYearInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaturityMonthYearSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) MaturityMonthYearDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MaturityMonthYearMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) CurrencyId() uint16 {
	return 15
}

func (*MDInstrumentDefinitionFuture54) CurrencySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) CurrencyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CurrencySinceVersion()
}

func (*MDInstrumentDefinitionFuture54) CurrencyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) CurrencyMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) CurrencyMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture54) CurrencyMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture54) CurrencyNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) CurrencyCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture54) SettlCurrencyId() uint16 {
	return 120
}

func (*MDInstrumentDefinitionFuture54) SettlCurrencySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SettlCurrencyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SettlCurrencySinceVersion()
}

func (*MDInstrumentDefinitionFuture54) SettlCurrencyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) SettlCurrencyMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) SettlCurrencyMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture54) SettlCurrencyMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture54) SettlCurrencyNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SettlCurrencyCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture54) MatchAlgorithmId() uint16 {
	return 1142
}

func (*MDInstrumentDefinitionFuture54) MatchAlgorithmSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) MatchAlgorithmInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchAlgorithmSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) MatchAlgorithmDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MatchAlgorithmMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) MatchAlgorithmMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture54) MatchAlgorithmMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture54) MatchAlgorithmNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MinTradeVolId() uint16 {
	return 562
}

func (*MDInstrumentDefinitionFuture54) MinTradeVolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) MinTradeVolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinTradeVolSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) MinTradeVolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MinTradeVolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) MinTradeVolMinValue() uint32 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MinTradeVolMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDInstrumentDefinitionFuture54) MinTradeVolNullValue() uint32 {
	return math.MaxUint32
}

func (*MDInstrumentDefinitionFuture54) MaxTradeVolId() uint16 {
	return 1140
}

func (*MDInstrumentDefinitionFuture54) MaxTradeVolSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) MaxTradeVolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaxTradeVolSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) MaxTradeVolDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MaxTradeVolMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) MaxTradeVolMinValue() uint32 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MaxTradeVolMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MDInstrumentDefinitionFuture54) MaxTradeVolNullValue() uint32 {
	return math.MaxUint32
}

func (*MDInstrumentDefinitionFuture54) MinPriceIncrementId() uint16 {
	return 969
}

func (*MDInstrumentDefinitionFuture54) MinPriceIncrementSinceVersion() uint16 {
	return 9
}

func (m *MDInstrumentDefinitionFuture54) MinPriceIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinPriceIncrementSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) MinPriceIncrementDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MinPriceIncrementMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) DisplayFactorId() uint16 {
	return 9787
}

func (*MDInstrumentDefinitionFuture54) DisplayFactorSinceVersion() uint16 {
	return 9
}

func (m *MDInstrumentDefinitionFuture54) DisplayFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DisplayFactorSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) DisplayFactorDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) DisplayFactorMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) MainFractionId() uint16 {
	return 37702
}

func (*MDInstrumentDefinitionFuture54) MainFractionSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) MainFractionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MainFractionSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) MainFractionDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MainFractionMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) MainFractionMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MainFractionMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionFuture54) MainFractionNullValue() uint8 {
	return 255
}

func (*MDInstrumentDefinitionFuture54) SubFractionId() uint16 {
	return 37703
}

func (*MDInstrumentDefinitionFuture54) SubFractionSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SubFractionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SubFractionSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) SubFractionDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) SubFractionMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) SubFractionMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) SubFractionMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionFuture54) SubFractionNullValue() uint8 {
	return 255
}

func (*MDInstrumentDefinitionFuture54) PriceDisplayFormatId() uint16 {
	return 9800
}

func (*MDInstrumentDefinitionFuture54) PriceDisplayFormatSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) PriceDisplayFormatInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceDisplayFormatSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) PriceDisplayFormatDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) PriceDisplayFormatMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) PriceDisplayFormatMinValue() uint8 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) PriceDisplayFormatMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MDInstrumentDefinitionFuture54) PriceDisplayFormatNullValue() uint8 {
	return 255
}

func (*MDInstrumentDefinitionFuture54) UnitOfMeasureId() uint16 {
	return 996
}

func (*MDInstrumentDefinitionFuture54) UnitOfMeasureSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) UnitOfMeasureInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnitOfMeasureSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) UnitOfMeasureDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) UnitOfMeasureMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) UnitOfMeasureMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture54) UnitOfMeasureMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture54) UnitOfMeasureNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) UnitOfMeasureCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture54) UnitOfMeasureQtyId() uint16 {
	return 1147
}

func (*MDInstrumentDefinitionFuture54) UnitOfMeasureQtySinceVersion() uint16 {
	return 9
}

func (m *MDInstrumentDefinitionFuture54) UnitOfMeasureQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnitOfMeasureQtySinceVersion()
}

func (*MDInstrumentDefinitionFuture54) UnitOfMeasureQtyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) UnitOfMeasureQtyMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) TradingReferencePriceId() uint16 {
	return 1150
}

func (*MDInstrumentDefinitionFuture54) TradingReferencePriceSinceVersion() uint16 {
	return 9
}

func (m *MDInstrumentDefinitionFuture54) TradingReferencePriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradingReferencePriceSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) TradingReferencePriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) TradingReferencePriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) SettlPriceTypeId() uint16 {
	return 731
}

func (*MDInstrumentDefinitionFuture54) SettlPriceTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) SettlPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SettlPriceTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) SettlPriceTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) SettlPriceTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) OpenInterestQtyId() uint16 {
	return 5792
}

func (*MDInstrumentDefinitionFuture54) OpenInterestQtySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) OpenInterestQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OpenInterestQtySinceVersion()
}

func (*MDInstrumentDefinitionFuture54) OpenInterestQtyDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) OpenInterestQtyMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) OpenInterestQtyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionFuture54) OpenInterestQtyMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionFuture54) OpenInterestQtyNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionFuture54) ClearedVolumeId() uint16 {
	return 5791
}

func (*MDInstrumentDefinitionFuture54) ClearedVolumeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) ClearedVolumeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ClearedVolumeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) ClearedVolumeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) ClearedVolumeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) ClearedVolumeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionFuture54) ClearedVolumeMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionFuture54) ClearedVolumeNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionFuture54) HighLimitPriceId() uint16 {
	return 1149
}

func (*MDInstrumentDefinitionFuture54) HighLimitPriceSinceVersion() uint16 {
	return 9
}

func (m *MDInstrumentDefinitionFuture54) HighLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.HighLimitPriceSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) HighLimitPriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) HighLimitPriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) LowLimitPriceId() uint16 {
	return 1148
}

func (*MDInstrumentDefinitionFuture54) LowLimitPriceSinceVersion() uint16 {
	return 9
}

func (m *MDInstrumentDefinitionFuture54) LowLimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LowLimitPriceSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) LowLimitPriceDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) LowLimitPriceMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) MaxPriceVariationId() uint16 {
	return 1143
}

func (*MDInstrumentDefinitionFuture54) MaxPriceVariationSinceVersion() uint16 {
	return 9
}

func (m *MDInstrumentDefinitionFuture54) MaxPriceVariationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaxPriceVariationSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) MaxPriceVariationDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MaxPriceVariationMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) DecayQuantityId() uint16 {
	return 5818
}

func (*MDInstrumentDefinitionFuture54) DecayQuantitySinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) DecayQuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DecayQuantitySinceVersion()
}

func (*MDInstrumentDefinitionFuture54) DecayQuantityDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) DecayQuantityMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) DecayQuantityMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionFuture54) DecayQuantityMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionFuture54) DecayQuantityNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionFuture54) DecayStartDateId() uint16 {
	return 5819
}

func (*MDInstrumentDefinitionFuture54) DecayStartDateSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) DecayStartDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DecayStartDateSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) DecayStartDateDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) DecayStartDateMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) DecayStartDateMinValue() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) DecayStartDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MDInstrumentDefinitionFuture54) DecayStartDateNullValue() uint16 {
	return 65535
}

func (*MDInstrumentDefinitionFuture54) OriginalContractSizeId() uint16 {
	return 5849
}

func (*MDInstrumentDefinitionFuture54) OriginalContractSizeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) OriginalContractSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OriginalContractSizeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) OriginalContractSizeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) OriginalContractSizeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) OriginalContractSizeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionFuture54) OriginalContractSizeMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionFuture54) OriginalContractSizeNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionFuture54) ContractMultiplierId() uint16 {
	return 231
}

func (*MDInstrumentDefinitionFuture54) ContractMultiplierSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) ContractMultiplierInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ContractMultiplierSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) ContractMultiplierDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) ContractMultiplierMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) ContractMultiplierMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MDInstrumentDefinitionFuture54) ContractMultiplierMaxValue() int32 {
	return math.MaxInt32
}

func (*MDInstrumentDefinitionFuture54) ContractMultiplierNullValue() int32 {
	return 2147483647
}

func (*MDInstrumentDefinitionFuture54) ContractMultiplierUnitId() uint16 {
	return 1435
}

func (*MDInstrumentDefinitionFuture54) ContractMultiplierUnitSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) ContractMultiplierUnitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ContractMultiplierUnitSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) ContractMultiplierUnitDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) ContractMultiplierUnitMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) ContractMultiplierUnitMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionFuture54) ContractMultiplierUnitMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionFuture54) ContractMultiplierUnitNullValue() int8 {
	return 127
}

func (*MDInstrumentDefinitionFuture54) FlowScheduleTypeId() uint16 {
	return 1439
}

func (*MDInstrumentDefinitionFuture54) FlowScheduleTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) FlowScheduleTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.FlowScheduleTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) FlowScheduleTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) FlowScheduleTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) FlowScheduleTypeMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionFuture54) FlowScheduleTypeMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionFuture54) FlowScheduleTypeNullValue() int8 {
	return 127
}

func (*MDInstrumentDefinitionFuture54) MinPriceIncrementAmountId() uint16 {
	return 1146
}

func (*MDInstrumentDefinitionFuture54) MinPriceIncrementAmountSinceVersion() uint16 {
	return 9
}

func (m *MDInstrumentDefinitionFuture54) MinPriceIncrementAmountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinPriceIncrementAmountSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) MinPriceIncrementAmountDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) MinPriceIncrementAmountMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) UserDefinedInstrumentId() uint16 {
	return 9779
}

func (*MDInstrumentDefinitionFuture54) UserDefinedInstrumentSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) UserDefinedInstrumentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UserDefinedInstrumentSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) UserDefinedInstrumentDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) UserDefinedInstrumentMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) UserDefinedInstrumentMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture54) UserDefinedInstrumentMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture54) UserDefinedInstrumentNullValue() byte {
	return 0
}

func (*MDInstrumentDefinitionFuture54) TradingReferenceDateId() uint16 {
	return 5796
}

func (*MDInstrumentDefinitionFuture54) TradingReferenceDateSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) TradingReferenceDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradingReferenceDateSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) TradingReferenceDateDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) TradingReferenceDateMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) TradingReferenceDateMinValue() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54) TradingReferenceDateMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MDInstrumentDefinitionFuture54) TradingReferenceDateNullValue() uint16 {
	return 65535
}

func (*MDInstrumentDefinitionFuture54NoEvents) EventTypeId() uint16 {
	return 865
}

func (*MDInstrumentDefinitionFuture54NoEvents) EventTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54NoEvents) EventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EventTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54NoEvents) EventTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54NoEvents) EventTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54NoEvents) EventTimeId() uint16 {
	return 1145
}

func (*MDInstrumentDefinitionFuture54NoEvents) EventTimeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54NoEvents) EventTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EventTimeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54NoEvents) EventTimeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54NoEvents) EventTimeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54NoEvents) EventTimeMinValue() uint64 {
	return 0
}

func (*MDInstrumentDefinitionFuture54NoEvents) EventTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MDInstrumentDefinitionFuture54NoEvents) EventTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MDFeedTypeId() uint16 {
	return 1022
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MDFeedTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54NoMDFeedTypes) MDFeedTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDFeedTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MDFeedTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MDFeedTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MDFeedTypeMinValue() byte {
	return byte(32)
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MDFeedTypeMaxValue() byte {
	return byte(126)
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MDFeedTypeNullValue() byte {
	return 0
}

func (m *MDInstrumentDefinitionFuture54NoMDFeedTypes) MDFeedTypeCharacterEncoding() string {
	return "US-ASCII"
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MarketDepthId() uint16 {
	return 264
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MarketDepthSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54NoMDFeedTypes) MarketDepthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MarketDepthSinceVersion()
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MarketDepthDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MarketDepthMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MarketDepthMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MarketDepthMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) MarketDepthNullValue() int8 {
	return math.MinInt8
}

func (*MDInstrumentDefinitionFuture54NoInstAttrib) InstAttribTypeId() uint16 {
	return 871
}

func (*MDInstrumentDefinitionFuture54NoInstAttrib) InstAttribTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54NoInstAttrib) InstAttribTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstAttribTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54NoInstAttrib) InstAttribTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54NoInstAttrib) InstAttribTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54NoInstAttrib) InstAttribTypeMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionFuture54NoInstAttrib) InstAttribTypeMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionFuture54NoInstAttrib) InstAttribTypeNullValue() int8 {
	return math.MinInt8
}

func (*MDInstrumentDefinitionFuture54NoInstAttrib) InstAttribValueId() uint16 {
	return 872
}

func (*MDInstrumentDefinitionFuture54NoInstAttrib) InstAttribValueSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54NoInstAttrib) InstAttribValueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstAttribValueSinceVersion()
}

func (*MDInstrumentDefinitionFuture54NoInstAttrib) InstAttribValueDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54NoInstAttrib) InstAttribValueMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54NoLotTypeRules) LotTypeId() uint16 {
	return 1093
}

func (*MDInstrumentDefinitionFuture54NoLotTypeRules) LotTypeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54NoLotTypeRules) LotTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LotTypeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54NoLotTypeRules) LotTypeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54NoLotTypeRules) LotTypeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54NoLotTypeRules) LotTypeMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MDInstrumentDefinitionFuture54NoLotTypeRules) LotTypeMaxValue() int8 {
	return math.MaxInt8
}

func (*MDInstrumentDefinitionFuture54NoLotTypeRules) LotTypeNullValue() int8 {
	return math.MinInt8
}

func (*MDInstrumentDefinitionFuture54NoLotTypeRules) MinLotSizeId() uint16 {
	return 1231
}

func (*MDInstrumentDefinitionFuture54NoLotTypeRules) MinLotSizeSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54NoLotTypeRules) MinLotSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinLotSizeSinceVersion()
}

func (*MDInstrumentDefinitionFuture54NoLotTypeRules) MinLotSizeDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54NoLotTypeRules) MinLotSizeMetaAttribute(meta int) string {
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

func (*MDInstrumentDefinitionFuture54) NoEventsId() uint16 {
	return 864
}

func (*MDInstrumentDefinitionFuture54) NoEventsSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) NoEventsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoEventsSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) NoEventsDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54NoEvents) SbeBlockLength() (blockLength uint) {
	return 9
}

func (*MDInstrumentDefinitionFuture54NoEvents) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionFuture54) NoMDFeedTypesId() uint16 {
	return 1141
}

func (*MDInstrumentDefinitionFuture54) NoMDFeedTypesSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) NoMDFeedTypesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoMDFeedTypesSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) NoMDFeedTypesDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) SbeBlockLength() (blockLength uint) {
	return 4
}

func (*MDInstrumentDefinitionFuture54NoMDFeedTypes) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionFuture54) NoInstAttribId() uint16 {
	return 870
}

func (*MDInstrumentDefinitionFuture54) NoInstAttribSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) NoInstAttribInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoInstAttribSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) NoInstAttribDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54NoInstAttrib) SbeBlockLength() (blockLength uint) {
	return 4
}

func (*MDInstrumentDefinitionFuture54NoInstAttrib) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*MDInstrumentDefinitionFuture54) NoLotTypeRulesId() uint16 {
	return 1234
}

func (*MDInstrumentDefinitionFuture54) NoLotTypeRulesSinceVersion() uint16 {
	return 0
}

func (m *MDInstrumentDefinitionFuture54) NoLotTypeRulesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NoLotTypeRulesSinceVersion()
}

func (*MDInstrumentDefinitionFuture54) NoLotTypeRulesDeprecated() uint16 {
	return 0
}

func (*MDInstrumentDefinitionFuture54NoLotTypeRules) SbeBlockLength() (blockLength uint) {
	return 5
}

func (*MDInstrumentDefinitionFuture54NoLotTypeRules) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
