package dataService

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
)

// LocalFileStorage provided.
type LocalFileStorage struct {
	filename  string
}

func NewLocalFileStorage(filename string) *LocalFileStorage {
	return &LocalFileStorage{
		filename: filename,
	}
}

func (l *LocalFileStorage) StoreData(r *Record) error {

	f, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	logger := log.New(f, "", 0)
	// write record as a line in the file, whitespace field separator.
	logger.Printf("%s %s \n", base64.StdEncoding.EncodeToString(r.ID), base64.StdEncoding.EncodeToString(r.Data))
	return nil

}

func (l *LocalFileStorage) RetrieveData(id []byte) (*Record, error) {

	record := &Record{}

	// read file contents
	f, err := os.OpenFile(l.filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewScanner(f)
	for reader.Scan() {
		s := strings.Split(reader.Text(), " ")
		rowID, _ := base64.StdEncoding.DecodeString(s[0])
		if bytes.Equal(id, rowID)  {
			data, err := base64.StdEncoding.DecodeString(s[1])
			if err != nil {
				fmt.Printf("error decoding string %s ", err.Error())
				break
			}
			record.ID = rowID
			record.Data = data
			return record, nil
		}
	}

	return record, fmt.Errorf("record not found for provided ID: %s", string(id))
}

func (l *LocalFileStorage) TearDown() {
	os.Remove(l.filename)
}
