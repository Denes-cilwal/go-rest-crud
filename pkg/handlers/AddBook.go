package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest-crud/pkg/mocks"
	"go-rest-crud/pkg/models"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	// steps

	defer r.Body.Close()
	// Read to request body

	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(body, "read from body is ---->")
	fmt.Println(r.Body, "read from body iss ---->")
	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	fmt.Println("books is", book)

	// body data from request  is map to pointer(address) book struct
	json.Unmarshal(body, &book)
	fmt.Println("books is ----", book)

	book.Id = rand.Intn(100)
	// Append to book mocks
	fmt.Println(mocks.Books, "before append is ---")
	mocks.Books = append(mocks.Books, book)
	fmt.Println(mocks.Books, "After append is ---")

	// Send a 201 created response

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("created")

}
