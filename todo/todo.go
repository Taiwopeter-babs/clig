package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

type Item struct {
	Text     string `json:"text"`
	Priority int
	Position int
	Done     bool
}

// ByPri implements sort.Interface for []Item based on the
// Priority and Position fields.
type ByPriority []Item

func (sequence ByPriority) Len() int { return len(sequence) }

func (sequence ByPriority) Swap(i, j int) { sequence[i], sequence[j] = sequence[j], sequence[i] }

func (sequence ByPriority) Less(i, j int) bool {

	if sequence[i].Done != sequence[j].Done {
		return sequence[i].Done
	}

	if sequence[i].Priority != sequence[j].Priority {
		return sequence[i].Priority < sequence[j].Priority
	}

	return sequence[i].Position < sequence[j].Position
}

// Save Items to a json file in the provided directory
func SaveItems(filename string, items []Item) error {
	file := createFile(filename)

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

// Read Todo items from the file. If the files does not exist
// an error will be returned with an empty Item list
func ReadItems(filename string) ([]Item, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return []Item{}, err
	}

	var items []Item

	if err := json.Unmarshal(data, &items); err != nil {
		return []Item{}, err
	}

	for index := range items {
		items[index].Position = index + 1
	}

	return items, nil
}

// Sets the priority on the Todo [Item]
func (item *Item) SetPriority(priority int) {
	switch priority {
	case 1:
		item.Priority = 1
	case 3:
		item.Priority = 3
	default:
		item.Priority = 2
	}
}

func (item *Item) PrettyP() string {
	switch item.Priority {
	case 1:
		return "(1)"
	case 3:
		return "(3)"
	default:
		return " "
	}
}

func (item *Item) Label() string {
	return strconv.Itoa(item.Position) + "."
}

func (item *Item) PrettyDone() string {
	switch item.Done {
	case true:
		return "X"
	default:
		return ""
	}
}

func createFile(filename string) *os.File {
	var file *os.File

	file, fileErr := os.Open(filename)

	if fileErr != nil && errors.Is(fileErr, fs.ErrNotExist) {
		file, _ = os.Create(filename)

		return file
	}

	return file
}

func init() {
	createFile(*AllConstants.Filename)
}
