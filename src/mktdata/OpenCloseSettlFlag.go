// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type OpenCloseSettlFlagEnum uint8
type OpenCloseSettlFlagValues struct {
	DailyOpenPrice         OpenCloseSettlFlagEnum
	IndicativeOpeningPrice OpenCloseSettlFlagEnum
	NullValue              OpenCloseSettlFlagEnum
}

var OpenCloseSettlFlag = OpenCloseSettlFlagValues{0, 5, 255}

func (o OpenCloseSettlFlagEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(o)); err != nil {
		return err
	}
	return nil
}

func (o *OpenCloseSettlFlagEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(o)); err != nil {
		return err
	}
	return nil
}

func (o OpenCloseSettlFlagEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OpenCloseSettlFlag)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OpenCloseSettlFlag, unknown enumeration value %d", o)
}

func (*OpenCloseSettlFlagEnum) EncodedLength() int64 {
	return 1
}

func (*OpenCloseSettlFlagEnum) DailyOpenPriceSinceVersion() uint16 {
	return 0
}

func (o *OpenCloseSettlFlagEnum) DailyOpenPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DailyOpenPriceSinceVersion()
}

func (*OpenCloseSettlFlagEnum) DailyOpenPriceDeprecated() uint16 {
	return 0
}

func (*OpenCloseSettlFlagEnum) IndicativeOpeningPriceSinceVersion() uint16 {
	return 0
}

func (o *OpenCloseSettlFlagEnum) IndicativeOpeningPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.IndicativeOpeningPriceSinceVersion()
}

func (*OpenCloseSettlFlagEnum) IndicativeOpeningPriceDeprecated() uint16 {
	return 0
}
