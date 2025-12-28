package events

import "time"

const USER_EVENTS = "user.events"

const USER_CREATED = "user.created"

const USER_DELETED = "user.deleted"

const USER_PROFILE_UPDATED = "user.profile.updated"

type UserCreatedEventPayload struct {
	ID                       string    `json:"id"`
	Email                    string    `json:"email"`
	EmailVerifyLink          string    `json:"email_verify_link"`
	EmailVerifyLinkExpiresAt time.Time `json:"email_verify_link_expires_at"`
	Locale                   string    `json:"locale"`
}

type UserProfileUpdatedEventPayload struct {
	NewLastName   *string `json:"new_last_name,omitempty"`
	NewFirstName  *string `json:"new_first_name,omitempty"`
	NewMiddleName *string `json:"new_middle_name,omitempty"`
	NewAvatarURL  *string `json:"new_avatar_url,omitempty"`
	ID            string  `json:"id"`
	Email         string  `json:"email"`
	Locale        string  `json:"locale"`
}

type UserDeletedEventPayload struct {
	ID string `json:"id"`
}
