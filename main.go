package main

import (
	"github.com/bonfy/bugo/utils"
	"fmt"
	"log"
	"github.com/bonfy/bugo/tpl"
)

func main() {
	h, err := utils.Md5File("README.md")
	if err != nil {
		log.Fatal("Error occur:", err)
	}
	fmt.Println("hash:", h)

	tpl.ConvertMdFile("bb.md")
}


