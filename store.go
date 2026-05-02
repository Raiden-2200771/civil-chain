package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func chainFilePath(port string) string {
	return fmt.Sprintf("data/chain_%s.json", port)
}

func saveChain(bc Blockchain, path string) error {
	data, err := json.Marshal(bc)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
