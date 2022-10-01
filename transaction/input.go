package transaction

import "crowdfunding/user"

type GetTransactionCampaignInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
