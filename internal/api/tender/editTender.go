package tender

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi"
)

func (h *Implementation) EditTenderHandler(w http.ResponseWriter, r *http.Request) {
	tenderID := chi.URLParam(r, "tenderId") 
	username := r.URL.Query().Get("username") 
	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	tender, err := h.tenderService.EditTender(tenderID, updateData, username)
	if err != nil {
		http.Error(w, "Failed to update tender", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tender)
}
