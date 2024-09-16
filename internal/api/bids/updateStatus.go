package bids

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	//"github.com/gin-gonic/gin"
)

func (h *Implementation) UpdateBidStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "bidId")
	status := r.URL.Query()["status"]
	username := r.URL.Query().Get("username")

	bids, err := h.bidService.UpdateBidStatus(id, status, username)
	if err != nil {
		http.Error(w, "Failed to get bids", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bids)
}
