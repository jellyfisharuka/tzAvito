package tender

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	//"github.com/gin-gonic/gin"
)

func (h *Implementation) RollbackToVersion(w http.ResponseWriter, r *http.Request) {
	tenderId := chi.URLParam(r, "tenderId")
	version := chi.URLParam(r, "version")
	username := r.URL.Query().Get("username")

	if tenderId == "" || version == "" {
		http.Error(w, "tenderId and version are required", http.StatusBadRequest)
		return
	}
	versionInt, err := strconv.Atoi(version)
	if err != nil {
		http.Error(w, "Invalid version", http.StatusBadRequest)
		return
	}
	tender, err := h.tenderService.RollbackToVersion(tenderId, versionInt, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tender)
}
