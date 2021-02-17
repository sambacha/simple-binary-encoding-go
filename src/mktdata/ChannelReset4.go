// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ChannelReset4 struct {
	TransactTime        uint64
	MatchEventIndicator MatchEventIndicator
	NoMDEntries         []ChannelReset4NoMDEntries
}
type ChannelReset4NoMDEntries struct {
	MDUpdateAction int8
	MDEntryType    [1]byte
	ApplID         int16
}

func (c *ChannelReset4) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := c.RangeCheck(c.SbeSchemaVersion(), c.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, c.TransactTime); err != nil {
		return err
	}
	if err := c.MatchEventIndicator.Encode(_m, _w); err != nil {
		return err
	}
	var NoMDEntriesBlockLength uint16 = 2
	if err := _m.WriteUint16(_w, NoMDEntriesBlockLength); err != nil {
		return err
	}
	var NoMDEntriesNumInGroup uint8 = uint8(len(c.NoMDEntries))
	if err := _m.WriteUint8(_w, NoMDEntriesNumInGroup); err != nil {
		return err
	}
	for _, prop := range c.NoMDEntries {
		if err := prop.Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (c *ChannelReset4) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !c.TransactTimeInActingVersion(actingVersion) {
		c.TransactTime = c.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &c.TransactTime); err != nil {
			return err
		}
	}
	if c.MatchEventIndicatorInActingVersion(actingVersion) {
		if err := c.MatchEventIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > c.SbeSchemaVersion() && blockLength > c.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-c.SbeBlockLength()))
	}

	if c.NoMDEntriesInActingVersion(actingVersion) {
		var NoMDEntriesBlockLength uint16
		if err := _m.ReadUint16(_r, &NoMDEntriesBlockLength); err != nil {
			return err
		}
		var NoMDEntriesNumInGroup uint8
		if err := _m.ReadUint8(_r, &NoMDEntriesNumInGroup); err != nil {
			return err
		}
		if cap(c.NoMDEntries) < int(NoMDEntriesNumInGroup) {
			c.NoMDEntries = make([]ChannelReset4NoMDEntries, NoMDEntriesNumInGroup)
		}
		c.NoMDEntries = c.NoMDEntries[:NoMDEntriesNumInGroup]
		for i, _ := range c.NoMDEntries {
			if err := c.NoMDEntries[i].Decode(_m, _r, actingVersion, uint(NoMDEntriesBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := c.RangeCheck(actingVersion, c.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (c *ChannelReset4) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if c.TransactTimeInActingVersion(actingVersion) {
		if c.TransactTime < c.TransactTimeMinValue() || c.TransactTime > c.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on c.TransactTime (%v < %v > %v)", c.TransactTimeMinValue(), c.TransactTime, c.TransactTimeMaxValue())
		}
	}
	for _, prop := range c.NoMDEntries {
		if err := prop.RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func ChannelReset4Init(c *ChannelReset4) {
	return
}

func (c *ChannelReset4NoMDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt16(_w, c.ApplID); err != nil {
		return err
	}
	return nil
}

func (c *ChannelReset4NoMDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	c.MDUpdateAction = 0
	c.MDEntryType[0] = 74
	if !c.ApplIDInActingVersion(actingVersion) {
		c.ApplID = c.ApplIDNullValue()
	} else {
		if err := _m.ReadInt16(_r, &c.ApplID); err != nil {
			return err
		}
	}
	if actingVersion > c.SbeSchemaVersion() && blockLength > c.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-c.SbeBlockLength()))
	}
	return nil
}

func (c *ChannelReset4NoMDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if c.ApplIDInActingVersion(actingVersion) {
		if c.ApplID < c.ApplIDMinValue() || c.ApplID > c.ApplIDMaxValue() {
			return fmt.Errorf("Range check failed on c.ApplID (%v < %v > %v)", c.ApplIDMinValue(), c.ApplID, c.ApplIDMaxValue())
		}
	}
	return nil
}

func ChannelReset4NoMDEntriesInit(c *ChannelReset4NoMDEntries) {
	c.MDUpdateAction = 0
	c.MDEntryType[0] = 74
	return
}

func (*ChannelReset4) SbeBlockLength() (blockLength uint16) {
	return 9
}

func (*ChannelReset4) SbeTemplateId() (templateId uint16) {
	return 4
}

func (*ChannelReset4) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*ChannelReset4) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}

