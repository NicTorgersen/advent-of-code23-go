package file

import (
	"log"
	"os"
)

func OpenOrCreateFile(filepath string) *os.File {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}

	return f
}
