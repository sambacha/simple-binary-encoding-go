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

type TestMessage1 struct {
	Tag1    uint32
	Entries []TestMessage1Entries
}
type TestMessage1Entries struct {
	TagGroup1    [9]byte
	TagGroup2    int64
	VarDataField []byte
}

func (t *TestMessage1) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	var EntriesBlockLength uint16 = 17
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

func (t *TestMessage1) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
			t.Entries = make([]TestMessage1Entries, EntriesNumInGroup)
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

func (t *TestMessage1) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func TestMessage1Init(t *TestMessage1) {
	return
}

func (t *TestMessage1Entries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, t.TagGroup1[:]); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, t.TagGroup2); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(t.VarDataField))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.VarDataField); err != nil {
		return err
	}
	return nil
}

func (t *TestMessage1Entries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !t.TagGroup1InActingVersion(actingVersion) {
		for idx := 0; idx < 9; idx++ {
			t.TagGroup1[idx] = t.TagGroup1NullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, t.TagGroup1[:]); err != nil {
			return err
		}
	}
	if !t.TagGroup2InActingVersion(actingVersion) {
		t.TagGroup2 = t.TagGroup2NullValue()
	} else {
		if err := _m.ReadInt64(_r, &t.TagGroup2); err != nil {
			return err
		}
	}
	if actingVersion > t.SbeSchemaVersion() && blockLength > t.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-t.SbeBlockLength()))
	}

	if t.VarDataFieldInActingVersion(actingVersion) {
		var VarDataFieldLength uint8
		if err := _m.ReadUint8(_r, &VarDataFieldLength); err != nil {
			return err
		}
		if cap(t.VarDataField) < int(VarDataFieldLength) {
			t.VarDataField = make([]byte, VarDataFieldLength)
		}
		t.VarDataField = t.VarDataField[:VarDataFieldLength]
		if err := _m.ReadBytes(_r, t.VarDataField); err != nil {
			return err
		}
	}
	return nil
}

func (t *TestMessage1Entries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if t.TagGroup1InActingVersion(actingVersion) {
		for idx := 0; idx < 9; idx++ {
			if t.TagGroup1[idx] < t.TagGroup1MinValue() || t.TagGroup1[idx] > t.TagGroup1MaxValue() {
				return fmt.Errorf("Range check failed on t.TagGroup1[%d] (%v < %v > %v)", idx, t.TagGroup1MinValue(), t.TagGroup1[idx], t.TagGroup1MaxValue())
			}
		}
	}
	if t.TagGroup2InActingVersion(actingVersion) {
		if t.TagGroup2 < t.TagGroup2MinValue() || t.TagGroup2 > t.TagGroup2MaxValue() {
			return fmt.Errorf("Range check failed on t.TagGroup2 (%v < %v > %v)", t.TagGroup2MinValue(), t.TagGroup2, t.TagGroup2MaxValue())
		}
	}
	if !utf8.Valid(t.VarDataField[:]) {
		return errors.New("t.VarDataField failed UTF-8 validation")
	}
	return nil
}

func TestMessage1EntriesInit(t *TestMessage1Entries) {
	return
}

func (*TestMessage1) SbeBlockLength() (blockLength uint16) {
	return 16
}

func (*TestMessage1) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*TestMessage1) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*TestMessage1) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*TestMessage1) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*TestMessage1) Tag1Id() uint16 {
	return 1
}

func (*TestMessage1) Tag1SinceVersion() uint16 {
	return 0
}

func (t *TestMessage1) Tag1InActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.Tag1SinceVersion()
}

func (*TestMessage1) Tag1Deprecated() uint16 {
	return 0
}

func (*TestMessage1) Tag1MetaAttribute(meta int) string {
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

func (*TestMessage1) Tag1MinValue() uint32 {
	return 0
}

func (*TestMessage1) Tag1MaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*TestMessage1) Tag1NullValue() uint32 {
	return math.MaxUint32
}

func (*TestMessage1Entries) TagGroup1Id() uint16 {
	return 3
}

func (*TestMessage1Entries) TagGroup1SinceVersion() uint16 {
	return 0
}

func (t *TestMessage1Entries) TagGroup1InActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TagGroup1SinceVersion()
}

func (*TestMessage1Entries) TagGroup1Deprecated() uint16 {
	return 0
}

func (*TestMessage1Entries) TagGroup1MetaAttribute(meta int) string {
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

func (*TestMessage1Entries) TagGroup1MinValue() byte {
	return byte(32)
}

func (*TestMessage1Entries) TagGroup1MaxValue() byte {
	return byte(126)
}

func (*TestMessage1Entries) TagGroup1NullValue() byte {
	return 0
}

func (t *TestMessage1Entries) TagGroup1CharacterEncoding() string {
	return "US-ASCII"
}

func (*TestMessage1Entries) TagGroup2Id() uint16 {
	return 4
}

func (*TestMessage1Entries) TagGroup2SinceVersion() uint16 {
	return 0
}

func (t *TestMessage1Entries) TagGroup2InActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TagGroup2SinceVersion()
}

func (*TestMessage1Entries) TagGroup2Deprecated() uint16 {
	return 0
}

func (*TestMessage1Entries) TagGroup2MetaAttribute(meta int) string {
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

func (*TestMessage1Entries) TagGroup2MinValue() int64 {
	return math.MinInt64 + 1
}

func (*TestMessage1Entries) TagGroup2MaxValue() int64 {
	return math.MaxInt64
}

func (*TestMessage1Entries) TagGroup2NullValue() int64 {
	return math.MinInt64
}

func (*TestMessage1Entries) VarDataFieldMetaAttribute(meta int) string {
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

func (*TestMessage1Entries) VarDataFieldSinceVersion() uint16 {
	return 0
}

func (t *TestMessage1Entries) VarDataFieldInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.VarDataFieldSinceVersion()
}

func (*TestMessage1Entries) VarDataFieldDeprecated() uint16 {
	return 0
}

func (TestMessage1Entries) VarDataFieldCharacterEncoding() string {
	return "UTF-8"
}

func (TestMessage1Entries) VarDataFieldHeaderLength() uint64 {
	return 1
}

func (*TestMessage1) EntriesId() uint16 {
	return 2
}

func (*TestMessage1) EntriesSinceVersion() uint16 {
	return 0
}

func (t *TestMessage1) EntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.EntriesSinceVersion()
}

func (*TestMessage1) EntriesDeprecated() uint16 {
	return 0
}

func (*TestMessage1Entries) SbeBlockLength() (blockLength uint) {
	return 17
}

func (*TestMessage1Entries) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}
