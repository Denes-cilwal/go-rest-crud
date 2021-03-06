package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest-crud/pkg/mocks"
	"go-rest-crud/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	fmt.Println(vars, "is---")
	id, _ := strconv.Atoi(vars["id"])

	fmt.Println("id is -====")
	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Iterate over all the mock books
	for _, book := range mocks.Books {
		// if id match that in dynamic id parameter
		if book.Id == id {
			// If ids are equal send book as response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
