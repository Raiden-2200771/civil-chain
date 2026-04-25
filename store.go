package main

import (
	"encoding/json"
	"os"
)

func saveChain(bc Blockchain, path string) error {
	data, err := json.Marshal(bc)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
