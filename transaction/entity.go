package transaction

import (
	"crowdfunding/campaign"
	"crowdfunding/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	PaymentURL string
	User       user.User
	Campaign   campaign.Campaign
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
