// Generated SBE (Simple Binary Encoding) message codec

package issue488

import (
	"io"
	"io/ioutil"
)

type Issue488 struct {
	VarData []uint8
}

func (i *Issue488) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := i.RangeCheck(i.SbeSchemaVersion(), i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, uint32(len(i.VarData))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, i.VarData); err != nil {
		return err
	}
	return nil
}

func (i *Issue488) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if actingVersion > i.SbeSchemaVersion() && blockLength > i.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-i.SbeBlockLength()))
	}

	if i.VarDataInActingVersion(actingVersion) {
		var VarDataLength uint32
		if err := _m.ReadUint32(_r, &VarDataLength); err != nil {
			return err
		}
		if cap(i.VarData) < int(VarDataLength) {
			i.VarData = make([]uint8, VarDataLength)
		}
		i.VarData = i.VarData[:VarDataLength]
		if err := _m.ReadBytes(_r, i.VarData); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := i.RangeCheck(actingVersion, i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (i *Issue488) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func Issue488Init(i *Issue488) {
	return
}

func (*Issue488) SbeBlockLength() (blockLength uint16) {
	return 0
}

func (*Issue488) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*Issue488) SbeSchemaId() (schemaId uint16) {
	return 488
}

func (*Issue488) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Issue488) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Issue488) VarDataMetaAttribute(meta int) string {
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

func (*Issue488) VarDataSinceVersion() uint16 {
	return 0
}

func (i *Issue488) VarDataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.VarDataSinceVersion()
}

func (*Issue488) VarDataDeprecated() uint16 {
	return 0
}

func (Issue488) VarDataCharacterEncoding() string {
	return "null"
}

func (Issue488) VarDataHeaderLength() uint64 {
	return 4
}
