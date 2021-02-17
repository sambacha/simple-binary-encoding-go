// Generated SBE (Simple Binary Encoding) message codec

package issue483

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type Issue483 struct {
	Unset    uint8
	Required uint8
	Constant uint8
	Optional uint8
}

func (i *Issue483) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := i.RangeCheck(i.SbeSchemaVersion(), i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, i.Unset); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, i.Required); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, i.Optional); err != nil {
		return err
	}
	return nil
}

func (i *Issue483) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !i.UnsetInActingVersion(actingVersion) {
		i.Unset = i.UnsetNullValue()
	} else {
		if err := _m.ReadUint8(_r, &i.Unset); err != nil {
			return err
		}
	}
	if !i.RequiredInActingVersion(actingVersion) {
		i.Required = i.RequiredNullValue()
	} else {
		if err := _m.ReadUint8(_r, &i.Required); err != nil {
			return err
		}
	}
	i.Constant = 1
	if !i.OptionalInActingVersion(actingVersion) {
		i.Optional = i.OptionalNullValue()
	} else {
		if err := _m.ReadUint8(_r, &i.Optional); err != nil {
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

func (i *Issue483) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if i.UnsetInActingVersion(actingVersion) {
		if i.Unset < i.UnsetMinValue() || i.Unset > i.UnsetMaxValue() {
			return fmt.Errorf("Range check failed on i.Unset (%v < %v > %v)", i.UnsetMinValue(), i.Unset, i.UnsetMaxValue())
		}
	}
	if i.RequiredInActingVersion(actingVersion) {
		if i.Required < i.RequiredMinValue() || i.Required > i.RequiredMaxValue() {
			return fmt.Errorf("Range check failed on i.Required (%v < %v > %v)", i.RequiredMinValue(), i.Required, i.RequiredMaxValue())
		}
	}
	if i.OptionalInActingVersion(actingVersion) {
		if i.Optional != i.OptionalNullValue() && (i.Optional < i.OptionalMinValue() || i.Optional > i.OptionalMaxValue()) {
			return fmt.Errorf("Range check failed on i.Optional (%v < %v > %v)", i.OptionalMinValue(), i.Optional, i.OptionalMaxValue())
		}
	}
	return nil
}

func Issue483Init(i *Issue483) {
	i.Constant = 1
	i.Optional = math.MaxUint8
	return
}

func (*Issue483) SbeBlockLength() (blockLength uint16) {
	return 3
}

func (*Issue483) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*Issue483) SbeSchemaId() (schemaId uint16) {
	return 483
}

func (*Issue483) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Issue483) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Issue483) UnsetId() uint16 {
	return 2
}

func (*Issue483) UnsetSinceVersion() uint16 {
	return 0
}

func (i *Issue483) UnsetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.UnsetSinceVersion()
}

func (*Issue483) UnsetDeprecated() uint16 {
	return 0
}

func (*Issue483) UnsetMetaAttribute(meta int) string {
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

func (*Issue483) UnsetMinValue() uint8 {
	return 0
}

func (*Issue483) UnsetMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*Issue483) UnsetNullValue() uint8 {
	return math.MaxUint8
}

func (*Issue483) RequiredId() uint16 {
	return 3
}

func (*Issue483) RequiredSinceVersion() uint16 {
	return 0
}

func (i *Issue483) RequiredInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.RequiredSinceVersion()
}

func (*Issue483) RequiredDeprecated() uint16 {
	return 0
}

func (*Issue483) RequiredMetaAttribute(meta int) string {
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

func (*Issue483) RequiredMinValue() uint8 {
	return 0
}

func (*Issue483) RequiredMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*Issue483) RequiredNullValue() uint8 {
	return math.MaxUint8
}

func (*Issue483) ConstantId() uint16 {
	return 4
}

func (*Issue483) ConstantSinceVersion() uint16 {
	return 0
}

func (i *Issue483) ConstantInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.ConstantSinceVersion()
}

func (*Issue483) ConstantDeprecated() uint16 {
	return 0
}

func (*Issue483) ConstantMetaAttribute(meta int) string {
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

func (*Issue483) ConstantMinValue() uint8 {
	return 0
}

func (*Issue483) ConstantMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*Issue483) ConstantNullValue() uint8 {
	return math.MaxUint8
}

func (*Issue483) OptionalId() uint16 {
	return 5
}

func (*Issue483) OptionalSinceVersion() uint16 {
	return 0
}

func (i *Issue483) OptionalInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.OptionalSinceVersion()
}

func (*Issue483) OptionalDeprecated() uint16 {
	return 0
}

func (*Issue483) OptionalMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*Issue483) OptionalMinValue() uint8 {
	return 0
}

func (*Issue483) OptionalMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*Issue483) OptionalNullValue() uint8 {
	return math.MaxUint8
}
