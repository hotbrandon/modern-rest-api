package main

import (
	"encoding/json"
	"net/http"
)

type ShoppingList struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

var shoppingLists []ShoppingList

func getShoppingList(w http.ResponseWriter, _ *http.Request) {
	if shoppingLists == nil {
		shoppingLists = make([]ShoppingList, 0)
	}
	jsonData, err := json.Marshal(shoppingLists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func createShoppingList(w http.ResponseWriter, r *http.Request) {
	list := ShoppingList{}
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	list.ID = len(shoppingLists) + 1
	shoppingLists = append(shoppingLists, list)

	jsonData, err := json.Marshal(shoppingLists)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func handleList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getShoppingList(w, r)
	case http.MethodPost:
		createShoppingList(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
