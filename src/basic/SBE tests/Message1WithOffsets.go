// Generated SBE (Simple Binary Encoding) message codec

package sbe_tests

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type Message1WithOffsets struct {
	Header     MessageHeader
	EDTField   [20]byte
	ENUMField  ENUMEnum
	SETField   SET
	Int64Field int64
}

func (m *Message1WithOffsets) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := m.Header.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.EDTField[:]); err != nil {
		return err
	}

	for i := 0; i < 4; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	if err := m.ENUMField.Encode(_m, _w); err != nil {
		return err
	}

	for i := 0; i < 95; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	if err := m.SETField.Encode(_m, _w); err != nil {
		return err
	}

	for i := 0; i < 4; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, m.Int64Field); err != nil {
		return err
	}
	return nil
}

func (m *Message1WithOffsets) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if m.HeaderInActingVersion(actingVersion) {
		if err := m.Header.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.EDTFieldInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			m.EDTField[idx] = m.EDTFieldNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.EDTField[:]); err != nil {
			return err
		}
	}
	io.CopyN(ioutil.Discard, _r, 4)
	if m.ENUMFieldInActingVersion(actingVersion) {
		if err := m.ENUMField.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	io.CopyN(ioutil.Discard, _r, 95)
	if m.SETFieldInActingVersion(actingVersion) {
		if err := m.SETField.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	io.CopyN(ioutil.Discard, _r, 4)
	if !m.Int64FieldInActingVersion(actingVersion) {
		m.Int64Field = m.Int64FieldNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.Int64Field); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *Message1WithOffsets) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.EDTFieldInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if m.EDTField[idx] < m.EDTFieldMinValue() || m.EDTField[idx] > m.EDTFieldMaxValue() {
				return fmt.Errorf("Range check failed on m.EDTField[%d] (%v < %v > %v)", idx, m.EDTFieldMinValue(), m.EDTField[idx], m.EDTFieldMaxValue())
			}
		}
	}
	if err := m.ENUMField.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.Int64FieldInActingVersion(actingVersion) {
		if m.Int64Field < m.Int64FieldMinValue() || m.Int64Field > m.Int64FieldMaxValue() {
			return fmt.Errorf("Range check failed on m.Int64Field (%v < %v > %v)", m.Int64FieldMinValue(), m.Int64Field, m.Int64FieldMaxValue())
		}
	}
	return nil
}

func Message1WithOffsetsInit(m *Message1WithOffsets) {
	return
}

func (*Message1WithOffsets) SbeBlockLength() (blockLength uint16) {
	return 144
}

func (*Message1WithOffsets) SbeTemplateId() (templateId uint16) {
	return 2
}

func (*Message1WithOffsets) SbeSchemaId() (schemaId uint16) {
	return 3
}

func (*Message1WithOffsets) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Message1WithOffsets) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Message1WithOffsets) HeaderId() uint16 {
	return 41
}

func (*Message1WithOffsets) HeaderSinceVersion() uint16 {
	return 0
}

func (m *Message1WithOffsets) HeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.HeaderSinceVersion()
}

func (*Message1WithOffsets) HeaderDeprecated() uint16 {
	return 0
}

func (*Message1WithOffsets) HeaderMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Message1WithOffsets) EDTFieldId() uint16 {
	return 42
}

func (*Message1WithOffsets) EDTFieldSinceVersion() uint16 {
	return 0
}

func (m *Message1WithOffsets) EDTFieldInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EDTFieldSinceVersion()
}

func (*Message1WithOffsets) EDTFieldDeprecated() uint16 {
	return 0
}

func (*Message1WithOffsets) EDTFieldMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Message1WithOffsets) EDTFieldMinValue() byte {
	return byte(32)
}

func (*Message1WithOffsets) EDTFieldMaxValue() byte {
	return byte(126)
}

func (*Message1WithOffsets) EDTFieldNullValue() byte {
	return 0
}

func (m *Message1WithOffsets) EDTFieldCharacterEncoding() string {
	return "US-ASCII"
}

func (*Message1WithOffsets) ENUMFieldId() uint16 {
	return 43
}

func (*Message1WithOffsets) ENUMFieldSinceVersion() uint16 {
	return 0
}

func (m *Message1WithOffsets) ENUMFieldInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ENUMFieldSinceVersion()
}

func (*Message1WithOffsets) ENUMFieldDeprecated() uint16 {
	return 0
}

func (*Message1WithOffsets) ENUMFieldMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Message1WithOffsets) SETFieldId() uint16 {
	return 44
}

func (*Message1WithOffsets) SETFieldSinceVersion() uint16 {
	return 0
}

func (m *Message1WithOffsets) SETFieldInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SETFieldSinceVersion()
}

func (*Message1WithOffsets) SETFieldDeprecated() uint16 {
	return 0
}

func (*Message1WithOffsets) SETFieldMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Message1WithOffsets) Int64FieldId() uint16 {
	return 45
}

func (*Message1WithOffsets) Int64FieldSinceVersion() uint16 {
	return 0
}

func (m *Message1WithOffsets) Int64FieldInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.Int64FieldSinceVersion()
}

func (*Message1WithOffsets) Int64FieldDeprecated() uint16 {
	return 0
}

func (*Message1WithOffsets) Int64FieldMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*Message1WithOffsets) Int64FieldMinValue() int64 {
	return math.MinInt64 + 1
}

func (*Message1WithOffsets) Int64FieldMaxValue() int64 {
	return math.MaxInt64
}

func (*Message1WithOffsets) Int64FieldNullValue() int64 {
	return math.MinInt64
}
