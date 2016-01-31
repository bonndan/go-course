package main

import (
	"os"
	"strings"
)

type Op struct {
	key   string
	value string
	mode  string
}

const (
	OP_READ  = "read"
	OP_WRITE = "write"
)

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
			tmp := strings.SplitN(s, "=", 2)
			op := Op{
				key:   tmp[0],
				value: tmp[1],
				mode:  OP_WRITE,
			}
			m = append(m, op)
		} else {
			op := Op{
				key:  s,
				mode: OP_READ,
			}
			m = append(m, op)
		}
	}

	//fmt.Println("args map: %s ", m) // For debug
	return m
}
