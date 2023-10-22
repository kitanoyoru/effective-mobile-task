package events

const (
	PersonUpdatedEventType = EventType("Person was updated")
	PersonDeletedEventType = EventType("Person was deleted")
)

type EventType string

type PersonUpdatedEvent struct {
	Type    EventType
	Payload PersonUpdatedEventPayload
}

type PersonUpdatedEventPayload struct {
	ID string
}

type PersonDeletedEvent struct {
	Type    EventType
	Payload PersonDeletedEventPayload
}

type PersonDeletedEventPayload struct {
	ID string
}
