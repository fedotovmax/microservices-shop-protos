package events

import "time"

const SESSION_EVENTS = "session.events"

const SESSION_BLACKLIST_ADDED = "session.blacklist.added"

const SESSION_BYPASS_ADDED = "session.bypass.added"

type SessionBlacklistAddedEventPayload struct {
	UID           string    `json:"uid"`
	Email         string    `json:"email"`
	Code          string    `json:"code"`
	CodeExpiresAt time.Time `json:"code_expires_at"`
}

type SessionBypassAddedEventPayload struct {
	UID             string    `json:"uid"`
	Email           string    `json:"email"`
	Code            string    `json:"code"`
	BypassExpiresAt time.Time `json:"code_expires_at"`
}
