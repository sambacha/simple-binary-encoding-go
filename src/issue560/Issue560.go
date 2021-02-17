// Generated SBE (Simple Binary Encoding) message codec

package issue560

import (
	"io"
	"io/ioutil"
)

type Issue560 struct {
	DiscountedModel ModelEnum
}

func (i *Issue560) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := i.RangeCheck(i.SbeSchemaVersion(), i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (i *Issue560) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	i.DiscountedModel = Model.C
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

func (i *Issue560) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := i.DiscountedModel.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func Issue560Init(i *Issue560) {
	i.DiscountedModel = Model.C
	return
}

func (*Issue560) SbeBlockLength() (blockLength uint16) {
	return 0
}

func (*Issue560) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*Issue560) SbeSchemaId() (schemaId uint16) {
	return 560
}

func (*Issue560) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Issue560) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Issue560) DiscountedModelId() uint16 {
	return 8
}

func (*Issue560) DiscountedModelSinceVersion() uint16 {
	return 0
}

func (i *Issue560) DiscountedModelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.DiscountedModelSinceVersion()
}

func (*Issue560) DiscountedModelDeprecated() uint16 {
	return 0
}

func (*Issue560) DiscountedModelMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "constant"
	}
	return ""
}
