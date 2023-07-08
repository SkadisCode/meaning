package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/SkadisCode/meaning/entity"
	"github.com/SkadisCode/meaning/usecase"
)

type DictionaryHandler struct {
	dictionaryUsecase *usecase.DictionaryUsecase
}

func NewDictionaryHandler(dictionaryUsecase *usecase.DictionaryUsecase) *DictionaryHandler {
	return &DictionaryHandler{dictionaryUsecase: dictionaryUsecase}
}

func (h *DictionaryHandler) SaveDictionary(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the dictionary entry
	dictionary := &entity.Dictionary{}
	err := json.NewDecoder(r.Body).Decode(dictionary)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the use case to save the dictionary entry
	err = h.dictionaryUsecase.SaveDictionary(dictionary)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *DictionaryHandler) GetDictionary(w http.ResponseWriter, r *http.Request) {
	// Get the word parameter from the URL
	word := r.URL.Query().Get("word")

	// Call the use case to get the dictionary entry
	dictionary, err := h.dictionaryUsecase.GetDictionary(word)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the dictionary entry as JSON
	json.NewEncoder(w).Encode(dictionary)
}
