package bids

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
)

func (h *Implementation) GetBidStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "bidId")
	username := r.URL.Query().Get("username")

	tender, err := h.bidService.GetBidStatus(id, username)
	if err != nil {
		fmt.Printf("Error finding bid: %v\n", err)
		http.Error(w, "bid not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tender.Status)
}
