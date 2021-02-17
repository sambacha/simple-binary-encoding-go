// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type LegSideEnum uint8
type LegSideValues struct {
	BuySide   LegSideEnum
	SellSide  LegSideEnum
	NullValue LegSideEnum
}

var LegSide = LegSideValues{1, 2, 255}

func (l LegSideEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(l)); err != nil {
		return err
	}
	return nil
}

func (l *LegSideEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(l)); err != nil {
		return err
	}
	return nil
}

func (l LegSideEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(LegSide)
	for idx := 0; idx < value.NumField(); idx++ {
		if l == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on LegSide, unknown enumeration value %d", l)
}

func (*LegSideEnum) EncodedLength() int64 {
	return 1
}

func (*LegSideEnum) BuySideSinceVersion() uint16 {
	return 0
}

func (l *LegSideEnum) BuySideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.BuySideSinceVersion()
}

func (*LegSideEnum) BuySideDeprecated() uint16 {
	return 0
}

func (*LegSideEnum) SellSideSinceVersion() uint16 {
	return 0
}

func (l *LegSideEnum) SellSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.SellSideSinceVersion()
}

func (*LegSideEnum) SellSideDeprecated() uint16 {
	return 0
}
