// Generated SBE (Simple Binary Encoding) message codec

package issue435

import (
	"io"
)

type ExampleRef struct {
	E EnumRefEnum
}

func (e *ExampleRef) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := e.E.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (e *ExampleRef) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if e.EInActingVersion(actingVersion) {
		if err := e.E.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	return nil
}

func (e *ExampleRef) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func ExampleRefInit(e *ExampleRef) {
	return
}

func (*ExampleRef) EncodedLength() int64 {
	return 1
}

func (*ExampleRef) ESinceVersion() uint16 {
	return 0
}

func (e *ExampleRef) EInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ESinceVersion()
}

func (*ExampleRef) EDeprecated() uint16 {
	return 0
}
