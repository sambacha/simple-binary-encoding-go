// Generated SBE (Simple Binary Encoding) message codec

package issue435

import (
	"fmt"
	"io"
	"reflect"
)

type EnumRefEnum uint8
type EnumRefValues struct {
	One       EnumRefEnum
	Two       EnumRefEnum
	NullValue EnumRefEnum
}

var EnumRef = EnumRefValues{0, 1, 255}

func (e EnumRefEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(e)); err != nil {
		return err
	}
	return nil
}

func (e *EnumRefEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(e)); err != nil {
		return err
	}
	return nil
}

func (e EnumRefEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(EnumRef)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on EnumRef, unknown enumeration value %d", e)
}

func (*EnumRefEnum) EncodedLength() int64 {
	return 1
}

func (*EnumRefEnum) OneSinceVersion() uint16 {
	return 0
}

func (e *EnumRefEnum) OneInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OneSinceVersion()
}

func (*EnumRefEnum) OneDeprecated() uint16 {
	return 0
}

func (*EnumRefEnum) TwoSinceVersion() uint16 {
	return 0
}

func (e *EnumRefEnum) TwoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TwoSinceVersion()
}

func (*EnumRefEnum) TwoDeprecated() uint16 {
	return 0
}
