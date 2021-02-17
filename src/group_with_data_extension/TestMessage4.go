// Generated SBE (Simple Binary Encoding) message codec

package group_with_data_extension

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"unicode/utf8"
)

type TestMessage4 struct {
	Tag1    uint32
	Entries []TestMessage4Entries
}
type TestMessage4Entries struct {
	VarDataField1 []byte
	VarDataField2 []byte
}

func (t *TestMessage4) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := t.RangeCheck(t.SbeSchemaVersion(), t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, t.Tag1); err != nil {
		return err
	}

	for i := 0; i < 12; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	var EntriesBlockLength uint16 = 0
	if err := _m.WriteUint16(_w, EntriesBlockLength); err != nil {
		return err
	}
	var EntriesNumInGroup uint8 = uint8(len(t.Entries))
	if err := _m.WriteUint8(_w, EntriesNumInGroup); err != nil {
		return err
	}
	for _, prop := range t.Entries {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (t *TestMessage4) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !t.Tag1InActingVersion(actingVersion) {
		t.Tag1 = t.Tag1NullValue()
	} else {
		if err := _m.ReadUint32(_r, &t.Tag1); err != nil {
			return err
		}
	}
	if actingVersion > t.SbeSchemaVersion() && blockLength > t.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-t.SbeBlockLength()))
	}
	io.CopyN(ioutil.Discard, _r, 12)

	if t.EntriesInActingVersion(actingVersion) {
		var EntriesBlockLength uint16
		if err := _m.ReadUint16(_r, &EntriesBlockLength); err != nil {
			return err
		}
		var EntriesNumInGroup uint8
		if err := _m.ReadUint8(_r, &EntriesNumInGroup); err != nil {
			return err
		}
		if cap(t.Entries) < int(EntriesNumInGroup) {
			t.Entries = make([]TestMessage4Entries, EntriesNumInGroup)
		}
		t.Entries = t.Entries[:EntriesNumInGroup]
		for i, _ := range t.Entries {
			if err := t.Entries[i].Decode(_m, _r, actingVersion, uint(EntriesBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := t.RangeCheck(actingVersion, t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (t *TestMessage4) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if t.Tag1InActingVersion(actingVersion) {
		if t.Tag1 < t.Tag1MinValue() || t.Tag1 > t.Tag1MaxValue() {
			return fmt.Errorf("Range check failed on t.Tag1 (%v < %v > %v)", t.Tag1MinValue(), t.Tag1, t.Tag1MaxValue())
		}
	}
	for _, prop := range t.Entries {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func TestMessage4Init(t *TestMessage4) {
	return
}

func (t *TestMessage4Entries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(len(t.VarDataField1))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.VarDataField1); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(t.VarDataField2))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.VarDataField2); err != nil {
		return err
	}
	return nil
}

func (t *TestMessage4Entries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if actingVersion > t.SbeSchemaVersion() && blockLength > t.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-t.SbeBlockLength()))
	}

	if t.VarDataField1InActingVersion(actingVersion) {
		var VarDataField1Length uint8
		if err := _m.ReadUint8(_r, &VarDataField1Length); err != nil {
			return err
		}
		if cap(t.VarDataField1) < int(VarDataField1Length) {
			t.VarDataField1 = make([]byte, VarDataField1Length)
		}
		t.VarDataField1 = t.VarDataField1[:VarDataField1Length]
		if err := _m.ReadBytes(_r, t.VarDataField1); err != nil {
			return err
		}
	}

	if t.VarDataField2InActingVersion(actingVersion) {
		var VarDataField2Length uint8
		if err := _m.ReadUint8(_r, &VarDataField2Length); err != nil {
			return err
		}
		if cap(t.VarDataField2) < int(VarDataField2Length) {
			t.VarDataField2 = make([]byte, VarDataField2Length)
		}
		t.VarDataField2 = t.VarDataField2[:VarDataField2Length]
		if err := _m.ReadBytes(_r, t.VarDataField2); err != nil {
			return err
		}
	}
	return nil
}

func (t *TestMessage4Entries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if !utf8.Valid(t.VarDataField1[:]) {
		return errors.New("t.VarDataField1 failed UTF-8 validation")
	}
	if !utf8.Valid(t.VarDataField2[:]) {
		return errors.New("t.VarDataField2 failed UTF-8 validation")
	}
	return nil
}

func TestMessage4EntriesInit(t *TestMessage4Entries) {
	return
}

func (*TestMessage4) SbeBlockLength() (blockLength uint16) {
	return 16
}

func (*TestMessage4) SbeTemplateId() (templateId uint16) {
	return 4
}

func (*TestMessage4) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*TestMessage4) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*TestMessage4) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*TestMessage4) Tag1Id() uint16 {
	return 1
}

func (*TestMessage4) Tag1SinceVersion() uint16 {
	return 0
}

func (t *TestMessage4) Tag1InActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.Tag1SinceVersion()
}

func (*TestMessage4) Tag1Deprecated() uint16 {
	return 0
}

func (*TestMessage4) Tag1MetaAttribute(meta int) string {
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

func (*TestMessage4) Tag1MinValue() uint32 {
	return 0
}

func (*TestMessage4) Tag1MaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*TestMessage4) Tag1NullValue() uint32 {
	return math.MaxUint32
}

func (*TestMessage4Entries) VarDataField1MetaAttribute(meta int) string {
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

func (*TestMessage4Entries) VarDataField1SinceVersion() uint16 {
	return 0
}

func (t *TestMessage4Entries) VarDataField1InActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.VarDataField1SinceVersion()
}

func (*TestMessage4Entries) VarDataField1Deprecated() uint16 {
	return 0
}

func (TestMessage4Entries) VarDataField1CharacterEncoding() string {
	return "UTF-8"
}

func (TestMessage4Entries) VarDataField1HeaderLength() uint64 {
	return 1
}

func (*TestMessage4Entries) VarDataField2MetaAttribute(meta int) string {
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

func (*TestMessage4Entries) VarDataField2SinceVersion() uint16 {
	return 0
}

func (t *TestMessage4Entries) VarDataField2InActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.VarDataField2SinceVersion()
}

func (*TestMessage4Entries) VarDataField2Deprecated() uint16 {
	return 0
}

func (TestMessage4Entries) VarDataField2CharacterEncoding() string {
	return "UTF-8"
}

func (TestMessage4Entries) VarDataField2HeaderLength() uint64 {
	return 1
}

func (*TestMessage4) EntriesId() uint16 {
	return 2
}

func (*TestMessage4) EntriesSinceVersion() uint16 {
	return 0
}

func (t *TestMessage4) EntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.EntriesSinceVersion()
}

func (*TestMessage4) EntriesDeprecated() uint16 {
	return 0
}

func (*TestMessage4Entries) SbeBlockLength() (blockLength uint) {
	return 0
}

func (*TestMessage4Entries) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}
