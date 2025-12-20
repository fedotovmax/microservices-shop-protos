package events

import "time"

const USER_EVENTS = "user.events"

const USER_CREATED = "user.created"

const USER_DELETED = "user.deleted"

const USER_PROFILE_UPDATED = "user.profile.updated"

type UserCreatedEventPayload struct {
	EmailVerifyLinkValidityPeriod time.Time `json:"email_verify_link_validity_period"`
	ID                            string    `json:"id"`
	Email                         string    `json:"email"`
	EmailVerifyLink               string    `json:"email_verify_link"`
	Locale                        string    `json:"locale"`
}

type UserUpdatedEventPayload struct {
	ID     string `json:"id"`
	Locale string `json:"locale"`
}
