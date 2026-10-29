package main

import (
	"flag"
	"fmt"
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
		fmt.Print(err)
		return
	}

	executor := executor.NewExecutor()
	results, err := executor.Run(tests)
	if err != nil {
		fmt.Print(err)
		return
	}

	view := view.NewView()
	view.Show(results)
}

func InitConfig() *Config {
	filename := flag.String("filename", "t", "file with tests")
	config := &Config{
		Filename: *filename,
	}
	flag.Parse()
	return config
}
