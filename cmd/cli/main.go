package main

import (
	"flag"
	"os"

	"github.com/bushmv/H/internal/executor"
	"github.com/bushmv/H/internal/parser"
	"github.com/bushmv/H/internal/view"
)

type Config struct {
	Filename string
}

func main() {
	config := InitConfig()
	file, err := os.Open(config.Filename)
	if err != nil {
		panic(err)
	}
	parser := parser.NewParser()
	tests, err := parser.Parse(file, config.Filename)
	if err != nil {
		panic(err)
	}

	storage := executor.NewStorage()
	templateProcessor := executor.NewTemplateProcessor()
	executor := executor.NewExecutor(templateProcessor)
	results, err := executor.Run(tests, storage)
	if err != nil {
		panic(err)
	}

	view := view.NewView()
	view.Show(results)

}

func InitConfig() *Config {
	filename := flag.String("filename", "test", "file with tests")
	config := &Config{
		Filename: *filename,
	}
	flag.Parse()
	return config
}
