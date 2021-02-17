// Generated SBE (Simple Binary Encoding) message codec

package simple

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type Simple0 struct {
	U64nv        uint64
	U64          uint64
	U32          uint32
	U16          uint16
	U8           uint8
	S8           int8
	S16          int16
	S32          int32
	S64          int64
	F32          float32
	D64          float64
	String6ASCII [6]byte
	String1ASCII byte
	Int2         [2]int32
}

func (s *Simple0) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := s.RangeCheck(s.SbeSchemaVersion(), s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, s.U64nv); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.U64); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, s.U32); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, s.U16); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, s.U8); err != nil {
		return err
	}

	for i := 0; i < 8; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	if err := _m.WriteInt8(_w, s.S8); err != nil {
		return err
	}
	if err := _m.WriteInt16(_w, s.S16); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.S32); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, s.S64); err != nil {
		return err
	}
	if err := _m.WriteFloat32(_w, s.F32); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, s.D64); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.String6ASCII[:]); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, s.String1ASCII); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt32(_w, s.Int2[idx]); err != nil {
			return err
		}
	}

	for i := 0; i < 3; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	return nil
}

func (s *Simple0) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint8, doRangeCheck bool) error {
	if !s.U64nvInActingVersion(actingVersion) {
		s.U64nv = s.U64nvNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.U64nv); err != nil {
			return err
		}
	}
	if !s.U64InActingVersion(actingVersion) {
		s.U64 = s.U64NullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.U64); err != nil {
			return err
		}
	}
	if !s.U32InActingVersion(actingVersion) {
		s.U32 = s.U32NullValue()
	} else {
		if err := _m.ReadUint32(_r, &s.U32); err != nil {
			return err
		}
	}
	if !s.U16InActingVersion(actingVersion) {
		s.U16 = s.U16NullValue()
	} else {
		if err := _m.ReadUint16(_r, &s.U16); err != nil {
			return err
		}
	}
	if !s.U8InActingVersion(actingVersion) {
		s.U8 = s.U8NullValue()
	} else {
		if err := _m.ReadUint8(_r, &s.U8); err != nil {
			return err
		}
	}
	io.CopyN(ioutil.Discard, _r, 8)
	if !s.S8InActingVersion(actingVersion) {
		s.S8 = s.S8NullValue()
	} else {
		if err := _m.ReadInt8(_r, &s.S8); err != nil {
			return err
		}
	}
	if !s.S16InActingVersion(actingVersion) {
		s.S16 = s.S16NullValue()
	} else {
		if err := _m.ReadInt16(_r, &s.S16); err != nil {
			return err
		}
	}
	if !s.S32InActingVersion(actingVersion) {
		s.S32 = s.S32NullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.S32); err != nil {
			return err
		}
	}
	if !s.S64InActingVersion(actingVersion) {
		s.S64 = s.S64NullValue()
	} else {
		if err := _m.ReadInt64(_r, &s.S64); err != nil {
			return err
		}
	}
	if !s.F32InActingVersion(actingVersion) {
		s.F32 = s.F32NullValue()
	} else {
		if err := _m.ReadFloat32(_r, &s.F32); err != nil {
			return err
		}
	}
	if !s.D64InActingVersion(actingVersion) {
		s.D64 = s.D64NullValue()
	} else {
		if err := _m.ReadFloat64(_r, &s.D64); err != nil {
			return err
		}
	}
	if !s.String6ASCIIInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			s.String6ASCII[idx] = s.String6ASCIINullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.String6ASCII[:]); err != nil {
			return err
		}
	}
	if !s.String1ASCIIInActingVersion(actingVersion) {
		s.String1ASCII = s.String1ASCIINullValue()
	} else {
		if err := _m.ReadUint8(_r, &s.String1ASCII); err != nil {
			return err
		}
	}
	if !s.Int2InActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			s.Int2[idx] = s.Int2NullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt32(_r, &s.Int2[idx]); err != nil {
				return err
			}
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
	io.CopyN(ioutil.Discard, _r, 3)
	return nil
}

