package handlers

import (
	"encoding/json"
	"go-rest-crud/pkg/mocks"
	"go-rest-crud/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()

	// Read updated data
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	//
	var updatedBook models.Book
	// request ma aako updated data lai updatedBook ma bind gardiyo
	json.Unmarshal(body, &updatedBook)

	/*
		   - working with database
			var book models.Book
			// check if book with requested id exist or not 
			if result := h.DB.First(&book, id); result.Error != nil {
				fmt.Println(result.Error)
			}

	*/
	// Iterate over all the mock Books

	for index, book := range mocks.Books {

		if book.Id == id {
			// Update and send response when book Id matches dynamic Id
			book.Title = updatedBook.Title
			book.Author = updatedBook.Author
			book.Desc = updatedBook.Desc

			mocks.Books[index] = book

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
