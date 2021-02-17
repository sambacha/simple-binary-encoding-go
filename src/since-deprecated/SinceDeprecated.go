// Generated SBE (Simple Binary Encoding) message codec

package since_deprecated

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type SinceDeprecated struct {
	V1 uint64
	V2 uint64
	V3 uint64
}

func (s *SinceDeprecated) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := s.RangeCheck(s.SbeSchemaVersion(), s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, s.V1); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.V2); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.V3); err != nil {
		return err
	}
	return nil
}

func (s *SinceDeprecated) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !s.V1InActingVersion(actingVersion) {
		s.V1 = s.V1NullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.V1); err != nil {
			return err
		}
	}
	if !s.V2InActingVersion(actingVersion) {
		s.V2 = s.V2NullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.V2); err != nil {
			return err
		}
	}
	if !s.V3InActingVersion(actingVersion) {
		s.V3 = s.V3NullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.V3); err != nil {
			return err
		}
	}
	if actingVersion > s.SbeSchemaVersion() && blockLength > s.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-s.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := s.RangeCheck(actingVersion, s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (s *SinceDeprecated) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.V1InActingVersion(actingVersion) {
		if s.V1 < s.V1MinValue() || s.V1 > s.V1MaxValue() {
			return fmt.Errorf("Range check failed on s.V1 (%v < %v > %v)", s.V1MinValue(), s.V1, s.V1MaxValue())
		}
	}
	if s.V2InActingVersion(actingVersion) {
		if s.V2 < s.V2MinValue() || s.V2 > s.V2MaxValue() {
			return fmt.Errorf("Range check failed on s.V2 (%v < %v > %v)", s.V2MinValue(), s.V2, s.V2MaxValue())
		}
	}
	if s.V3InActingVersion(actingVersion) {
		if s.V3 < s.V3MinValue() || s.V3 > s.V3MaxValue() {
			return fmt.Errorf("Range check failed on s.V3 (%v < %v > %v)", s.V3MinValue(), s.V3, s.V3MaxValue())
		}
	}
	return nil
}

func SinceDeprecatedInit(s *SinceDeprecated) {
	return
}

func (*SinceDeprecated) SbeBlockLength() (blockLength uint16) {
	return 24
}

func (*SinceDeprecated) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*SinceDeprecated) SbeSchemaId() (schemaId uint16) {
	return 108
}

func (*SinceDeprecated) SbeSchemaVersion() (schemaVersion uint16) {
	return 4
}

func (*SinceDeprecated) SbeSemanticType() (semanticType []byte) {
	return []byte("n/a")
}

func (*SinceDeprecated) V1Id() uint16 {
	return 1
}

func (*SinceDeprecated) V1SinceVersion() uint16 {
	return 0
}

func (s *SinceDeprecated) V1InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.V1SinceVersion()
}

func (*SinceDeprecated) V1Deprecated() uint16 {
	return 0
}

func (*SinceDeprecated) V1MetaAttribute(meta int) string {
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

func (*SinceDeprecated) V1MinValue() uint64 {
	return 0
}

func (*SinceDeprecated) V1MaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SinceDeprecated) V1NullValue() uint64 {
	return math.MaxUint64
}

func (*SinceDeprecated) V2Id() uint16 {
	return 2
}

func (*SinceDeprecated) V2SinceVersion() uint16 {
	return 2
}

func (s *SinceDeprecated) V2InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.V2SinceVersion()
}

func (*SinceDeprecated) V2Deprecated() uint16 {
	return 0
}

func (*SinceDeprecated) V2MetaAttribute(meta int) string {
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

func (*SinceDeprecated) V2MinValue() uint64 {
	return 0
}

func (*SinceDeprecated) V2MaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SinceDeprecated) V2NullValue() uint64 {
	return math.MaxUint64
}

func (*SinceDeprecated) V3Id() uint16 {
	return 3
}

func (*SinceDeprecated) V3SinceVersion() uint16 {
	return 3
}

func (s *SinceDeprecated) V3InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.V3SinceVersion()
}

func (*SinceDeprecated) V3Deprecated() uint16 {
	return 4
}

func (*SinceDeprecated) V3MetaAttribute(meta int) string {
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

func (*SinceDeprecated) V3MinValue() uint64 {
	return 0
}

func (*SinceDeprecated) V3MaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SinceDeprecated) V3NullValue() uint64 {
	return math.MaxUint64
}
