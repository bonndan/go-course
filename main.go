package main

import (
	"fmt"
	"os"
)

func main() {

	backend := FileBackend{Filename:"/tmp/kv-go.gob"}
	store := NewStore(backend)

	if len(os.Args) > 1 {
		for _, op := range parseArgs() {
			op.handle(store)
		}
	} else {
		//dump the whole map
		for key, entry := range store.GetAll() {
			fmt.Printf("%s = %s\n", key, entry)
		}
	}
}
