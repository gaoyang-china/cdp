// Code generated by cdpgen. DO NOT EDIT.

package backgroundservice

import (
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/serviceworker"
)

// ServiceName The Background Service that will be associated with the
// commands/events. Every Background Service operates independently, but they
// share the same API.
type ServiceName string

// ServiceName as enums.
const (
	ServiceNameNotSet          ServiceName = ""
	ServiceNameBackgroundFetch ServiceName = "backgroundFetch"
	ServiceNameBackgroundSync  ServiceName = "backgroundSync"
	ServiceNamePushMessaging   ServiceName = "pushMessaging"
	ServiceNameNotifications   ServiceName = "notifications"
)

func (e ServiceName) Valid() bool {
	switch e {
	case "backgroundFetch", "backgroundSync", "pushMessaging", "notifications":
		return true
	default:
		return false
	}
}

func (e ServiceName) String() string {
	return string(e)
}

// EventMetadata A key-value pair for additional event information to pass
// along.
type EventMetadata struct {
	Key   string `json:"key"`   // No description.
	Value string `json:"value"` // No description.
}

// Event
type Event struct {
	Timestamp                   network.TimeSinceEpoch       `json:"timestamp"`                   // Timestamp of the event (in seconds).
	Origin                      string                       `json:"origin"`                      // The origin this event belongs to.
	ServiceWorkerRegistrationID serviceworker.RegistrationID `json:"serviceWorkerRegistrationId"` // The Service Worker ID that initiated the event.
	Service                     ServiceName                  `json:"service"`                     // The Background Service this event belongs to.
	EventName                   string                       `json:"eventName"`                   // A description of the event.
	InstanceID                  string                       `json:"instanceId"`                  // An identifier that groups related events together.
	EventMetadata               []EventMetadata              `json:"eventMetadata"`               // A list of event-specific information.
}
