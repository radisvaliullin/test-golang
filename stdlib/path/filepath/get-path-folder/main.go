package main

import (
	"log"
	"path"
	"path/filepath"
)

func main() {

	fpath := "./asdf/qwer.csv"

	log.Print(filepath.Base(fpath))
	log.Print(filepath.Dir(fpath))

	log.Print(path.Base(fpath))
	log.Print(path.Dir(fpath))
}
