package main

import (
	"fmt"
)

type BackendInterface interface {
	
	/**
	 * stores the map
	 */
	StoreMap(kvmap map[string]string)

	/**
	 * Reads the map
	 */
	ReadMap() map[string]string
}

type Store struct {
	backend BackendInterface
	kvmap map[string]string  
}

/**
 * Factory method
 */
func NewStore(backend BackendInterface) *Store {
	return &Store {
		backend: backend,
		kvmap: backend.ReadMap(),
	}
}

/**
 * Handles an operation on the map (read/write).
 */
func (store Store) handle(op Op) {

	if op.mode == OP_READ {
		val, exist := store.kvmap[op.key]
		if !exist {
			fmt.Printf("%s is an invalid key\n", op.key)
		}
		fmt.Printf("%s\n", val)
	}

	if op.mode == OP_WRITE {
		fmt.Printf("Adding %s as %s.\n", op.key, op.value)
		store.set(op.key, op.value)
	}
}

/**
 * key-value setter
 */
func (store Store) set(key string, value string) {
	store.kvmap[key] = value
	store.backend.StoreMap(store.kvmap)
}

/**
 * Returns the whole map.
 */
func (store Store) GetAll() map[string]string {
	return store.kvmap
}