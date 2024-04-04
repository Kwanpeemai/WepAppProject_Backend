package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web-project/models"
)

var sizes []models.Size

func CreateSize(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var size models.Size
	_ = json.NewDecoder(r.Body).Decode(&size)
	sizes = append(sizes, size)
	json.NewEncoder(w).Encode(size)
}


func GetSizes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(sizes)
}

func GetSize(w http.ResponseWriter, r *http.Request) {
	sizeID := r.URL.Query().Get("size_id")
	for _, size := range sizes {
		if fmt.Sprint(size.Size_ID) == sizeID {
			json.NewEncoder(w).Encode(size)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Size{})
}

func UpdateSize(w http.ResponseWriter, r *http.Request) {
	var updatedSize models.Size
	json.NewDecoder(r.Body).Decode(&updatedSize)
	for i, size := range sizes {
		if size.Size_ID == updatedSize.Size_ID {
			sizes[i] = updatedSize
			json.NewEncoder(w).Encode(updatedSize)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Size{})
}

func DeleteSize(w http.ResponseWriter, r *http.Request) {
	sizeID := r.URL.Query().Get("size_id")
	for i, size := range sizes {
		if fmt.Sprint(size.Size_ID) == sizeID {
			sizes = append(sizes[:i], sizes[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(sizes)
}
