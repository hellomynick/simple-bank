package event_sourcing

import (
	"simplebank/internal/common"

	pb "simplebank/internal/proto"

	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Aggregate struct {
	version int64
	changes []*pb.EventEnvelope
}

func (a *Aggregate) TrackChanges(aggregateID string, e proto.Message) error {
	payloadBytes, err := protojson.Marshal(e)
	if err != nil {
		return err
	}

	envelope := &pb.EventEnvelope{
		EventId:     uuid.New().String(),
		AggregateId: aggregateID,
		TypeName:    common.GetEventName(e),
		Payload:     payloadBytes,
	}

	a.changes = append(a.changes, envelope)
	a.version++

	return nil
}

func (a *Aggregate) GetChanges() []*pb.EventEnvelope {
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
