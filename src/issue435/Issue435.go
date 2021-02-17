// Generated SBE (Simple Binary Encoding) message codec

package issue435

import (
	"io"
	"io/ioutil"
)

type Issue435 struct {
	Example ExampleRef
}

func (i *Issue435) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := i.RangeCheck(i.SbeSchemaVersion(), i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := i.Example.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (i *Issue435) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if i.ExampleInActingVersion(actingVersion) {
		if err := i.Example.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > i.SbeSchemaVersion() && blockLength > i.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-i.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := i.RangeCheck(actingVersion, i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (i *Issue435) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func Issue435Init(i *Issue435) {
	return
}

func (*Issue435) SbeBlockLength() (blockLength uint16) {
	return 1
}

func (*Issue435) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*Issue435) SbeSchemaId() (schemaId uint16) {
	return 435
}

func (*Issue435) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Issue435) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Issue435) ExampleId() uint16 {
	return 10
}

func (*Issue435) ExampleSinceVersion() uint16 {
	return 0
}

func (i *Issue435) ExampleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.ExampleSinceVersion()
}

func (*Issue435) ExampleDeprecated() uint16 {
	return 0
}

func (*Issue435) ExampleMetaAttribute(meta int) string {
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
