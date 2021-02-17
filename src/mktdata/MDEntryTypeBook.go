// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type MDEntryTypeBookEnum byte
type MDEntryTypeBookValues struct {
	Bid          MDEntryTypeBookEnum
	Offer        MDEntryTypeBookEnum
	ImpliedBid   MDEntryTypeBookEnum
	ImpliedOffer MDEntryTypeBookEnum
	BookReset    MDEntryTypeBookEnum
	NullValue    MDEntryTypeBookEnum
}

var MDEntryTypeBook = MDEntryTypeBookValues{48, 49, 69, 70, 74, 0}

func (m MDEntryTypeBookEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(m)); err != nil {
		return err
	}
	return nil
}

func (m *MDEntryTypeBookEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(m)); err != nil {
		return err
	}
	return nil
}

func (m MDEntryTypeBookEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MDEntryTypeBook)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MDEntryTypeBook, unknown enumeration value %d", m)
}

func (*MDEntryTypeBookEnum) EncodedLength() int64 {
	return 1
}

func (*MDEntryTypeBookEnum) BidSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeBookEnum) BidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BidSinceVersion()
}

func (*MDEntryTypeBookEnum) BidDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeBookEnum) OfferSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeBookEnum) OfferInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OfferSinceVersion()
}

func (*MDEntryTypeBookEnum) OfferDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeBookEnum) ImpliedBidSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeBookEnum) ImpliedBidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ImpliedBidSinceVersion()
}

func (*MDEntryTypeBookEnum) ImpliedBidDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeBookEnum) ImpliedOfferSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeBookEnum) ImpliedOfferInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ImpliedOfferSinceVersion()
}

func (*MDEntryTypeBookEnum) ImpliedOfferDeprecated() uint16 {
	return 0
}

func (*MDEntryTypeBookEnum) BookResetSinceVersion() uint16 {
	return 0
}

func (m *MDEntryTypeBookEnum) BookResetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BookResetSinceVersion()
}

func (*MDEntryTypeBookEnum) BookResetDeprecated() uint16 {
	return 0
}
