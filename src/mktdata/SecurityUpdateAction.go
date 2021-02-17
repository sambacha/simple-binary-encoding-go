// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type SecurityUpdateActionEnum byte
type SecurityUpdateActionValues struct {
	Add       SecurityUpdateActionEnum
	Delete    SecurityUpdateActionEnum
	Modify    SecurityUpdateActionEnum
	NullValue SecurityUpdateActionEnum
}

var SecurityUpdateAction = SecurityUpdateActionValues{65, 68, 77, 0}

func (s SecurityUpdateActionEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(s)); err != nil {
		return err
	}
	return nil
}

func (s *SecurityUpdateActionEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(s)); err != nil {
		return err
	}
	return nil
}

func (s SecurityUpdateActionEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SecurityUpdateAction)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SecurityUpdateAction, unknown enumeration value %d", s)
}

func (*SecurityUpdateActionEnum) EncodedLength() int64 {
	return 1
}

func (*SecurityUpdateActionEnum) AddSinceVersion() uint16 {
	return 0
}

func (s *SecurityUpdateActionEnum) AddInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.AddSinceVersion()
}

func (*SecurityUpdateActionEnum) AddDeprecated() uint16 {
	return 0
}

func (*SecurityUpdateActionEnum) DeleteSinceVersion() uint16 {
	return 0
}

func (s *SecurityUpdateActionEnum) DeleteInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.DeleteSinceVersion()
}

func (*SecurityUpdateActionEnum) DeleteDeprecated() uint16 {
	return 0
}

func (*SecurityUpdateActionEnum) ModifySinceVersion() uint16 {
	return 0
}

func (s *SecurityUpdateActionEnum) ModifyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ModifySinceVersion()
}

func (*SecurityUpdateActionEnum) ModifyDeprecated() uint16 {
	return 0
}
