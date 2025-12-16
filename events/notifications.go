package events

const NOTIFICATIONS_EVENTS = "notifications.events"

const NOTIFICATIONS_EMAIL = "notifications.email"

const NOTIFICATIONS_TELEGRAM = "notifications.telegram"

type EmailVerifyNotificationPayload struct {
	Email       string
	Title       string
	Description string
	Link        string
	Locale      string
}
