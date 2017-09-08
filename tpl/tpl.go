package tpl

import (
	"io/ioutil"
	"log"
	"github.com/russross/blackfriday"
)

func ConvertMdFile(filename string) {
	source_path := get_source_path(filename)
	dist_path := get_dist_path(filename)
	text, err := ioutil.ReadFile(source_path)

	if err != nil {
		log.Fatal("Read File Error: ", source_path)
	}

	output := blackfriday.Run(text)
	err = ioutil.WriteFile(dist_path, output, 0644)
	if err != nil {
		log.Fatal("Write File Error: ", dist_path)
	}
}
