// Generated SBE (Simple Binary Encoding) message codec

package group_with_data

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"unicode/utf8"
)

type TestMessage3 struct {
	Tag1    uint32
	Entries []TestMessage3Entries
}
type TestMessage3Entries struct {
	TagGroup1     [9]byte
	NestedEntries []TestMessage3EntriesNestedEntries
	VarDataField  []byte
}
type TestMessage3EntriesNestedEntries struct {
	TagGroup2          int64
	VarDataFieldNested []byte
}

func (t *TestMessage3) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	var EntriesBlockLength uint16 = 9
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

func (t *TestMessage3) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
			t.Entries = make([]TestMessage3Entries, EntriesNumInGroup)
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

func (t *TestMessage3) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func TestMessage3Init(t *TestMessage3) {
	return
}

func (t *TestMessage3Entries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteBytes(_w, t.TagGroup1[:]); err != nil {
		return err
	}
	var NestedEntriesBlockLength uint16 = 8
	if err := _m.WriteUint16(_w, NestedEntriesBlockLength); err != nil {
		return err
	}
	var NestedEntriesNumInGroup uint8 = uint8(len(t.NestedEntries))
	if err := _m.WriteUint8(_w, NestedEntriesNumInGroup); err != nil {
		return err
	}
	for _, prop := range t.NestedEntries {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, uint8(len(t.VarDataField))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.VarDataField); err != nil {
		return err
	}
	return nil
}

func (t *TestMessage3Entries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !t.TagGroup1InActingVersion(actingVersion) {
		for idx := 0; idx < 9; idx++ {
			t.TagGroup1[idx] = t.TagGroup1NullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, t.TagGroup1[:]); err != nil {
			return err
		}
	}
	if actingVersion > t.SbeSchemaVersion() && blockLength > t.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-t.SbeBlockLength()))
	}

	if t.NestedEntriesInActingVersion(actingVersion) {
		var NestedEntriesBlockLength uint16
		if err := _m.ReadUint16(_r, &NestedEntriesBlockLength); err != nil {
			return err
		}
		var NestedEntriesNumInGroup uint8
		if err := _m.ReadUint8(_r, &NestedEntriesNumInGroup); err != nil {
			return err
		}
		if cap(t.NestedEntries) < int(NestedEntriesNumInGroup) {
			t.NestedEntries = make([]TestMessage3EntriesNestedEntries, NestedEntriesNumInGroup)
		}
		t.NestedEntries = t.NestedEntries[:NestedEntriesNumInGroup]
		for i, _ := range t.NestedEntries {
			if err := t.NestedEntries[i].Decode(_m, _r, actingVersion, uint(NestedEntriesBlockLength)); err != nil {
				return err
			}
		}
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

func (t *TestMessage3Entries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if t.TagGroup1InActingVersion(actingVersion) {
		for idx := 0; idx < 9; idx++ {
			if t.TagGroup1[idx] < t.TagGroup1MinValue() || t.TagGroup1[idx] > t.TagGroup1MaxValue() {
				return fmt.Errorf("Range check failed on t.TagGroup1[%d] (%v < %v > %v)", idx, t.TagGroup1MinValue(), t.TagGroup1[idx], t.TagGroup1MaxValue())
			}
		}
	}
	for _, prop := range t.NestedEntries {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	if !utf8.Valid(t.VarDataField[:]) {
		return errors.New("t.VarDataField failed UTF-8 validation")
	}
	return nil
}

func TestMessage3EntriesInit(t *TestMessage3Entries) {
	return
}

func (t *TestMessage3EntriesNestedEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, t.TagGroup2); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(t.VarDataFieldNested))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.VarDataFieldNested); err != nil {
		return err
	}
	return nil
}

func (t *TestMessage3EntriesNestedEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

	if t.VarDataFieldNestedInActingVersion(actingVersion) {
		var VarDataFieldNestedLength uint8
		if err := _m.ReadUint8(_r, &VarDataFieldNestedLength); err != nil {
			return err
		}
		if cap(t.VarDataFieldNested) < int(VarDataFieldNestedLength) {
			t.VarDataFieldNested = make([]byte, VarDataFieldNestedLength)
		}
		t.VarDataFieldNested = t.VarDataFieldNested[:VarDataFieldNestedLength]
		if err := _m.ReadBytes(_r, t.VarDataFieldNested); err != nil {
			return err
		}
	}
	return nil
}

func (t *TestMessage3EntriesNestedEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if t.TagGroup2InActingVersion(actingVersion) {
		if t.TagGroup2 < t.TagGroup2MinValue() || t.TagGroup2 > t.TagGroup2MaxValue() {
			return fmt.Errorf("Range check failed on t.TagGroup2 (%v < %v > %v)", t.TagGroup2MinValue(), t.TagGroup2, t.TagGroup2MaxValue())
		}
	}
	if !utf8.Valid(t.VarDataFieldNested[:]) {
		return errors.New("t.VarDataFieldNested failed UTF-8 validation")
	}
	return nil
}

func TestMessage3EntriesNestedEntriesInit(t *TestMessage3EntriesNestedEntries) {
	return
}

