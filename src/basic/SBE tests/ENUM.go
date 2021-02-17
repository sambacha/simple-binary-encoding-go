// Generated SBE (Simple Binary Encoding) message codec

package sbe_tests

import (
	"fmt"
	"io"
	"reflect"
)

type ENUMEnum uint8
type ENUMValues struct {
	Value1    ENUMEnum
	Value10   ENUMEnum
	NullValue ENUMEnum
}

var ENUM = ENUMValues{1, 10, 255}

func (e ENUMEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(e)); err != nil {
		return err
	}
	return nil
}

func (e *ENUMEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(e)); err != nil {
		return err
	}
	return nil
}

func (e ENUMEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ENUM)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ENUM, unknown enumeration value %d", e)
}

func (*ENUMEnum) EncodedLength() int64 {
	return 1
}

func (*ENUMEnum) Value1SinceVersion() uint16 {
	return 0
}

func (e *ENUMEnum) Value1InActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.Value1SinceVersion()
}

func (*ENUMEnum) Value1Deprecated() uint16 {
	return 0
}

func (*ENUMEnum) Value10SinceVersion() uint16 {
	return 0
}

func (e *ENUMEnum) Value10InActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.Value10SinceVersion()
}

func (*ENUMEnum) Value10Deprecated() uint16 {
	return 0
}
