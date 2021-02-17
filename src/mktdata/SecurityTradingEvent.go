// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type SecurityTradingEventEnum uint8
type SecurityTradingEventValues struct {
	NoEvent            SecurityTradingEventEnum
	NoCancel           SecurityTradingEventEnum
	ResetStatistics    SecurityTradingEventEnum
	ImpliedMatchingON  SecurityTradingEventEnum
	ImpliedMatchingOFF SecurityTradingEventEnum
	NullValue          SecurityTradingEventEnum
}

var SecurityTradingEvent = SecurityTradingEventValues{0, 1, 4, 5, 6, 255}

func (s SecurityTradingEventEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(s)); err != nil {
		return err
	}
	return nil
}

func (s *SecurityTradingEventEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(s)); err != nil {
		return err
	}
	return nil
}

func (s SecurityTradingEventEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SecurityTradingEvent)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SecurityTradingEvent, unknown enumeration value %d", s)
}

func (*SecurityTradingEventEnum) EncodedLength() int64 {
	return 1
}

func (*SecurityTradingEventEnum) NoEventSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingEventEnum) NoEventInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NoEventSinceVersion()
}

func (*SecurityTradingEventEnum) NoEventDeprecated() uint16 {
	return 0
}

func (*SecurityTradingEventEnum) NoCancelSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingEventEnum) NoCancelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NoCancelSinceVersion()
}

func (*SecurityTradingEventEnum) NoCancelDeprecated() uint16 {
	return 0
}

func (*SecurityTradingEventEnum) ResetStatisticsSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingEventEnum) ResetStatisticsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ResetStatisticsSinceVersion()
}

func (*SecurityTradingEventEnum) ResetStatisticsDeprecated() uint16 {
	return 0
}

func (*SecurityTradingEventEnum) ImpliedMatchingONSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingEventEnum) ImpliedMatchingONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ImpliedMatchingONSinceVersion()
}

func (*SecurityTradingEventEnum) ImpliedMatchingONDeprecated() uint16 {
	return 0
}

func (*SecurityTradingEventEnum) ImpliedMatchingOFFSinceVersion() uint16 {
	return 0
}

func (s *SecurityTradingEventEnum) ImpliedMatchingOFFInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ImpliedMatchingOFFSinceVersion()
}

func (*SecurityTradingEventEnum) ImpliedMatchingOFFDeprecated() uint16 {
	return 0
}