func (s *Simple0) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.U64nvInActingVersion(actingVersion) {
		if s.U64nv != s.U64nvNullValue() && (s.U64nv < s.U64nvMinValue() || s.U64nv > s.U64nvMaxValue()) {
			return fmt.Errorf("Range check failed on s.U64nv (%v < %v > %v)", s.U64nvMinValue(), s.U64nv, s.U64nvMaxValue())
		}
	}
	if s.U64InActingVersion(actingVersion) {
		if s.U64 < s.U64MinValue() || s.U64 > s.U64MaxValue() {
			return fmt.Errorf("Range check failed on s.U64 (%v < %v > %v)", s.U64MinValue(), s.U64, s.U64MaxValue())
		}
	}
	if s.U32InActingVersion(actingVersion) {
		if s.U32 < s.U32MinValue() || s.U32 > s.U32MaxValue() {
			return fmt.Errorf("Range check failed on s.U32 (%v < %v > %v)", s.U32MinValue(), s.U32, s.U32MaxValue())
		}
	}
	if s.U16InActingVersion(actingVersion) {
		if s.U16 < s.U16MinValue() || s.U16 > s.U16MaxValue() {
			return fmt.Errorf("Range check failed on s.U16 (%v < %v > %v)", s.U16MinValue(), s.U16, s.U16MaxValue())
		}
	}
	if s.U8InActingVersion(actingVersion) {
		if s.U8 < s.U8MinValue() || s.U8 > s.U8MaxValue() {
			return fmt.Errorf("Range check failed on s.U8 (%v < %v > %v)", s.U8MinValue(), s.U8, s.U8MaxValue())
		}
	}
	if s.S8InActingVersion(actingVersion) {
		if s.S8 < s.S8MinValue() || s.S8 > s.S8MaxValue() {
			return fmt.Errorf("Range check failed on s.S8 (%v < %v > %v)", s.S8MinValue(), s.S8, s.S8MaxValue())
		}
	}
	if s.S16InActingVersion(actingVersion) {
		if s.S16 < s.S16MinValue() || s.S16 > s.S16MaxValue() {
			return fmt.Errorf("Range check failed on s.S16 (%v < %v > %v)", s.S16MinValue(), s.S16, s.S16MaxValue())
		}
	}
	if s.S32InActingVersion(actingVersion) {
		if s.S32 < s.S32MinValue() || s.S32 > s.S32MaxValue() {
			return fmt.Errorf("Range check failed on s.S32 (%v < %v > %v)", s.S32MinValue(), s.S32, s.S32MaxValue())
		}
	}
	if s.S64InActingVersion(actingVersion) {
		if s.S64 < s.S64MinValue() || s.S64 > s.S64MaxValue() {
			return fmt.Errorf("Range check failed on s.S64 (%v < %v > %v)", s.S64MinValue(), s.S64, s.S64MaxValue())
		}
	}
	if s.F32InActingVersion(actingVersion) {
		if s.F32 < s.F32MinValue() || s.F32 > s.F32MaxValue() {
			return fmt.Errorf("Range check failed on s.F32 (%v < %v > %v)", s.F32MinValue(), s.F32, s.F32MaxValue())
		}
	}
	if s.D64InActingVersion(actingVersion) {
		if s.D64 < s.D64MinValue() || s.D64 > s.D64MaxValue() {
			return fmt.Errorf("Range check failed on s.D64 (%v < %v > %v)", s.D64MinValue(), s.D64, s.D64MaxValue())
		}
	}
	if s.String6ASCIIInActingVersion(actingVersion) {
		for idx := 0; idx < 6; idx++ {
			if s.String6ASCII[idx] < s.String6ASCIIMinValue() || s.String6ASCII[idx] > s.String6ASCIIMaxValue() {
				return fmt.Errorf("Range check failed on s.String6ASCII[%d] (%v < %v > %v)", idx, s.String6ASCIIMinValue(), s.String6ASCII[idx], s.String6ASCIIMaxValue())
			}
		}
	}
	for idx, ch := range s.String6ASCII {
		if ch > 127 {
			return fmt.Errorf("s.String6ASCII[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if s.String1ASCIIInActingVersion(actingVersion) {
		if s.String1ASCII < s.String1ASCIIMinValue() || s.String1ASCII > s.String1ASCIIMaxValue() {
			return fmt.Errorf("Range check failed on s.String1ASCII (%v < %v > %v)", s.String1ASCIIMinValue(), s.String1ASCII, s.String1ASCIIMaxValue())
		}
	}
	if s.Int2InActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if s.Int2[idx] < s.Int2MinValue() || s.Int2[idx] > s.Int2MaxValue() {
				return fmt.Errorf("Range check failed on s.Int2[%d] (%v < %v > %v)", idx, s.Int2MinValue(), s.Int2[idx], s.Int2MaxValue())
			}
		}
	}
	return nil
}

