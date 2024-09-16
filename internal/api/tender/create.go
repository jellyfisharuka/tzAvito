package tender

import (
	"encoding/json"
	"net/http"
	"tzAvito/internal/model"
	"tzAvito/internal/repository"
)

func getCurrentUserFromContext(r *http.Request) (*model.User, bool) {
	user := r.Context().Value("currentUser")
	if user == nil {
		return nil, false
	}
	currentUser, ok := user.(*model.User)
	return currentUser, ok
}

func (h *Implementation) CreateTender(w http.ResponseWriter, r *http.Request) {
	var tender model.Tender
	if err := json.NewDecoder(r.Body).Decode(&tender); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if err := h.tenderService.CreateTender(&tender); err != nil {
		switch err.Error() {
		case "Пользователь не существует или некорректен":
			http.Error(w, err.Error(), http.StatusUnauthorized)
		case "Недостаточно прав для выполнения действия":
			http.Error(w, err.Error(), http.StatusForbidden)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	response := repository.FormatTenderResponse(&tender)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}