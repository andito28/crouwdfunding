package transaction

import (
	"crowdfunding/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	User       user.User
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
