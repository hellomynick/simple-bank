package event_sourcing

type Aggregate struct {
	version int64
	changes []Event
}

func (a *Aggregate) TrackChanges(e Event) {
	a.changes = append(a.changes, e)
	a.version++
}

func (a *Aggregate) GetChanges() []Event {
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
