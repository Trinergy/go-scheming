package main

import (
	"encoding/json"
	"os"

	"github.com/Trinergy/gologger"
)

type Schema struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

var (
	log    = gologger.SetupLogFile("debug_log.txt")
	logger = gologger.SetupLogger(log)
)

func ParseFile(fp string) {
	f, err := os.Open(fp)
	if err != nil {
		panic("ParseFile: something went wrong")
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	logger.Println(decoder)
}

func main() {
	ParseFile("example-schema.json")
}
