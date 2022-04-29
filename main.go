package main

import (
	"encoding/json"
	"fmt"

	"github.com/a-poor/data-stream-viz/lib/dsviz"
	"github.com/niemeyer/pretty"
)

func main() {
	fmt.Println("Hello, world!")

	raw := `{
		"foo": 1,
		"baz": [
			true,
			{"name": "Test"},
			{}
		],
		"info": { "faveColor": null }
	}`
	var d map[string]any
	err := json.Unmarshal([]byte(raw), &d)
	if err != nil {
		panic(err)
	}
	fmt.Println("Source data:")
	pretty.Println(d)
	fmt.Println()

	root := dsviz.NewObject()
	err = root.Add(d)
	if err != nil {
		panic(err)
	}
	fmt.Println("Schema:")
	pretty.Println(root)
}
