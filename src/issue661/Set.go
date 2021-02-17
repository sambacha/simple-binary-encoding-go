// Generated SBE (Simple Binary Encoding) message codec

package issue661

import (
	"io"
)

type Set [8]bool
type SetChoiceValue uint8
type SetChoiceValues struct {
	One SetChoiceValue
	Two SetChoiceValue
}

var SetChoice = SetChoiceValues{0, 1}

func (s *Set) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint8 = 0
	for k, v := range s {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint8(_w, wireval)
}

func (s *Set) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	var wireval uint8

	if err := _m.ReadUint8(_r, &wireval); err != nil {
		return err
	}

	var idx uint
	for idx = 0; idx < 8; idx++ {
		s[idx] = (wireval & (1 << idx)) > 0
	}
	return nil
}

func (Set) EncodedLength() int64 {
	return 1
}

func (*Set) OneSinceVersion() uint16 {
	return 0
}

func (s *Set) OneInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.OneSinceVersion()
}

func (*Set) OneDeprecated() uint16 {
	return 0
}

func (*Set) TwoSinceVersion() uint16 {
	return 0
}

func (s *Set) TwoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TwoSinceVersion()
}

func (*Set) TwoDeprecated() uint16 {
	return 0
}
