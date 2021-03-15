package storage

import (
	"context"
	"sync"

	"github.com/MicahParks/bookstore/models"
)

// MemBook is a BookStore implementation that stores everything in a Go map in memory.
type MemBook struct {
	books map[string]*models.Book
	mux   sync.RWMutex
}

// NewMemBook creates a new MemBook.
func NewMemBook() (bookStore BookStore) {
	return &MemBook{
		books: make(map[string]*models.Book),
	}
}

func (m *MemBook) Close(_ context.Context) (err error) {

	// Lock the Book data for async safe use.
	m.mux.Lock()
	defer m.mux.Unlock()

	// Delete all Book data.
	m.deleteAll()

	return nil
}

func (m *MemBook) Delete(_ context.Context, isbns []string) (err error) {

	// Lock the Book data for async safe use.
	m.mux.Lock()
	defer m.mux.Unlock()

	// Check for the empty case.
	if len(isbns) == 0 {

		// Delete all Book data.
		m.deleteAll()
	} else {

		// Iterate through the given ISBNs and delete their statuses.
		for _, isbn := range isbns {
			delete(m.books, isbn)
		}
	}

	return nil
}

func (m *MemBook) Read(_ context.Context, isbns []string) (books map[string]models.Book, err error) {

	// Create the return map.
	books = make(map[string]models.Book, len(isbns))

	// Lock the book data for async safe use.
	m.mux.RLock()
	defer m.mux.RUnlock()

	// Check for the empty case.
	if len(isbns) == 0 {

		// Copy all book data.
		for isbn, book := range m.books {
			books[isbn] = *book
		}
	} else {

		// Iterate through the give ISBNs. Copy the requested ones.
		for _, isbn := range isbns {
			book, ok := m.books[isbn]
			if !ok {
				return nil, ErrISBNNotFound
			}
			books[isbn] = *book
		}
	}

	return books, nil
}

func (m *MemBook) Write(_ context.Context, books []models.Book, operation WriteOperation) (err error) {

	// Lock the Book data for async safe use.
	m.mux.Lock()
	defer m.mux.Unlock()

	// Iterate through the given statuses.
	for _, book := range books {

		// Check to see if a book with that ISBN already exists.
		if operation == Insert || operation == Update {
			_, ok := m.books[book.ISBN]
			if ok && operation == Insert {
				return ErrISBNExists
			}
			if !ok && operation == Update {
				return ErrISBNNotFound
			}
		}

		// Assign the Book data to the ISBN.
		m.books[book.ISBN] = &book
	}

	return nil
}

// deleteAll deletes all Book data. It does not lock, so must be locked for async safe use.
func (m *MemBook) deleteAll() {

	// Reassign the Book data so it's take by the garbage collector.
	m.books = make(map[string]*models.Book)
}
