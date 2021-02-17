// Generated SBE (Simple Binary Encoding) message codec

package issue435

import (
	"io"
)

type SetRef [8]bool
type SetRefChoiceValue uint8
type SetRefChoiceValues struct {
	One SetRefChoiceValue
	Two SetRefChoiceValue
}

var SetRefChoice = SetRefChoiceValues{0, 1}

func (s *SetRef) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint8 = 0
	for k, v := range s {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint8(_w, wireval)
}

func (s *SetRef) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
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

func (SetRef) EncodedLength() int64 {
	return 1
}

func (*SetRef) OneSinceVersion() uint16 {
	return 0
}

func (s *SetRef) OneInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.OneSinceVersion()
}

func (*SetRef) OneDeprecated() uint16 {
	return 0
}

func (*SetRef) TwoSinceVersion() uint16 {
	return 0
}

func (s *SetRef) TwoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TwoSinceVersion()
}

func (*SetRef) TwoDeprecated() uint16 {
	return 0
}
