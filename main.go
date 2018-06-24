package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Trinergy/gologger"
)

// Schema represents a queryable object that respresents a JSON schema file
type Schema struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

var (
	log    = gologger.SetupLogFile("debug_log.txt")
	logger = gologger.SetupLogger(log)
)

// ParseJSONToSchema takes a filename and converts it to a queryable Schema struct
func ParseJSONToSchema(name string) *Schema {
	f, err := os.Open(name)
	if err != nil {
		logger.Fatal(err)
	}
	return decodeJSON(f)
}

func decodeJSON(f io.Reader) *Schema {
	var s Schema
	dec := json.NewDecoder(f)

	for {
		if err := dec.Decode(&s); err == io.EOF {
			break
		} else if err != nil {
			logger.Fatal(err)
		}
	}

	return &s
}

func main() {
	s := ParseJSONToSchema("example-schema.json")
	logger.Printf("%s", s.Title)
}
