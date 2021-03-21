package generator

import "testing"

func TestGenerateModels(t *testing.T) {
	file := "schema.sample.json"
	sample, err := getSchema(file)
	if err != nil {
		t.Fatalf("error [%v] while parsing file [%s]", err, file)
	}

	entities, err := generateGolangModels(sample)
	if err != nil {
		t.Fatalf("error creating entities: %v", err)
	}

	expectedInt := 2
	if len(entities) != expectedInt {
		fail(t, expectedInt, len(entities))
	}
}
