// Generated SBE (Simple Binary Encoding) message codec

package composite_elements

import (
	"io"
	"io/ioutil"
)

type Msg2 struct {
	Structure OuterWithOffsets
}

func (m *Msg2) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := m.Structure.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *Msg2) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if m.StructureInActingVersion(actingVersion) {
		if err := m.Structure.Decode(_m, _r, actingVersion); err != nil {
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

func (m *Msg2) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func Msg2Init(m *Msg2) {
	return
}

func (*Msg2) SbeBlockLength() (blockLength uint16) {
	return 32
}

func (*Msg2) SbeTemplateId() (templateId uint16) {
	return 2
}

func (*Msg2) SbeSchemaId() (schemaId uint16) {
	return 3
}

func (*Msg2) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Msg2) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Msg2) StructureId() uint16 {
	return 43
}

func (*Msg2) StructureSinceVersion() uint16 {
	return 0
}

func (m *Msg2) StructureInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.StructureSinceVersion()
}

func (*Msg2) StructureDeprecated() uint16 {
	return 0
}

func (*Msg2) StructureMetaAttribute(meta int) string {
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
