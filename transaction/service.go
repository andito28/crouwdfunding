package transaction

import (
	"crowdfunding/campaign"
	"errors"
)

type Service interface {
	GetTransactionByCampaignID(input GetTransactionCampaignInput) ([]Transaction, error)
	GetTransactionByUserID(userID int) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignID(input GetTransactionCampaignInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}

	transaction, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) GetTransactionByUserID(userID int) ([]Transaction, error) {
	newTransaction, err := s.repository.GetUserByID(userID)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}
