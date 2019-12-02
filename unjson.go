package unjson

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

// LoadFile reads file from disk
func LoadFile(filename string) string {
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	jsonBytes, _ := ioutil.ReadAll(jsonFile)
	return string(jsonBytes[:])
}

// Get returns object found in path
func Get(data string, path string) interface{} {
	dotPath := parsePath2DotNotation(path)
	tmp := gjson.Get(data, dotPath)
	return tmp.Value()
}

func parsePath2DotNotation(path string) (result string) {
	result = strings.ReplaceAll(path, "[", ".")
	result = strings.ReplaceAll(result, "]", "")

	return
}

func deeper(data interface{}, params ...string) interface{} {
	if len(params) == 0 {
		return data
	}
	var next interface{}

	index, err := strconv.Atoi(params[0])
	if err != nil { // case string
		sub, ok := data.(map[string]interface{})
		if !ok {
			fmt.Printf("conversion error: interface is %T, not map[string]interface{}\n", data)
		}
		next = sub[params[0]]

	} else { // case int
		sub, ok := data.([]interface{})
		if !ok {
			fmt.Printf("conversion error: interface is %T, not []interface{}\n", data)
		}
		next = sub[index]
		if next == nil {
			fmt.Printf("invalid path: %q is nil\n", params[0])
			return ""
		}
	}

	if next == nil {
		fmt.Printf("invalid path: %q is nil\n", params[0])
		return ""
	}
	return deeper(next, params[1:]...)
}
