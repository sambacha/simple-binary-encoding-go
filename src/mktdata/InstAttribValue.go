// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"io"
)

type InstAttribValue [32]bool
type InstAttribValueChoiceValue uint8
type InstAttribValueChoiceValues struct {
	ElectronicMatchEligible            InstAttribValueChoiceValue
	OrderCrossEligible                 InstAttribValueChoiceValue
	BlockTradeEligible                 InstAttribValueChoiceValue
	EFPEligible                        InstAttribValueChoiceValue
	EBFEligible                        InstAttribValueChoiceValue
	EFSEligible                        InstAttribValueChoiceValue
	EFREligible                        InstAttribValueChoiceValue
	OTCEligible                        InstAttribValueChoiceValue
	ILinkIndicativeMassQuotingEligible InstAttribValueChoiceValue
	NegativeStrikeEligible             InstAttribValueChoiceValue
	NegativePriceOutrightEligible      InstAttribValueChoiceValue
	IsFractional                       InstAttribValueChoiceValue
	VolatilityQuotedOption             InstAttribValueChoiceValue
	RFQCrossEligible                   InstAttribValueChoiceValue
	ZeroPriceOutrightEligible          InstAttribValueChoiceValue
	DecayingProductEligibility         InstAttribValueChoiceValue
	VariableProductEligibility         InstAttribValueChoiceValue
	DailyProductEligibility            InstAttribValueChoiceValue
	GTOrdersEligibility                InstAttribValueChoiceValue
	ImpliedMatchingEligibility         InstAttribValueChoiceValue
	TriangulationEligible              InstAttribValueChoiceValue
	VariableCabEligible                InstAttribValueChoiceValue
}

var InstAttribValueChoice = InstAttribValueChoiceValues{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}

func (i *InstAttribValue) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint32 = 0
	for k, v := range i {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint32(_w, wireval)
}

func (i *InstAttribValue) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	var wireval uint32

	if err := _m.ReadUint32(_r, &wireval); err != nil {
		return err
	}

	var idx uint
	for idx = 0; idx < 32; idx++ {
		i[idx] = (wireval & (1 << idx)) > 0
	}
	return nil
}

func (InstAttribValue) EncodedLength() int64 {
	return 4
}

func (*InstAttribValue) ElectronicMatchEligibleSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) ElectronicMatchEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.ElectronicMatchEligibleSinceVersion()
}

func (*InstAttribValue) ElectronicMatchEligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) OrderCrossEligibleSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) OrderCrossEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.OrderCrossEligibleSinceVersion()
}

func (*InstAttribValue) OrderCrossEligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) BlockTradeEligibleSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) BlockTradeEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.BlockTradeEligibleSinceVersion()
}

func (*InstAttribValue) BlockTradeEligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) EFPEligibleSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) EFPEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.EFPEligibleSinceVersion()
}

func (*InstAttribValue) EFPEligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) EBFEligibleSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) EBFEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.EBFEligibleSinceVersion()
}

func (*InstAttribValue) EBFEligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) EFSEligibleSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) EFSEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.EFSEligibleSinceVersion()
}

func (*InstAttribValue) EFSEligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) EFREligibleSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) EFREligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.EFREligibleSinceVersion()
}

func (*InstAttribValue) EFREligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) OTCEligibleSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) OTCEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.OTCEligibleSinceVersion()
}

func (*InstAttribValue) OTCEligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) iLinkIndicativeMassQuotingEligibleSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) iLinkIndicativeMassQuotingEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.iLinkIndicativeMassQuotingEligibleSinceVersion()
}

func (*InstAttribValue) iLinkIndicativeMassQuotingEligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) NegativeStrikeEligibleSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) NegativeStrikeEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.NegativeStrikeEligibleSinceVersion()
}

func (*InstAttribValue) NegativeStrikeEligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) NegativePriceOutrightEligibleSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) NegativePriceOutrightEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.NegativePriceOutrightEligibleSinceVersion()
}

func (*InstAttribValue) NegativePriceOutrightEligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) IsFractionalSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) IsFractionalInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.IsFractionalSinceVersion()
}

func (*InstAttribValue) IsFractionalDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) VolatilityQuotedOptionSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) VolatilityQuotedOptionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.VolatilityQuotedOptionSinceVersion()
}

func (*InstAttribValue) VolatilityQuotedOptionDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) RFQCrossEligibleSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) RFQCrossEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.RFQCrossEligibleSinceVersion()
}

func (*InstAttribValue) RFQCrossEligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) ZeroPriceOutrightEligibleSinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) ZeroPriceOutrightEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.ZeroPriceOutrightEligibleSinceVersion()
}

func (*InstAttribValue) ZeroPriceOutrightEligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) DecayingProductEligibilitySinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) DecayingProductEligibilityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.DecayingProductEligibilitySinceVersion()
}

func (*InstAttribValue) DecayingProductEligibilityDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) VariableProductEligibilitySinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) VariableProductEligibilityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.VariableProductEligibilitySinceVersion()
}

func (*InstAttribValue) VariableProductEligibilityDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) DailyProductEligibilitySinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) DailyProductEligibilityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.DailyProductEligibilitySinceVersion()
}

func (*InstAttribValue) DailyProductEligibilityDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) GTOrdersEligibilitySinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) GTOrdersEligibilityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.GTOrdersEligibilitySinceVersion()
}

func (*InstAttribValue) GTOrdersEligibilityDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) ImpliedMatchingEligibilitySinceVersion() uint16 {
	return 0
}

func (i *InstAttribValue) ImpliedMatchingEligibilityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.ImpliedMatchingEligibilitySinceVersion()
}

func (*InstAttribValue) ImpliedMatchingEligibilityDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) TriangulationEligibleSinceVersion() uint16 {
	return 9
}

func (i *InstAttribValue) TriangulationEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.TriangulationEligibleSinceVersion()
}

func (*InstAttribValue) TriangulationEligibleDeprecated() uint16 {
	return 0
}

func (*InstAttribValue) VariableCabEligibleSinceVersion() uint16 {
	return 9
}

func (i *InstAttribValue) VariableCabEligibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.VariableCabEligibleSinceVersion()
}

func (*InstAttribValue) VariableCabEligibleDeprecated() uint16 {
	return 0
}
