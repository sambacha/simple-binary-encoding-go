// Generated SBE (Simple Binary Encoding) message codec

package composite

import (
	"fmt"
	"io"
	"reflect"
)

type BooleanEnumEnum uint8
type BooleanEnumValues struct {
	F         BooleanEnumEnum
	T         BooleanEnumEnum
	NullValue BooleanEnumEnum
}

var BooleanEnum = BooleanEnumValues{0, 1, 255}

func (b BooleanEnumEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(b)); err != nil {
		return err
	}
	return nil
}

func (b *BooleanEnumEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(b)); err != nil {
		return err
	}
	return nil
}

func (b BooleanEnumEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(BooleanEnum)
	for idx := 0; idx < value.NumField(); idx++ {
		if b == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on BooleanEnum, unknown enumeration value %d", b)
}

func (*BooleanEnumEnum) EncodedLength() int64 {
	return 1
}

func (*BooleanEnumEnum) FSinceVersion() uint16 {
	return 0
}

func (b *BooleanEnumEnum) FInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.FSinceVersion()
}

func (*BooleanEnumEnum) FDeprecated() uint16 {
	return 0
}

func (*BooleanEnumEnum) TSinceVersion() uint16 {
	return 0
}

func (b *BooleanEnumEnum) TInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.TSinceVersion()
}

func (*BooleanEnumEnum) TDeprecated() uint16 {
	return 0
}
