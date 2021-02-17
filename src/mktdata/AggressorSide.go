// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type AggressorSideEnum uint8
type AggressorSideValues struct {
	NoAggressor AggressorSideEnum
	Buy         AggressorSideEnum
	Sell        AggressorSideEnum
	NullValue   AggressorSideEnum
}

var AggressorSide = AggressorSideValues{0, 1, 2, 255}

func (a AggressorSideEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(a)); err != nil {
		return err
	}
	return nil
}

func (a *AggressorSideEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(a)); err != nil {
		return err
	}
	return nil
}

func (a AggressorSideEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(AggressorSide)
	for idx := 0; idx < value.NumField(); idx++ {
		if a == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on AggressorSide, unknown enumeration value %d", a)
}

func (*AggressorSideEnum) EncodedLength() int64 {
	return 1
}

func (*AggressorSideEnum) NoAggressorSinceVersion() uint16 {
	return 0
}

func (a *AggressorSideEnum) NoAggressorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.NoAggressorSinceVersion()
}

func (*AggressorSideEnum) NoAggressorDeprecated() uint16 {
	return 0
}

func (*AggressorSideEnum) BuySinceVersion() uint16 {
	return 0
}

func (a *AggressorSideEnum) BuyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.BuySinceVersion()
}

func (*AggressorSideEnum) BuyDeprecated() uint16 {
	return 0
}

func (*AggressorSideEnum) SellSinceVersion() uint16 {
	return 0
}

func (a *AggressorSideEnum) SellInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.SellSinceVersion()
}

func (*AggressorSideEnum) SellDeprecated() uint16 {
	return 0
}
