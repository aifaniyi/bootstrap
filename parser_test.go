package main

import (
	"testing"
)

func TestParser(t *testing.T) {

	file := "schema.sample.json"
	sample, err := getSchema(file)
	if err != nil {
		t.Fatalf("error [%v] while parsing file [%s]", err, file)
	}

	expectedString := "golang"
	if sample.Lang != expectedString {
		fail(t, expectedString, sample.Lang)
	}

	expectedInt := 1
	if len(sample.Entities) != expectedInt {
		fail(t, expectedInt, len(sample.Entities))
	}

	expectedInt = 7
	if len(sample.Entities[0].Properties) != expectedInt {
		fail(t, expectedInt, len(sample.Entities[0].Properties))
	}
}
