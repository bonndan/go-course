package main

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
 * key-value setter
 */
func (store Store) get(key string) string {

	value, exists := store.kvmap[key]
	if exists == false {
		panic("Invalid key")
	}
	return value
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