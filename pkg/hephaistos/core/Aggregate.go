package core

import (
	"simplebank/internal/common"
	"simplebank/pkg/hephaistos/core/event_sourcing"

	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Aggregate struct {
	version int64
	changes []*event_sourcing.EventEnvelope
}

func (a *Aggregate) TrackChange(aggregateID string, e proto.Message) error {
	payloadBytes, err := protojson.Marshal(e)
	if err != nil {
		return err
	}

	nextVersion := a.version + 1

	envelope := &event_sourcing.EventEnvelope{
		Id:               uuid.New().String(),
		AggregateId:      aggregateID,
		AggregateVersion: nextVersion,
		TypeName:         common.GetEventName(e),
		Payload:          payloadBytes,
	}

	a.changes = append(a.changes, envelope)
	a.version = nextVersion

	return nil
}

func (a *Aggregate) GetChanges() []*event_sourcing.EventEnvelope {
	return a.changes
}

func (a *Aggregate) GetVersion() int64 {
	return a.version
}

func (a *Aggregate) ClearChanges() {
	a.changes = nil
}

func (a *Aggregate) IncreaseVersion() {
	a.version++
}
