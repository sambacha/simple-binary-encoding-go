// Generated SBE (Simple Binary Encoding) message codec

package mktdata

import (
	"fmt"
	"io"
	"reflect"
)

type HaltReasonEnum uint8
type HaltReasonValues struct {
	GroupSchedule            HaltReasonEnum
	SurveillanceIntervention HaltReasonEnum
	MarketEvent              HaltReasonEnum
	InstrumentActivation     HaltReasonEnum
	InstrumentExpiration     HaltReasonEnum
	Unknown                  HaltReasonEnum
	RecoveryInProcess        HaltReasonEnum
	NullValue                HaltReasonEnum
}

var HaltReason = HaltReasonValues{0, 1, 2, 3, 4, 5, 6, 255}

func (h HaltReasonEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(h)); err != nil {
		return err
	}
	return nil
}

func (h *HaltReasonEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(h)); err != nil {
		return err
	}
	return nil
}

func (h HaltReasonEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(HaltReason)
	for idx := 0; idx < value.NumField(); idx++ {
		if h == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on HaltReason, unknown enumeration value %d", h)
}

func (*HaltReasonEnum) EncodedLength() int64 {
	return 1
}

func (*HaltReasonEnum) GroupScheduleSinceVersion() uint16 {
	return 0
}

func (h *HaltReasonEnum) GroupScheduleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= h.GroupScheduleSinceVersion()
}

func (*HaltReasonEnum) GroupScheduleDeprecated() uint16 {
	return 0
}

func (*HaltReasonEnum) SurveillanceInterventionSinceVersion() uint16 {
	return 0
}

func (h *HaltReasonEnum) SurveillanceInterventionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= h.SurveillanceInterventionSinceVersion()
}

func (*HaltReasonEnum) SurveillanceInterventionDeprecated() uint16 {
	return 0
}

func (*HaltReasonEnum) MarketEventSinceVersion() uint16 {
	return 0
}

func (h *HaltReasonEnum) MarketEventInActingVersion(actingVersion uint16) bool {
	return actingVersion >= h.MarketEventSinceVersion()
}

func (*HaltReasonEnum) MarketEventDeprecated() uint16 {
	return 0
}

func (*HaltReasonEnum) InstrumentActivationSinceVersion() uint16 {
	return 0
}

func (h *HaltReasonEnum) InstrumentActivationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= h.InstrumentActivationSinceVersion()
}

func (*HaltReasonEnum) InstrumentActivationDeprecated() uint16 {
	return 0
}

func (*HaltReasonEnum) InstrumentExpirationSinceVersion() uint16 {
	return 0
}

func (h *HaltReasonEnum) InstrumentExpirationInActingVersion(actingVersion uint16) bool {
	return actingVersion >= h.InstrumentExpirationSinceVersion()
}

func (*HaltReasonEnum) InstrumentExpirationDeprecated() uint16 {
	return 0
}

func (*HaltReasonEnum) UnknownSinceVersion() uint16 {
	return 0
}

func (h *HaltReasonEnum) UnknownInActingVersion(actingVersion uint16) bool {
	return actingVersion >= h.UnknownSinceVersion()
}

func (*HaltReasonEnum) UnknownDeprecated() uint16 {
	return 0
}

func (*HaltReasonEnum) RecoveryInProcessSinceVersion() uint16 {
	return 3
}

func (h *HaltReasonEnum) RecoveryInProcessInActingVersion(actingVersion uint16) bool {
	return actingVersion >= h.RecoveryInProcessSinceVersion()
}

func (*HaltReasonEnum) RecoveryInProcessDeprecated() uint16 {
	return 0
}
