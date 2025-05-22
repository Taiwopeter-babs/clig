package todo

import (
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

type TodoConstants struct {
	Filename *string
}

var (
	AllConstants *TodoConstants
)

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory.")
	}

	todoFile := home + string(os.PathSeparator) + "clig-todos.json"

	AllConstants = &TodoConstants{
		Filename: &todoFile,
	}

}
