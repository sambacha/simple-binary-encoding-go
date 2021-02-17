// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type MDUpdateActionEnum uint8
type MDUpdateActionValues struct {
	New        MDUpdateActionEnum
	Change     MDUpdateActionEnum
	Delete     MDUpdateActionEnum
	DeleteThru MDUpdateActionEnum
	DeleteFrom MDUpdateActionEnum
	Overlay    MDUpdateActionEnum
	NullValue  MDUpdateActionEnum
}

var MDUpdateAction = MDUpdateActionValues{0, 1, 2, 3, 4, 5, 255}

func (m MDUpdateActionEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MDUpdateActionEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MDUpdateActionEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MDUpdateAction)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MDUpdateAction, unknown enumeration value %d", m)
}

func (*MDUpdateActionEnum) EncodedLength() int64 {
	return 1
}

func (*MDUpdateActionEnum) NewSinceVersion() uint16 {
	return 0
}

func (m *MDUpdateActionEnum) NewInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NewSinceVersion()
}

func (*MDUpdateActionEnum) NewDeprecated() uint16 {
	return 0
}

func (*MDUpdateActionEnum) ChangeSinceVersion() uint16 {
	return 0
}

func (m *MDUpdateActionEnum) ChangeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ChangeSinceVersion()
}

func (*MDUpdateActionEnum) ChangeDeprecated() uint16 {
	return 0
}

func (*MDUpdateActionEnum) DeleteSinceVersion() uint16 {
	return 0
}

func (m *MDUpdateActionEnum) DeleteInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DeleteSinceVersion()
}

func (*MDUpdateActionEnum) DeleteDeprecated() uint16 {
	return 0
}

func (*MDUpdateActionEnum) DeleteThruSinceVersion() uint16 {
	return 0
}

func (m *MDUpdateActionEnum) DeleteThruInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DeleteThruSinceVersion()
}

func (*MDUpdateActionEnum) DeleteThruDeprecated() uint16 {
	return 0
}

func (*MDUpdateActionEnum) DeleteFromSinceVersion() uint16 {
	return 0
}

func (m *MDUpdateActionEnum) DeleteFromInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DeleteFromSinceVersion()
}

func (*MDUpdateActionEnum) DeleteFromDeprecated() uint16 {
	return 0
}

func (*MDUpdateActionEnum) OverlaySinceVersion() uint16 {
	return 0
}

func (m *MDUpdateActionEnum) OverlayInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OverlaySinceVersion()
}

func (*MDUpdateActionEnum) OverlayDeprecated() uint16 {
	return 0
}
