package benchmark

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ppp225/unjson"
	"github.com/tidwall/gjson"
)

func Benchmark_LargeFile_UntypedJSON(b *testing.B) {
	jsonFile, err := os.Open("large.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var mydata interface{}
	json.Unmarshal(byteValue, &mydata)

	path1 := "categories.performance.score"
	path2 := "audits.screenshot-thumbnails.details.items"
	path3 := "audits.screenshot-thumbnails.details.items.3.timing"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = unjson.Get(mydata, path1)
		_ = unjson.Get(mydata, path2)
		_ = unjson.Get(mydata, path3)
	}
	b.StopTimer()
}

func Benchmark_LargeFile_GJson(b *testing.B) {
	jsonFile, err := os.Open("large.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var mydata interface{}
	json.Unmarshal(byteValue, &mydata)

	path1 := "categories.performance.score"
	path2 := "audits.screenshot-thumbnails.details.items"
	path3 := "audits.screenshot-thumbnails.details.items.3.timing"
	myString := string(byteValue[:])

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = gjson.Get(myString, path1)
		_ = gjson.Get(myString, path2)
		_ = gjson.Get(myString, path3)
	}
	b.StopTimer()
}

func Benchmark_SmallFile_UntypedJSON(b *testing.B) {
	jsonFile, err := os.Open("small.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var mydata interface{}
	json.Unmarshal(byteValue, &mydata)

	path1 := "widget.window.name"
	path2 := "widget.image.hOffset"
	path3 := "widget.text.onMouseUp"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = unjson.Get(mydata, path1)
		_ = unjson.Get(mydata, path2)
		_ = unjson.Get(mydata, path3)
	}
	b.StopTimer()
}

func Benchmark_SmallFile_GJson(b *testing.B) {
	jsonFile, err := os.Open("small.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var mydata interface{}
	json.Unmarshal(byteValue, &mydata)

	path1 := "widget.window.name"
	path2 := "widget.image.hOffset"
	path3 := "widget.text.onMouseUp"
	myString := string(byteValue[:])

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = gjson.Get(myString, path1)
		_ = gjson.Get(myString, path2)
		_ = gjson.Get(myString, path3)
	}
	b.StopTimer()
}
