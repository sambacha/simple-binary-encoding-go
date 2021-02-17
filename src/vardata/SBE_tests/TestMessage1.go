// Generated SBE (Simple Binary Encoding) message codec

package sbe_tests

import (
	"io"
	"io/ioutil"
)

type TestMessage1 struct {
	EncryptedNewPassword []byte
}

func (t *TestMessage1) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := t.RangeCheck(t.SbeSchemaVersion(), t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, uint8(len(t.EncryptedNewPassword))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.EncryptedNewPassword); err != nil {
		return err
	}
	return nil
}

func (t *TestMessage1) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if actingVersion > t.SbeSchemaVersion() && blockLength > t.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-t.SbeBlockLength()))
	}

	if t.EncryptedNewPasswordInActingVersion(actingVersion) {
		var EncryptedNewPasswordLength uint8
		if err := _m.ReadUint8(_r, &EncryptedNewPasswordLength); err != nil {
			return err
		}
		if cap(t.EncryptedNewPassword) < int(EncryptedNewPasswordLength) {
			t.EncryptedNewPassword = make([]byte, EncryptedNewPasswordLength)
		}
		t.EncryptedNewPassword = t.EncryptedNewPassword[:EncryptedNewPasswordLength]
		if err := _m.ReadBytes(_r, t.EncryptedNewPassword); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := t.RangeCheck(actingVersion, t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (t *TestMessage1) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func TestMessage1Init(t *TestMessage1) {
	return
}

func (*TestMessage1) SbeBlockLength() (blockLength uint16) {
	return 0
}

func (*TestMessage1) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*TestMessage1) SbeSchemaId() (schemaId uint16) {
	return 4
}

func (*TestMessage1) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*TestMessage1) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*TestMessage1) EncryptedNewPasswordMetaAttribute(meta int) string {
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

func (*TestMessage1) EncryptedNewPasswordSinceVersion() uint16 {
	return 0
}

func (t *TestMessage1) EncryptedNewPasswordInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.EncryptedNewPasswordSinceVersion()
}

func (*TestMessage1) EncryptedNewPasswordDeprecated() uint16 {
	return 0
}

func (TestMessage1) EncryptedNewPasswordCharacterEncoding() string {
	return "US-ASCII"
}

func (TestMessage1) EncryptedNewPasswordHeaderLength() uint64 {
	return 1
}
