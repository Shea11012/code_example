package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	err := client.SetLanguage("eng", "chi_sim")
	if err != nil {
		log.Fatalf("set language  err: %v", err)
	}

	err = filepath.Walk("./data/img", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		err = client.SetImage(path)
		if err != nil {
			return fmt.Errorf("set image err %w", err)
		}

		text, err := client.Text()
		if err != nil {
			return fmt.Errorf("text err: %w", err)
		}
		fmt.Printf("%s\t\t%s\n", info.Name(), text)

		return nil
	})

	if err != nil {
		log.Fatalf("filepath walk err: %v", err)
	}
}
