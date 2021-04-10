package dataService

// Record contains the ID and Data stored as a unit.
type Record struct {
	ID   []byte
	Data []byte
}

// Records is a slice of Record.
type Records []*Record

// EncryptedDataStorage provides storage for encrypted data
type EncryptedDataStorage interface {
	StoreData(r *Record) error
	RetrieveData(id []byte) (*Record, error)
}
