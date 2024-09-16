package tender

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Implementation) GetUserTender(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "5"
	}
	offset := r.URL.Query().Get("offset")
	if offset == "" {
		offset = "0"
	}
	username := r.URL.Query().Get("username")

	if username == "" {
		http.Error(w, "Пользователь не существует или некорректен", http.StatusUnauthorized)
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "Неверный формат запроса или его параметры.", http.StatusBadRequest)
		return
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		http.Error(w, "Неверный формат запроса или его параметры.", http.StatusBadRequest)
		return
	}

	tenders, err := h.tenderService.GetUserTender(limitInt, offsetInt, username)
	if err != nil {
		http.Error(w, "Сервер не готов обрабатывать запросы", http.StatusInternalServerError)
		return
	}

	if len(tenders) == 0 {
		http.Error(w, "Пользователь не найден или у пользователя нет тендеров", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tenders)
}