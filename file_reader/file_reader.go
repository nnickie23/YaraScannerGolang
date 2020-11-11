package file_reader

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func yara_scanner(path string, info os.FileInfo, err error) {
	println(path)
}

func Is_php(path string, info os.FileInfo, err error) error {
	extension := filepath.Ext(path)
	if extension == ".php" {
		yara_scanner(path, info, err)
	}
	return (err)
}

func Create_file(path string) *os.File{
	var file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if IsError(err) {
		return nil
	}

	data, err := ioutil.ReadFile(path)
	if IsError(err) {
		return nil
	}
	fmt.Println("Contents of file:", string(data))

	return file
}
