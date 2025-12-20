package events

const NOTIFICATIONS_EVENTS = "notifications.events"

const NOTIFICATIONS_EMAIL = "notifications.email"

const NOTIFICATIONS_TELEGRAM = "notifications.telegram"

type EmailVerifyNotificationPayload struct {
	Email       string `json:"email"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Locale      string `json:"locale"`
}

type TelegramNotificationPayload struct {
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}
