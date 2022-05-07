package main

import (
	"./internal"

	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	sourceUrl = "https://raw.githubusercontent.com/HsinChang/pkms/main/STL-evaluation.md"

	inputFile = "./static/STL-evaluation.md"

	outputFile = "./static/resource.js"

	pictureFile = "./static/picture.txt"
)

func main() {
	err := internal.DownloadMarkdown(sourceUrl, inputFile)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	arr, err := internal.ParseSTLMarkdown(inputFile)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	internal.FillPicture(arr, pictureFile)

	_res, err := json.Marshal(arr)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	_res = append([]byte("export var out = "), _res...)
	err = ioutil.WriteFile(outputFile, _res, 0666)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
}
