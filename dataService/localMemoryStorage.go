package dataService

import (
	"bytes"
	"fmt"
)

// LocalInMemoryStorage provided.
type LocalInMemoryStorage struct {
	dataStore Records
}

func (l *LocalInMemoryStorage) StoreData(r *Record) error {

	l.dataStore = append(l.dataStore, r)

	return nil

}

func (l *LocalInMemoryStorage) RetrieveData(id []byte) (*Record, error) {

	for _, v := range l.dataStore {
		if bytes.Equal(v.ID, id) {
			return v, nil
		}
	}
	return nil, fmt.Errorf("record not found for provided ID: %s", string(id))
}

