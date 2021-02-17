// Generated SBE (Simple Binary Encoding) message codec

package issue472

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type Issue472 struct {
	Optional uint64
}

func (i *Issue472) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := i.RangeCheck(i.SbeSchemaVersion(), i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, i.Optional); err != nil {
		return err
	}
	return nil
}

func (i *Issue472) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !i.OptionalInActingVersion(actingVersion) {
		i.Optional = i.OptionalNullValue()
	} else {
		if err := _m.ReadUint64(_r, &i.Optional); err != nil {
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

func (i *Issue472) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if i.OptionalInActingVersion(actingVersion) {
		if i.Optional != i.OptionalNullValue() && (i.Optional < i.OptionalMinValue() || i.Optional > i.OptionalMaxValue()) {
			return fmt.Errorf("Range check failed on i.Optional (%v < %v > %v)", i.OptionalMinValue(), i.Optional, i.OptionalMaxValue())
		}
	}
	return nil
}

func Issue472Init(i *Issue472) {
	i.Optional = math.MaxUint64
	return
}

func (*Issue472) SbeBlockLength() (blockLength uint16) {
	return 8
}

func (*Issue472) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*Issue472) SbeSchemaId() (schemaId uint16) {
	return 472
}

func (*Issue472) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Issue472) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Issue472) OptionalId() uint16 {
	return 2
}

func (*Issue472) OptionalSinceVersion() uint16 {
	return 0
}

func (i *Issue472) OptionalInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.OptionalSinceVersion()
}

func (*Issue472) OptionalDeprecated() uint16 {
	return 0
}

func (*Issue472) OptionalMetaAttribute(meta int) string {
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

func (*Issue472) OptionalMinValue() uint64 {
	return 0
}

func (*Issue472) OptionalMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*Issue472) OptionalNullValue() uint64 {
	return math.MaxUint64
}
