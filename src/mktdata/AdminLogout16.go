// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
)

type AdminLogout16 struct {
	Text [180]byte
}

func (a *AdminLogout16) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := a.RangeCheck(a.SbeSchemaVersion(), a.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, a.Text[:]); err != nil {
		return err
	}
	return nil
}

func (a *AdminLogout16) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !a.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 180; idx++ {
			a.Text[idx] = a.TextNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, a.Text[:]); err != nil {
			return err
		}
	}
	if actingVersion > a.SbeSchemaVersion() && blockLength > a.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-a.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := a.RangeCheck(actingVersion, a.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (a *AdminLogout16) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if a.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 180; idx++ {
			if a.Text[idx] < a.TextMinValue() || a.Text[idx] > a.TextMaxValue() {
				return fmt.Errorf("Range check failed on a.Text[%d] (%v < %v > %v)", idx, a.TextMinValue(), a.Text[idx], a.TextMaxValue())
			}
		}
	}
	return nil
}

func AdminLogout16Init(a *AdminLogout16) {
	return
}

func (*AdminLogout16) SbeBlockLength() (blockLength uint16) {
	return 180
}

func (*AdminLogout16) SbeTemplateId() (templateId uint16) {
	return 16
}

func (*AdminLogout16) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*AdminLogout16) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*AdminLogout16) SbeSemanticType() (semanticType []byte) {
	return []byte("5")
}

func (*AdminLogout16) TextId() uint16 {
	return 58
}

func (*AdminLogout16) TextSinceVersion() uint16 {
	return 0
}

func (a *AdminLogout16) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.TextSinceVersion()
}

func (*AdminLogout16) TextDeprecated() uint16 {
	return 0
}

func (*AdminLogout16) TextMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*AdminLogout16) TextMinValue() byte {
	return byte(32)
}

func (*AdminLogout16) TextMaxValue() byte {
	return byte(126)
}

func (*AdminLogout16) TextNullValue() byte {
	return 0
}

func (a *AdminLogout16) TextCharacterEncoding() string {
	return "US-ASCII"
}