func Simple0Init(s *Simple0) {
	s.U64nv = 18446744073709551615
	return
}

func (*Simple0) SbeBlockLength() (blockLength uint8) {
	return 76
}

func (*Simple0) SbeTemplateId() (templateId uint16) {
	return 11
}

func (*Simple0) SbeSchemaId() (schemaId uint16) {
	return 3
}

func (*Simple0) SbeSchemaVersion() (schemaVersion uint16) {
	return 2
}

func (*Simple0) SbeSemanticType() (semanticType []byte) {
	return []byte("womble")
}

func (*Simple0) U64nvId() uint16 {
	return 10
}

func (*Simple0) U64nvSinceVersion() uint16 {
	return 1
}

func (s *Simple0) U64nvInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.U64nvSinceVersion()
}

func (*Simple0) U64nvDeprecated() uint16 {
	return 0
}

func (*Simple0) U64nvMetaAttribute(meta int) string {
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

func (*Simple0) U64nvMinValue() uint64 {
	return 0
}

func (*Simple0) U64nvMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*Simple0) U64nvNullValue() uint64 {
	return 18446744073709551615
}

func (*Simple0) U64Id() uint16 {
	return 20
}

func (*Simple0) U64SinceVersion() uint16 {
	return 1
}

func (s *Simple0) U64InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.U64SinceVersion()
}

func (*Simple0) U64Deprecated() uint16 {
	return 2
}

func (*Simple0) U64MetaAttribute(meta int) string {
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

func (*Simple0) U64MinValue() uint64 {
	return 0
}

func (*Simple0) U64MaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*Simple0) U64NullValue() uint64 {
	return math.MaxUint64
}

func (*Simple0) U32Id() uint16 {
	return 30
}

func (*Simple0) U32SinceVersion() uint16 {
	return 0
}

func (s *Simple0) U32InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.U32SinceVersion()
}

func (*Simple0) U32Deprecated() uint16 {
	return 0
}

func (*Simple0) U32MetaAttribute(meta int) string {
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

func (*Simple0) U32MinValue() uint32 {
	return 0
}

func (*Simple0) U32MaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*Simple0) U32NullValue() uint32 {
	return math.MaxUint32
}

func (*Simple0) U16Id() uint16 {
	return 40
}

func (*Simple0) U16SinceVersion() uint16 {
	return 0
}

func (s *Simple0) U16InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.U16SinceVersion()
}

func (*Simple0) U16Deprecated() uint16 {
	return 0
}

func (*Simple0) U16MetaAttribute(meta int) string {
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

func (*Simple0) U16MinValue() uint16 {
	return 0
}

func (*Simple0) U16MaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*Simple0) U16NullValue() uint16 {
	return math.MaxUint16
}

func (*Simple0) U8Id() uint16 {
	return 50
}

func (*Simple0) U8SinceVersion() uint16 {
	return 0
}

func (s *Simple0) U8InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.U8SinceVersion()
}

func (*Simple0) U8Deprecated() uint16 {
	return 0
}

func (*Simple0) U8MetaAttribute(meta int) string {
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

func (*Simple0) U8MinValue() uint8 {
	return 0
}

func (*Simple0) U8MaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*Simple0) U8NullValue() uint8 {
	return math.MaxUint8
}

func (*Simple0) S8Id() uint16 {
	return 60
}

func (*Simple0) S8SinceVersion() uint16 {
	return 0
}

func (s *Simple0) S8InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.S8SinceVersion()
}

func (*Simple0) S8Deprecated() uint16 {
	return 0
}

func (*Simple0) S8MetaAttribute(meta int) string {
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

func (*Simple0) S8MinValue() int8 {
	return math.MinInt8 + 1
}

func (*Simple0) S8MaxValue() int8 {
	return math.MaxInt8
}

func (*Simple0) S8NullValue() int8 {
	return math.MinInt8
}

func (*Simple0) S16Id() uint16 {
	return 70
}

func (*Simple0) S16SinceVersion() uint16 {
	return 0
}

func (s *Simple0) S16InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.S16SinceVersion()
}

func (*Simple0) S16Deprecated() uint16 {
	return 0
}

func (*Simple0) S16MetaAttribute(meta int) string {
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

func (*Simple0) S16MinValue() int16 {
	return math.MinInt16 + 1
}

func (*Simple0) S16MaxValue() int16 {
	return math.MaxInt16
}

func (*Simple0) S16NullValue() int16 {
	return math.MinInt16
}

func (*Simple0) S32Id() uint16 {
	return 80
}

func (*Simple0) S32SinceVersion() uint16 {
	return 0
}

func (s *Simple0) S32InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.S32SinceVersion()
}

