// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type AdminLogin15 struct {
	HeartBtInt int8
}

func (a *AdminLogin15) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := a.RangeCheck(a.SbeSchemaVersion(), a.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt8(_w, a.HeartBtInt); err != nil {
		return err
	}
	return nil
}

func (a *AdminLogin15) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !a.HeartBtIntInActingVersion(actingVersion) {
		a.HeartBtInt = a.HeartBtIntNullValue()
	} else {
		if err := _m.ReadInt8(_r, &a.HeartBtInt); err != nil {
			return err
		}
	}
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

func (a *AdminLogin15) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if a.HeartBtIntInActingVersion(actingVersion) {
		if a.HeartBtInt < a.HeartBtIntMinValue() || a.HeartBtInt > a.HeartBtIntMaxValue() {
			return fmt.Errorf("Range check failed on a.HeartBtInt (%v < %v > %v)", a.HeartBtIntMinValue(), a.HeartBtInt, a.HeartBtIntMaxValue())
		}
	}
	return nil
}

func AdminLogin15Init(a *AdminLogin15) {
	return
}

func (*AdminLogin15) SbeBlockLength() (blockLength uint16) {
	return 1
}

func (*AdminLogin15) SbeTemplateId() (templateId uint16) {
	return 15
}

func (*AdminLogin15) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*AdminLogin15) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*AdminLogin15) SbeSemanticType() (semanticType []byte) {
	return []byte("A")
}

func (*AdminLogin15) HeartBtIntId() uint16 {
	return 108
}

func (*AdminLogin15) HeartBtIntSinceVersion() uint16 {
	return 0
}

func (a *AdminLogin15) HeartBtIntInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.HeartBtIntSinceVersion()
}

func (*AdminLogin15) HeartBtIntDeprecated() uint16 {
	return 0
}

func (*AdminLogin15) HeartBtIntMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*AdminLogin15) HeartBtIntMinValue() int8 {
	return math.MinInt8 + 1
}

func (*AdminLogin15) HeartBtIntMaxValue() int8 {
	return math.MaxInt8
}

func (*AdminLogin15) HeartBtIntNullValue() int8 {
	return math.MinInt8
}
