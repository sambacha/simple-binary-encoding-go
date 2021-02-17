// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"io"
)

type MatchEventIndicator [8]bool
type MatchEventIndicatorChoiceValue uint8
type MatchEventIndicatorChoiceValues struct {
	LastTradeMsg   MatchEventIndicatorChoiceValue
	LastVolumeMsg  MatchEventIndicatorChoiceValue
	LastQuoteMsg   MatchEventIndicatorChoiceValue
	LastStatsMsg   MatchEventIndicatorChoiceValue
	LastImpliedMsg MatchEventIndicatorChoiceValue
	RecoveryMsg    MatchEventIndicatorChoiceValue
	Reserved       MatchEventIndicatorChoiceValue
	EndOfEvent     MatchEventIndicatorChoiceValue
}

var MatchEventIndicatorChoice = MatchEventIndicatorChoiceValues{0, 1, 2, 3, 4, 5, 6, 7}

func (m *MatchEventIndicator) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint8 = 0
	for k, v := range m {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint8(_w, wireval)
}

func (m *MatchEventIndicator) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	var wireval uint8

	if err := _m.ReadUint8(_r, &wireval); err != nil {
		return err
	}

	var idx uint
	for idx = 0; idx < 8; idx++ {
		m[idx] = (wireval & (1 << idx)) > 0
	}
	return nil
}

func (MatchEventIndicator) EncodedLength() int64 {
	return 1
}

func (*MatchEventIndicator) LastTradeMsgSinceVersion() uint16 {
	return 0
}

func (m *MatchEventIndicator) LastTradeMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastTradeMsgSinceVersion()
}

func (*MatchEventIndicator) LastTradeMsgDeprecated() uint16 {
	return 0
}

func (*MatchEventIndicator) LastVolumeMsgSinceVersion() uint16 {
	return 0
}

func (m *MatchEventIndicator) LastVolumeMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastVolumeMsgSinceVersion()
}

func (*MatchEventIndicator) LastVolumeMsgDeprecated() uint16 {
	return 0
}

func (*MatchEventIndicator) LastQuoteMsgSinceVersion() uint16 {
	return 0
}

func (m *MatchEventIndicator) LastQuoteMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastQuoteMsgSinceVersion()
}

func (*MatchEventIndicator) LastQuoteMsgDeprecated() uint16 {
	return 0
}

func (*MatchEventIndicator) LastStatsMsgSinceVersion() uint16 {
	return 0
}

func (m *MatchEventIndicator) LastStatsMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastStatsMsgSinceVersion()
}

func (*MatchEventIndicator) LastStatsMsgDeprecated() uint16 {
	return 0
}

func (*MatchEventIndicator) LastImpliedMsgSinceVersion() uint16 {
	return 0
}

func (m *MatchEventIndicator) LastImpliedMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastImpliedMsgSinceVersion()
}

func (*MatchEventIndicator) LastImpliedMsgDeprecated() uint16 {
	return 0
}

func (*MatchEventIndicator) RecoveryMsgSinceVersion() uint16 {
	return 0
}

func (m *MatchEventIndicator) RecoveryMsgInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RecoveryMsgSinceVersion()
}

func (*MatchEventIndicator) RecoveryMsgDeprecated() uint16 {
	return 0
}

func (*MatchEventIndicator) ReservedSinceVersion() uint16 {
	return 0
}

func (m *MatchEventIndicator) ReservedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ReservedSinceVersion()
}

func (*MatchEventIndicator) ReservedDeprecated() uint16 {
	return 0
}

func (*MatchEventIndicator) EndOfEventSinceVersion() uint16 {
	return 0
}

func (m *MatchEventIndicator) EndOfEventInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EndOfEventSinceVersion()
}

func (*MatchEventIndicator) EndOfEventDeprecated() uint16 {
	return 0
}
