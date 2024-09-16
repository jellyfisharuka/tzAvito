package tender

import (
	"encoding/json"
	"net/http"

	//"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
)
func (h *Implementation) UpdateTenderStatus(w http.ResponseWriter, r *http.Request) {
	// Get URL and query parameters
	id := chi.URLParam(r, "tenderId")
	status := r.URL.Query()["status"] // QueryArray equivalent
	username := r.URL.Query().Get("username")

	// Call service to update tender status
	tenders, err := h.tenderService.UpdateTenderStatus(id, status, username)
	if err != nil {
		http.Error(w, "Failed to update tender status", http.StatusInternalServerError)
		return
	}

	// Respond with updated tenders
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tenders)
}
