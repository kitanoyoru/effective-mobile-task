package events

import (
	"context"

	evbus "github.com/asaskevich/EventBus"
	log "github.com/sirupsen/logrus"
)

type EventBusSession struct {
	bus evbus.Bus
}

func NewEventBusSession() *EventBusSession {
	bus := evbus.New()

	return &EventBusSession{
		bus,
	}
}

func (s *EventBusSession) PublishEvent(ctx context.Context, topic string, event any) {
	select {
	case <-ctx.Done():
		log.Debug("PublishEvent timeout")
		return
	default:
		s.bus.Publish(topic, event)
	}
}

func (s *EventBusSession) AsyncConsumeEvents(ctx context.Context, topic string, handler any) {
	s.bus.SubscribeAsync(topic, handler, false)

	for {
		select {
		case <-ctx.Done():
			log.Debugf("Ending consuming topic: %s", topic)
			s.bus.WaitAsync()
			s.bus.Unsubscribe(topic, handler)
			return
		}
	}
}
