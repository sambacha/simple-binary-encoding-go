// Generated SBE (Simple Binary Encoding) message codec

package composite

import (
	"io"
	"io/ioutil"
)

type Composite struct {
	Start Point
	End   Point
}

func (c *Composite) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := c.RangeCheck(c.SbeSchemaVersion(), c.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := c.Start.Encode(_m, _w); err != nil {
		return err
	}
	if err := c.End.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (c *Composite) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if c.StartInActingVersion(actingVersion) {
		if err := c.Start.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if c.EndInActingVersion(actingVersion) {
		if err := c.End.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > c.SbeSchemaVersion() && blockLength > c.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-c.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := c.RangeCheck(actingVersion, c.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (c *Composite) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func CompositeInit(c *Composite) {
	return
}

func (*Composite) SbeBlockLength() (blockLength uint16) {
	return 42
}

func (*Composite) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*Composite) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*Composite) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Composite) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Composite) StartId() uint16 {
	return 2
}

func (*Composite) StartSinceVersion() uint16 {
	return 0
}

func (c *Composite) StartInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.StartSinceVersion()
}

func (*Composite) StartDeprecated() uint16 {
	return 0
}

func (*Composite) StartMetaAttribute(meta int) string {
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

func (*Composite) EndId() uint16 {
	return 3
}

func (*Composite) EndSinceVersion() uint16 {
	return 0
}

func (c *Composite) EndInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.EndSinceVersion()
}

func (*Composite) EndDeprecated() uint16 {
	return 0
}

func (*Composite) EndMetaAttribute(meta int) string {
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
