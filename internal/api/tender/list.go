package tender

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Implementation) GetTenders(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "5"
	}
	offset := r.URL.Query().Get("offset")
	if offset == "" {
		offset = "0"
	}
	serviceType := r.URL.Query()["service_type"] 

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

	tenders, err := h.tenderService.GetTenders(limitInt, offsetInt, serviceType)
	if err != nil {
		http.Error(w, "Failed to get tenders", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tenders)
}
