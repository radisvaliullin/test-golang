package ufiles

import (
	"log"
	"os"
)

const (
	pkgLogName = "ufiles"
)

// CreateWrite creates file and write data or fatal.
func CreateWrite(name string, data []byte) {

	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("%v: create file err - %v", pkgLogName, err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		log.Fatalf("%v: create file, write err - %v", pkgLogName, err)
	}
}
