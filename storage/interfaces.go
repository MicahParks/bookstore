package storage

import (
	"context"
	"errors"

	"github.com/MicahParks/bookstore/models"
)

var (

	// ErrISBNExists indicates a book with the given ISBN was already in the library.
	ErrISBNExists = errors.New("the ISBN was already in the library")

	// ErrISBNNotFound indicates a book with the given ISBN was not found in the library.
	ErrISBNNotFound = errors.New("the ISBN was not found in the library")
)

// BookStore is the Book data storage interface. It allows for Book data storage operations without needing to know how
// the Book data are stored.
type BookStore interface {

	// Close closes the connection to the underlying storage.
	Close(ctx context.Context) (err error)

	// Delete deletes the Book data for the given ISBNs. There will be no error if the ISBNs are not found.
	Delete(ctx context.Context, isbns []string) (err error)

	// Read reads the Book data for the given ISBNs. The error ErrISBNNotFound will be given if an ISBNs is not found.
	Read(ctx context.Context, isbns []string) (books map[string]models.Book, err error)

	// Write writes the given Book data. It will return ErrISBNExists for in Insert operation where the ISBN already
	// exists and an ErrISBNNotFound if an Update operation has an ISBN that doesn't exist.
	Write(ctx context.Context, books []models.Book, operation WriteOperation) (err error)
}

// StoreStore is the Status data storage interface. It allows for Book data storage operations without needing to know how
// the Book data are stored.
type StatusStore interface {

	// Close closes the connection to the underlying storage.
	Close(ctx context.Context) (err error)

	// Delete deletes the Status data for the ISBNs. There will be no error if the ISBNs are not found.
	Delete(ctx context.Context, isbns []string) (err error)

	// Read reads the Status data for the given ISBNs. The error ErrISBNNotFound will be given if an ISBN is not found.
	Read(ctx context.Context, isbns []string) (books map[string]models.Status, err error)

	// Write writes the given Status data. It will return ErrISBNExists for in Insert operation where the ISBN already
	// exists and an ErrISBNNotFound if an Update operation has an ISBN that doesn't exist.
	Write(ctx context.Context, statuses map[string]models.Status, operation WriteOperation) (err error)
}
