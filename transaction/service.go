package transaction

import (
	"backendstartup/campaign"
	"errors"
)

type service struct {
	repository Repository
	//inject data campaign di service transaction
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTrasactionByUserID(userID int) ([]Transaction, error)
}

//Harus mendeklarasikan bahwa campaignRepositoy juga digunakan di transaction service
func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	//get campaign
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}
	//check campaign user_id != user id login
	//Hanya user pembuat campaign yang berhak melihat transaksi campaign
	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) GetTrasactionByUserID(userID int) ([]Transaction, error) {
	transaction, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
