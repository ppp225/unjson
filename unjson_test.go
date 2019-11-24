package unjson

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func BenchmarkParser(b *testing.B) {
	path := "audits.screenshot-thumbnails.details.items[3].timing"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = parseJsonPath(path)
	}
	b.StopTimer()
}

func BenchmarkCaster(b *testing.B) {
	jsonFile, err := os.Open("benchmark/large.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var mydata interface{}
	json.Unmarshal(byteValue, &mydata)

	path := "audits.screenshot-thumbnails.details.items[3].timing"
	s := parseJsonPath(path)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = deeper(mydata, s...)
	}
	b.StopTimer()
}
