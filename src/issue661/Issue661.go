// Generated SBE (Simple Binary Encoding) message codec

package issue661

import (
	"io"
	"io/ioutil"
)

type Issue661 struct {
	Set0 set
	Set1 set
}

func (i *Issue661) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := i.RangeCheck(i.SbeSchemaVersion(), i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := i.Set0.Encode(_m, _w); err != nil {
		return err
	}
	if err := i.Set1.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (i *Issue661) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if i.Set0InActingVersion(actingVersion) {
		if err := i.Set0.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if i.Set1InActingVersion(actingVersion) {
		if err := i.Set1.Decode(_m, _r, actingVersion); err != nil {
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

func (i *Issue661) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func Issue661Init(i *Issue661) {
	return
}

func (*Issue661) SbeBlockLength() (blockLength uint16) {
	return 2
}

func (*Issue661) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*Issue661) SbeSchemaId() (schemaId uint16) {
	return 661
}

func (*Issue661) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*Issue661) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Issue661) Set0Id() uint16 {
	return 10
}

func (*Issue661) Set0SinceVersion() uint16 {
	return 0
}

func (i *Issue661) Set0InActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.Set0SinceVersion()
}

func (*Issue661) Set0Deprecated() uint16 {
	return 0
}

func (*Issue661) Set0MetaAttribute(meta int) string {
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

func (*Issue661) Set1Id() uint16 {
	return 11
}

func (*Issue661) Set1SinceVersion() uint16 {
	return 1
}

func (i *Issue661) Set1InActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.Set1SinceVersion()
}

func (*Issue661) Set1Deprecated() uint16 {
	return 0
}

func (*Issue661) Set1MetaAttribute(meta int) string {
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
