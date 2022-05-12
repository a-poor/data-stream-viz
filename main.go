package main

import (
	"encoding/json"
	"fmt"

	"github.com/a-poor/data-stream-viz/pkg/dsviz"
	"github.com/niemeyer/pretty"
)

func toJSON(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	raw := `{
		"foo": 1,
		"baz": [
			true,
			{"name": "Test"},
			{}
		],
		"info": { "faveColor": [ null ] }
	}`

	var d map[string]any
	err := json.Unmarshal([]byte(raw), &d)
	if err != nil {
		panic(err)
	}
	fmt.Println("Source data:")
	pretty.Println(toJSON(d))
	fmt.Println()

	root := dsviz.NewObject()
	err = root.Add(d)
	if err != nil {
		panic(err)
	}
	fmt.Println("Schema:")
	// pretty.Println(root)
	fmt.Println(toJSON(root))
}
