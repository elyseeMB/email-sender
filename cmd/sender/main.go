package main

import (
	"email-sender/packages/emails"
	"fmt"
	"io/fs"
	"log"
)

func main() {
	err := fs.WalkDir(emails.Templates, "dist", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		fmt.Println(d)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
