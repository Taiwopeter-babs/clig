package todo

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
)

type Item struct {
	Text string `json:"text"`
}

// Save Items to a json file in the `/tmp` directory of the linux fs
// If the file is not already created, it
func SaveItems(filename string, items []Item) error {
	var file *os.File

	file, fileErr := os.Open(filename)

	if fileErr != nil && errors.Is(fileErr, fs.ErrNotExist) {
		os.Create(filename)
	}

	defer file.Close()

	data, err := json.Marshal(items)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return []Item{}, err
	}

	var items []Item

	if err := json.Unmarshal(data, &items); err != nil {
		return []Item{}, err
	}

	return items, nil
}
