package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"github.com/nickbrockett/yoti/dataService"
)

// Server is an encryption-server which implements the Client interface.
type Server struct {

	// needs a data service
	DataService dataService.EncryptedDataStorage
}

// NewEncryptionServer creates new server with provided dataStorage.
func NewEncryptionServer(ds dataService.EncryptedDataStorage) Server {
	return Server{DataService: ds}
}

// Encrypt converts plainText to cipherText.
func (s *Server) Encrypt(data []byte) ([]byte, []byte, error) {

	// obtain aesKey
	aesKey, err := getAESKey()
	if err != nil {
		return nil, nil, err
	}

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, nil, err
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	//Encrypt the data using aesGCM.Seal within the record
	data = aesGCM.Seal(nonce, nonce, data, nil)

	return data, aesKey, nil

}

// Decrypt converts cipherText to plainText.
func (s *Server) Decrypt(encryptedData, aesKey []byte) ([]byte, error) {

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(encryptedData) < nonceSize {
		return nil, errors.New("data length incompatible with nonceSize")
	}
	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil

}

//Store implements the Client.Store interface.
func (s Server) Store(id, payload []byte) (aesKey []byte, err error) {

	var encryptedData []byte

	encryptedData, aesKey, err = s.Encrypt(payload)
	if err != nil {
		return nil, err
	}

	record := &dataService.Record{
		ID:   id,
		Data: encryptedData}

	return aesKey, s.DataService.StoreData(record)

}

//Retrieve implements the Client.Retrieve interface.
func (s Server) Retrieve(id, aesKey []byte) (payload []byte, err error) {

	record, err := s.DataService.RetrieveData(id)
	if err != nil {
		return nil, err
	}

	return s.Decrypt(record.Data, aesKey)

}

// getAESKey acquires a random 32 byte key.
func getAESKey() ([]byte, error) {
	aesKey := make([]byte, 32)
	_, err := rand.Read(aesKey)
	return aesKey, err
}
