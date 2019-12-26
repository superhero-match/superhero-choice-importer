package main

import (
	"fmt"

	"github.com/superhero-choice-importer/cmd/importer/importer"
	"github.com/superhero-choice-importer/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	i, err := importer.NewImporter(cfg)
	if err != nil {
		panic(err)
	}

	err = i.Import()
	if err != nil {
		panic(err)
	}

	fmt.Println("Import succeeded!!!")
}