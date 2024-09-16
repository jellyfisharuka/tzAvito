package service

import (
	"tzAvito/internal/model"
	"tzAvito/internal/repository"
	//"tzAvito/internal/repository"
)

type TenderService interface {
	CreateTender(tender *model.Tender) error
	FindTenderByID(id string, username string) (*model.Tender, error)
	GetTenders(limit, offset int, serviceTypes []string) ([]*model.Tender, error)
	GetUserTender(limit, offset int, username string) ([]*model.Tender, error)
	UpdateTenderStatus(id string, status []string, username string) (*model.Tender, error)
	EditTender(tenderid string, updateData map[string]interface{}, username string) (*model.Tender, error)
	RollbackToVersion(tenderID string, version int, username string) (*model.Tender, error)
}
type tenderService struct {
	tenderRepository repository.TenderRepository
}

func NewTenderService(tenderRepo repository.TenderRepository) TenderService {
	return &tenderService{tenderRepository: tenderRepo}
}
func (t *tenderService) CreateTender(tender *model.Tender) error {
	return t.tenderRepository.CreateTender(tender)
}
func (t *tenderService) FindTenderByID(id string, username string) (*model.Tender, error) {
	return t.tenderRepository.FindTenderByID(id, username)
}
func (t *tenderService) GetTenders(limit, offset int, serviceTypes []string) ([]*model.Tender, error) {
	if limit <= 0 {
		limit = 5 // Значение по умолчанию для limit
	}
	if offset < 0 {
		offset = 0 // Значение по умолчанию для offset
	}
	return t.tenderRepository.GetTenders(limit, offset, serviceTypes)

}
func (t *tenderService) GetUserTender(limit, offset int, username string) ([]*model.Tender, error) {
	if limit <= 0 {
		limit = 5 // Значение по умолчанию для limit
	}
	if offset < 0 {
		offset = 0 // Значение по умолчанию для offset
	}
	return t.tenderRepository.GetUserTender(limit, offset, username)
}

func (t *tenderService) UpdateTenderStatus(id string, status []string, username string) (*model.Tender, error) {
	return t.tenderRepository.UpdateTenderStatus(id, status, username)
}

func (t *tenderService) EditTender(tenderid string, updateData map[string]interface{}, username string) (*model.Tender, error) {
	return t.tenderRepository.EditTender(tenderid, updateData, username)
}
func (t *tenderService) RollbackToVersion(tenderID string, version int, username string) (*model.Tender, error) {
	return t.tenderRepository.RollbackToVersion(tenderID, version, username)
}
