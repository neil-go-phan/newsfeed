package entities

type FcmNotification struct {
	Username      string `json:"username" gorm:"foreignKey:Username;constraint:OnDelete:CASCADE"`
	FirebaseToken string   `json:"firebase_token"`
}

type NotificationMessage struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
