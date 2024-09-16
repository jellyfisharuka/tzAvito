package bids

import (
	"encoding/json"
	"net/http"
	"tzAvito/internal/model"
	"tzAvito/internal/repository"
	"tzAvito/internal/service"

	//"github.com/gin-gonic/gin"
)

type Implementation struct {
	bidService service.BidService
}

func NewImplementation(bidService service.BidService) *Implementation {
	return &Implementation{
		bidService: bidService,
	}
}

func (h *Implementation) CreateBid(w http.ResponseWriter, r *http.Request) {
    var bid model.Bid
    if err := json.NewDecoder(r.Body).Decode(&bid); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    if err := h.bidService.CreateBid(&bid); err != nil {
        http.Error(w, "Тендер не найден", http.StatusInternalServerError)
        return
    }
    bidResponse := repository.FormatBidResponse(&bid)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(bidResponse); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}
