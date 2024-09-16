package bids

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tzAvito/internal/repository"

	"github.com/go-chi/chi"
)

func (h *Implementation) TenderList(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "tenderId")
	username := r.URL.Query().Get("username")
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "Invalid limit", http.StatusBadRequest)
		return
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		http.Error(w, "Invalid offset", http.StatusBadRequest)
		return
	}

	bid, err := h.bidService.TenderList(id, username, limitInt, offsetInt)
	if err != nil {
		http.Error(w, "bid not found", http.StatusNotFound)
		return
	}
	bidResponse := repository.FormatBidResponses(bid)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bidResponse)
}