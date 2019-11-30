package benchmark

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/ppp225/unjson"
	"github.com/tidwall/gjson"
)

func Benchmark_LargeFile_unjson(b *testing.B) {
	data := unjson.LoadFile("large.json")

	path1 := "categories.performance.score"
	path2 := "audits.screenshot-thumbnails.details.items"
	path3 := "audits.screenshot-thumbnails.details.items.3.timing"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = unjson.Get(data, path1)
		_ = unjson.Get(data, path2)
		_ = unjson.Get(data, path3)
	}
	b.StopTimer()
}

func Benchmark_LargeFile_GJson(b *testing.B) {
	jsonFile, err := os.Open("large.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	path1 := "categories.performance.score"
	path2 := "audits.screenshot-thumbnails.details.items"
	path3 := "audits.screenshot-thumbnails.details.items.3.timing"
	myString := string(byteValue[:])

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tmp := gjson.Get(myString, path1)
		_ = tmp.Value()
		tmp = gjson.Get(myString, path2)
		_ = tmp.Value()
		tmp = gjson.Get(myString, path3)
		_ = tmp.Value()
	}
	b.StopTimer()
}

func Benchmark_SmallFile_unjson(b *testing.B) {
	data := unjson.LoadFile("large.json")

	path1 := "widget.window.name"
	path2 := "widget.image.hOffset"
	path3 := "widget.text.onMouseUp"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = unjson.Get(data, path1)
		_ = unjson.Get(data, path2)
		_ = unjson.Get(data, path3)
	}
	b.StopTimer()
}

func Benchmark_SmallFile_GJson(b *testing.B) {
	jsonFile, err := os.Open("small.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	path1 := "widget.window.name"
	path2 := "widget.image.hOffset"
	path3 := "widget.text.onMouseUp"
	myString := string(byteValue[:])

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tmp := gjson.Get(myString, path1)
		_ = tmp.Value()
		tmp = gjson.Get(myString, path2)
		_ = tmp.Value()
		tmp = gjson.Get(myString, path3)
		_ = tmp.Value()
	}
	b.StopTimer()
}
