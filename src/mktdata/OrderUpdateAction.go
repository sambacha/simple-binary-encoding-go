// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type OrderUpdateActionEnum uint8
type OrderUpdateActionValues struct {
	New       OrderUpdateActionEnum
	Update    OrderUpdateActionEnum
	Delete    OrderUpdateActionEnum
	NullValue OrderUpdateActionEnum
}

var OrderUpdateAction = OrderUpdateActionValues{0, 1, 2, 255}

func (o OrderUpdateActionEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrderUpdateActionEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrderUpdateActionEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrderUpdateAction)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrderUpdateAction, unknown enumeration value %d", o)
}

func (*OrderUpdateActionEnum) EncodedLength() int64 {
	return 1
}

func (*OrderUpdateActionEnum) NewSinceVersion() uint16 {
	return 0
}

func (o *OrderUpdateActionEnum) NewInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NewSinceVersion()
}

func (*OrderUpdateActionEnum) NewDeprecated() uint16 {
	return 0
}

func (*OrderUpdateActionEnum) UpdateSinceVersion() uint16 {
	return 0
}

func (o *OrderUpdateActionEnum) UpdateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.UpdateSinceVersion()
}

func (*OrderUpdateActionEnum) UpdateDeprecated() uint16 {
	return 0
}

func (*OrderUpdateActionEnum) DeleteSinceVersion() uint16 {
	return 0
}

func (o *OrderUpdateActionEnum) DeleteInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DeleteSinceVersion()
}

func (*OrderUpdateActionEnum) DeleteDeprecated() uint16 {
	return 0
}
