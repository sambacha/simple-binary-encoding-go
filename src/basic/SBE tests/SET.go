// Generated SBE (Simple Binary Encoding) message codec

package sbe_tests

import (
	"io"
)

type SET [32]bool
type SETChoiceValue uint8
type SETChoiceValues struct {
	Bit0  SETChoiceValue
	Bit16 SETChoiceValue
	Bit26 SETChoiceValue
}

var SETChoice = SETChoiceValues{0, 16, 26}

func (s *SET) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint32 = 0
	for k, v := range s {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint32(_w, wireval)
}

func (s *SET) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
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

func (SET) EncodedLength() int64 {
	return 4
}

func (*SET) Bit0SinceVersion() uint16 {
	return 0
}

func (s *SET) Bit0InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.Bit0SinceVersion()
}

func (*SET) Bit0Deprecated() uint16 {
	return 0
}

func (*SET) Bit16SinceVersion() uint16 {
	return 0
}

func (s *SET) Bit16InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.Bit16SinceVersion()
}

func (*SET) Bit16Deprecated() uint16 {
	return 0
}

func (*SET) Bit26SinceVersion() uint16 {
	return 0
}

func (s *SET) Bit26InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.Bit26SinceVersion()
}

func (*SET) Bit26Deprecated() uint16 {
	return 0
}
