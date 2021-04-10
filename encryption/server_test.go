package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"priva.te/yoti/dataService"
)

func TestEncryptionServer(t *testing.T) {

	ds := dataService.NewLocalFileStorage("data.txt")
	s := NewEncryptionServer(ds)

	for _, test := range []struct {
		name    string
		idStr   string
		dataStr string
	}{
		{name: "add record 1",
			idStr:   "test ID1",
			dataStr: "test Data1",
		},
		{name: "add record 2",
			idStr:   "test ID2",
			dataStr: "test Data2",
		},
	} {
		t.Run(test.name, func(t *testing.T) {

			// test for encrypted storage of provided plaintext
			aesKey, err := s.Store([]byte(test.idStr), []byte(test.dataStr))
			assert.NoError(t, err)

			// test for decryption and retrieval of the encrypted data
			retrieved, err := s.Retrieve([]byte(test.idStr), aesKey)
			assert.NoError(t, err)

			assert.Equal(t, test.dataStr, string(retrieved))
		})
	}

	// Test for non-existence of record
	_, err := s.DataService.RetrieveData([]byte("9999"))
	assert.Error(t, err)
	assert.EqualError(t, err, "record not found for provided ID: 9999")

	ds.TearDown()
}
