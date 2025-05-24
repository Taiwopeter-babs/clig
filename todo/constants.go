package todo

import (
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

type TodoConstants struct {
	Filename           *string
	ConfigDataFileName *string
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
	configDatafileName := "datafile"

	AllConstants = &TodoConstants{
		Filename:           &todoFile,
		ConfigDataFileName: &configDatafileName,
	}

}
