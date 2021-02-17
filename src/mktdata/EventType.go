// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type EventTypeEnum uint8
type EventTypeValues struct {
	Activation            EventTypeEnum
	LastEligibleTradeDate EventTypeEnum
	NullValue             EventTypeEnum
}

var EventType = EventTypeValues{5, 7, 255}

func (e EventTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(e)); err != nil {
		return err
	}
	return nil
}

func (e *EventTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(e)); err != nil {
		return err
	}
	return nil
}

func (e EventTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(EventType)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on EventType, unknown enumeration value %d", e)
}

func (*EventTypeEnum) EncodedLength() int64 {
	return 1
}

func (*EventTypeEnum) ActivationSinceVersion() uint16 {
	return 0
}

func (e *EventTypeEnum) ActivationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ActivationSinceVersion()
}

func (*EventTypeEnum) ActivationDeprecated() uint16 {
	return 0
}

func (*EventTypeEnum) LastEligibleTradeDateSinceVersion() uint16 {
	return 0
}

func (e *EventTypeEnum) LastEligibleTradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastEligibleTradeDateSinceVersion()
}

func (*EventTypeEnum) LastEligibleTradeDateDeprecated() uint16 {
	return 0
}
