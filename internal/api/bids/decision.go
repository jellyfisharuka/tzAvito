package bids

import (
	"encoding/json"
	"net/http"
	"tzAvito/internal/model"

	"github.com/go-chi/chi"
)

type SubmitDecisionRequest struct {
	Decision []string `json:"decision"`
}
type SubmitDecisionResponse struct {
	Bid *model.Bid `json:"bid"`
}

func (h *Implementation) SubmitDecision(w http.ResponseWriter, r *http.Request) {
	bidId := chi.URLParam(r, "bidId")
	username := r.URL.Query().Get("username")
	decision := r.URL.Query().Get("decision")

	if decision == "" {
		http.Error(w, "Decision is required", http.StatusBadRequest)
		return
	}
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	bid, err := h.bidService.SubmitDecision(bidId, []string{decision}, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id":         bid.UUID,
		"name":       bid.Name,
		"status":     bid.Status,
		"authorType": bid.AuthorType,
		"authorId":   bid.AuthorID,
		"version":    bid.Version,
		"createdAt":  bid.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}