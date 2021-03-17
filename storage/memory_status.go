package storage

import (
	"context"
	"sync"

	"github.com/MicahParks/bookstore/models"
)

// MemStatus is a StatusStore implementation that stores everything in a Go map in memory.
type MemStatus struct {
	statuses map[string]*models.History
	mux      sync.RWMutex
}

// NewMemStatus creates a new MemStatus.
func NewMemStatus() (statusStore StatusStore) {
	return &MemStatus{
		statuses: make(map[string]*models.History),
	}
}

// Close closes the connection to the underlying storage.
func (m *MemStatus) Close(_ context.Context) (err error) {

	// Lock the Status data for async safe use.
	m.mux.Lock()
	defer m.mux.Unlock()

	// Delete all Status data.
	m.deleteAll()

	return nil
}

// Delete deletes the Status data for the ISBNs. There will be no error if the ISBNs are not found.
func (m *MemStatus) Delete(_ context.Context, isbns []string) (err error) {

	// Lock the Status data for async safe use.
	m.mux.Lock()
	defer m.mux.Unlock()

	// Check for the empty case.
	if len(isbns) == 0 {

		// Delete all Status data.
		m.deleteAll()
	} else {

		// Iterate through the given ISBNs and delete their statuses.
		for _, isbn := range isbns {
			delete(m.statuses, isbn)
		}
	}

	return nil
}

// Read reads the Status data for the given ISBNs. The error ErrISBNNotFound will be given if an ISBN is not found.
func (m *MemStatus) Read(_ context.Context, isbns []string) (statuses map[string]models.History, err error) {

	// Create the return map.
	statuses = make(map[string]models.History, len(isbns))

	// Lock the Status data for async safe use.
	m.mux.RLock()
	defer m.mux.RUnlock()

	// Check for the empty case.
	if len(isbns) == 0 {

		// Copy all Status data.
		for isbn, status := range m.statuses {
			statuses[isbn] = *status
		}
	} else {

		// Iterate through the give ISBNs. Copy the requested ones.
		for _, isbn := range isbns {
			status, ok := m.statuses[isbn]
			if !ok {
				return nil, ErrISBNNotFound
			}
			statuses[isbn] = *status
		}
	}

	return statuses, nil
}

// Write writes the given Status data. It will return ErrISBNExists for in Insert operation where the ISBN already
// exists and an ErrISBNNotFound if an Update operation has an ISBN that doesn't exist.
func (m *MemStatus) Write(_ context.Context, statuses map[string]models.History, operation WriteOperation) (err error) {

	// Lock the Status data for async safe use.
	m.mux.Lock()
	defer m.mux.Unlock()

	// Iterate through the given statuses.
	for isbn, status := range statuses {

		// Check to see if a status with that ISBN already exists.
		if operation == Insert || operation == Update {
			_, ok := m.statuses[isbn]
			if ok && operation == Insert {
				return ErrISBNExists
			}
			if !ok && operation == Update {
				return ErrISBNNotFound
			}
		}

		// Prevent a leaky buffer.
		status := status

		// Assign the Status data to the ISBN.
		m.statuses[isbn] = &status
	}

	return nil
}

// deleteAll deletes all Status data. It does not lock, so must be locked for async safe use.
func (m *MemStatus) deleteAll() {

	// Reassign the Status data so it's take by the garbage collector.
	m.statuses = make(map[string]*models.History)
}
