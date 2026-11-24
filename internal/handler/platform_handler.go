package handler

import (
	"encoding/json"
	"iamcodingabit/go-track-game-objectives/internal/store"
	"log"
	"net/http"
	"strconv"
)

type PlatformHandler struct {
	store *store.PlatformStore
}

func NewPlatformHandler(s *store.Store) *PlatformHandler {
	return &PlatformHandler{
		store: s.Platform,
	}
}

func (ph *PlatformHandler) GetAllPlatformHandler(w http.ResponseWriter, r *http.Request) {
	platforms, err := ph.store.GetAllPlatforms(r.Context())
	if err != nil {
		log.Println("fetch games error:", err)
		http.Error(w, "Fail to retrieve game", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(platforms)
}

func (ph *PlatformHandler) AddNewPlatformHandler(w http.ResponseWriter, r *http.Request){
	var newPlatform store.Platform

	err := json.NewDecoder(r.Body).Decode(&newPlatform)
	if err != nil {
		log.Println("decode error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	res, err := ph.store.AddNewPlatform(r.Context(), newPlatform)
	if err != nil {
		log.Println("decode error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "platform successfully inserted",
		"id": strconv.Itoa(res.ID),
		"name": res.Name,
	})
}

func(ph *PlatformHandler) DeletePlatformHandler(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}

	deletedName, err := ph.store.Delete(r.Context(), id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "deletion failed",
			"error":  err.Error(),
		})	
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "deleted successfully",
		"title":   deletedName,
	})
}

func(ph *PlatformHandler) UpdatePlatformHandler(w http.ResponseWriter, r *http.Request){
	var updatedPlatform store.Platform
	err := json.NewDecoder(r.Body).Decode(&updatedPlatform)
	if err != nil {
		log.Println("decode error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}

	res, err := ph.store.Update(r.Context(), updatedPlatform, id)
	if err != nil {
		log.Println("decode error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "new game successfully added",
		"id": strconv.Itoa(res.ID),
		"name": res.Name,
	})
}