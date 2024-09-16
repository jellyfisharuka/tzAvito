package repository

import (
	"fmt"
	"log"
	"tzAvito/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FormatBidResponse(bid *model.Bid) *model.BidResponse {
	return &model.BidResponse{
		ID:         bid.UUID,
		Name:       bid.Name,
		Status:     bid.Status,
		AuthorType: bid.AuthorType,
		AuthorID:   bid.AuthorID,
		Version:    bid.Version + 1,
		//CreatedAt:     tender.CreatedAt.Time.Format(time.RFC3339),
	}
}
func FormatBidResponses(bids []*model.Bid) []*model.BidResponse {
	var responses []*model.BidResponse
	for _, bid := range bids {
		responses = append(responses, FormatBidResponse(bid))
	}
	return responses
}

type BidRepository interface {
	CreateBid(bid *model.Bid) error
	GetUserBid(limit, offset int, username string) ([]*model.Bid, error)
	TenderList(tenderId, username string, limit, offset int) ([]*model.Bid, error)
	UpdateBidStatus(id string, status []string, username string) (*model.Bid, error)
	GetBidStatus(id, username string) (*model.Bid, error)
	EditBid(bidId, username, name, description string) (*model.Bid, error)
	SubmitDecision(bidId string, decision []string, username string) (*model.Bid, error)
	//PutFeedBack (bidId string, feedback string, username string)(*model.Bid, error)
}

var bid model.Bid

type bidRepository struct {
	db *gorm.DB
}

func NewBidRepository(db *gorm.DB) BidRepository {
	return &bidRepository{db: db}
}
func (t *bidRepository) CreateBid(bid *model.Bid) error {
	if bid.UUID == "" {
		bid.UUID = uuid.New().String()
	}

	bid.Status = model.StatusTypeCREATED
	if err := t.db.Create(bid).Error; err != nil {
		return err
	}
	return nil
}

func (t *bidRepository) GetBidStatus(id, username string) (*model.Bid, error) {
	query :=
		`SELECT bids.status, tenders.creator_username
    FROM bids
    JOIN tenders ON bids.tender_id = tenders.uuid
    WHERE bids.uuid = ? AND tenders.creator_username = ? 
    `
	// //вот так запросы прямо в коде конечно же делать не стоит но это повышает риск sql injection вообще лучше бы использовать библиотеку sqlx но опять таки не успела и мне было вот так проще
	var result struct {
		Status          string
		CreatorUsername string
	}

	if err := t.db.Raw(query, id, username).Scan(&result).Error; err != nil {
		log.Printf("Error fetching bid status: %v", err)
		return nil, err
	}

	if username != result.CreatorUsername {
		return nil, fmt.Errorf("user %s is not authorized to view bid status", username)
	}

	bid := &model.Bid{
		UUID:   id,
		Status: model.StatusType(result.Status),
	}

	return bid, nil

}

