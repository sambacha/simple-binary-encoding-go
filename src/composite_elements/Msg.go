// Generated SBE (Simple Binary Encoding) message codec

package composite_elements

import (
	"io"
	"io/ioutil"
)

type Msg struct {
	Structure Outer
}

func (m *Msg) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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

func (m *Msg) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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

func (m *Msg) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func MsgInit(m *Msg) {
	return
}

func (*Msg) SbeBlockLength() (blockLength uint16) {
	return 22
}

func (*Msg) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*Msg) SbeSchemaId() (schemaId uint16) {
	return 3
}

func (*Msg) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Msg) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Msg) StructureId() uint16 {
	return 42
}

func (*Msg) StructureSinceVersion() uint16 {
	return 0
}

func (m *Msg) StructureInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.StructureSinceVersion()
}

func (*Msg) StructureDeprecated() uint16 {
	return 0
}

func (*Msg) StructureMetaAttribute(meta int) string {
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
