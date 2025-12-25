package event_sourcing

type Event interface {
	TypeName() string
}
