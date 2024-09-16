package tender

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
)

func (h *Implementation) FindTenderByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "tenderId")
	username := r.URL.Query().Get("username")
	tender, err := h.tenderService.FindTenderByID(id, username)
	if err != nil {
		fmt.Printf("Error finding tender: %v\n", err)
		http.Error(w, "Tender not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tender.Status)
}