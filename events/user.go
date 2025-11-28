package events

const USER_EVENTS = "user.events"

const USER_CREATED = "user.created"

const USER_DELETED = "user.deleted"

const USER_UPDATED = "user.updated"

type UserCreatedEventPayload struct {
	ID string `json:"id"`
}
