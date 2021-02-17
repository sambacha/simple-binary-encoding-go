// Generated SBE (Simple Binary Encoding) message codec

package composite_elements

import (
	"io"
)

type SetOne [32]bool
type SetOneChoiceValue uint8
type SetOneChoiceValues struct {
	Bit0  SetOneChoiceValue
	Bit16 SetOneChoiceValue
	Bit26 SetOneChoiceValue
}

var SetOneChoice = SetOneChoiceValues{0, 16, 26}

func (s *SetOne) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint32 = 0
	for k, v := range s {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint32(_w, wireval)
}

func (s *SetOne) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	var wireval uint32

	if err := _m.ReadUint32(_r, &wireval); err != nil {
		return err
	}

	var idx uint
	for idx = 0; idx < 32; idx++ {
		s[idx] = (wireval & (1 << idx)) > 0
	}
	return nil
}

func (SetOne) EncodedLength() int64 {
	return 4
}

func (*SetOne) Bit0SinceVersion() uint16 {
	return 0
}

func (s *SetOne) Bit0InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.Bit0SinceVersion()
}

func (*SetOne) Bit0Deprecated() uint16 {
	return 0
}

func (*SetOne) Bit16SinceVersion() uint16 {
	return 0
}

func (s *SetOne) Bit16InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.Bit16SinceVersion()
}

func (*SetOne) Bit16Deprecated() uint16 {
	return 0
}

func (*SetOne) Bit26SinceVersion() uint16 {
	return 0
}

func (s *SetOne) Bit26InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.Bit26SinceVersion()
}

func (*SetOne) Bit26Deprecated() uint16 {
	return 0
}
