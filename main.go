package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/gorules/zen-go"
)

func workingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		panic("Unable to determine working directory")
	}

	return wd
}

func decisionLoader(key string) ([]byte, error) {
	fmt.Println(key)
	decisionPath := path.Join(workingDirectory(), "test-data", key)
	fmt.Println(decisionPath)
	decisionContent, err := os.ReadFile(decisionPath)
	if err != nil {
		return nil, errors.New("file not found")
	}

	return decisionContent, nil
}

func main() {
	engine := zen.NewEngine(zen.EngineConfig{
		Loader: decisionLoader,
	})

	response, err := engine.Evaluate("signup/graph.json", map[string]any{"residence": "es", "account": "real"})
	if err != nil {
		panic(err)
	}

	panic(response)
}
