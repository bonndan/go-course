package main

import (
	"os"
	"io"
	"encoding/gob"
)

type FileBackend struct {
	Filename string
	file *os.File
}

/**
 * Returns the file
 */
func (backend FileBackend) getFile() *os.File {
	file, err := os.OpenFile(backend.Filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	return file
}

/**
 * Encodes the maps and writes it to the file
 */
func (backend FileBackend) StoreMap(kvmap map[string]string) {

	file := backend.getFile()
	file.Seek(0, 0)
	enc := gob.NewEncoder(file)
	err := enc.Encode(kvmap)

	if err != nil {
		panic(err)
	}

	file.Close()
}

/**
 * Reads the map from the file
 */
func (backend FileBackend) ReadMap() map[string]string {
	
	file := backend.getFile()
	file.Seek(0, 0)
	decoder := gob.NewDecoder(file)
	kvmap := make(map[string]string)

	// Decode -- We need to pass a pointer otherwise kvmap isn't modified
	err := decoder.Decode(&kvmap)
	if err != nil {
		if (err == io.EOF) {
			return kvmap
		}
		panic(err)
	}

	file.Close()
	return kvmap
}
