package repository

import (
	"fmt"
	"time"
	//"time"
	"tzAvito/internal/model"

	//"github.com/google/uuid"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (t *tenderRepository) SaveCurrentVersion(tender *model.Tender) error {
	version := &model.TenderVersion{
		TenderID:    tender.UUID,
		Version:     tender.Version,
		Name:        tender.Name,
		Description: tender.Description,
		ServiceType: tender.ServiceType,
		CreatedAt:   time.Now(),
	}

	if err := t.db.Create(version).Error; err != nil {
		return fmt.Errorf("failed to save tender version: %w", err)
	}
	return nil

}

func FormatTenderResponse(tender *model.Tender) model.TenderResponse {
	return model.TenderResponse{
		ID:          tender.UUID,
		Name:        tender.Name,
		Description: tender.Description,
		Status:      tender.Status,
		ServiceType: tender.ServiceType,
		Version:     tender.Version,
		//CreatedAt:     tender.CreatedAt.Time.Format(time.RFC3339),
	}
}

type TenderRepository interface {
	CreateTender(tender *model.Tender) error
	FindTenderByID(id string, username string) (*model.Tender, error)
	GetTenders(limit, offset int, serviceTypes []string) ([]*model.Tender, error)
	GetUserTender(limit, offset int, username string) ([]*model.Tender, error)
	UpdateTenderStatus(id string, status []string, username string) (*model.Tender, error)
	EditTender(tenderid string, updateData map[string]interface{}, username string) (*model.Tender, error)
	RollbackToVersion(tenderID string, version int, username string) (*model.Tender, error)
}

var tender model.Tender

type tenderRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) TenderRepository {
	return &tenderRepository{db: db}
}
func (t *tenderRepository) CreateTender(tender *model.Tender) error {
	if tender.UUID == "" {
		tender.UUID = uuid.New().String()
	}
	tender.Status = model.StatusTypeCREATED
	if err := t.db.Create(tender).Error; err != nil {
		return err
	}
	return nil
}
func (t *tenderRepository) FindTenderByID(id string, username string) (*model.Tender, error) {
	tender := &model.Tender{}
	query := t.db.Where("uuid = ?", id)
	if username != "" {
		query = query.Where("creator_username = ?", username)
	}

	err := query.First(tender).Error
	if err != nil {
		return nil, fmt.Errorf("error finding tender: %v", err)
	}
	return tender, nil
}
func (t *tenderRepository) GetTenders(limit, offset int, serviceTypes []string) ([]*model.Tender, error) {
	query := t.db.Table("tenders").Limit(limit).Offset(offset)
	if len(serviceTypes) > 0 {
		query = query.Where("service_type IN ?", serviceTypes)
	}
	tenders := []*model.Tender{}
	if err := query.Find(&tenders).Error; err != nil {
		return nil, err
	}
	return tenders, nil
}
func (t *tenderRepository) GetUserTender(limit, offset int, username string) ([]*model.Tender, error) {
	query := t.db.Table("tenders").Limit(limit).Offset(offset)
	if username != "" {
		query = query.Where("creator_username = ?", username)
	}
	tenders := []*model.Tender{}
	if err := query.Find(&tenders).Error; err != nil {
		return nil, err
	}
	return tenders, nil
}

func (t *tenderRepository) UpdateTenderStatus(id string, status []string, username string) (*model.Tender, error) {
	tender := &model.Tender{}
	query := t.db.Table("tenders").Where("uuid = ?", id)
	if username != "" {
		query = query.Where("creator_username = ?", username)
	}
	if err := query.First(tender).Error; err != nil {
		return nil, err
	}
	if len(status) > 0 {
		validStatuses := map[model.StatusType]bool{
			model.StatusTypeCREATED:   true,
			model.StatusTypePUBLISHED: true,
			model.StatusTypeCLOSED:    true,
		}

		newStatus := model.StatusType(status[0])
		if _, ok := validStatuses[newStatus]; !ok {
			return nil, fmt.Errorf("invalid status: %s", newStatus)
		}
		if err := t.db.Model(tender).Where("uuid = ?", id).Update("status", newStatus).Error; err != nil {
			return nil, err
		}

		tender.Status = newStatus

	}
	return tender, nil

}
func (t *tenderRepository) EditTender(tenderid string, updateData map[string]interface{}, username string) (*model.Tender, error) {
	tender := &model.Tender{}
	query := t.db.Table("tenders").Where("uuid = ?", tenderid)
	if username != "" {
		query = query.Where("creator_username = ?", username)
	}
	if err := query.First(tender).Error; err != nil {
		return nil, err
	}
	for key, value := range updateData {
		switch key {
		case "name":
			tender.Name, _ = value.(string)
		case "description":
			tender.Description, _ = value.(string)
		case "serviceType":
			tender.ServiceType, _ = value.(string)
		}
	}
	if err := t.SaveCurrentVersion(tender); err != nil {
		return nil, err
	}
	if len(updateData) > 0 {
		if err := t.db.Model(tender).Where("uuid = ?", tenderid).Updates(updateData).Error; err != nil {
			return nil, err
		}
		tender.Version += 1

	}

	return tender, nil

}

func (t *tenderRepository) RollbackToVersion(tenderID string, version int, usename string) (*model.Tender, error) {
	versionData := &model.TenderVersion{}
	if err := t.db.Table("tender_versions").Where("tender_id = ? AND version = ?", tenderID, version).First(versionData).Error; err != nil {
		return nil, err
	}

	tender := &model.Tender{
		UUID:        tenderID,
		Name:        versionData.Name,
		Description: versionData.Description,
		ServiceType: versionData.ServiceType,
		Version:     versionData.Version + 1, // Инкремент версии
	}

	if err := t.db.Model(&model.Tender{}).Where("uuid = ?", tenderID).Updates(tender).Error; err != nil {
		return nil, err
	}

	return tender, nil
}