func (*ChannelReset4) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*ChannelReset4) TransactTimeId() uint16 {
	return 60
}

func (*ChannelReset4) TransactTimeSinceVersion() uint16 {
	return 0
}

func (c *ChannelReset4) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.TransactTimeSinceVersion()
}

func (*ChannelReset4) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ChannelReset4) TransactTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "UTCTimestamp"
	case 4:
		return "required"
	}
	return ""
}

func (*ChannelReset4) TransactTimeMinValue() uint64 {
	return 0
}

func (*ChannelReset4) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ChannelReset4) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*ChannelReset4) MatchEventIndicatorId() uint16 {
	return 5799
}

func (*ChannelReset4) MatchEventIndicatorSinceVersion() uint16 {
	return 0
}

func (c *ChannelReset4) MatchEventIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.MatchEventIndicatorSinceVersion()
}

func (*ChannelReset4) MatchEventIndicatorDeprecated() uint16 {
	return 0
}

func (*ChannelReset4) MatchEventIndicatorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "MultipleCharValue"
	case 4:
		return "required"
	}
	return ""
}

func (*ChannelReset4NoMDEntries) MDUpdateActionId() uint16 {
	return 279
}

func (*ChannelReset4NoMDEntries) MDUpdateActionSinceVersion() uint16 {
	return 2
}

func (c *ChannelReset4NoMDEntries) MDUpdateActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.MDUpdateActionSinceVersion()
}

func (*ChannelReset4NoMDEntries) MDUpdateActionDeprecated() uint16 {
	return 0
}

func (*ChannelReset4NoMDEntries) MDUpdateActionMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "constant"
	}
	return ""
}

func (*ChannelReset4NoMDEntries) MDUpdateActionMinValue() int8 {
	return math.MinInt8 + 1
}

func (*ChannelReset4NoMDEntries) MDUpdateActionMaxValue() int8 {
	return math.MaxInt8
}

func (*ChannelReset4NoMDEntries) MDUpdateActionNullValue() int8 {
	return math.MinInt8
}

func (*ChannelReset4NoMDEntries) MDEntryTypeId() uint16 {
	return 269
}

func (*ChannelReset4NoMDEntries) MDEntryTypeSinceVersion() uint16 {
	return 0
}

func (c *ChannelReset4NoMDEntries) MDEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.MDEntryTypeSinceVersion()
}

func (*ChannelReset4NoMDEntries) MDEntryTypeDeprecated() uint16 {
	return 0
}

func (*ChannelReset4NoMDEntries) MDEntryTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "char"
	case 4:
		return "constant"
	}
	return ""
}

func (*ChannelReset4NoMDEntries) MDEntryTypeMinValue() byte {
	return byte(32)
}

func (*ChannelReset4NoMDEntries) MDEntryTypeMaxValue() byte {
	return byte(126)
}

func (*ChannelReset4NoMDEntries) MDEntryTypeNullValue() byte {
	return 0
}

func (*ChannelReset4NoMDEntries) ApplIDId() uint16 {
	return 1180
}

func (*ChannelReset4NoMDEntries) ApplIDSinceVersion() uint16 {
	return 3
}

func (c *ChannelReset4NoMDEntries) ApplIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ApplIDSinceVersion()
}

func (*ChannelReset4NoMDEntries) ApplIDDeprecated() uint16 {
	return 0
}

func (*ChannelReset4NoMDEntries) ApplIDMetaAttribute(meta int) string {
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

func (*ChannelReset4NoMDEntries) ApplIDMinValue() int16 {
	return math.MinInt16 + 1
}

func (*ChannelReset4NoMDEntries) ApplIDMaxValue() int16 {
	return math.MaxInt16
}

func (*ChannelReset4NoMDEntries) ApplIDNullValue() int16 {
	return math.MinInt16
}

func (*ChannelReset4) NoMDEntriesId() uint16 {
	return 268
}

func (*ChannelReset4) NoMDEntriesSinceVersion() uint16 {
	return 0
}

func (c *ChannelReset4) NoMDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.NoMDEntriesSinceVersion()
}

func (*ChannelReset4) NoMDEntriesDeprecated() uint16 {
	return 0
}

func (*ChannelReset4NoMDEntries) SbeBlockLength() (blockLength uint) {
	return 2
}

func (*ChannelReset4NoMDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 9
}
