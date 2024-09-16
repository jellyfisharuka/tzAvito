package service

import (
	//"fmt"
	"tzAvito/internal/model"
	"tzAvito/internal/repository"
	//"tzAvito/internal/repository"
)

type BidService interface {
	CreateBid(bid *model.Bid) error
	GetUserBid(limit, offset int, username string) ([]*model.Bid, error)
	TenderList(tenderId, username string, limit, offset int) ([]*model.Bid, error)
	UpdateBidStatus(id string, status []string, username string) (*model.Bid, error)
	GetBidStatus(id, username string) (*model.Bid, error)
	EditBid(bidId, username, name, description string) (*model.Bid, error)
	SubmitDecision(bidId string, decision []string, username string) (*model.Bid, error)
	//GetBidStatus(bidId, username string)()
	//PutFeedBack (bidId string, feedback string, username string)(*model.Bid, error)

}
type bidService struct {
	bidRepository repository.BidRepository
}

func NewBidService(bidRepo repository.BidRepository) BidService {
	return &bidService{bidRepository: bidRepo}
}
func (t *bidService) CreateBid(bid *model.Bid) error {

	return t.bidRepository.CreateBid(bid)
}
func (t *bidService) GetUserBid(limit, offset int, username string) ([]*model.Bid, error) {
	if limit <= 0 {
		limit = 5 // Значение по умолчанию для limit
	}
	if offset < 0 {
		offset = 0 // Значение по умолчанию для offset
	}
	return t.bidRepository.GetUserBid(limit, offset, username)
}
func (t *bidService) TenderList(tenderId, username string, limit, offset int) ([]*model.Bid, error) {
	if limit <= 0 {
		limit = 5 // Значение по умолчанию для limit
	}
	if offset < 0 {
		offset = 0 // Значение по умолчанию для offset
	}
	return t.bidRepository.TenderList(tenderId, username, limit, offset)

}
func (t *bidService) UpdateBidStatus(id string, status []string, username string) (*model.Bid, error) {
	return t.bidRepository.UpdateBidStatus(id, status, username)
}

func (t *bidService) GetBidStatus(id, username string) (*model.Bid, error) {
	return t.bidRepository.GetBidStatus(id, username)
}
func (t *bidService) EditBid(bidId, username, name, description string) (*model.Bid, error) {
	return t.bidRepository.EditBid(bidId, username, name, description)
}
func (t *bidService) SubmitDecision(bidId string, decision []string, username string) (*model.Bid, error) {
	return t.bidRepository.SubmitDecision(bidId, decision, username)
}

//func(t *bidService) PutFeedBack (bidId string, feedback string, username string)(*model.Bid, error) {
// return t.bidRepository.SubmitDecision(bidId, feedback, username)
//}