func (*Simple0) S32Deprecated() uint16 {
	return 0
}

func (*Simple0) S32MetaAttribute(meta int) string {
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

func (*Simple0) S32MinValue() int32 {
	return math.MinInt32 + 1
}

func (*Simple0) S32MaxValue() int32 {
	return math.MaxInt32
}

func (*Simple0) S32NullValue() int32 {
	return math.MinInt32
}

func (*Simple0) S64Id() uint16 {
	return 90
}

func (*Simple0) S64SinceVersion() uint16 {
	return 0
}

func (s *Simple0) S64InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.S64SinceVersion()
}

func (*Simple0) S64Deprecated() uint16 {
	return 0
}

func (*Simple0) S64MetaAttribute(meta int) string {
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

func (*Simple0) S64MinValue() int64 {
	return math.MinInt64 + 1
}

func (*Simple0) S64MaxValue() int64 {
	return math.MaxInt64
}

func (*Simple0) S64NullValue() int64 {
	return math.MinInt64
}

func (*Simple0) F32Id() uint16 {
	return 100
}

func (*Simple0) F32SinceVersion() uint16 {
	return 0
}

func (s *Simple0) F32InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.F32SinceVersion()
}

func (*Simple0) F32Deprecated() uint16 {
	return 0
}

func (*Simple0) F32MetaAttribute(meta int) string {
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

func (*Simple0) F32MinValue() float32 {
	return -math.MaxFloat32
}

func (*Simple0) F32MaxValue() float32 {
	return math.MaxFloat32
}

func (*Simple0) F32NullValue() float32 {
	return float32(math.NaN())
}

func (*Simple0) D64Id() uint16 {
	return 110
}

func (*Simple0) D64SinceVersion() uint16 {
	return 0
}

func (s *Simple0) D64InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.D64SinceVersion()
}

func (*Simple0) D64Deprecated() uint16 {
	return 0
}

func (*Simple0) D64MetaAttribute(meta int) string {
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

func (*Simple0) D64MinValue() float64 {
	return -math.MaxFloat64
}

func (*Simple0) D64MaxValue() float64 {
	return math.MaxFloat64
}

func (*Simple0) D64NullValue() float64 {
	return math.NaN()
}

func (*Simple0) String6ASCIIId() uint16 {
	return 120
}

func (*Simple0) String6ASCIISinceVersion() uint16 {
	return 0
}

func (s *Simple0) String6ASCIIInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.String6ASCIISinceVersion()
}

func (*Simple0) String6ASCIIDeprecated() uint16 {
	return 0
}

func (*Simple0) String6ASCIIMetaAttribute(meta int) string {
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

func (*Simple0) String6ASCIIMinValue() byte {
	return byte(32)
}

func (*Simple0) String6ASCIIMaxValue() byte {
	return byte(126)
}

func (*Simple0) String6ASCIINullValue() byte {
	return 0
}

func (s *Simple0) String6ASCIICharacterEncoding() string {
	return "ASCII"
}

func (*Simple0) String1ASCIIId() uint16 {
	return 130
}

func (*Simple0) String1ASCIISinceVersion() uint16 {
	return 0
}

func (s *Simple0) String1ASCIIInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.String1ASCIISinceVersion()
}

func (*Simple0) String1ASCIIDeprecated() uint16 {
	return 0
}

func (*Simple0) String1ASCIIMetaAttribute(meta int) string {
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

func (*Simple0) String1ASCIIMinValue() byte {
	return byte(32)
}

func (*Simple0) String1ASCIIMaxValue() byte {
	return byte(126)
}

func (*Simple0) String1ASCIINullValue() byte {
	return 0
}

func (*Simple0) Int2Id() uint16 {
	return 140
}

func (*Simple0) Int2SinceVersion() uint16 {
	return 0
}

func (s *Simple0) Int2InActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.Int2SinceVersion()
}

func (*Simple0) Int2Deprecated() uint16 {
	return 0
}

func (*Simple0) Int2MetaAttribute(meta int) string {
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

func (*Simple0) Int2MinValue() int32 {
	return math.MinInt32 + 1
}

func (*Simple0) Int2MaxValue() int32 {
	return math.MaxInt32
}

func (*Simple0) Int2NullValue() int32 {
	return math.MinInt32
}
