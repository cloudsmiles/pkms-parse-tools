package main

import (
	"./internal"

	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	input_file = "../pkms/STL-evaluation.md"

	output_file = "./resource.js"
)

func main() {
	arr, err := internal.ParseSTLMarkdown(input_file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	internal.FillPicture(arr)

	_res, err := json.Marshal(arr)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	_res = append([]byte("export var out = "), _res...)
	err = ioutil.WriteFile(output_file, _res, 0666)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
}
