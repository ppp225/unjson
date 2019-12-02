package unjson

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func BenchmarkParserBracket(b *testing.B) {
	path := "audits.screenshot-thumbnails.details.items[3].timing"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = parsePath2DotNotation(path)
	}
	b.StopTimer()
}

func BenchmarkParserDot(b *testing.B) {
	path := "audits.screenshot-thumbnails.details.items.3.timing"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = parsePath2DotNotation(path)
	}
	b.StopTimer()
}

func BenchmarkGJson(b *testing.B) {
	data := LoadFile("benchmark/large.json")

	path := "audits.screenshot-thumbnails.details.items[3].timing"
	s := parsePath2DotNotation(path)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = Get(data, s)
	}
	b.StopTimer()
}

func BenchmarkRecurse(b *testing.B) {
	path := "audits.screenshot-thumbnails.details.items[3].timing"
	s := parsePath2DotNotation(path)

	jsonFile, err := os.Open("benchmark/large.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data interface{}

	s2 := strings.Split(s, ".")

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		json.Unmarshal(byteValue, &data)
		_ = deeper(data, s2...)
	}
	b.StopTimer()
}
