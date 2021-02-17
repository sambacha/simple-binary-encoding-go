// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"io"
	"io/ioutil"
)

type AdminHeartbeat12 struct {
}

func (a *AdminHeartbeat12) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := a.RangeCheck(a.SbeSchemaVersion(), a.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (a *AdminHeartbeat12) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if actingVersion > a.SbeSchemaVersion() && blockLength > a.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-a.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := a.RangeCheck(actingVersion, a.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (a *AdminHeartbeat12) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func AdminHeartbeat12Init(a *AdminHeartbeat12) {
	return
}

func (*AdminHeartbeat12) SbeBlockLength() (blockLength uint16) {
	return 0
}

func (*AdminHeartbeat12) SbeTemplateId() (templateId uint16) {
	return 12
}

func (*AdminHeartbeat12) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*AdminHeartbeat12) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*AdminHeartbeat12) SbeSemanticType() (semanticType []byte) {
	return []byte("0")
}