func (*TestMessage3) SbeBlockLength() (blockLength uint16) {
	return 16
}

func (*TestMessage3) SbeTemplateId() (templateId uint16) {
	return 3
}

func (*TestMessage3) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*TestMessage3) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*TestMessage3) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*TestMessage3) Tag1Id() uint16 {
	return 1
}

func (*TestMessage3) Tag1SinceVersion() uint16 {
	return 0
}

func (t *TestMessage3) Tag1InActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.Tag1SinceVersion()
}

func (*TestMessage3) Tag1Deprecated() uint16 {
	return 0
}

func (*TestMessage3) Tag1MetaAttribute(meta int) string {
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

func (*TestMessage3) Tag1MinValue() uint32 {
	return 0
}

func (*TestMessage3) Tag1MaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*TestMessage3) Tag1NullValue() uint32 {
	return math.MaxUint32
}

func (*TestMessage3Entries) TagGroup1Id() uint16 {
	return 3
}

func (*TestMessage3Entries) TagGroup1SinceVersion() uint16 {
	return 0
}

func (t *TestMessage3Entries) TagGroup1InActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TagGroup1SinceVersion()
}

func (*TestMessage3Entries) TagGroup1Deprecated() uint16 {
	return 0
}

func (*TestMessage3Entries) TagGroup1MetaAttribute(meta int) string {
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

func (*TestMessage3Entries) TagGroup1MinValue() byte {
	return byte(32)
}

func (*TestMessage3Entries) TagGroup1MaxValue() byte {
	return byte(126)
}

func (*TestMessage3Entries) TagGroup1NullValue() byte {
	return 0
}

func (t *TestMessage3Entries) TagGroup1CharacterEncoding() string {
	return "US-ASCII"
}

func (*TestMessage3EntriesNestedEntries) TagGroup2Id() uint16 {
	return 5
}

func (*TestMessage3EntriesNestedEntries) TagGroup2SinceVersion() uint16 {
	return 0
}

func (t *TestMessage3EntriesNestedEntries) TagGroup2InActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TagGroup2SinceVersion()
}

func (*TestMessage3EntriesNestedEntries) TagGroup2Deprecated() uint16 {
	return 0
}

func (*TestMessage3EntriesNestedEntries) TagGroup2MetaAttribute(meta int) string {
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

func (*TestMessage3EntriesNestedEntries) TagGroup2MinValue() int64 {
	return math.MinInt64 + 1
}

func (*TestMessage3EntriesNestedEntries) TagGroup2MaxValue() int64 {
	return math.MaxInt64
}

func (*TestMessage3EntriesNestedEntries) TagGroup2NullValue() int64 {
	return math.MinInt64
}

func (*TestMessage3EntriesNestedEntries) VarDataFieldNestedMetaAttribute(meta int) string {
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

func (*TestMessage3EntriesNestedEntries) VarDataFieldNestedSinceVersion() uint16 {
	return 0
}

func (t *TestMessage3EntriesNestedEntries) VarDataFieldNestedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.VarDataFieldNestedSinceVersion()
}

func (*TestMessage3EntriesNestedEntries) VarDataFieldNestedDeprecated() uint16 {
	return 0
}

func (TestMessage3EntriesNestedEntries) VarDataFieldNestedCharacterEncoding() string {
	return "UTF-8"
}

func (TestMessage3EntriesNestedEntries) VarDataFieldNestedHeaderLength() uint64 {
	return 1
}

func (*TestMessage3Entries) VarDataFieldMetaAttribute(meta int) string {
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

func (*TestMessage3Entries) VarDataFieldSinceVersion() uint16 {
	return 0
}

func (t *TestMessage3Entries) VarDataFieldInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.VarDataFieldSinceVersion()
}

func (*TestMessage3Entries) VarDataFieldDeprecated() uint16 {
	return 0
}

func (TestMessage3Entries) VarDataFieldCharacterEncoding() string {
	return "UTF-8"
}

func (TestMessage3Entries) VarDataFieldHeaderLength() uint64 {
	return 1
}

func (*TestMessage3) EntriesId() uint16 {
	return 2
}

func (*TestMessage3) EntriesSinceVersion() uint16 {
	return 0
}

func (t *TestMessage3) EntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.EntriesSinceVersion()
}

func (*TestMessage3) EntriesDeprecated() uint16 {
	return 0
}

func (*TestMessage3Entries) SbeBlockLength() (blockLength uint) {
	return 9
}

func (*TestMessage3Entries) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*TestMessage3Entries) NestedEntriesId() uint16 {
	return 4
}

func (*TestMessage3Entries) NestedEntriesSinceVersion() uint16 {
	return 0
}

func (t *TestMessage3Entries) NestedEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.NestedEntriesSinceVersion()
}

func (*TestMessage3Entries) NestedEntriesDeprecated() uint16 {
	return 0
}

func (*TestMessage3EntriesNestedEntries) SbeBlockLength() (blockLength uint) {
	return 8
}

func (*TestMessage3EntriesNestedEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}
