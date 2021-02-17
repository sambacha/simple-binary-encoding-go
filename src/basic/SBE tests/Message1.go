// Generated SBE (Simple Binary Encoding) message codec

package sbe_tests

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type Message1 struct {
	Header     MessageHeader
	EDTField   [20]byte
	ENUMField  ENUMEnum
	SETField   SET
	Int64Field int64
}

func (m *Message1) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := m.ENUMField.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.SETField.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.Int64Field); err != nil {
		return err
	}
	return nil
}

func (m *Message1) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if m.ENUMFieldInActingVersion(actingVersion) {
		if err := m.ENUMField.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.SETFieldInActingVersion(actingVersion) {
		if err := m.SETField.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
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

func (m *Message1) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func Message1Init(m *Message1) {
	return
}

func (*Message1) SbeBlockLength() (blockLength uint16) {
	return 41
}

func (*Message1) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*Message1) SbeSchemaId() (schemaId uint16) {
	return 3
}

func (*Message1) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Message1) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Message1) HeaderId() uint16 {
	return 41
}

func (*Message1) HeaderSinceVersion() uint16 {
	return 0
}

func (m *Message1) HeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.HeaderSinceVersion()
}

func (*Message1) HeaderDeprecated() uint16 {
	return 0
}

func (*Message1) HeaderMetaAttribute(meta int) string {
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

func (*Message1) EDTFieldId() uint16 {
	return 42
}

func (*Message1) EDTFieldSinceVersion() uint16 {
	return 0
}

func (m *Message1) EDTFieldInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.EDTFieldSinceVersion()
}

func (*Message1) EDTFieldDeprecated() uint16 {
	return 0
}

func (*Message1) EDTFieldMetaAttribute(meta int) string {
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

func (*Message1) EDTFieldMinValue() byte {
	return byte(32)
}

func (*Message1) EDTFieldMaxValue() byte {
	return byte(126)
}

func (*Message1) EDTFieldNullValue() byte {
	return 0
}

func (m *Message1) EDTFieldCharacterEncoding() string {
	return "US-ASCII"
}

func (*Message1) ENUMFieldId() uint16 {
	return 43
}

func (*Message1) ENUMFieldSinceVersion() uint16 {
	return 0
}

func (m *Message1) ENUMFieldInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ENUMFieldSinceVersion()
}

func (*Message1) ENUMFieldDeprecated() uint16 {
	return 0
}

func (*Message1) ENUMFieldMetaAttribute(meta int) string {
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

func (*Message1) SETFieldId() uint16 {
	return 44
}

func (*Message1) SETFieldSinceVersion() uint16 {
	return 0
}

func (m *Message1) SETFieldInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SETFieldSinceVersion()
}

func (*Message1) SETFieldDeprecated() uint16 {
	return 0
}

func (*Message1) SETFieldMetaAttribute(meta int) string {
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

func (*Message1) Int64FieldId() uint16 {
	return 45
}

func (*Message1) Int64FieldSinceVersion() uint16 {
	return 0
}

func (m *Message1) Int64FieldInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.Int64FieldSinceVersion()
}

func (*Message1) Int64FieldDeprecated() uint16 {
	return 0
}

func (*Message1) Int64FieldMetaAttribute(meta int) string {
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

func (*Message1) Int64FieldMinValue() int64 {
	return math.MinInt64 + 1
}

func (*Message1) Int64FieldMaxValue() int64 {
	return math.MaxInt64
}

func (*Message1) Int64FieldNullValue() int64 {
	return math.MinInt64
}
