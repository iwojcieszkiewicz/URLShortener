package handler

import (
	"encoding/json"
	"net/http"
	"url-shortener/store"

	"github.com/gorilla/mux"
)

type Handler struct {
	store *store.Store
}

func New(s *store.Store) *Handler {
	return &Handler{store: s}
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	var body struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "nieprawidłowy JSON", http.StatusBadRequest)
		return
	}

	code := h.store.Save(body.URL)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(code)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]

	url, err := h.store.Get(code)

	if err != nil {
		http.Error(w, "nie znaleziono", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
