package bids

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tzAvito/internal/repository"
)

func (h *Implementation) GetUserBid(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	log.Printf("Fetching bids with username: %s, limit: %d, offset: %d", username, limit, offset)

	bids, err := h.bidService.GetUserBid(limit, offset, username)
	if err != nil {
		log.Printf("Error fetching bids: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(bids) == 0 {
		log.Printf("No bids found for username: %s", username)
		http.Error(w, "No bids found", http.StatusNotFound)
		return
	}
	bidResponse := repository.FormatBidResponses(bids)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bidResponse)
}
