// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type PutOrCallEnum uint8
type PutOrCallValues struct {
	Put       PutOrCallEnum
	Call      PutOrCallEnum
	NullValue PutOrCallEnum
}

var PutOrCall = PutOrCallValues{0, 1, 255}

func (p PutOrCallEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(p)); err != nil {
		return err
	}
	return nil
}

func (p *PutOrCallEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(p)); err != nil {
		return err
	}
	return nil
}

func (p PutOrCallEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(PutOrCall)
	for idx := 0; idx < value.NumField(); idx++ {
		if p == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on PutOrCall, unknown enumeration value %d", p)
}

func (*PutOrCallEnum) EncodedLength() int64 {
	return 1
}

func (*PutOrCallEnum) PutSinceVersion() uint16 {
	return 0
}

func (p *PutOrCallEnum) PutInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PutSinceVersion()
}

func (*PutOrCallEnum) PutDeprecated() uint16 {
	return 0
}

func (*PutOrCallEnum) CallSinceVersion() uint16 {
	return 0
}

func (p *PutOrCallEnum) CallInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.CallSinceVersion()
}

func (*PutOrCallEnum) CallDeprecated() uint16 {
	return 0
}
