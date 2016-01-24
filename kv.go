package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"strings"
)

func main() {

	file := getFile()

	kvmap := readMap(file)
	if len(os.Args) > 1 {
		for _, op := range parseArgs() {
			handle(op, kvmap)
		}
		//save to file after modifying
		storeMap(kvmap, file)
	} else {
		//dump the whole map
		for key, entry := range kvmap {
			fmt.Printf("%s = %s\n", key, entry)
		}
	}

	file.Close()
}

type Op struct {
	key   string
	value string
	mode  string
}

/**
 * Parses the arguments into a list of read/write operations.
 */
func parseArgs() []Op {

	m := make([]Op, len(os.Args)-1)

	for i, s := range os.Args {
		if i == 0 {
			continue
		}

		if strings.Contains(s, "=") {
			tmp := strings.Split(s, "=")
			op := Op{
				key:   tmp[0],
				value: tmp[1],
				mode:  "write",
			}
			m = append(m, op)
		} else {
			op := Op{
				key:  s,
				mode: "read",
			}
			m = append(m, op)
		}
	}

	//fmt.Println("args map: %s ", m) // For debug
	return m
}

/**
 * Handles an operation on the map (read/write).
 */
func handle(op Op, kvmap map[string]string) {

	if op.mode == "read" {
		val, exist := kvmap[op.key]
		if !exist {
			fmt.Printf("%s is an invalid key\n", op.key)
		}
		fmt.Printf("%s\n", val)
	}

	if op.mode == "write" {
		fmt.Printf("Adding %s as %s.\n", op.key, op.value)
		kvmap[op.key] = op.value
	}

}

/**
 * Returns a file handler. Creates the file if necessary.
 */
func getFile() (handler *os.File) {

	fileName := "/tmp/kv-go.gob"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	//Do not defer file.Close() here, called when function is left
	return file
}

/**
 * Encodes the maps and writes it to the file
 */
func storeMap(kvmap map[string]string, handler *os.File) {

	handler.Seek(0, 0)
	enc := gob.NewEncoder(handler)
	err := enc.Encode(kvmap)

	if err != nil {
		panic(err)
	}

	//	for key, value := range kvmap {
	//		fmt.Printf("writing %s %s\n", key, value)
	//	}
}

/**
 * Reads the map from the file
 */
func readMap(handler *os.File) map[string]string {
	handler.Seek(0, 0)
	decoder := gob.NewDecoder(handler)
	kvmap := make(map[string]string)

	// Decode -- We need to pass a pointer otherwise kvmap isn't modified
	err := decoder.Decode(&kvmap)
	if err != nil {
		panic(err)
	}

	//	for _, entry := range kvmap {
	//		fmt.Printf("read %s", entry)
	//	}
	return kvmap
}