func (t *bidRepository) UpdateBidStatus(id string, status []string, username string) (*model.Bid, error) {
	query :=
		`SELECT bids.status, tenders.creator_username
    FROM bids
    JOIN tenders ON bids.tender_id = tenders.uuid
    WHERE bids.uuid = ? AND tenders.creator_username = ?
    `
	var result struct {
		Status          string
		CreatorUsername string
	}

	if err := t.db.Raw(query, id, username).Scan(&result).Error; err != nil {
		log.Printf("Error fetching bid status: %v", err)
		return nil, err
	}
	if username != result.CreatorUsername {
		return nil, fmt.Errorf("user %s is not authorized to view bid status", username)
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
		var bid model.Bid
		if err := t.db.Model(&bid).Where("uuid = ?", id).Update("status", newStatus).Error; err != nil {
			return nil, err
		}
		bid.Status = newStatus

		if err := t.db.Preload("Tender").Preload("Creator").First(&bid, "uuid = ?", id).Error; err != nil {
			return nil, fmt.Errorf("failed to fetch updated bid: %v", err)
		}

		return &bid, nil
	}

	return nil, fmt.Errorf("no status provided")
}
func (t *bidRepository) GetUserBid(limit, offset int, username string) ([]*model.Bid, error) {
	query := `
		SELECT bids.*, tenders.creator_username
		FROM bids
		JOIN tenders ON bids.tender_id = tenders.uuid
		WHERE tenders.creator_username = ?
		LIMIT ? OFFSET ?
	`
	bids := []*model.Bid{}
	if err := t.db.Raw(query, username, limit, offset).Scan(&bids).Error; err != nil {
		log.Printf("Error fetching bids: %v", err)
		return nil, err
	}
	log.Printf("Bids fetched: %v", bids)
	return bids, nil

}
func (t *bidRepository) TenderList(tenderId, username string, limit, offset int) ([]*model.Bid, error) {
	var bids []*model.Bid

	query := `
		SELECT bids.*
		FROM bids
		JOIN tenders ON tenders.uuid = bids.tender_id
		WHERE bids.tender_id = ?
	`
	if username != "" {
		query += " AND tenders.creator_username = ?"
	}
	query += " LIMIT ? OFFSET ?"
	args := []interface{}{tenderId}
	if username != "" {
		args = append(args, username)
	}
	args = append(args, limit, offset)

	if err := t.db.Raw(query, args...).Scan(&bids).Error; err != nil {
		log.Printf("Error fetching bids: %v", err)
		return nil, err
	}

	log.Printf("Bids fetched: %v", bids)
	return bids, nil
}

func (t *bidRepository) EditBid(bidId, username, name, description string) (*model.Bid, error) {
	updateQuery := t.db.Table("bids").
		Joins("JOIN tenders ON tenders.uuid = bids.tender_id").
		Where("bids.uuid = ? AND tenders.creator_username = ?", bidId, username).
		Updates(map[string]interface{}{
			"bids.name":        gorm.Expr("COALESCE(NULLIF(?, ''), bids.name)", name),
			"bids.description": gorm.Expr("COALESCE(NULLIF(?, ''), bids.description)", description),
		})

	if updateQuery.Error != nil {
		return nil, fmt.Errorf("failed to update bid: %w", updateQuery.Error)
	}

	rowsAffected := updateQuery.RowsAffected
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no bid found with id %s and username %s", bidId, username)
	}
	updatedBid := &model.Bid{}
	if err := t.db.Where("bids.uuid = ?", bidId).Preload("Creator").Preload("Tender").First(updatedBid).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve updated bid: %w", err)
	}
	return updatedBid, nil
}

func (t *bidRepository) SubmitDecision(bidId string, decision []string, username string) (*model.Bid, error) {
	if len(decision) == 0 {
		return nil, fmt.Errorf("no decision provided")
	}
	var creatorUsername string
	checkQuery := `
        SELECT tenders.creator_username
        FROM bids
        JOIN tenders ON bids.tender_id = tenders.uuid
        WHERE bids.uuid = ?
    `
	if err := t.db.Raw(checkQuery, bidId).Scan(&creatorUsername).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch creator_username: %w", err)
	}

	if creatorUsername != username {
		return nil, fmt.Errorf("user %s is not authorized to make decisions for this bid", username)
	}
	validDecisions := map[model.StatusType]bool{
		model.StatusTypeAPPROVED: true,
		model.StatusTypeREJECTED: true,
	}

	newDecision := model.StatusType(decision[0])
	if _, ok := validDecisions[newDecision]; !ok {
		return nil, fmt.Errorf("invalid decision: %s", newDecision)
	}
	var bid model.Bid
	if err := t.db.Model(&model.Bid{}).
		Where("uuid = ?", bidId).
		Update("status", newDecision).Error; err != nil {
		return nil, fmt.Errorf("failed to update bid decision: %w", err)
	}
	if err := t.db.Where("uuid = ?", bidId).Preload("Tender").Preload("Creator").First(&bid).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve updated bid: %w", err)
	}

	return &bid, nil

}

//func (t *bidRepository) PutFeedBack (bidId string, feedback string, username string)(*model.Bid, error)
