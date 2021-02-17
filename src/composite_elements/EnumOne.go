// Generated SBE (Simple Binary Encoding) message codec

package composite_elements

import (
	"fmt"
	"io"
	"reflect"
)

type EnumOneEnum uint8
type EnumOneValues struct {
	Value1    EnumOneEnum
	Value10   EnumOneEnum
	NullValue EnumOneEnum
}

var EnumOne = EnumOneValues{1, 10, 255}

func (e EnumOneEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(e)); err != nil {
		return err
	}
	return nil
}

func (e *EnumOneEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(e)); err != nil {
		return err
	}
	return nil
}

func (e EnumOneEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(EnumOne)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on EnumOne, unknown enumeration value %d", e)
}

func (*EnumOneEnum) EncodedLength() int64 {
	return 1
}

func (*EnumOneEnum) Value1SinceVersion() uint16 {
	return 0
}

func (e *EnumOneEnum) Value1InActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.Value1SinceVersion()
}

func (*EnumOneEnum) Value1Deprecated() uint16 {
	return 0
}

func (*EnumOneEnum) Value10SinceVersion() uint16 {
	return 0
}

func (e *EnumOneEnum) Value10InActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.Value10SinceVersion()
}

func (*EnumOneEnum) Value10Deprecated() uint16 {
	return 0
}
