package events

import "github.com/kitanoyoru/effective-mobile-task/internal/models"

const (
	PersonGetEventTopic     = "person:get"
	PersonDeletedEventTopic = "person:delete"
	PersonUpdatedEventTopic = "person:update"
	PersonPostEventTopic    = "person:post"
)

const (
	PersonUpdatedEventType = EventType("Person was updated")
	PersonDeletedEventType = EventType("Person was deleted")
	PersonGetEventType     = EventType("Get person request processed")
	PersonPostEventType    = EventType("Post person request processed")
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

type PersonGetEvent struct {
	Type    EventType
	Payload PersonGetEventPayload
}

type PersonGetEventPayload struct {
	Person models.Person
}

type PersonPostEvent struct {
	Type    EventType
	Payload PersonPostEventPayload
}

type PersonPostEventPayload struct {
	ID   int
	Name string
}
