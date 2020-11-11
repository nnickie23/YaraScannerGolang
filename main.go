package main

import (
	//"fmt"
	"os"

	"github.com/nnickie23/YaraScannerGolang/file_reader"
	//"io/ioutil"
	"path/filepath"
)

func main() {

	args := os.Args

	err := filepath.Walk(args[1], file_reader.Is_php)
	if file_reader.IsError(err) {
		return
	}

	var file *os.File
	if len(args) == 4 {
		file = file_reader.Create_file(args[3])
	} else {
		file = os.Stdout
	}

	d2 := []byte{115, 111, 109, 101, 10}
	_, err = file.Write(d2)
	if file_reader.IsError(err) {
		return
	}
    //fmt.Printf("wrote %d bytes\n", n2)
	if file_reader.IsError(err) {
		return
	}

	defer file.Close()
}
